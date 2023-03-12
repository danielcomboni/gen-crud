# gen-crud - a golang cli to generate models/entities, and it's CRUD related operation

# samples
to generate a model;

```cli
gen-crud model User p_Name-string p_Age-int p_Email-string
```

this creates a file  user.go and produces the struct below plus the package name and with 

```go
package models

type User struct {
	Name	string	`json:"name"`
	Age	int	`json:"age"`
	Email	string	`json:"email"`
}
```

to generate a controller;

```cli
gen-crud controller User
```

this produces the code below

```go
package controllers

func CreateUser(shouldAuthenticate bool) gin.HandlerFunc {
			return Create[models.User](CreateHandler[models.User]{
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
		}
func GetAllUsers(shouldAuthenticate bool) gin.HandlerFunc {
			return GetAll[models.User](GetAllHandler[models.User]{
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
			
func GetAllUsersByClientId(shouldAuthenticate bool) gin.HandlerFunc {
					return GetAllByClientId[models.User](GetAllHandler[models.User]{
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

			
func UpdateUserById(shouldAuthenticate bool) gin.HandlerFunc {
					return UpdateById[models.User](UpdateHandler[models.User]{
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


			
func DeleteUserById(shouldAuthenticate bool) gin.HandlerFunc {
						return DeleteById[models.User](DeleteHandler[models.User]{
							EntityName:               "",
							ShouldAuthenticate:       false,
							ShouldRecordAuditHistory: false,
							})
						}
			
```

and to generate routes;

```cli
gen-crud route User
```

this produces;

```go
package routes

func UserRoutes(baseResourceURL string, router *gin.Engine) {
		endpointGroupV1 := router.Group("")
		{		
			
routes_utils.Post(endpointGroupV1, baseResourceURL, "users/v1", controllers.CreateUser())
routes_utils.Get(endpointGroupV1, baseResourceURL, "users/v1/", controllers.GetUsers())
routes_utils.Get(endpointGroupV1, baseResourceURL, "users/v1/:customerId", controllers.GetUsersByClientId())
routes_utils.Get(endpointGroupV1, baseResourceURL, "users/v1/getById/:id", controllers.GetUserById())
routes_utils.Put(endpointGroupV1, baseResourceURL, "users/v1/:id", controllers.UpdateUserById())
routes_utils.Del(endpointGroupV1, baseResourceURL, "users/v1/:id", controllers.DeleteUserById())
		}
	}

```