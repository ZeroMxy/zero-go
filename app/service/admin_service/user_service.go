package admin_service

import (
	"zero-go/app/model"
	"xorm.io/xorm"
)

// 用户列表
func Get_user_list (user_type, status, username, nickname, phone string) *xorm.Session {

	var user_list_sql = model.DB().
						Table("user").
						Join("inner", "role", "role.id = user.role_id").
						Select("user.*, role.id as role_id, role.name as role_name").
						Desc("user.id")

	// 条件
	if (user_type != "-1") {
		user_list_sql = user_list_sql.Where("user.user_type = ?", user_type)
	}

	if (status != "-1") {
		user_list_sql = user_list_sql.Where("user.status = ?", status)
	}

	if (username != "") {
		user_list_sql = user_list_sql.Where("user.username like ?", "%" + username + "%")
	}

	if (nickname != "") {
		user_list_sql = user_list_sql.Where("user.nickname like ?", "%" + nickname + "%")
	}

	if (phone != "") {
		user_list_sql = user_list_sql.Where("user.phone = ?", phone)
	}

	return user_list_sql

}

// 用户详情
func Get_user_info (id int, username string) *model.User_extend {

	var user_extend = model.User_extend {}

	var user_info_sql = model.DB().Table("user").
						Join("inner", "role", "role.id = user.role_id").
						Select("user.*, role.id as role_id, role.name as role_name")

	if (id > 0) {
		user_info_sql = user_info_sql.Where("user.id = ?", id)
	}

	if (username != "") {
		user_info_sql = user_info_sql.Where("user.username = ?", username)
	}

	user_info_sql.Get(&user_extend)

	return &user_extend

}

// 添加用户
func Create_user (user *model.User) bool {

	var affected, _ = model.DB().Table("user").Insert(user)

	if (affected > 0) {
		return true
	}

	return false

}

// 删除用户
func Delete_user (id int) bool {

	var affected, _ = model.DB().Table("user").Where("id = ?", id).Delete(&model.User {})

	if (affected > 0) {
		return true
	}

	return false

}

// 更新用户
func Update_user (user *model.User) bool {

	var affected, _ = model.DB().Table("user").Where("id = ?", user.Id).Update(user)

	if (affected > 0) {
		return true
	}

	return false

}