package models

import (
	"fmt"
	"github.com/iancoleman/strcase"
	//"github.com/samber/lo"
	"log"
	"strings"
)

// sample:  cli-1 model PaySheet p_Name-string p_Department-string p_PaySlip-[]PaySlip  package_models

type GenerateModel struct {
	Args        []string
	Struct      string
	PackageName string
}

func (g *GenerateModel) GetModelName() string {
	return g.Args[0]
}

func (g *GenerateModel) SetModelName(str string) *GenerateModel {
	g.Args[0] = strcase.ToCamel(str)
	return g
}

func (g *GenerateModel) SetModelArgs(args []string) *GenerateModel {
	g.Args = args
	if len(args) > 0 {
		g.SetModelName(args[0])
	} else {
		log.Fatalf(">>> ERROR >>>: please pass in the struct name and its fields\n")
	}
	return g
}

func trimLeftChar(s string) string {
	for i := range s {
		if i > 0 {
			return s[i:]
		}
	}
	return s[:0]
}

func setPackageName(g *GenerateModel) string {
	packageName := "models"
	//for _, arg := range g.Args {
	//	if strings.HasPrefix(arg, "package_") {
	//		packageLine := strings.Split(arg, "_")
	//		packageName = packageLine[1]
	//		g.SetPackageName(packageName)
	//		return packageName
	//	}
	//}
	return packageName
}

func (g *GenerateModel) SetPackageName(packageName string) *GenerateModel {
	g.PackageName = packageName
	return g
}

func (g *GenerateModel) GetPackageName() string {
	return "models"
}

func (g *GenerateModel) SetStruct() *GenerateModel {
	var sb strings.Builder
	// add package name
	setPackageName(g)
	sb.WriteString(fmt.Sprintf("package models"))
	sb.WriteString(fmt.Sprintf("\n"))
	// struct opening
	sb.WriteString(fmt.Sprintf("\ntype %v struct {\n", g.GetModelName()))

	for _, arg := range g.Args {
		if strings.HasPrefix(arg, "p_") || strings.HasPrefix(arg, "P_") {
			removeP := trimLeftChar(arg)
			remove_ := trimLeftChar(removeP)
			property := strings.Split(remove_, "-")
			fieldName := property[0]
			fieldType := property[1]
			jsonMap := fmt.Sprintf("`json:\"%v\"`", strcase.ToLowerCamel(property[0]))
			//`json:"rateId"
			sb.WriteString(fmt.Sprintf("\t%v\t%v\t%v\n", fieldName, fieldType, jsonMap))
		}
	}

	// closing brace
	sb.WriteString(fmt.Sprintf("}"))
	g.Struct = sb.String()
	return g
}

func (g *GenerateModel) GetStruct() string {
	return g.Struct
}
