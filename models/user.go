package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	Model

	Nickname string `json:"nickname"`
	HeadUrl string `json:"head_url"`
	State int `json:"state"`
	CreateTime int `json:"create_time"`
	UpdateTime int  `json:"update_time"`
}

func GetUsers(pageNum int, pageSize int, maps interface {}) (Users [] User) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&Users)

	return
}

func GetUserTotal(maps interface {}) (count int){
	db.Model(&User{}).Where(maps).Count(&count)

	return
}

func ExistUserByName(name string) bool {
	var user User
	db.Select("id").Where("nickname = ?", name).First(&user)
	if user.ID > 0 {
		return true
	}

	return false
}

func AddUser(name string, state int) bool{
	db.Create(&User {
		Nickname : name,
		State : state,
	})
	return true
}

func (user *User) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreateTime", time.Now().Unix())
	return nil
}

func (user *User) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("UpdateTime", time.Now().Unix())
	return nil
}

func ExistUserByID(id int) bool {
	var user User
	db.Select("id").Where("id = ?", id).First(&user)
	if user.ID > 0 {
		return true
	}

	return false
}

func DeleteUser(id int) bool {
	db.Where("id = ?", id).Delete(&User{})

	return true
}

func EditUser(id int, data interface {}) bool {
	db.Model(&User{}).Where("id = ?", id).Updates(data)

	return true
}