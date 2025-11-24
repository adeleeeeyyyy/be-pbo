package models

import "gorm.io/gorm"

type Books struct{
	gorm.Model
	Title	string
	Author	string
	Publisher	string
	Category	string
	Stock	int
	Description	string	`gorm:"type:text"`
	Image	 string
	Status 	string `gorm:"type:enum('tersedia', 'dipinjam')"`
}