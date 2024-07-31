package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID              uint   `gorm:"column:id;primaryKey"`
	Name            string `gorm:"column:name;primaryKey"`
	Email           string `gorm:"column:email;payload:varchar(100);uniqueIndex"`
	EncodedPassword string `gorm:"column:password"`
}
