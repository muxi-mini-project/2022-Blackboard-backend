package model

import (
	"errors"
	"github.com/jinzhu/gorm"
)

func ChangeName(user User) error {
	return DB.Model(&user).Where("student_id = ?", user.StudentID).Update(user).Error
}
