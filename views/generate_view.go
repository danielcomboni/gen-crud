package views

import (
	"fmt"
	"gen-crud/utils"
	"github.com/iancoleman/strcase"
	"github.com/samber/lo"
	"strings"
)

func AddView(args []string) (dir, modelName, content string) {
	dir = ""
	modelName = ""

	lo.ForEach(args, func(item string, index int) {

		if strings.HasPrefix(item, "dir_") {
			dir = strings.Split(item, "_")[1]
		}

		if strings.HasPrefix(item, "model_") {
			modelName = strings.Split(item, "_")[1]
		}

	})

	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("<script lang=\"ts\">\n"))
	sb.WriteString(fmt.Sprintf("\tlet all = new Array<I%v>();\n", modelName))
	sb.WriteString(fmt.Sprintf("\tconst columns: IColumnProps[] = [\n\t\t{\n\t\t\tname: '#',\n\t\t\tvaluePath: 'id'\n\t\t},\n\t\t{\n\t\t\tname: 'Created On',\n\t\t\tvaluePath: 'createdAt',\n\t\t\texpressedFormatter: true,\n\t\t\tformat: (row: I%v) => {\n\t\t\t\treturn `${new Date(row.createdAt).toLocaleDateString()} ${new Date(\n\t\t\t\t\trow.createdAt\n\t\t\t\t).toLocaleTimeString()}`;\n\t\t\t}\n\t\t},\n];\n\n", modelName))
	sb.WriteString(fmt.Sprintf("\tconst deleteRow = (config: IRequestResponseVariable, id: number) => {\n\t\t%vHttpHandler.getInstance('deleting...').deleteById(config, id);\n\t};", utils.ToPlural(modelName)))
	sb.WriteString(fmt.Sprintf("\tconst loadOrReload = () => {\n"))
	sb.WriteString(fmt.Sprintf("\t\t%vHttpHandler.getInstance('loading currencies...').getAll({\n", utils.ToPlural(modelName)))
	sb.WriteString(fmt.Sprintf("\t\t\turl: endpoints.%v.getAllByClientId(\n", utils.ToPlural(strcase.ToLowerCamel(modelName))))
	sb.WriteString(fmt.Sprintf("\t\t\t\tNumber(ConstantsLocalStorage.getCustomerId())\n"))
	sb.WriteString(fmt.Sprintf("\t\t\t),\n"))
	sb.WriteString(fmt.Sprintf("\t\t\tsetData: (d: any) => {\n\t\t\t\tconsole.log(' to tabulate', d);\n\t\t\t}\n"))
	sb.WriteString(fmt.Sprintf("\t\t\tsetData: (d: any) => {\n\t\t\t\tconsole.log(' to tabulate', d);\n\t\t\t}\n"))
	sb.WriteString(fmt.Sprintf("\t\t});\n\t};\n"))

	sb.WriteString(fmt.Sprintf("</script>\n"))

	sb.WriteString(fmt.Sprintf("<NoRights entityName=\"currencies\" action=\"view\">\n"))
	sb.WriteString(fmt.Sprintf("\t<GenericTableWithDeleteAndReload\n"))
	sb.WriteString(fmt.Sprintf("\t\t{all}\n"))
	sb.WriteString(fmt.Sprintf("\t\tmapOfEntitiesToIDs={mapOf%vToIDs}\n", utils.ToPlural(modelName)))
	sb.WriteString(fmt.Sprintf("\t\tentityToEdit={%vToEdit}\n", strcase.ToLowerCamel(modelName)))
	sb.WriteString(fmt.Sprintf("\t\t{columns}\n"))
	sb.WriteString(fmt.Sprintf("\t\tdeleteByIdEndpoint={endpoints.%v.getAllByClientId}\n", utils.ToPlural(strcase.ToLowerCamel(modelName))))
	sb.WriteString(fmt.Sprintf("\t\t{loadOrReload}\n"))
	sb.WriteString(fmt.Sprintf("\t\tdeleteFn={deleteRow}\n"))
	sb.WriteString(fmt.Sprintf("\t/>\n</NoRights>\n"))
	content = sb.String()
	return dir, modelName, content
}
