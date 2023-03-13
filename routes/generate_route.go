package routes

import (
	"fmt"
	"gen-crud/utils"
	"github.com/iancoleman/strcase"
	"strings"
)

type GenerateRoute struct {
	Args        []string
	RouteName   string
	PackageName string
	ModelName   string
}

func (g *GenerateRoute) SetArgs(args []string) *GenerateRoute {
	var arguments []string
	arguments = args
	g.Args = arguments
	if len(g.Args) > 0 {
		g.ModelName = g.Args[0]
	}
	return g
}

func (g *GenerateRoute) GenerateRouteString() string {
	var sb strings.Builder

	endpoint := strcase.ToLowerCamel(utils.ToPlural(g.ModelName))
	modelNamePlural := utils.ToPlural(g.ModelName)
	create := fmt.Sprintf("\t\troutes_utils.Post(endpointGroupV1, baseResourceURL, \"%v/v1\", controllers.Create%v())", endpoint, g.ModelName)
	//create := fmt.Sprintf(`routes_utils.Post(endpointGroupV1, baseResourceURL, "%v/v1", controllers.Create%v())`, endpoint, g.ModelName)
	getAllByClientId := fmt.Sprintf("\t\troutes_utils.Get(endpointGroupV1, baseResourceURL, \"%v/v1/:customerId\", controllers.Get%vByClientId())", endpoint, modelNamePlural)
	getAll := fmt.Sprintf("\t\troutes_utils.Get(endpointGroupV1, baseResourceURL, \"%v/v1/\", controllers.Get%v())", endpoint, modelNamePlural)
	getById := fmt.Sprintf("\t\troutes_utils.Get(endpointGroupV1, baseResourceURL, \"%v/v1/getById/:id\", controllers.Get%vById())", endpoint, g.ModelName)
	updateById := fmt.Sprintf("\t\troutes_utils.Put(endpointGroupV1, baseResourceURL, \"%v/v1/:id\", controllers.Update%vById())", endpoint, g.ModelName)
	deleteById := fmt.Sprintf("\t\troutes_utils.Del(endpointGroupV1, baseResourceURL, \"%v/v1/:id\", controllers.Delete%vById())", endpoint, g.ModelName)

	sb.WriteString("\n")
	sb.WriteString(create)
	sb.WriteString("\n")
	sb.WriteString(getAll)
	sb.WriteString("\n")
	sb.WriteString(getAllByClientId)
	sb.WriteString("\n")
	sb.WriteString(getById)
	sb.WriteString("\n")
	sb.WriteString(updateById)
	sb.WriteString("\n")
	sb.WriteString(deleteById)

	var sb2 strings.Builder
	sb2.WriteString("package routes")
	sb2.WriteString("\n")
	sb2.WriteString("\n")

	return fmt.Sprintf("%vfunc %vRoutes(baseResourceURL string, router *gin.Engine) {\n\n\tendpointGroupV1 := router.Group(\"\")\n\t{%v\n\t}}", sb2.String(), g.ModelName, sb.String())
}
