

export class UsersHttpHandler extends GeneralHttpHandler<IUser> {
	public static getInstance(progressMessage = "please wait...") {
		return new UsersHttpHandler(allUsersStore, mapOfUsersToIDs, userToEdit, progressMessage);
	}
}