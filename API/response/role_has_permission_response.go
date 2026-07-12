package response

type PermissionWithAssignedRole struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	ModuleName string `json:"module_name"`
	Assigned   bool   `json:"assigned"`
}
