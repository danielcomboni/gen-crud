package models

import "testing"
import "github.com/iancoleman/strcase"

func TestGenerateModel_GetModelName(t *testing.T) {
	println(strcase.ToCamel("onetwo"))

	v := new(GenerateModel).SetModelArgs([]string{"one", "two"})
	println(v.GetModelName())
	s := "hello"
	println(trimLeftChar(s))
	println(trimLeftChar(s))
}
