package models

import (
	"fmt"
	"github.com/iancoleman/strcase"
	//"github.com/samber/lo"
	"log"
	"strings"
)

// sample:  cli-1 model PaySheet p_Name-string p_Department-string p_PaySlip-[]PaySlip  package_models

type GenerateModelUI struct {
	Args        []string
	Model       string
	PackageName string
	ModelName   string
}

func (g *GenerateModelUI) GetModelName() string {
	return g.ModelName
}

func (g *GenerateModelUI) SetModelName(str string) *GenerateModelUI {
	g.ModelName = strcase.ToCamel(str)
	return g
}

func (g *GenerateModelUI) SetModelArgs(args []string) *GenerateModelUI {
	g.Args = args
	if len(args) > 0 {
		g.SetModelName(args[0])
		g.SetModel()
	} else {
		log.Fatalf(">>> ERROR >>>: please pass in the struct name and its fields\n")
	}
	return g
}

func (g *GenerateModelUI) AddInstantiation() string {
	var sbInstantiate strings.Builder
	sbInstantiate.WriteString("\n")
	sbInstantiate.WriteString(fmt.Sprintf("export const instantiate%v = (m?: Partial<I%v>): I%v => {\n", g.GetModelName(), g.GetModelName(), g.GetModelName()))

	var defaults strings.Builder

	defaults.WriteString(fmt.Sprintf("const defaults: I%v = {\n", g.GetModelName()))

	addDefaults := func(fieldName, fieldType string) {
		switch strings.ToLower(fieldType) {
		case "number":
			defaults.WriteString(fmt.Sprintf("\t%v:%v\n", fieldName, 0))
		case "string":
			defaults.WriteString(fmt.Sprintf("\t%v:%v\n", fieldName, `""`))
		default:
			if strings.HasPrefix(strings.ToLower(fieldType), "i") {
				instantiate := fmt.Sprintf(`instantiate%v()`, trimLeftChar(fieldType))
				defaults.WriteString(fmt.Sprintf("\t%v:%v\n", fieldName, instantiate))
			}
			// has array
			if strings.HasPrefix(strings.ToLower(fieldType), "[]") {
				instantiate := fmt.Sprintf(`new Array<%v>()`, strings.ReplaceAll(fieldType, "[]", ""))
				defaults.WriteString(fmt.Sprintf("\t%v:%v\n", fieldName, instantiate))
			}
		}
		//defaults.WriteString("\n")
	}

	for _, arg := range g.Args {
		if strings.HasPrefix(arg, "p_") || strings.HasPrefix(arg, "P_") {
			removeP := trimLeftChar(arg)
			remove_ := trimLeftChar(removeP)
			property := strings.Split(remove_, "-")
			fieldName := property[0]
			fieldType := property[1]
			//sbInstantiate.WriteString(fmt.Sprintf("\t%v:%v\n", fieldName, fieldType))
			addDefaults(fieldName, fieldType)
		}
	}

	defaults.WriteString("}")
	defaults.WriteString("\n")
	defaults.WriteString("return {")
	defaults.WriteString(fmt.Sprintf("\n...defaults,\n"))
	defaults.WriteString(fmt.Sprintf("...m\n"))
	defaults.WriteString(" }\n") // close return bracket

	sbInstantiate.WriteString(defaults.String())
	sbInstantiate.WriteString(fmt.Sprintf("}\n"))                                  // closing brace (bracket)
	sbInstantiate.WriteString(fmt.Sprintf("export default I%v", g.GetModelName())) // closing brace (bracket)

	return sbInstantiate.String()
}

func (g *GenerateModelUI) SetModel() *GenerateModelUI {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("\n"))
	// model opening
	sb.WriteString(fmt.Sprintf("\ninterface I%v {\n", g.GetModelName()))

	for _, arg := range g.Args {
		if strings.HasPrefix(arg, "p_") || strings.HasPrefix(arg, "P_") {
			removeP := trimLeftChar(arg)
			remove_ := trimLeftChar(removeP)
			property := strings.Split(remove_, "-")
			fieldName := property[0]
			fieldType := property[1]
			//`json:"rateId"

			sb.WriteString(fmt.Sprintf("\t%v:%v\n", fieldName, fieldType))
		}
	}

	// closing brace
	sb.WriteString(fmt.Sprintf("}"))
	g.Model = sb.String()
	return g
}

func (g *GenerateModelUI) GetModel() string {
	return g.Model
}
