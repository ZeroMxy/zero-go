package bootstrap

import (
	"zero-go/app/provider"
	"zero-go/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type App struct {}

// 启动服务
func (this *App) Run () {

	var app = fiber.New()

	// 配置跨域
	app.Use(cors.New())

	// 静态文件
	app.Static("/storage/upload", "storage/upload")

	// 加载路由服务
	(&provider.Provider {}).Route_server(app)

	app.Listen(config.App["host"])

}
