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
	}}

```

# for ui (svelte/svelte-kit)

to add a model;

```cli
gen-crud uimodel User p_userName-string p_age-number p_business-IBusiness p_branches-[]IBranch dir_models
```
this generates

```typescript
interface IUser {
    userName:string
    age:number
    business:IBusiness
    branches:[]IBranch
}

export const instantiateUser = (m?: Partial<IUser>): IUser => {
    const defaults: IUser = {
        userName:"",
        age:0,
        business:instantiateBusiness(),
        branches:new Array<IBranch>(),
    }
    return {
        ...defaults,
        ...m
    }
}
export default IUser

```

to add a store;

```cli
gen-crud store User dir_stores
```

# NOTE: dir_store means the directory is store... e.g dir_stores/users

this generates

```typescript
export const allUserDetailsStore = writable(new List<IUserDetail>());

export const userDetailToEdit = writable(instantiateUserDetail());

export const mapOfUserDetailsToIDs = writable(new KeyValueMap<number, IUserDetail>());

export let listOfUserDetails = new List<IUserDetail>();

allUserDetailsStore.subscribe((list) => {
        listOfUserDetails = list;
    }
);

export let mapOfUserDetails = new KeyValueMap<number, IUserDetail>();

mapOfUserDetailsToIDs.subscribe((map) => {
    mapOfUserDetails = map;
});

export let aUserDetail = instantiateUserDetail();

userDetailToEdit.subscribe((cat) => {
    aUserDetail = cat;
});

export const getUserDetailById = (id: number): IUserDetail | undefined => {
    let val = instantiateUserDetail();
    mapOfUserDetailsToIDs.subscribe(c => {
        if (c.get(Number(id))) {
            val = c.get(Number(id))!
        }
    });
    return val.id > 0 ? val : undefined
}


```

# generate a create component

```cli
gen-crud uicreate component/test Test f_input::t_text::n_username f_input::t_checkbox::n_isAdmin
```

this produces Test.svelte file in the directory path component/test

```typescript
<script lang="ts">

	async function handleSubmit() {
		// todo...please add your here to submit...
	}

</script>

<NoRights entityName="" action="">
	<form on:submit|preventDefault={handleSubmit}>
	

	

	<input  type="text" name="username" id="username" />

	<input  type="checkbox" name="isAdmin" id="isAdmin" />

	<FourthColumnGridOfFour>
		<IosButtonSmallGreen btnLabel="save" btnType="submit" />
	</FourthColumnGridOfFour>

	</form>

</NoRights>
```

# to add a view

```cli
 gen-crud view dir_views model_User
```

produces a file ViewUsers.svelte in the directory views

```typescript
<script lang="ts">
	let all = new Array<IUser>();
	const columns: IColumnProps[] = [
		{
			name: '#',
			valuePath: 'id'
		},
		{
			name: 'Created On',
			valuePath: 'createdAt',
			expressedFormatter: true,
			format: (row: IUser) => {
				return `${new Date(row.createdAt).toLocaleDateString()} ${new Date(
					row.createdAt
				).toLocaleTimeString()}`;
			}
		},
];

	const deleteRow = (config: IRequestResponseVariable, id: number) => {
		UsersHttpHandler.getInstance('deleting...').deleteById(config, id);
	};	const loadOrReload = () => {
		UsersHttpHandler.getInstance('loading currencies...').getAll({
			url: endpoints.users.getAllByClientId(
				Number(ConstantsLocalStorage.getCustomerId())
			),
			setData: (d: any) => {
				console.log(' to tabulate', d);
			}
			setData: (d: any) => {
				console.log(' to tabulate', d);
			}
		});
	};
</script>

<NoRights entityName="currencies" action="view">
	<GenericTableWithDeleteAndReload
		{all}
		mapOfEntitiesToIDs={mapOfUsersToIDs}
		entityToEdit={userToEdit}
		{columns}
		deleteByIdEndpoint={endpoints.users.getAllByClientId}
		{loadOrReload}
		deleteFn={deleteRow}
	/>
</NoRights>

```