package manageweb

import (
	"net/http"
	"time"

	webconfig "github.com/it234/goapp/internal/app/manageweb/config"
	"github.com/it234/goapp/internal/pkg/config"
	"github.com/it234/goapp/internal/app/manageweb/middleware"
	"github.com/it234/goapp/internal/app/manageweb/routers"
	"github.com/it234/goapp/internal/pkg/models"
	"github.com/it234/goapp/pkg/logger"
	"github.com/it234/goapp/internal/app/manageweb/controllers/common"
	"github.com/it234/goapp/pkg/convert"
	
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// 运行
func Run(configPath string) {
	if configPath==""{
		configPath="./config.yaml"
	}
	// 加载配置
	config, err := webconfig.LoadConfig(configPath)
	if err != nil {
		panic(err)
	}
	logger.InitLog("debug","./log/logb.log")
	initDB(config)
	common.InitCsbinEnforcer()
	initWeb(config)
	logger.Debug(config.Web.Domain+"站点已启动...")
}

func initDB(config *config.Config){
	models.InitDB(config)
	models.Migration()
}

func initWeb(config *config.Config){
	gin.SetMode(gin.DebugMode) //调试模式
	app := gin.New()
	app.NoRoute(middleware.NoRouteHandler())
	// 崩溃恢复
	app.Use(middleware.RecoveryMiddleware())
	app.LoadHTMLGlob(config.Web.StaticPath+"dist/*.html")
	app.Static("/static", config.Web.StaticPath+"dist/static")
	app.Static("/resource", config.Web.StaticPath+"resource")
	app.StaticFile("/favicon.ico", config.Web.StaticPath+"dist/favicon.ico")
	// 注册路由
	routers.RegisterRouter(app)
  go initHTTPServer(config,app)
}

// InitHTTPServer 初始化http服务
func initHTTPServer(config *config.Config,handler http.Handler) {
	srv := &http.Server{
		Addr:         ":"+convert.ToString(config.Web.Port),
		Handler:      handler,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
	srv.ListenAndServe()
}
