package model

import "github.com/jinzhu/gorm"

func Migrate(DB *gorm.DB) {
	DB.AutoMigrate(&User{}, &Collection{}, &FollowingOrganization{}, &Organization{}, &Grouping{}, &Announcement{})
}
