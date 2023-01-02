package models

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	Name    string
	Age     string
	Class   string
	Section string
}
