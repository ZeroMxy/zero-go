package admin_controller

import (
	"strconv"
	"strings"

	"zero-go/app/http/controller"
	"zero-go/app/model"
	"zero-go/app/service/admin_service"
	"github.com/gofiber/fiber/v2"
)

type Menu_controller struct {
	controller.Response
}

// 菜单列表
func (this *Menu_controller) Menu_list (context *fiber.Ctx) error {

	var role_id, _ = strconv.Atoi(context.Query("role_id", "0"))

	// 获取角色绑定的菜单 id 集合进行分割
	var menu_id_array = admin_service.Get_role_info(role_id, "").Menu_id
	var menu_id = strings.Split(menu_id_array, ",")

	var menus = admin_service.Get_menu_list(menu_id)

	var menus_tree = admin_service.Menu_to_tree(*menus, 0)

	this.Success(context, menus_tree)

	return nil

}

// 菜单详情
func (this *Menu_controller) Menu_info (context *fiber.Ctx) error {

	var id, _ = strconv.Atoi(context.Query("id"))

	if (id <= 0) {
		this.Fail(context, "缺少参数")
		return nil
	}

	this.Success(context, admin_service.Get_menu_info(id, ""))

	return nil

}

// 添加菜单
func (this *Menu_controller) Create_menu (context *fiber.Ctx) error {

	var status, _ 		= strconv.Atoi(context.Query("status", "1"))
	var parent_id, _ 	= strconv.Atoi(context.Query("parent_id", "0"))
	var name 			= context.Query("name")
	var icon 			= context.Query("icon")
	var path 			= context.Query("path")
	var redirect 		= context.Query("redirect")
	var component 		= context.Query("component")
	var key 			= context.Query("key")

	if (name == "") {
		this.Fail(context, "请添加菜单名")
		return nil
	}

	var menu = admin_service.Get_menu_info(0, name)

	if (menu != nil) {
		this.Fail(context, "菜单名已存在")
		return nil
	}

	var result = admin_service.Create_menu(&model.Menu {
					Parent_id: 	parent_id,
					Name: 		name,
					Icon: 		icon,
					Path: 		path,
					Redirect: 	redirect,
					Component: 	component,
					Key: 		key,
					Status: 	status,
				})

	if (!result) {
		this.Fail(context, "添加失败")
		return nil
	}

	this.Success(context, nil)

	return nil

}

// 删除菜单
func (this *Menu_controller) Delete_menu (context *fiber.Ctx) error {

	var id, _ = strconv.Atoi(context.Query("id"))

	if (id <= 0) {
		this.Fail(context, "删除失败")
		return nil
	}

	var menu_children = admin_service.Menu_children(id)

	if (menu_children != nil) {
		this.Fail(context, "存在下级菜单")
		return nil
	}

	var result = admin_service.Delete_menu(id)

	if (!result) {
		this.Fail(context, "删除失败")
		return nil
	}

	this.Success(context, nil)

	return nil

}

// 更新菜单
func (this *Menu_controller) Update_menu (context *fiber.Ctx) error {

	var id, _ 			= strconv.Atoi(context.Query("id"))
	var status, _ 		= strconv.Atoi(context.Query("status", "1"))
	var parent_id, _ 	= strconv.Atoi(context.Query("parent_id", "0"))
	var name 			= context.Query("name")
	var icon 			= context.Query("icon")
	var path 			= context.Query("path")
	var redirect 		= context.Query("redirect")
	var component 		= context.Query("component")
	var key 			= context.Query("key")

	var menu_info = admin_service.Get_menu_info(0, name)

	if (menu_info != nil && menu_info.Id != id) {
		this.Fail(context, "菜单名已存在")
		return nil
	}

	var menu = model.Menu {
		Parent_id: 	parent_id,
		Name: 		name,
		Icon: 		icon,
		Path: 		path,
		Redirect: 	redirect,
		Component: 	component,
		Key: 		key,
		Status: 	status,
	}

	menu.Id = id

	var result = admin_service.Update_menu(&menu)

	if (!result) {
		this.Fail(context, "修改失败")
		return nil
	}

	this.Success(context, nil)

	return nil

}

// 权限绑定至菜单
func (this *Menu_controller) Menu_operate (context *fiber.Ctx) error {

	var operate_ids = context.Query("operate_ids")
	var id, _ 		= strconv.Atoi(context.Query("id", "0"))

	// 操作权限 id 集合进行分割
	var operate_ids_array = strings.Split(operate_ids, ",")

	if (len(operate_ids_array) <= 0 || id <= 0) {
		this.Fail(context, "缺少参数")
		return nil
	}

	var result = admin_service.Operate_menu(operate_ids_array, id)

	if (!result) {
		this.Fail(context, "绑定失败")
		return nil
	}

	this.Success(context, nil)

	return nil

}