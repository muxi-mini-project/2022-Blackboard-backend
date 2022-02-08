package model

import (


	"github.com/jinzhu/gorm"
)

//用户信息
type User struct {
	gorm.Model
	StudentID string `json:"student_id" binding:"required" gorm:"student_id"`
	PassWord  string `json:"password" binding:"required" gorm:"column:password"`
	NickName  string `json:"nickname" gorm:"nickname"`
	Avatar    string
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
	StudentID      string `json:"student_id" gorm:"student_id"`
	AnnouncementID uint   `json:"announcement_id" binding:"required" gorm:"announcement_id"`
	Announcement   string `json:"announcement" gorm:"announcement"`
}

//用户关注的组织
type FollowingOrganization struct {
	gorm.Model
	StudentID        string `json:"student_id" gorm:"student_id"`
	OrganizationID   uint   `json:"organization_id" gorm:"org_id"`
	OrganizationName string `json:"organization_name" gorm:"org_name"`
}

//组织信息
type Organization struct {
	gorm.Model
	FounderID         string `json:"founder_id"  gorm:"founder_id"`
	OrganizationName  string `json:"organization_name" binding:"required" gorm:"organization_name"`
	OrganizationIntro string `json:"intro" gorm:"intro"`
	Avatar            string
	Sha               string
	Path              string
}

//组织公告分组
type Grouping struct {
	gorm.Model
	OrganizationID   uint   `json:"organization_id" gorm:"organization_id"`
	OrganizationName string `json:"organization_name" binding:"required" gorm:"organization_name"`
	GroupName        string `json:"group_name" gorm:"group_name"`
}

//通告
type Announcement struct {
	gorm.Model
	PublisherID      string `json:"publisher_id" gorm:"publisher_id"`
	OrganizationID   uint   `json:"organization_id" gorm:"org_id"`
	OrganizationName string `json:"organization_name" gorm:"org_name"`
	GroupID          uint   `json:"group_id" gorm:"group_id"`
	GroupName        string `json:"group_name" gorm:"group_name"`
	Contents         string `json:"contents" gorm:"contents"`
}
