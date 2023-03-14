package httpclient

import (
	"fmt"
	"gen-crud/utils"
	"github.com/iancoleman/strcase"
	"github.com/samber/lo"
	"strings"
)

func AddHttpClient(args []string) (dir, modelName, content string) {

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

	sb.WriteString(fmt.Sprintf("\n\nexport class %vHttpHandler extends GeneralHttpHandler<I%v> {", utils.ToPlural(modelName), modelName))
	sb.WriteString(fmt.Sprintf("\n\tpublic static getInstance(progressMessage = \"please wait...\") {"))
	sb.WriteString(fmt.Sprintf("\n\t\treturn new %vHttpHandler(all%vStore, mapOf%vToIDs, %vToEdit, progressMessage);", utils.ToPlural(modelName), utils.ToPlural(modelName), utils.ToPlural(modelName), strcase.ToLowerCamel(modelName)))
	sb.WriteString(fmt.Sprintf("\n\t}"))
	sb.WriteString(fmt.Sprintf("\n}"))
	content = sb.String()
	return dir, modelName, content
}
