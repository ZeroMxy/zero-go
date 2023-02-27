package route

import (
	"zero-go/app/http/controller/tool_controller"
	"github.com/gofiber/fiber/v2"
)

func Tool_route (app *fiber.App) {

	var tool_app = app.Group("api")

	{
		tool_app.Post("tool/captcha", (&tool_controller.Tool_controller {}).Captcha).Name("图形验证码")
		tool_app.Post("tool/upload", (&tool_controller.Tool_controller {}).Upload).Name("文件上传")
		
	}

}