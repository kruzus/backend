package models

import (
	"github.com/jinzhu/gorm"
	"gorm.io/gorm"
)

//this has no purpoe; will be deleted.
type D struct {
	test gorm.ConnPool
}

type User struct {
	ID       uint64 `json: "id" gorm:"primary_key"`
	Name     string `json: "name"`
	Password string `json: "password"`
}
