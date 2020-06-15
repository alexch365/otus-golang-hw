package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"strings"
)

type (
	TypeDecl struct {
		Name        string
		IsArrayType bool
	}
	ValidationFuncDecl struct {
		Name string
		Args interface{}
	}
	FieldDecl struct {
		Name            string
		Type            TypeDecl
		ValidationFuncs []ValidationFuncDecl
	}
	StructDecl struct {
		Name   string
		Fields []FieldDecl
	}
)

var generatedFile *os.File
var usedFuncTemplates []string

func main() {
	fileName := os.Args[1]
	if fileName == "" {
		log.Fatalln("Missing file name")
	}

	fileSet := token.NewFileSet()
	fileParsed, err := parser.ParseFile(fileSet, fileName, nil, 0)
	if err != nil {
		log.Fatalln(err)
	}

	generatedFile, err = os.Create(fileParsed.Name.Name + "_validation_generated.go")
	if err != nil {
		log.Fatalln(err)
	}
	defer generatedFile.Close()

	fileTemplate, err := prepareTemplates()
	if err != nil {
		log.Fatalln(err)
	}
	err = fileTemplate.ExecuteTemplate(generatedFile, "beginning", fileParsed.Name.Name)
	if err != nil {
		log.Fatalln(err)
	}

	structsToValidate := inspectFile(fileParsed)
	for _, structDecl := range structsToValidate {
		err = fileTemplate.ExecuteTemplate(generatedFile, "func", structDecl)
		if err != nil {
			log.Fatalln(err)
		}
	}

	for _, funcName := range usedFuncTemplates {
		err = fileTemplate.ExecuteTemplate(generatedFile, funcName, nil)
		if err != nil {
			log.Fatalln(err)
		}
	}

	os.Exit(0)
}

func inspectFile(file *ast.File) (structsToValidate []StructDecl) {
	var structToValidate StructDecl
	typeAliases = make(map[string]TypeDecl)

	ast.Inspect(file, func(node ast.Node) bool {
		typeSpec, ok := node.(*ast.TypeSpec)
		if ok {
			switch currentType := typeSpec.Type.(type) {
			case *ast.StructType:
				structToValidate = StructDecl{Name: typeSpec.Name.Name}
			default:
				if typeDecl := constructTypeDecl(currentType); typeDecl != (TypeDecl{}) {
					typeAliases[typeSpec.Name.Name] = typeDecl
				}
			}
			return true
		}

		st, ok := node.(*ast.StructType)
		if !ok {
			return true
		}

		for _, field := range st.Fields.List {
			if field.Tag == nil {
				continue
			}
			validationTag := parseValidationTag(field.Tag)
			if validationTag == "" {
				continue
			}

			fieldToValidate := FieldDecl{
				Name: field.Names[0].Name,
				Type: constructTypeDecl(field.Type),
			}
			if fieldToValidate.Type == (TypeDecl{}) {
				continue // unsupported field type
			}

			validationStrings := strings.Split(validationTag, "|")
			for _, validationStr := range validationStrings {
				funcDecl, err := constructValidationFuncDecl(validationStr)
				if err != nil {
					continue
				}

				fieldToValidate.ValidationFuncs = append(fieldToValidate.ValidationFuncs, funcDecl)
				addToUsedFuncTemplates(funcDecl.Name)
			}
			structToValidate.Fields = append(structToValidate.Fields, fieldToValidate)
		}
		if len(structToValidate.Fields) > 0 {
			structsToValidate = append(structsToValidate, structToValidate)
		}
		return true
	})

	return
}

func addToUsedFuncTemplates(funcName string) {
	isNewFunc := true
	for _, f := range usedFuncTemplates {
		if f == funcName {
			isNewFunc = false
			break
		}
	}

	if isNewFunc {
		usedFuncTemplates = append(usedFuncTemplates, funcName)
	}
}
