package model

import (
	"fiy/common/models"
)

/*
  @Author : lanyulei
*/

// 模型信息
type Info struct {
	Id         int    `gorm:"column:id; primary_key;AUTO_INCREMENT" json:"id"`
	Identifies string `gorm:"column:identifies; type:varchar(128)" json:"identifies" binding:"required"` // 分组唯一标识，只能是英文
	Name       string `gorm:"column:name; type:varchar(128)" json:"name" binding:"required"`             // 分组名称
	Icon       string `gorm:"column:icon; type:varchar(128)" json:"icon" binding:"required"`             // 模型图标
	IsUsable   bool   `gorm:"column:is_usable; type:tinyint(1)" json:"is_usable"`                        // 是否可用
	IsInternal bool   `gorm:"column:is_internal; type:tinyint(1)" json:"is_internal"`                    // 是否内置
	GroupId    int    `gorm:"column:group_id; type:int(11)" json:"group_id" binding:"required"`          // 模型分组ID
	models.BaseModel
}

func (Info) TableName() string {
	return "cmdb_model_info"
}