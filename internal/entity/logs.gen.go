// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

import (
	"time"
)

const TableNameLog = "logs"

// Log mapped from table <logs>
type Log struct {
	ID         int32     `gorm:"column:id;type:int(11);primaryKey" json:"id"`
	Qid        int32     `gorm:"column:qid;type:int(11);not null;comment:被修改的qid" json:"qid"`       // 被修改的qid
	OldAnswer  string    `gorm:"column:old_answer;type:longtext;comment:修改之前的答案" json:"old_answer"` // 修改之前的答案
	NewAnswer  string    `gorm:"column:new_answer;type:longtext;comment:修改之后的答案" json:"new_answer"` // 修改之后的答案
	UserID     int32     `gorm:"column:user_id;type:int(11);comment:谁修改的答案" json:"user_id"`         // 谁修改的答案
	CreateTime time.Time `gorm:"column:create_time;type:datetime;comment:修改日期" json:"create_time"`  // 修改日期
}

// TableName Log's table name
func (*Log) TableName() string {
	return TableNameLog
}
