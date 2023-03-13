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
	sb.WriteString(fmt.Sprintf("export const all%vStore = writable(new List<%v>());\n", pluralModelName, modelName))
	sb.WriteString(fmt.Sprintf("export const %vToEdit = writable(instantiate%v());\n", strcase.ToLowerCamel(modelName), modelName))
	sb.WriteString(fmt.Sprintf("export const mapOf%vToIDs = writable(new KeyValueMap<number, I%v>());\n", pluralModelName, modelName))
	sb.WriteString(fmt.Sprintf("export let listOf%v = new List<I%v>();\n\n", pluralModelName, modelName))
	sb.WriteString(fmt.Sprintf(`all%vStore.subscribe((list) => {
											listOf%v = list;
										});
`, pluralModelName, pluralModelName))

	sb.WriteString(fmt.Sprintf("export let mapOf%v = new KeyValueMap<number, I%v>();\n", pluralModelName, modelName))
	sb.WriteString(fmt.Sprintf(`mapOf%vToIDs.subscribe((map) => {
										    mapOf%v = map;
										});

						`, pluralModelName, pluralModelName))

	sb.WriteString(fmt.Sprintf("export let a%v = instantiate%v();\n", modelName, modelName))
	sb.WriteString(fmt.Sprintf("export let a%v = instantiate%v();\n", modelName, modelName))
	sb.WriteString(fmt.Sprintf(`%vToEdit.subscribe((cat) => {
											a%v = cat;
										});

									`, strcase.ToLowerCamel(modelName), modelName))
	sb.WriteString(fmt.Sprintf(`export const get%vById = (id: number): I%v | undefined => {
										let val = instantiate%v()
										mapOf%vToIDs.subscribe(c => {
											if (c.get(Number(id))) {
												val = c.get(Number(id))!
											}
										})
										return val.id > 0 ? val : undefined
									}
`, strcase.ToLowerCamel(modelName), modelName, modelName, pluralModelName))

	return sb.String()

}
