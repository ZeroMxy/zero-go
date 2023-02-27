package admin_controller

import (
	"strconv"

	"zero-go/app/http/controller"
	"zero-go/app/model"
	"zero-go/app/service/admin_service"
	"github.com/gofiber/fiber/v2"
)

type Operate_controller struct {
	controller.Response
}

// 操作权限列表
func (this *Operate_controller) Operate_list (context *fiber.Ctx) error {

	var current, _ 	= strconv.Atoi(context.Query("current", "1"))
	var size, _ 	= strconv.Atoi(context.Query("size", "10"))

	var operates = []model.Operate {}

	var total, _ = admin_service.Get_operate_list().
					Limit(size, (current - 1) * size).
					FindAndCount(&operates)

	this.Pager(context, operates, int(total), current, size)

	return nil
	
}

// 操作权限详情
func (this *Operate_controller) Operate_info (context *fiber.Ctx) error {

	var id, _ = strconv.Atoi(context.Query("id", "0"))

	if (id <= 0) {
		this.Fail(context, "缺少参数")
		return nil
	}

	var operate = admin_service.Get_operate_info(id, "", "", "")

	this.Success(context, operate)

	return nil
	
}

// 添加操作权限
func (this *Operate_controller) Create_operate (context *fiber.Ctx) error {

	var name 	= context.Query("name")
	var path 	= context.Query("path")
	var method 	= context.Query("method")

	if (name == "") {
		this.Fail(context, "请添加操作权限名称")
		return nil
	}

	if (path == "") {
		this.Fail(context, "请添加操作权限地址")
		return nil
	}

	if (method == "") {
		this.Fail(context, "请选择操作权限请求方式")
		return nil
	}

	var operate = admin_service.Get_operate_info(0, "", path, method)

	if (operate != nil) {
		this.Fail(context, "操作权限已存在")
		return nil
	}

	var result = admin_service.Create_operate(&model.Operate {
					Name: 	name,
					Path: 	path,
					Method: method,
				})

	if (!result) {
		this.Fail(context, "添加失败")
		return nil
	}

	this.Success(context, nil)

	return nil
	
}

// 删除操作权限
func (this *Operate_controller) Delete_operate (context *fiber.Ctx) error {

	var id, _ = strconv.Atoi(context.Query("id"))

	if (id <= 0) {
		this.Fail(context, "缺少参数")
		return nil
	}

	var result = admin_service.Delete_operate(id)

	if (!result) {
		this.Fail(context, "删除失败")
		return nil
	}

	this.Success(context, nil)

	return nil
	
}

// 更新操作权限
func (this *Operate_controller) Update_operate (context *fiber.Ctx) error {

	var id, _ 	= strconv.Atoi(context.Query("id"))
	var name 	= context.Query("name")
	var path 	= context.Query("path")
	var method 	= context.Query("method")

	if (id <= 0) {
		this.Fail(context, "缺少参数")
		return nil
	}

	var operate_info = admin_service.Get_operate_info(0, "", path, method)

	if (operate_info != nil && operate_info.Id != id) {
		this.Fail(context, "操作权限已存在")
		return nil
	}

	var operate = model.Operate {
		Name: 	name,
		Path: 	path,
		Method: method,
	}

	operate.Id = id

	var result = admin_service.Update_operate(&operate)

	if (!result) {
		this.Fail(context, "修改失败")
		return nil
	}

	this.Success(context, nil)

	return nil
	
}

// 操作权限列表无分页
func (this *Operate_controller) Select_operate (context *fiber.Ctx) error {

	var name = context.Query("name")

	this.Success(context, admin_service.Select_operate(name))

	return nil

}

// 同步操作权限
func (this *Operate_controller) Sync_operate (context *fiber.Ctx) error {
	
	// 获取所有路由
	var routes = context.App().GetRoutes()

	if (len(routes) <= 0) {
		this.Fail(context, "同步失败")
		return nil
	}

	for _, route := range routes {

		if (route.Name != "" && admin_service.Get_operate_info(0, "", route.Path, route.Method) == nil) {
			// 路由名不为空&库里不存在同步到数据库
			admin_service.Create_operate(&model.Operate{
				Menu_id: 	0,
				Name: 		route.Name,
				Path: 		route.Path,
				Method: 	route.Method,
			});

		}

	}

	this.Success(context, nil)

	return nil
	
}