package session

import (
	"sync"

	"zero-go/framework/log"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type Session struct {}

var (
	sess 		*session.Session
	sess_err 	error
	sess_once 	sync.Once
)

// 初始化 session
func session_initialize (context *fiber.Ctx) {

	sess_once.Do(func () {

		var store = session.New()

		sess, sess_err = store.Get(context)

		if (sess_err != nil) {
			log.Error(sess_err)
			return
		}
		
	})
	
	return

}

// 设置值
func Set (context *fiber.Ctx, key string, value interface {}) bool {

	session_initialize(context)
	
	sess.Set(key, value)

	return true

}

// 获取值
func Get (context *fiber.Ctx, key string) string {

	session_initialize(context)

	var value = sess.Get(key)

	if (value == nil) {
		return ""
	}

	return value.(string)

}

// 删除key
func Delete (context *fiber.Ctx, key string) bool {

	session_initialize(context)

	sess.Delete(key)
	
	return true

}