package models

import (
	"gorm.io/gorm"
)

type Fact struct {
	gorm.Model
	Questions string `json:"questions" gorm:"text;not null;defoult:null"`
	Answers   string `json:"answers" gorm:"text;not null;defoult:null"`
}
