package model

import (
	"github.com/jinzhu/gorm"
)

//用户信息
type User struct {
	gorm.Model
	StudentID string `json:"student_id" binding:"required" gorm:"column:student_id;not null"`
	PassWord  string `json:"password" binding:"required" gorm:"column:password;not null"`
	NickName  string `json:"nickname" gorm:"column:nick_name;not null"`
	Avatar    string `gorm:"not null"`
	Sha       string
	Path      string
}

type Info struct {
	gorm.Model
	StudentID string
	NickName  string `json:"nickname"`
}

//用户收藏
type Collection struct {
	gorm.Model
	StudentID      string `json:"student_id" gorm:"column:student_id;not null"`
	AnnouncementID uint   `json:"announcement_id" binding:"required" gorm:"column:announcement_id;not null"`
	Announcement   string `json:"announcement" gorm:"size:500;not null"`
}

//用户关注的组织
type FollowingOrganization struct {
	gorm.Model
	StudentID        string `json:"student_id" gorm:"column:student_id;not null"`
	OrganizationID   uint   `json:"organization_id" gorm:"column:organization_id;not null"`
	OrganizationName string `json:"organization_name" gorm:"column:organization_name;not null"`
}

//组织信息
type Organization struct {
	gorm.Model
	FounderID         string `json:"founder_id"  gorm:"column:founder_id;not null"`
	OrganizationName  string `json:"organization_name" binding:"required" gorm:"column:organization_name;not null"`
	OrganizationIntro string `json:"intro" gorm:"column:intro;not null"`
	Avatar            string
	Sha               string
	Path              string
}

//组织公告分组
type Grouping struct {
	gorm.Model
	OrganizationID   uint   `json:"organization_id" gorm:"column:organization_id;not null"`
	OrganizationName string `json:"organization_name" binding:"required" gorm:"column:organization_name;not null"`
	GroupName        string `json:"group_name" gorm:"column:group_name;not null"`
}

//通告
type Announcement struct {
	gorm.Model
	PublisherID      string `json:"publisher_id" gorm:"column:publisher_id;not null"`
	OrganizationID   uint   `json:"organization_id" gorm:"column:organization_id;not null"`
	OrganizationName string `json:"organization_name" gorm:"column:organization_name;not null"`
	GroupID          uint   `json:"group_id" gorm:"column:group_id;not null"`
	GroupName        string `json:"group_name" gorm:"column:group_name;not null"`
	Contents         string `json:"contents" gorm:"size:500;not null"`
}
