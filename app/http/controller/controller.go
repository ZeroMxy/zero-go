// 响应
package controller

import (
	"github.com/ZeroMxy/fastgo"
	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Time 	string 			`json:"time"`
	Code    int         	`json:"code"`
	Message string      	`json:"message"`
	Data    interface {} 	`json:"data"`
}

const (
	SUCCESS 	= 1 	// 成功
	FAIL 		= 0 	// 失败
	TOKEN_FAIL 	= -1 	// 缺少 token
)

// 成功返回
func (this *Response) Success (context *fiber.Ctx, data interface {}) {

	context.JSON(&Response {
		Time:		fastgo.Format_date_time(fastgo.YMD, ""),
		Code: 		SUCCESS,
		Message: 	"ok",
		Data: 		data,
	})

	return

}

// 失败返回
func (this *Response) Fail (context *fiber.Ctx, message string) {

	context.JSON(&Response {
		Time:		fastgo.Format_date_time(fastgo.YMD, ""),
		Code: 		FAIL,
		Message: 	message,
		Data: 		nil,
	})

	return

}

// token 失败返回
func (this *Response) Token_fail (context *fiber.Ctx) {

	context.JSON(&Response {
		Time:		fastgo.Format_date_time(fastgo.YMD, ""),
		Code: 		TOKEN_FAIL,
		Message: 	"token 失效",
		Data: 		nil,
	})

	return

}

// 分页返回
func (this *Response) Pager (context *fiber.Ctx, data interface {}, total, current, size int) {

	context.JSON(&Response {
		Time:		fastgo.Format_date_time(fastgo.YMD, ""),
		Code: 		SUCCESS,
		Message: 	"ok",
		Data: 		map[string] interface {}{
			"rows": 	data,
			"total": 	total,
			"current": 	current,
			"size": 	size,
		},
	})

	return

}