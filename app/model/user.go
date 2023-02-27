package model

// 用户
type User struct {
	Model				`xorm:"extends"`
	User_type 	int 	`json:"user_type"`
	Role_id		int 	`json:"role_id"`
	Username 	string 	`json:"username"`
	Password 	string 	`json:"password"`
	Nickname	string 	`json:"nickname"`
	Phone 		string 	`json:"phone"`
	Email		string 	`json:"email"`
	Sex			int		`json:"sex"`
	Age			int		`json:"age"`
	Avatar   	string 	`json:"avatar"`
	Status   	int    	`json:"status"`
}

// 用户扩展（联表查询用户角色）
type User_extend struct {
	User						`xorm:"extends"`
	Role_name 	string 			`json:"role_name"`
	Menu		[]Menu_extend	`json:"menu"`
}