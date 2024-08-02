package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID              uint   `gorm:"column:id;primaryKey"`
	Email           string `gorm:"column:email;payload:varchar(100);uniqueIndex"`
	Name            string `gorm:"column:name"`
	EncodedPassword string `gorm:"column:password"`

	ProfileUrl string `gorm:"column:profile_url"`
}
