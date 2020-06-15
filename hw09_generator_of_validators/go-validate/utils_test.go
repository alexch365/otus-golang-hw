package main

import (
	"go/ast"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseValidationTag(t *testing.T) {
	t.Run("correct validation tag", func(t *testing.T) {
		tag := ast.BasicLit{Value: `json:"id" validate:"len:36"`}
		parsedTag := parseValidationTag(&tag)

		require.Equal(t, "len:36", parsedTag)
	})

	t.Run("no validation tag", func(t *testing.T) {
		tag := ast.BasicLit{Value: `json:"id"`}
		parsedTag := parseValidationTag(&tag)

		require.Empty(t, parsedTag)
	})
}

func TestConstructTypeDecl(t *testing.T) {
	t.Run("type is string", func(t *testing.T) {
		typeDecl := constructTypeDecl(ast.NewIdent("string"))
		require.EqualValues(t, TypeDecl{"string", false}, typeDecl)
	})

	t.Run("type is strings slice", func(t *testing.T) {
		typeDecl := constructTypeDecl(&ast.ArrayType{Elt: ast.NewIdent("string")})
		require.EqualValues(t, TypeDecl{"string", true}, typeDecl)
	})

	t.Run("type is unsupported", func(t *testing.T) {
		typeDecl := constructTypeDecl(&ast.ArrayType{Elt: ast.NewIdent("bool")})
		require.Equal(t, TypeDecl{}, typeDecl)
	})
}

func TestConstructValidationFuncDecl(t *testing.T) {
	t.Run("validation string is correct `in` func", func(t *testing.T) {
		funcDecl, err := constructValidationFuncDecl("in:admin,stuff")
		require.Nil(t, err)
		require.EqualValues(t, ValidationFuncDecl{
			Name: "validateIn",
			Args: "admin,stuff",
		}, funcDecl)
	})

	t.Run("validation string is correct `regexp` func", func(t *testing.T) {
		funcDecl, err := constructValidationFuncDecl(`regexp:^\w+@\w+\.\w+$`)
		require.Nil(t, err)
		require.EqualValues(t, ValidationFuncDecl{
			Name: "validateRegexp",
			Args: `^\\w+@\\w+\\.\\w+$`,
		}, funcDecl)
	})

	t.Run("validation string is incorrect `len` func", func(t *testing.T) {
		funcDecl, err := constructValidationFuncDecl("len:e2")
		require.NotNil(t, err)
		require.Equal(t, ValidationFuncDecl{}, funcDecl)
	})
}
