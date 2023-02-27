package middleware

import (
	"strconv"
	"strings"

	"zero-go/app/http/controller"
	"zero-go/app/service/admin_service"
	"zero-go/framework/session"

	"github.com/gofiber/fiber/v2"
)

type Admin_middleware struct {
	controller.Response
}

func (this *Admin_middleware) Handle (context *fiber.Ctx) error {

	var token = context.Query("token")

	if (token == "") {
		// token 为空判断
		this.Token_fail(context)
		return nil
	}

	var user_id = session.Get(context, token)

	if (user_id == "") {
		// 用户信息不存在
		this.Token_fail(context)
		return nil
	}

	var uid, _ = strconv.Atoi(user_id)

	// 鉴权
	var authentication_result = authentication(context, uid)

	if (!authentication_result) {
		this.Fail(context, "权限不足")
		return nil
	}

	context.Next()
	
	return nil

}

// 鉴权
func authentication (context *fiber.Ctx, uid int) bool {

	var user = admin_service.Get_user_info(uid, "")

	if (user == nil || user.Status == 0 || user.User_type != 1) {
		// 非后台/禁用状态
		return false
	}

	if (user.Role_id == 1) {
		// 超级管理员
		return true
	}

	var role = admin_service.Get_role_info(user.Role_id, "")

	if (role == nil || role.Status == 0 || (role.Id != 1 && role.Menu_id == "")) {
		// 无角色/禁用状态/不是超管没绑定菜单
		return false
	}

	// 角色绑定的菜单 id 集合进行分割
	var menu_id_array = strings.Split(role.Menu_id, ",")

	var path 	= context.Path()
	var method 	= context.Method()

	var operate = admin_service.Get_operate_info(0, "", path, method)

	if (operate == nil)  {
		// 未找到操作权限
		return false
	}

	// 初始化 is_exist 存在 => true，不存在 => false
	var is_exist = false
	
	// 寻找菜单与访问的操作地址关系
	for _, value := range menu_id_array {

		if (value == strconv.Itoa(operate.Menu_id)) {
			is_exist = true
			break
		}

	}
	
	return is_exist

}
