package routers

import (
	"github.com/it234/goapp/internal/app/manageweb/controllers/sys"

	"github.com/gin-gonic/gin"
)

func RegisterRouterSys(app *gin.RouterGroup) {
	menu := sys.Menu{}
	app.GET("/menu/list", menu.List)
	app.GET("/menu/detail", menu.Detail)
	app.GET("/menu/allmenu", menu.AllMenu)
	app.GET("/menu/menubuttonlist", menu.MenuButtonList)
	app.POST("/menu/delete", menu.Delete)
	app.POST("/menu/update", menu.Update)
	app.POST("/menu/create", menu.Create)
	user := sys.User{}
	app.GET("/user/info", user.Info)
	app.POST("/user/login", user.Login)
	app.POST("/user/logout", user.Logout)
	app.POST("/user/editpwd", user.EditPwd)
	admins := sys.Admins{}
	app.GET("/admins/list", admins.List)
	app.GET("/admins/detail", admins.Detail)
	app.GET("/admins/adminsroleidlist", admins.AdminsRoleIDList)
	app.POST("/admins/delete", admins.Delete)
	app.POST("/admins/update", admins.Update)
	app.POST("/admins/create", admins.Create)
	app.POST("/admins/setrole", admins.SetRole)
	role := sys.Role{}
	app.GET("/role/list", role.List)
	app.GET("/role/detail", role.Detail)
	app.GET("/role/rolemenuidlist", role.RoleMenuIDList)
	app.GET("/role/allrole", role.AllRole)
	app.POST("/role/delete", role.Delete)
	app.POST("/role/update", role.Update)
	app.POST("/role/create", role.Create)
	app.POST("/role/setrole", role.SetRole)
}
