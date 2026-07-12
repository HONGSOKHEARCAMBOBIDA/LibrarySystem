package route

const (

	// Authentication
	Login   = "login"
	Refresh = "refresh"
	Logout  = "logout"

	RoleView             = "role.view"
	RoleCreate           = "role.create"
	RoleUpdate           = "role.update/:id"
	RolePermissionView   = "role.permission.view/:id"
	RolePermissionCreate = "role.permission.create"
	RolePermissionDelete = "role.permission.delete"
)
