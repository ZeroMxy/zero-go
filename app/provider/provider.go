package provider

import (
	"zero-go/config"
	"zero-go/framework/log"
	"zero-go/route"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Provider struct {}

// 路由服务
func (this *Provider) Route_server (app *fiber.App) {

	// 开启 api 请求日志
	if (config.Log["system"] == "true") {
		var log_file = log.Initialize("system")
	
		app.Use(logger.New(logger.Config {
			Format :		"[${time}] ${status} ${method} ${protocol}://${host}${url} ${resBody}\n",
			TimeZone : 		"Asia/Chongqing",
			TimeFormat : 	"2006-01-02 15:04:05",
			Output : 		log_file,
		}))
	}

	// 路由注册
	route.Tool_route(app)
	route.Admin_route(app)

}
