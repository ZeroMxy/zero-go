package upload

import (
	"mime/multipart"
	"os"
	"strconv"
	"time"

	"zero-go/config"
	"github.com/ZeroMxy/fastgo"
	"github.com/gofiber/fiber/v2"
)

type Upload struct {}

// 文件上传至本地
func Local (context *fiber.Ctx, file *multipart.FileHeader) string {

	// 文件保存路径
	var file_folder_path = "storage/upload/"

	// 不存在则创建目录
	if (!fastgo.Path_is_exist(file_folder_path)) {
		os.MkdirAll(file_folder_path, os.ModePerm)
	}

	// 当前时间戳 + 原名称
	var file_name = strconv.FormatInt(time.Now().Unix(), 10) + file.Filename

	// 路径 + 文件名 + 后缀
	var file_path = file_folder_path + file_name

	context.SaveFile(file, file_path)

	return config.App["host"] + "/" + file_path

}
