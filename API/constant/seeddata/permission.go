package seeddata

import "mysql/model"

var Permissions = []model.Permission{
	{
		Name:        "user.view",
		ModuleName:  "user",
		DisplayName: "View User",
		OrderNo:     1,
	},

	{
		Name:        "user.create",
		ModuleName:  "user",
		DisplayName: "Create User",
		OrderNo:     2,
	},

	{
		Name:        "user.update",
		ModuleName:  "user",
		DisplayName: "Update User",
		OrderNo:     3,
	},

	{
		Name:        "user.delete",
		ModuleName:  "user",
		DisplayName: "Delete User",
		OrderNo:     4,
	},

	{
		Name:        "role.view",
		ModuleName:  "role",
		DisplayName: "View Role",
		OrderNo:     1,
	},

	{
		Name:        "role.create",
		ModuleName:  "role",
		DisplayName: "Create Role",
		OrderNo:     2,
	},

	{
		Name:        "role.update",
		ModuleName:  "role",
		DisplayName: "Update Role",
		OrderNo:     3,
	},

	{
		Name:        "role.permission.view",
		ModuleName:  "role",
		DisplayName: "View Role Has Permission",
		OrderNo:     4,
	},

	{
		Name:        "role.permission.create",
		ModuleName:  "role",
		DisplayName: "Create Role Has Permission",
		OrderNo:     5,
	},

	{
		Name:        "role.permission.delete",
		ModuleName:  "role",
		DisplayName: "Delete Role Has Permission",
		OrderNo:     6,
	},

	{
		Name:        "author.create",
		ModuleName:  "author",
		DisplayName: "Create Author",
		OrderNo:     1,
	},

	{
		Name:        "author.view",
		ModuleName:  "author",
		DisplayName: "View Author",
		OrderNo:     2,
	},

	{
		Name:        "author.update",
		ModuleName:  "author",
		DisplayName: "Update Author",
		OrderNo:     3,
	},

	{
		Name:        "author.delete",
		ModuleName:  "author",
		DisplayName: "Delete Author",
		OrderNo:     4,
	},

	{
		Name:        "faculty.view",
		ModuleName:  "fdp",
		DisplayName: "View Faculty",
		OrderNo:     1,
	},

	{
		Name:        "department.view",
		ModuleName:  "fdp",
		DisplayName: "View Department",
		OrderNo:     2,
	},

	{
		Name:        "program.view",
		ModuleName:  "fdp",
		DisplayName: "View Program",
		OrderNo:     3,
	},
}
