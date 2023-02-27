package captcha

import (
	"bytes"
	"encoding/base64"
	"image/color"
	"image/png"
	"unsafe"

	"zero-go/framework/session"
	"github.com/afocus/captcha"
	"github.com/gofiber/fiber/v2"
)

type Captcha struct {}

// 创建图形验证码
func Create (context *fiber.Ctx, font_path string) string {

	// 初始化
	var captcha_object = captcha.New()
	// 设置字体
	captcha_object.SetFont(font_path)
	// 设置颜色
	captcha_object.SetFrontColor(color.Black, color.RGBA{255, 0, 0, 255}, color.RGBA{0, 0, 255, 255}, color.RGBA{0, 153, 0, 255})

	var captcha_image, captcha_context = captcha_object.Create(4, captcha.NUM)

	var session_captcha = session.Set(context, "captcha", captcha_context)

	if (!session_captcha) {
		return ""
	}

	// 开辟一个新的空 buff
	var base64_buff = bytes.NewBuffer(nil)

	png.Encode(base64_buff, captcha_image)
	// 开辟存储空间
	var base64_byte = make([]byte, 5000)
	// buff 转成 base64
	base64.StdEncoding.Encode(base64_byte, base64_buff.Bytes())
	// 去除未被填充完部分
	var index = bytes.IndexByte(base64_byte, 0)
	var base64_image = base64_byte[0: index]

	return "data:image/png;base64," + *(*string)(unsafe.Pointer(&base64_image))

}

// 验证图形验证码
func Verify (context *fiber.Ctx, captcha string) bool {

	// session 获取存入验证码内容
	var session_captcha = session.Get(context, "captcha")

	if (session_captcha == "" || session_captcha != captcha) {
		return false
	}

	return true

}