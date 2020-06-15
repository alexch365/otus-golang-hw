package main

import "text/template"

const (
	beginningTemplate = `{{define "beginning"}}// Code generated by go-validate. DO NOT EDIT.

package {{ . }}

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

type ValidationError struct {
	Field string
	Err   error
}

{{end}}
`
	validateRegexpTemplate = `{{define "validateRegexp"}}
func validateRegexp(value string, regex string) error {
	if re := regexp.MustCompile(regex); !re.MatchString(value) {
		return errors.New("value doesn't match regexp")
	}
	return nil
}
{{end}}
`
	validateInTemplate = `{{define "validateIn"}}
func validateIn(value interface{}, validValues string) error {
	validValuesArray := strings.Split(validValues, ",")
	for _, str := range validValuesArray {
		if fmt.Sprintf("%v", value) == str {
			return nil
		}		
	}
	return errors.New("value is not permitted")
}
{{end}}
`
	validateLenTemplate = `{{define "validateLen"}}
func validateLen(value string, validLen string) error {
	if fmt.Sprintf("%d", len(value)) != validLen {
		return errors.New("value length is not equal to " + validLen)
	}
	return nil
}
{{end}}
`
	validateMinTemplate = `{{define "validateMin"}}
func validateMin(value int, minValue string) error {
	if fmt.Sprintf("%d", value) < minValue {
		return errors.New("value must be higher than " + minValue)
	}
	return nil
}
{{end}}
`
	validateMaxTemplate = `{{define "validateMax"}}
func validateMax(value int, maxValue string) error {
	if fmt.Sprintf("%d", value) > maxValue {
		return errors.New("value must be lower than " + maxValue)
	}
	return nil
}
{{end}}
`
	funcTemplate = `{{define "func"}}
func (obj {{.Name}}) Validate() ([]ValidationError, error) {
	var err error
	var errors []ValidationError
	{{range .Fields -}}
		{{$fieldName := .Name -}}
		{{$fieldType := .Type -}}
		{{range .ValidationFuncs -}}
			{{if $fieldType.IsArrayType}}
	for _, val := range obj.{{$fieldName}} {
		err = {{.Name}}(val, "{{.Args}}")
		if err != nil {
			errors = append(errors, ValidationError{"{{$fieldName}}", err})
		}
	}
			{{else}}
	err = {{.Name}}(obj.{{$fieldName}}, "{{.Args}}")
	if err != nil {
		errors = append(errors, ValidationError{"{{$fieldName}}", err})
	}
			{{- end}}
		{{- end}}
	{{- end}}
	return errors, nil
}
{{end}}
`
)

func prepareTemplates() (tmpl *template.Template, err error) {
	tmpl = template.New("Generated template")
	templates := []string{beginningTemplate, validateLenTemplate, validateRegexpTemplate, validateInTemplate,
		validateMinTemplate, validateMaxTemplate, funcTemplate}
	for _, tmplText := range templates {
		_, err = tmpl.Parse(tmplText)
		if err != nil {
			return nil, err
		}
	}

	return
}
