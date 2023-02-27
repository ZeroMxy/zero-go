package model

import (
	"time"

	"zero-go/framework/database"
	"xorm.io/xorm"
)

type Model struct {
	Id         	int         	`json:"id"`
	Created_at 	Format_time   	`json:"created_at" xorm:"created"`
	Updated_at 	Format_time   	`json:"updated_at" xorm:"updated"`
	Deleted_at 	Format_time		`json:"deleted_at" xorm:"deleted"`
}

type Format_time time.Time

// 时间格式化
func (this Format_time) MarshalJSON () ([]byte, error) {

	//当返回时间为空时，需特殊处理
    if (time.Time(this).IsZero()) {
        return []byte(`""`), nil
    }

    return []byte(`"` + time.Time(this).Format("2006-01-02 15:04:05") + `"`), nil

}

// 返回实例
func DB () *xorm.Engine {

	return database.Conn()

}
