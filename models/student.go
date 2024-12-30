package models

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	Name string
	Age  int
	Height float32
	Weight float32
}