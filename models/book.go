package models

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title string
	Author string
	ISBN string
	PublishedDate time.Time
	Genre string
	Pages uint
}