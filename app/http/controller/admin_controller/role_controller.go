package admin_controller

import (
	"strconv"

	"zero-go/app/http/controller"
	"zero-go/app/model"
	"zero-go/app/service/admin_service"
	"github.com/gofiber/fiber/v2"
)

type Role_controller struct {
	controller.Response
}

// 角色列表
func (this *Role_controller) Role_list (context *fiber.Ctx) error {

	var current, _ 	= strconv.Atoi(context.Query("current", "1"))
	var size, _ 	= strconv.Atoi(context.Query("size", "10"))

	var roles = []model.Role {}

	var total, _ = admin_service.Get_role_list().Limit(size, (current - 1) * size).FindAndCount(&roles)

	this.Pager(context, roles, int(total), current, size)

	return nil

}

// 角色详情
func (this *Role_controller) Role_info (context *fiber.Ctx) error {

	var id, _ = strconv.Atoi(context.Query("id"))

	if (id <= 0) {
		this.Fail(context, "缺少参数")
		return nil
	}

	this.Success(context, admin_service.Get_role_info(id, ""))

	return nil

}

// 添加角色
func (this *Role_controller) Create_role (context *fiber.Ctx) error {

	var status, _ 	= strconv.Atoi(context.Query("status", "1"))
	var name 		= context.Query("name")
	var menu_id 	= context.Query("menu_id")

	if (name == "") {
		this.Fail(context, "请添加角色名")
		return nil
	}

	var role = admin_service.Get_role_info(0, name)

	if (role != nil) {
		this.Fail(context, "角色名已存在")
		return nil
	}

	var result = admin_service.Create_role(&model.Role {
					Name: 		name,
					Menu_id: 	menu_id,
					Status: 	status,
				})

	if (!result) {
		this.Fail(context, "添加失败")
		return nil
	}

	this.Success(context, nil)

	return nil

}

// 删除角色
func (this *Role_controller) Delete_role (context *fiber.Ctx) error {

	var id, _ = strconv.Atoi(context.Query("id"))

	if (id <= 1) {
		this.Fail(context, "超管角色不可删除")
		return nil
	}

	var result = admin_service.Delete_role(id)

	if (!result) {
		this.Fail(context, "删除失败")
		return nil
	}

	this.Success(context, nil)

	return nil

}

// 更新角色
func (this *Role_controller) Update_role (context *fiber.Ctx) error {

	var id, _ 		= strconv.Atoi(context.Query("id"))
	var status, _ 	= strconv.Atoi(context.Query("status", "1"))
	var name 		= context.Query("name")
	var menu_id 	= context.Query("menu_id")

	if (id <= 1) {
		this.Fail(context, "超管角色不可修改")
		return nil
	}

	var role = model.Role {
		Name: 		name,
		Menu_id: 	menu_id,
		Status: 	status,
	}

	role.Id = id

	var result = admin_service.Update_role(&role)

	if (!result) {
		this.Fail(context, "修改失败")
		return nil
	}

	this.Success(context, nil)

	return nil

}

// 角色列表无分页
func (this *Role_controller) Select_role (context *fiber.Ctx) error {

	var name = context.Query("name")

	this.Success(context, admin_service.Select_role(name))

	return nil

}