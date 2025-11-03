package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     			string `json:"name"`
	Email    			string `gorm:"unique" json:"email"`
	Password 			string `json:"-"`
	Role				string `json:"role" gorm:"type:enum('admin','guru','siswa','orang_tua');default:'siswa'"`
}
