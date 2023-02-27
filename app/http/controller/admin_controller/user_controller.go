package admin_controller

import (
	"strconv"
	"strings"

	"zero-go/app/http/controller"
	"zero-go/app/model"
	"zero-go/app/service/admin_service"
	"zero-go/framework/captcha"
	"zero-go/framework/cipher"
	"zero-go/framework/session"
	"zero-go/framework/token"
	"github.com/gofiber/fiber/v2"
)

type User_controller struct {
	controller.Response
}

// 用户列表
func (this *User_controller) User_list (context *fiber.Ctx) error {

	var current, _ 	= strconv.Atoi(context.Query("current", "1"))
	var size, _ 	= strconv.Atoi(context.Query("size", "10"))
	var user_type 	= context.Query("user_type", "-1")
	var status 		= context.Query("status", "-1")
	var username 	= context.Query("username")
	var nickname 	= context.Query("nickname")
	var phone 		= context.Query("phone")

	var users = []model.User_extend {}

	var total, _ = admin_service.Get_user_list(user_type, status, username, nickname, phone).
					Limit(size, (current - 1) * size).
					FindAndCount(&users)

	this.Pager(context, users, int(total), current, size)

	return nil

}

// 用户详情
func (this *User_controller) User_info (context *fiber.Ctx) error {

	var id, _ = strconv.Atoi(context.Query("id", "0"))
	var token = context.Query("token")

	// 未传属于登录从缓存中获取 
	if (id <= 0) {
		id, _ = strconv.Atoi(session.Get(context, token))
	}

	if (id <= 0) {
		this.Fail(context, "缺少参数")
		return nil
	}

	var user_extend = admin_service.Get_user_info(id, "")

	var menu_id = make([]string, 0)

	if (user_extend.Id != 1) {
		// 获取角色绑定的菜单 id 集合进行分割
		var menu_id_string = admin_service.Get_role_info(user_extend.Role_id, "").Menu_id
		menu_id = strings.Split(menu_id_string, ",")
	}

	// 根据角色查询菜单权限
	var menus = admin_service.Get_menu_list(menu_id)

	user_extend.Menu = admin_service.Menu_to_tree(*menus, 0)
	
	this.Success(context, user_extend)

	return nil

}

// 添加用户
func (this *User_controller) Create_user (context *fiber.Ctx) error {

	var user_type, _ 	= strconv.Atoi(context.Query("user_type", "0"))
	var role_id, _ 		= strconv.Atoi(context.Query("role_id", "9999"))
	var status, _ 		= strconv.Atoi(context.Query("status", "1"))
	var sex, _ 			= strconv.Atoi(context.Query("sex"))
	var age, _ 			= strconv.Atoi(context.Query("age"))
	var username 		= context.Query("username")
	var nickname 		= context.Query("nickname")
	var phone 			= context.Query("phone")
	var email 			= context.Query("email")
	var avatar 			= context.Query("avatar")

	if (username == "") {
		this.Fail(context, "请添加用户名")
		return nil
	}

	if (phone == "") {
		this.Fail(context, "请添加手机号")
		return nil
	}

	var result = admin_service.Create_user(&model.User {
					User_type: 	user_type,
					Role_id: 	role_id,
					Username: 	username,
					Password: 	cipher.Encrypt("123456"),
					Nickname: 	nickname,
					Phone:  	phone,
					Email: 		email,
					Sex: 		sex,
					Age: 		age,
					Avatar:  	avatar,
					Status: 	status,
				})

	if (!result) {
		this.Fail(context, "添加失败")
		return nil
	}

	this.Success(context, nil)

	return nil
	
}

// 删除用户
func (this *User_controller) Delete_user (context *fiber.Ctx) error {

	var id, _ = strconv.Atoi(context.Query("id"))

	if (id <= 1) {
		this.Fail(context, "超管不可删除")
		return nil
	}

	var result = admin_service.Delete_user(id)

	if (!result) {
		this.Fail(context, "删除失败")
		return nil
	}

	this.Success(context, nil)

	return nil

}

// 更新用户
func (this *User_controller) Update_user (context *fiber.Ctx) error {

	var id, _ 			= strconv.Atoi(context.Query("id"))
	var user_type, _ 	= strconv.Atoi(context.Query("user_type", "0"))
	var role_id, _ 		= strconv.Atoi(context.Query("role_id", "9999"))
	var status, _ 		= strconv.Atoi(context.Query("status"))
	var sex, _ 			= strconv.Atoi(context.Query("sex"))
	var age, _ 			= strconv.Atoi(context.Query("age"))
	var username 		= context.Query("username")
	var password 		= context.Query("password")
	var nickname 		= context.Query("nickname")
	var phone 			= context.Query("phone")
	var email 			= context.Query("email")
	var avatar 			= context.Query("avatar")

	if (id <= 1) {
		this.Fail(context, "暂不支持修改超管或不存在的用户")
		return nil
	}

	var user = model.User {
		User_type: 	user_type,
		Role_id: 	role_id,
		Username: 	username,
		Password: 	cipher.Encrypt(password),
		Nickname: 	nickname,
		Phone:  	phone,
		Email: 		email,
		Sex: 		sex,
		Age: 		age,
		Avatar:  	avatar,
		Status: 	status,
	}

	user.Id = id

	var result = admin_service.Update_user(&user)

	if (!result) {
		this.Fail(context, "修改失败")
		return nil
	}

	this.Success(context, nil)

	return nil

}

// 登录
func (this *User_controller) Login (context *fiber.Ctx) error {

	var username 		= context.Query("username")
	var password 		= context.Query("password")
	var captcha_base64	= context.Query("captcha")

	if (!captcha.Verify(context, captcha_base64)) {
		this.Fail(context, "验证码错误")
		return nil
	}

	if (username == "" || password == "") {
		this.Fail(context, "请输入有效用户名或密码")
		return nil
	}

	var user = admin_service.Get_user_info(0, username)
	
	if (user.Id > 0) {
		// 验证密码
		var verify = cipher.Verify(user.Password, password)

		if (!verify) {
			this.Fail(context, "密码错误")
			return nil
		}

		if (user.Status == 0) {
			this.Fail(context, "账户被禁用")
			return nil
		}

		if (user.User_type != 1) {
			this.Fail(context, "非后台用户")
			return nil
		}

		var token = token.Create(context, strconv.Itoa(user.Id))

		if (token == "") {
			this.Fail(context, "登录失败")
			return nil
		}

		this.Success(context, map[string] string {
			"token": token,
		})
		return nil
	}

	if (username == "admin") {
		var result = admin_service.Create_user(&model.User {
						User_type: 	1,
						Username: 	"admin",
						Password: 	cipher.Encrypt(password),
						Nickname: 	"超级管理员",
						Phone: 		"110",
						Status: 	1,
					})
		if (result) {
			this.Fail(context, "超管创建成功，请重新登陆")
			return nil
		}
	} else {
		this.Fail(context, "用户不存在")
	}

	return nil

}
