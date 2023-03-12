package controllers

import (
	"fmt"
	"gen-crud/utils"
	"github.com/iancoleman/strcase"
	"log"
	"strings"
)

type GenerateController struct {
	Args           []string
	ControllerName string
	PackageName    string
	ModelName      string
	CreateTemplate string
	GetAllTemplate string
	Templates      map[string]string
}

func (g *GenerateController) GetModelNamePlural() string {
	return utils.ToPlural(g.ModelName)
}

func (g *GenerateController) SetModelName(str string) *GenerateController {
	g.ModelName = strcase.ToCamel(str)
	return g
}

func (g *GenerateController) AddCreateTemplate() *GenerateController {
	g.Templates["Create"] = fmt.Sprintf(`func Create%v(shouldAuthenticate bool) gin.HandlerFunc {
			return Create[models.%v](CreateHandler[models.%v]{
				EntityName:                    "",
				ShouldAuthenticate:            false,
				ShouldRecordAuditHistory:      false,
				UseCustomValidator:            false,
				CustomValidatedModel:          nil,
				ShouldCheckDuplicates:         false,
				DuplicateParams:               nil,
				ShouldCheckDuplicatesManually: false,
				DuplicateParamsManually:       nil,
				ShouldValidateModelValues:     false,
				ValidateModelValues:           nil,
				ShouldModifyModelValue:        false,
				ModifyModelValue:              nil,
				CheckIdCreated:                false,
				AfterSuccessfulCreate:         nil,
				PriorActions:                  nil,
			})
		}`, g.ModelName, g.ModelName, g.ModelName)

	return g
}

func (g *GenerateController) GetCreateTemplate() string {
	return g.Templates["Create"]
}

func (g *GenerateController) AddGetAllTemplate() *GenerateController {
	g.Templates["GetAll"] = fmt.Sprintf(`func GetAll%v(shouldAuthenticate bool) gin.HandlerFunc {
			return GetAll[models.%v](GetAllHandler[models.%v]{
					EntityName:            "",
					ShouldAuthenticate:    false,
					Preloads:              nil,
					URLParamValidator:     nil,
					URLParams:             nil,
					QueryParams:           nil,
					PriorManipulations:    nil,
					RunPriorManipulations: false,
					})
				}
			`, g.GetModelNamePlural(), g.ModelName, g.ModelName)

	return g
}

func (g *GenerateController) GetGetAllTemplate() string {
	return g.Templates["GetAll"]
}

func (g *GenerateController) AddGetAllByClientIdTemplate() *GenerateController {
	g.Templates["GetAllByClientId"] = fmt.Sprintf(`func GetAll%vByClientId(shouldAuthenticate bool) gin.HandlerFunc {
					return GetAllByClientId[models.%v](GetAllHandler[models.%v]{
						EntityName:            "",
						ShouldAuthenticate:    false,
						Preloads:              nil,
						URLParamValidator:     nil,
						URLParams:             nil,
						QueryParams:           nil,
						PriorManipulations:    nil,
						RunPriorManipulations: false,
					})
				}

			`, g.GetModelNamePlural(), g.ModelName, g.ModelName)

	return g
}

func (g *GenerateController) GetGetAllByClientIdTemplate() string {
	return g.Templates["GetAllByClientId"]
}

func (g *GenerateController) AddUpdateByIdTemplate() *GenerateController {
	g.Templates["UpdateById"] = fmt.Sprintf(`func Update%vById(shouldAuthenticate bool) gin.HandlerFunc {
					return UpdateById[models.%v](UpdateHandler[models.%v]{
						EntityName:                "",
						ShouldAuthenticate:        false,
						ShouldRecordAuditHistory:  false,
						UseCustomValidator:        false,
						CustomValidatedModel:      nil,
						ShouldCheckDuplicates:     false,
						DuplicateParams:           nil,
						ShouldValidateModelValues: false,
						ValidateModelValues:       nil,
						ShouldModifyModelValue:    false,
						ModifyModelValue:          nil,
					})
				}


			`, g.ModelName, g.ModelName, g.ModelName)

	return g
}

func (g *GenerateController) GetUpdateByIdTemplate() string {
	return g.Templates["UpdateById"]
}

func (g *GenerateController) AddDeleteByIdTemplate() *GenerateController {
	g.Templates["DeleteById"] = fmt.Sprintf(`func Delete%vById(shouldAuthenticate bool) gin.HandlerFunc {
						return DeleteById[models.%v](DeleteHandler[models.%v]{
							EntityName:               "",
							ShouldAuthenticate:       false,
							ShouldRecordAuditHistory: false,
							})
						}
			`, g.ModelName, g.ModelName, g.ModelName)

	return g
}

func (g *GenerateController) GetDeleteByIdTemplate() string {
	return g.Templates["DeleteById"]
}

func (g *GenerateController) GetControllerName() string {
	return g.ControllerName
}

func (g *GenerateController) SetControllerName(str string) *GenerateController {
	g.ControllerName = str
	return g
}

func (g *GenerateController) SetControllerArgs(args []string) *GenerateController {
	g.Args = args
	if len(args) > 0 {
		g.SetControllerName(fmt.Sprintf("%v_controller", strcase.ToSnake(args[0])))
		g.SetModelName(args[0])
	} else {
		log.Fatalf(">>> ERROR >>>: please pass in the struct/model name to be handle by the controller\n")
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

func (g *GenerateController) SetPackageName(packageName string) *GenerateController {
	g.PackageName = packageName
	return g
}

func (g *GenerateController) GetPackageName() string {
	return g.PackageName
}

func (g *GenerateController) AddCRUD() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("package controllers"))

	g.Templates = map[string]string{}

	// add create
	sb.WriteString("\n")
	sb.WriteString("\n")
	g.AddCreateTemplate()
	sb.WriteString(g.GetCreateTemplate())

	// add get all
	sb.WriteString("\n")
	g.AddGetAllTemplate()
	sb.WriteString(g.GetGetAllTemplate())

	// add get all by client ID
	sb.WriteString("\n")
	g.AddGetAllByClientIdTemplate()
	sb.WriteString(g.GetGetAllByClientIdTemplate())

	// add update by ID
	sb.WriteString("\n")
	g.AddUpdateByIdTemplate()
	sb.WriteString(g.GetUpdateByIdTemplate())

	// add delete by ID
	sb.WriteString("\n")
	g.AddDeleteByIdTemplate()
	sb.WriteString(g.GetDeleteByIdTemplate())

	return sb.String()
}
