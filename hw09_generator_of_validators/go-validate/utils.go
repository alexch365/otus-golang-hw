package main

import (
	"go/ast"
	"reflect"
	"strconv"
	"strings"
)

var typeAliases map[string]TypeDecl

func parseValidationTag(tag *ast.BasicLit) string {
	tagValues := reflect.StructTag(strings.Trim(tag.Value, "`"))
	validationStr, ok := tagValues.Lookup("validate")
	if !ok {
		return ""
	}

	return validationStr
}

func constructTypeDecl(fieldType ast.Expr) TypeDecl {
	supportedTypes := []string{"int", "string"}
	typeDecl := TypeDecl{}
	switch currentType := fieldType.(type) {
	case *ast.Ident:
		typeDecl = TypeDecl{currentType.Name, false}
	case *ast.ArrayType:
		typeDecl = TypeDecl{currentType.Elt.(*ast.Ident).Name, true}
	}

	if typeAliasDecl := typeAliases[typeDecl.Name]; len(typeAliases) > 0 && typeAliasDecl != (TypeDecl{}) {
		if typeAliasDecl.IsArrayType && typeDecl.IsArrayType {
			return TypeDecl{} // double array types is unsupported
		}
		typeDecl.Name = typeAliasDecl.Name
		return typeDecl
	}

	for _, supportedType := range supportedTypes {
		typeSupported := typeDecl.Name == supportedType
		if typeSupported {
			return typeDecl
		}
	}

	return TypeDecl{}
}

func constructValidationFuncDecl(validationStr string) (ValidationFuncDecl, error) {
	funcDeclArr := strings.Split(validationStr, ":")
	funcName := "validate" + strings.Title(funcDeclArr[0])

	var funcArgs interface{}
	var err error
	switch funcName {
	case "validateRegexp":
		funcArgs = strings.ReplaceAll(funcDeclArr[1], `\`, `\\`)
	case "validateLen", "validateMin", "validateMax":
		funcArgs, err = strconv.Atoi(funcDeclArr[1])
		if err != nil {
			return ValidationFuncDecl{}, err
		}
	default:
		funcArgs = funcDeclArr[1]
	}

	return ValidationFuncDecl{funcName, funcArgs}, nil
}
