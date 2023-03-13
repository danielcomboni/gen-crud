package stores

import (
	"fmt"
	"gen-crud/utils"
	"github.com/iancoleman/strcase"
	"strings"
)

func AddStore(args []string) string {
	var sb strings.Builder

	modelName := args[0]
	pluralModelName := utils.ToPlural(modelName)
	sb.WriteString(fmt.Sprintf("export const all%vStore = writable(new List<I%v>());\n\n", pluralModelName, modelName))
	sb.WriteString(fmt.Sprintf("export const %vToEdit = writable(instantiate%v());\n\n", strcase.ToLowerCamel(modelName), modelName))
	sb.WriteString(fmt.Sprintf("export const mapOf%vToIDs = writable(new KeyValueMap<number, I%v>());\n\n", pluralModelName, modelName))
	sb.WriteString(fmt.Sprintf("export let listOf%v = new List<I%v>();\n\n", pluralModelName, modelName))
	sb.WriteString(fmt.Sprintf("all%vStore.subscribe((list) => {\n\tlistOf%v = list;\n}\n);\n\n", pluralModelName, pluralModelName))

	sb.WriteString(fmt.Sprintf("export let mapOf%v = new KeyValueMap<number, I%v>();\n\n", pluralModelName, modelName))
	sb.WriteString(fmt.Sprintf("mapOf%vToIDs.subscribe((map) => {\n\tmapOf%v = map;\n});\n\n", pluralModelName, pluralModelName))

	sb.WriteString(fmt.Sprintf("export let a%v = instantiate%v();\n\n", modelName, modelName))
	sb.WriteString(fmt.Sprintf("%vToEdit.subscribe((cat) => {\n\ta%v = cat;\n});\n\n", strcase.ToLowerCamel(modelName), modelName))
	sb.WriteString(fmt.Sprintf("export const get%vById = (id: number): I%v | undefined => {\n\tlet val = instantiate%v();\n\tmapOf%vToIDs.subscribe(c => {\n\t\tif (c.get(Number(id))) {\n\t\t\tval = c.get(Number(id))!\n\t\t}\n});\n\treturn val.id > 0 ? val : undefined\n}\n\n", strcase.ToCamel(modelName), modelName, modelName, pluralModelName))

	return sb.String()

}
