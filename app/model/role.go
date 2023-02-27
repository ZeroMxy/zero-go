package model

// 角色
type Role struct {
	Model			`xorm:"extends"`
	Name 	string 	`json:"name"`
	Menu_id string  `json:"menu_id"`
	Status 	int 	`json:"status"`
}

// 角色扩展（联表查询角色菜单权限）
type Role_extend struct {
	Model					`xorm:"extends"`
	Name 	string 			`json:"name"`
	Status 	int 			`json:"status"`
	Menus 	[]Menu_extend	`json:"menus"`
}