package model

import (
// "errors"
// "github.com/jinzhu/gorm"
)

func GetUserInfo(id string) (User, error) {
	var u User
	//// SELECT * FROM users ORDER BY id LIMIT 1;

	return u, DB.Where("student_id=?", id).First(&u).Error
}
func ChangeName(user User) error {
	return DB.Model(&user).Where("student_id = ?", user.StudentID).Update(user.NickName).Error
}

func GetCollection(id string) (Collection, error) {
	var collect Collection
	return collect, DB.Where("student_id = ?", id).Find(&collect).Error

}

func GetPublished(id string) (Announcement, error) {
	var announce Announcement
	return announce, DB.Where("publisher_id = ?", id).Find(&announce).Error
}

func GetCreated(id string) (Organization, error) {
	var created Organization
	return created, DB.Where("founder_id = ?", id).Find(&created).Error
}
