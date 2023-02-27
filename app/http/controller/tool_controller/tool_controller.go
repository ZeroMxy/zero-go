package tool_controller

import (
	"strings"

	"zero-go/app/http/controller"
	"zero-go/framework/captcha"
	"github.com/ZeroMxy/fastgo"
	"zero-go/framework/upload"
	"github.com/gofiber/fiber/v2"
)

type Tool_controller struct {
	controller.Response
}

// 图形验证码
func (this *Tool_controller) Captcha (context *fiber.Ctx) error {

	var base64_captcha = captcha.Create(context, "storage/font/Zaio.ttf")

	this.Success(context, map[string] string {
		"captcha": base64_captcha,
	})

	return nil

}

// 上传
func (this *Tool_controller) Upload (context *fiber.Ctx) error {

	var file, file_err = context.FormFile("file")

	if (file_err != nil || file == nil) {
		this.Fail(context, "上传失败")
		return nil
	}

	var file_name_array = strings.Split(file.Filename, ".")
	
	var suffix = []string { "jpg", "gif", "png", "jpeg", "docx", "doc", "xlsx", "xls" }
	
	// 文件后缀不存在/不在后缀数组中
	if (len(file_name_array) <= 1 || !fastgo.In_array(suffix, file_name_array[1])) {
		this.Fail(context, "文件格式错误")
		return nil
	}

	var url = upload.Local(context, file)

	this.Success(context, map[string] string {
		"url": url,
	})

	return nil

}