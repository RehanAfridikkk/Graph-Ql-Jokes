package model

import "gorm.io/gorm"

type Joke struct {
	gorm.Model
	Content string `gorm:"not null"`
}
