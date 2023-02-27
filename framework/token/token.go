package token

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"

	"zero-go/framework/session"
	"github.com/gofiber/fiber/v2"
)

type Token struct {}

// 生成token
func Create (context *fiber.Ctx, value string) string {

	var time_string = []byte(fmt.Sprint(time.Now().Unix()) + fmt.Sprint(value))

	var token_byte = md5.Sum(time_string)

	// [size]byte 转字符串
	var token = hex.EncodeToString((token_byte[:]))

	var session_bool = session.Set(context, token, value)

	if (session_bool) {
		return token
	}

	return ""

}