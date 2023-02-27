package model


// 操作权限
type Operate struct {
	Model
	Menu_id		int		`json:"menu_id"`
	Name		string	`json:"name"`
	Path		string	`json:"path"`
	Method		string	`json:"method"`
}