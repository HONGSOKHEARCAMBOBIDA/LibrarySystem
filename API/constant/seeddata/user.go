package seeddata

import "mysql/model"

var DeveloperUser = model.User{
	NameKh:   "Developer Team",
	NameEN:   "Developer Team",
	Email:    "developerteam168@gmail.com",
	Password: "$2a$04$0DZrjT3ivE4ipHnX9TKHfewx/eJHBTXLEVhAE.MYYWWZAvlCZrz2e",
	Level:    7,
	RoleId:   1,
	Gender:   1,
	Dob:      "2024-01-10",
	Isactive: true,
}
