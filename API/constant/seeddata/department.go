package seeddata

import "mysql/model"

var Departments = []model.Department{
	{
		Name:      "នាយកដ្ឋាន វិទ្យាសាស្រ្តអប់រំ",
		Code:      "DEP-EDU",
		Isactive:  true,
		FacultyID: 8,
	},
	{
		Name:      "នាយកដ្ឋាន មនុស្សសាស្រ្ត សិល្បៈ និងភាសាវិទ្យា",
		Code:      "DEP-HAL",
		Isactive:  true,
		FacultyID: 9,
	},
	{
		Name:      "នាយកដ្ឋាន គ្រប់គ្រងពាណិជ្ជកម្ម និងទេសចរណ៍",
		Code:      "DEP-BTM",
		Isactive:  true,
		FacultyID: 10,
	},
	{
		Name:      "នាយកដ្ឋាន នីតិសាស្រ្ត និងវិទ្យាសាស្រ្តសេដ្ឋកិច្ច",
		Code:      "DEP-LEC",
		Isactive:  true,
		FacultyID: 11,
	},
	{
		Name:      "នាយកដ្ឋាន វិទ្យាសាស្រ្តបច្ចេកវិទ្យា និងព័ត៌មានវិទ្យា",
		Code:      "DEP-STI",
		Isactive:  true,
		FacultyID: 12,
	},
	{
		Name:      "នាយកដ្ឋាន វិទ្យាសាស្រ្តនយោបាយ និងទំនាក់ទំនងអន្តរជាតិ",
		Code:      "DEP-PIR",
		Isactive:  true,
		FacultyID: 13,
	},
	{
		Name:      "នាយកដ្ឋាន វិទ្យាសាស្រ្តកសិកម្ម និងអភិវឌ្ឍន៍ជនបទ",
		Code:      "DEP-ARD",
		Isactive:  true,
		FacultyID: 14,
	},
}
