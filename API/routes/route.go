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
	}
}
