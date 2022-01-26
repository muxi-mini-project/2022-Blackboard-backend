package model

import (
	"database/sql"

	"github.com/jinzhu/gorm"
)

//用户信息
type User struct {
	gorm.Model
	StudentID string `json:"student_id" gorm:"student_id"`
	PassWord  string `json:"password" gorm:"column:password"`
	NickName  string `json:"nickname" gorm:"nickname"`
	Avatar    sql.NullString
}

//用户收藏
type Collection struct {
	gorm.Model
	StudentID      string `json:"student_id" gorm:"student_id"`
	AnnouncementID string `json:"announcement_id" gorm:"announcement_id"`
	Announcement   string `json:"announcement" gorm:"announcement"`
}

//用户关注的组织
type FollowingOrganization struct {
	gorm.Model
	StudentID        string `json:"student_id" gorm:"student_id"`
	OrganizationID   string `json:"organization_id" gorm:"org_id"`
	OrganizationName string `json:"organization_name" gorm:"org_name"`
}

//组织信息
type Organization struct {
	gorm.Model
	FounderID         string `json:"founder_id" gorm:"founder_id"`
	OrganizationName  string `json:"organization_name" gorm:"org_name"`
	OrganizationIntro string `json:"intro" gorm:"intro"`
	Avatar            sql.NullString
}

//组织公告分组
type Group struct {
	gorm.Model
	OrganizationID   string `json:"organization_id" gorm:"org_id"`
	OrganizationName string `json:"organization_name" gorm:"org_name"`
	GroupName        string `json:"group_name" gorm:"group_name"`
}

//通告
type Announcement struct {
	gorm.Model
	PublisherID      string `json:"publisher_id" gorm:"publisher_id"`
	OrganizationID   string `json:"organization_id" gorm:"org_id"`
	OrganizationName string `json:"organization_name" gorm:"org_name"`
	GroupID          string `json:"group_id" gorm:"group_id"`
	GroupName        string `json:"group_name" gorm:"group_name"`
	Contents         string `json:"contents" gorm:"contents"`
}
