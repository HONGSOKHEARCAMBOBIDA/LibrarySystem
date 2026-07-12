package seeddata

import "mysql/model"

var Roles = []model.Role{
	{
		Name:       "Developer Team",
		ModuleName: "Developer Team",
		IsActive:   true,
		Level:      7,
	},

	{
		Name:       "IT Support",
		ModuleName: "IT Support",
		IsActive:   true,
		Level:      5,
	},
}
