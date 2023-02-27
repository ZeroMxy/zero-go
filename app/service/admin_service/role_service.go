package admin_service

import (
	"zero-go/app/model"
	"xorm.io/xorm"
)

// 角色列表
func Get_role_list () *xorm.Session {

	return model.DB().Table("role").Desc("id")

}

// 角色详情
func Get_role_info (id int, name string) *model.Role {

	var role = model.Role {}

	var role_info_sql = model.DB().Table("role")

	if (id > 0) {
		role_info_sql = role_info_sql.Where("id = ?", id)
	}

	if (name != "") {
		role_info_sql = role_info_sql.Where("name = ?", name)
	}

	role_info_sql.Get(&role)

	return &role

}

// 添加角色
func Create_role (role *model.Role) bool {

	var affected, _ = model.DB().Table("role").Insert(role)

	if (affected > 0) {
		return true
	}

	return false

}

// 删除角色
func Delete_role (id int) bool {

	var affected, _ = model.DB().Table("role").Where("id = ?", id).Delete(&model.Role {})

	if (affected > 0) {
		return true
	}

	return false

}

// 更新角色
func Update_role (role *model.Role) bool {

	var affected, _ = model.DB().Table("role").Update(role)

	if (affected > 0) {
		return true
	}

	return false

}

// 角色列表无分页
func Select_role (name string) *[]model.Role {

	var role_select = []model.Role {}

	var role_select_sql = model.DB().Table("role").Desc("id").Where("status = 1")

	if (name != "") {
		role_select_sql = role_select_sql.Where("name like ?", "%" + name + "%")
	}

	role_select_sql.Find(&role_select)

	return &role_select

}