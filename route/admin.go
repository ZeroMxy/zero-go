package route

import (
	"zero-go/app/http/controller/admin_controller"
	"zero-go/app/http/middleware"
	"github.com/gofiber/fiber/v2"
)

func Admin_route (app *fiber.App) {

	var admin_app = app.Group("api/admin")

	{
		admin_app.Post("user/login", (&admin_controller.User_controller {}).Login)
	}

	// 使用中间件
	var admin_auth_app = app.Group("api/admin", (&middleware.Admin_middleware {}).Handle)

	{
		// 用户
		admin_auth_app.Post("user/list", (&admin_controller.User_controller {}).User_list).Name("用户列表")
		admin_auth_app.Post("user/info", (&admin_controller.User_controller {}).User_info).Name("用户详情")
		admin_auth_app.Post("user/create", (&admin_controller.User_controller {}).Create_user).Name("添加用户")
		admin_auth_app.Post("user/delete", (&admin_controller.User_controller {}).Delete_user).Name("删除用户")
		admin_auth_app.Post("user/update", (&admin_controller.User_controller {}).Update_user).Name("更新用户")

		// 角色
		admin_auth_app.Post("role/list", (&admin_controller.Role_controller {}).Role_list).Name("角色列表")
		admin_auth_app.Post("role/info", (&admin_controller.Role_controller {}).Role_info).Name("角色详情")
		admin_auth_app.Post("role/create", (&admin_controller.Role_controller {}).Create_role).Name("添加角色")
		admin_auth_app.Post("role/delete", (&admin_controller.Role_controller {}).Delete_role).Name("删除角色")
		admin_auth_app.Post("role/update", (&admin_controller.Role_controller {}).Update_role).Name("修改角色")
		admin_auth_app.Post("role/select", (&admin_controller.Role_controller {}).Select_role).Name("角色列表无分页")

		// 菜单
		admin_auth_app.Post("menu/list", (&admin_controller.Menu_controller {}).Menu_list).Name("菜单列表")
		admin_auth_app.Post("menu/info", (&admin_controller.Menu_controller {}).Menu_info).Name("菜单详情")
		admin_auth_app.Post("menu/create", (&admin_controller.Menu_controller {}).Create_menu).Name("添加菜单")
		admin_auth_app.Post("menu/delete", (&admin_controller.Menu_controller {}).Delete_menu).Name("删除菜单")
		admin_auth_app.Post("menu/update", (&admin_controller.Menu_controller {}).Update_menu).Name("更新菜单")
		admin_auth_app.Post("menu/update", (&admin_controller.Menu_controller {}).Menu_operate).Name("菜单绑定操作权限")

		// 操作权限
		admin_auth_app.Post("operate/list", (&admin_controller.Operate_controller {}).Operate_list).Name("操作权限列表")
		admin_auth_app.Post("operate/info", (&admin_controller.Operate_controller {}).Operate_info).Name("操作权限详情")
		admin_auth_app.Post("operate/create", (&admin_controller.Operate_controller {}).Create_operate).Name("添加操作权限")
		admin_auth_app.Post("operate/delete", (&admin_controller.Operate_controller {}).Delete_operate).Name("删除操作权限")
		admin_auth_app.Post("operate/update", (&admin_controller.Operate_controller {}).Update_operate).Name("更新操作权限")
		admin_auth_app.Post("operate/select", (&admin_controller.Operate_controller {}).Select_operate).Name("操作权限列表无分页")
		admin_auth_app.Post("operate/sync", (&admin_controller.Operate_controller {}).Sync_operate).Name("操作权限同步")
	}

}