package routes

import (
	"mysql/constant/permission"
	"mysql/constant/route"
	"mysql/controller"
	"mysql/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	authcontroller := controller.NewAuthController()
	rolecontroller := controller.NewRoleHasPermissionController()
	authorcontroller := controller.NewAuthorController()
	facultycontroller := controller.NewFacultyController()
	departmentcontroller := controller.NewDepartmentController()
	programcontroller := controller.NewProgramController()
	categorycontroller := controller.NewCategoryController()
	cabinetcontroller := controller.NewCabinetController()
	filingcabinetcontroller := controller.NewFilingCabinetController()
	public := r.Group("/api/v1")
	public.Use(middleware.APIKeyAuth())
	{
		public.POST(route.Login, authcontroller.Login)
		public.POST(route.Refresh, authcontroller.Refresh)
	}
	auth := r.Group("/api/v1")
	auth.Use(middleware.APIKeyAuth())
	auth.Use(middleware.AuthMiddleware())
	{
		// User
		auth.POST(route.Register, middleware.PermissionMiddleware(permission.UserCreate), authcontroller.Register)
		auth.GET(route.UserView, middleware.PermissionMiddleware(permission.UserView), authcontroller.GetUser)
		auth.PUT(route.UserUpdate, middleware.PermissionMiddleware(permission.UserUpdate), authcontroller.Update)

		// Role
		auth.GET(route.RoleView, middleware.PermissionMiddleware(permission.RoleView), rolecontroller.GetRole)
		auth.PUT(route.RoleUpdate, middleware.PermissionMiddleware(permission.RoleUpdate), rolecontroller.UpdateRole)
		auth.GET(route.RolePermissionView, middleware.PermissionMiddleware(permission.RolePermissionView), rolecontroller.GetRolePermission)
		auth.POST(route.RolePermissionCreate, middleware.PermissionMiddleware(permission.RolePermissionCreate), rolecontroller.CreateRoleHasPermission)
		auth.DELETE(route.RolePermissionDelete, middleware.PermissionMiddleware(permission.RolePermissionDelete), rolecontroller.DeleteRoleHasPermission)

		// Author
		auth.POST(route.AuthorCreate, middleware.PermissionMiddleware(permission.AuthorCreate), authorcontroller.CreateAuthor)
		auth.GET(route.AuthorView, middleware.PermissionMiddleware(permission.AuthorView), authorcontroller.GetAuthor)
		auth.PUT(route.AuthorUpdate, middleware.PermissionMiddleware(permission.AuthorUpdate), authorcontroller.UpdateAuthor)
		auth.PUT(route.AuthorToggleStatus, middleware.PermissionMiddleware(permission.AuthorUpdate), authorcontroller.ToggleStatusAuthor)

		// Faculty
		auth.GET(route.FacultyView, middleware.PermissionMiddleware(permission.FacultyView), facultycontroller.GetFaculty)
		auth.GET(route.DepartmentView, middleware.PermissionMiddleware(permission.DepartmentView), departmentcontroller.GetDepartment)
		auth.GET(route.ProgramView, middleware.PermissionMiddleware(permission.ProgramView), programcontroller.GetProgram)

		// Category
		auth.GET(route.CategoryView, middleware.PermissionMiddleware(permission.CategoryModify), categorycontroller.GetCategory)
		auth.POST(route.CategoryCreate, middleware.PermissionMiddleware(permission.CategoryModify), categorycontroller.CreateCategory)
		auth.PUT(route.CategoryUpdate, middleware.PermissionMiddleware(permission.CategoryModify), categorycontroller.UpdateCategory)
		auth.PUT(route.CategoryToggleStatus, middleware.PermissionMiddleware(permission.CategoryModify), categorycontroller.ToggleStatusCategory)

		// Cabinet
		auth.GET(route.CabinetView, middleware.PermissionMiddleware(permission.CabinetModify), cabinetcontroller.GetCabinet)

		// FilingCabinet
		auth.GET(route.FilingCabinetView, middleware.PermissionMiddleware(permission.FilingCabinetModify), filingcabinetcontroller.GetFilingCabinet)
	}
}
