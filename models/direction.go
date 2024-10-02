package models

import (
	"gorm.io/gorm"
)

type Direction struct {
	gorm.Model
	Name   string
	Number int
}
