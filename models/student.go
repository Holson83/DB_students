package models

import (
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	FullName string
	Age      uint
	//Direction Direction `gorm:"foreignKey:DirectionID"`
}
