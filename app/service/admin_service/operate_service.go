package admin_service

import (
	"zero-go/app/model"
	"xorm.io/xorm"
)

// 操作权限列表
func Get_operate_list () *xorm.Session {

	return model.DB().Table("operate").Desc("id")

}

// 操作权限详情
func Get_operate_info (id int, name, path, method string) *model.Operate {

	var operate = model.Operate {}

	var operate_info_sql = model.DB().Table("operate")

	if (id > 0) {
		operate_info_sql = operate_info_sql.Where("id = ?", id)
	}

	if (name != "") {
		operate_info_sql = operate_info_sql.Where("name = ?", name)
	}

	if (path != "" && method != "") {
		operate_info_sql = operate_info_sql.Where("path = ?", path).Where("method = ?", method)
	}

	operate_info_sql.Find(&operate)

	return &operate

}

// 添加操作权限
func Create_operate (operate *model.Operate) bool {

	var affected, _ = model.DB().Table("operate").Insert(operate)

	if (affected > 0) {
		return true
	}

	return false

}

// 删除操作权限
func Delete_operate (id int) bool {

	var affected, _ = model.DB().Table("operate").Where("id = ?", id).Delete(&model.Operate {})

	if (affected > 0) {
		return true
	}

	return false

}

// 更新操作权限
func Update_operate (operate *model.Operate) bool {

	var affected, _ = model.DB().Table("operate").Update(operate)

	if (affected > 0) {
		return true
	}

	return false

}

// 操作权限列表无分页
func Select_operate (name string) *[]model.Operate {

	var operate_select = []model.Operate {}

	var role_select_sql = model.DB().Table("operate").Desc("id").Where("status = 1")

	if (name != "") {
		role_select_sql = role_select_sql.Where("name like ?", "%" + name + "%")
	}

	role_select_sql.Find(&operate_select)

	return &operate_select

}

// 操作权限绑定至菜单
func Operate_menu (operate_ids_array []string, menu_id int) bool {

	var affected, _ = model.DB().Table("operate").
						In("id", operate_ids_array).
						Update(map[string] int { "menu_id": menu_id})

	if (affected > 0) {
		return true
	}

	return false

}