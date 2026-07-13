package route

const (

	// Authentication
	Login    = "login"
	Refresh  = "refresh"
	Logout   = "logout"
	Register = "register"

	// User
	UserView   = "user.view"
	UserCreate = "user.create"
	UserUpdate = "user.update"
	UserDelete = "user.delete"

	RoleView             = "role.view"
	RoleCreate           = "role.create"
	RoleUpdate           = "role.update/:id"
	RolePermissionView   = "role.permission.view/:id"
	RolePermissionCreate = "role.permission.create"
	RolePermissionDelete = "role.permission.delete"
)
