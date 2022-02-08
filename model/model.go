package model

import (
	"errors"
)

// "errors"
// "github.com/jinzhu/gorm"

//查询用户信息
func GetUserInfo(id string) (User, error) {
	var u User
	//// SELECT * FROM users ORDER BY id LIMIT 1;
	return u, DB.Where("student_id=?", id).First(&u).Error
}

//修改用户姓名
func ChangeName(info Info) error {
	var user User
	return DB.Model(&user).Where("student_id = ?", info.StudentID).Update("nick_name", info.NickName).Error
}

//修改用户头像
func UpdateUser(id string, url string, sha string, path string) error {
	return DB.Model(User{}).Where("student_id = ?", id).Updates(User{Avatar: url, Sha: sha, Path: path}).Error
}

//查询用户收藏
func GetCollection(id string) ([]Collection, error) {
	collects := []Collection{}
	return collects, DB.Where("student_id = ?", id).Find(&collects).Error

}

//查询已发布的通告
func GetPublished(id string) ([]Announcement, error) {
	announce := []Announcement{}
	return announce, DB.Where("publisher_id = ?", id).Find(&announce).Error
}

//查询创建的组织
func GetCreated(id string) ([]Organization, error) {
	created := []Organization{}
	return created, DB.Where("founder_id = ?", id).Find(&created).Error
}

//查询关注的组织
func GetFollowing(id string) ([]FollowingOrganization, error) {
	following := []FollowingOrganization{}
	return following, DB.Where("student_id = ?", id).Find(&following).Error
}

//关注新的组织
func Follow(follow FollowingOrganization) error {
	result := DB.Where("organization_name = ?", follow.OrganizationName).Find(&follow)
	if result.RowsAffected >= 1 {
		return errors.New("已经关注")
	} else {
		return DB.Create(&follow).Error
	}
}

//查询所有组织
func GetAllOrganizations(interface{}) ([]Organization, error) {
	org := []Organization{}
	return org, DB.Find(&org).Error
}

//查询组织ID
func GetOrgID(name string) uint {
	org := Organization{}
	DB.Where("organization_name = ?", name).First(&org)
	return org.ID
}

//查询指定组织的全部信息
func GetDetails(ID string, Name string) (Organization, error) {
	org := Organization{}
	if Name == "" {
		return org, DB.Where("id = ?", ID).First(&org).Error
	} else if ID == "" {
		return org, DB.Where("organization_name = ?", Name).First(&org).Error
	}
	return org, DB.Where("id = ? and organization_name = ?", ID, Name).First(&org).Error
}

//查看全部通知
func GetAnnouncements(interface{}) ([]Announcement, error) {
	announce := []Announcement{}

	return announce, DB.Find(&announce).Error
}

//查询特定通知
func CheckAnnouce(ID uint) string {
	announce := Announcement{}
	DB.Where("id = ?", ID).First(&announce)
	return announce.Contents
}

//创建分组
func CreateGroup(group Grouping) error {
	result := DB.Create(&group)
	return result.Error
}

//获得group id
func GetGroupID(name string, orgID uint) uint {
	group := Grouping{}
	DB.Where("organization_id = ? and group_name = ?", orgID, name).First(&group)
	return group.ID
}

//更新组织信息
func UpdateOrganization(org Organization) error {
	return DB.Model(Organization{}).Where("id = ?", org.ID).Updates(Organization{
		FounderID:         org.FounderID,
		OrganizationName:  org.OrganizationName,
		OrganizationIntro: org.OrganizationIntro,
		Avatar:            org.Avatar,
		Sha:               org.Sha,
		Path:              org.Path,
	}).Error
}

//验证身份是否为组织创建人
func JudgeFounder(PublisherId string, OrganizationId uint) bool {
	org := Organization{}
	DB.First(&org)
	return PublisherId == org.FounderID
}

//验证身份是否为通知发布人
func JudgePublisher(id string, ID string) bool {
	announcement := Announcement{}
	DB.Where("id = ?", ID).First(&announcement)
	return id == announcement.PublisherID
}

//删除通告的分类
func DeleteAnnoucement(id string) error {
	announcement := Announcement{}
	return DB.Where("id = ?", id).Delete(&announcement).Error
}

//取消收藏
func CancelCollect(id string) error {
	collect := Collection{}
	return DB.Where("id = ?", id).Delete(&collect).Error
}
