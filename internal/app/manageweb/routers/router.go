package routers

import (
	//"html/template"
	"net/http"

	"goapp/internal/app/manageweb/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(app *gin.Engine) {
	//首页
	app.GET("/", func(c *gin.Context) { c.HTML(http.StatusOK, "index.html", nil) })
	apiPrefix:="/api"
	g := app.Group(apiPrefix)
	// 登录验证 jwt token 验证 及信息提取
	var notCheckLoginUrlArr []string
	notCheckLoginUrlArr = append(notCheckLoginUrlArr, apiPrefix+"/user/login")
	notCheckLoginUrlArr = append(notCheckLoginUrlArr, apiPrefix+"/user/logout")
	g.Use(middleware.UserAuthMiddleware(
		middleware.AllowPathPrefixSkipper(notCheckLoginUrlArr...),
	))
	// 权限验证
	var notCheckPermissionUrlArr []string
	notCheckPermissionUrlArr = append(notCheckPermissionUrlArr, notCheckLoginUrlArr...)
	notCheckPermissionUrlArr = append(notCheckPermissionUrlArr, apiPrefix+"/menu/menubuttonlist")
	notCheckPermissionUrlArr = append(notCheckPermissionUrlArr, apiPrefix+"/menu/allmenu")
	notCheckPermissionUrlArr = append(notCheckPermissionUrlArr, apiPrefix+"/admins/adminsroleidlist")
	notCheckPermissionUrlArr = append(notCheckPermissionUrlArr, apiPrefix+"/user/info")
	notCheckPermissionUrlArr = append(notCheckPermissionUrlArr, apiPrefix+"/user/editpwd")
	notCheckPermissionUrlArr = append(notCheckPermissionUrlArr, apiPrefix+"/role/rolemenuidlist")
	notCheckPermissionUrlArr = append(notCheckPermissionUrlArr, apiPrefix+"/role/allrole")
	g.Use(middleware.CasbinMiddleware(
		middleware.AllowPathPrefixSkipper(notCheckPermissionUrlArr...),
	))
	//sys
	RegisterRouterSys(g)
}
