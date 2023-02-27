package admin_service

import (
	"zero-go/app/model"
)

// 菜单列表
func Get_menu_list (ids []string) *[]model.Menu {

	var menus = []model.Menu {}

	var menus_sql = model.DB().Table("menu").Select("menu.*").Where("menu.status = 1")

	if (len(ids) > 0) {
		menus_sql = menus_sql.In("id", ids)
	}

	menus_sql.Find(&menus)

	return &menus

}

// 菜单详情
func Get_menu_info (id int, name string) *model.Menu {

	var menu = model.Menu {}

	var menu_info_sql = model.DB().Table("menu")

	if (id > 0) {
		menu_info_sql = menu_info_sql.Where("id = ?", id)
	}

	if (name != "") {
		menu_info_sql = menu_info_sql.Where("name = ?", name)
	}

	menu_info_sql.Get(&menu)

	return &menu

}

// 添加菜单
func Create_menu (menu *model.Menu) bool {

	var affected, _ = model.DB().Table("menu").Insert(menu)

	if (affected > 0) {
		return true
	}

	return false

}

// 删除菜单
func Delete_menu (id int) bool {

	var affected, _ = model.DB().Table("menu").Where("id = ?", id).Delete(&model.Menu {})

	if (affected > 0) {
		return true
	}

	return false

}

// 更新菜单
func Update_menu (menu *model.Menu) bool {

	var affected, _ = model.DB().Table("menu").Update(menu)

	if (affected > 0) {
		return true
	}

	return false

}

// 子级菜单
func Menu_children (parent_id int) *[]model.Menu {

	var menus = []model.Menu {}

	model.DB().Table("menu").Where("parent_id = ?", parent_id).Find(&menus)

	return &menus

}

// 无限级 tree 类型菜单
func Menu_to_tree (menus []model.Menu, parent_id int) []model.Menu_extend {

	// 初始化
	var menus_tree = []model.Menu_extend {}

	for _, value := range menus {
		// 循环中找到子级
		if (value.Status == 1 && value.Parent_id == parent_id) {
			// 获取子级菜单
			var children = Menu_to_tree(menus, value.Id)

			var menu_tree = model.Menu_extend {}
			// 初始化赋值
			menu_tree.Id 			= value.Id
			menu_tree.Parent_id 	= value.Parent_id
			menu_tree.Name 			= value.Name
			menu_tree.Icon 			= value.Icon
			menu_tree.Path 			= value.Path
			menu_tree.Redirect 		= value.Redirect
			menu_tree.Component 	= value.Component
			menu_tree.Key 			= value.Key
			menu_tree.Status 		= value.Status
			menu_tree.Created_at 	= value.Created_at
			menu_tree.Updated_at 	= value.Updated_at
			menu_tree.Children 		= children
			// 追加至菜单列表
			menus_tree = append(menus_tree, menu_tree)
		}
	}

	return menus_tree
	
}