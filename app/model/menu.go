package model

// 菜单
type Menu struct {
	Model				`xorm:"extends"`
	Parent_id 	int 	`json:"parent_id"`
	Name 		string 	`json:"name"`
	Key			string 	`json:"key"`
	Icon		string	`json:"icon"`
	Path		string	`json:"path"`
	Redirect	string	`json:"redirect"`
	Component	string	`json:"component"`
	Status 		int 	`json:"status"`
}

// 菜单扩展
type Menu_extend struct {
	Menu
	Children []Menu_extend	`json:"children"`
}