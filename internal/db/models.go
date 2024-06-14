package database

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName string `gorm:"size:50; not null"`
	LastName  string `gorm:"size:50; not null"`
	Email     string `gorm:"size:50; not null"`
}
