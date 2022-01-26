package model

// "errors"
// "github.com/jinzhu/gorm"

//查询用户信息
func GetUserInfo(id string) (User, error) {
	var u User
	//// SELECT * FROM users ORDER BY id LIMIT 1;

	return u, DB.Where("student_id=?", id).First(&u).Error
}

//修改用户姓名
func ChangeName(user User) error {
	return DB.Model(&user).Where("student_id = ?", user.StudentID).Update("nick_name", user.NickName).Error
}

//修改用户头像
func (user *User) UpdateUser(id string) error {
	return DB.Model(User{}).Where("student_id = ?", id).Update("avatar", user.Avatar).Error
}

//查询用户收藏
func GetCollection(id string) ([]Collection, error) {
	var collects []Collection
	return collects, DB.Where("student_id = ?", id).Find(&collects).Error

}

//查询已发布的通告
func GetPublished(id string) ([]Announcement, error) {
	var announce []Announcement
	return announce, DB.Where("publisher_id = ?", id).Find(&announce).Error
}

//查询创建的组织
func GetCreated(id string) ([]Organization, error) {
	var created []Organization
	return created, DB.Where("founder_id = ?", id).Find(&created).Error
}

//查询关注的组织
func GetFollowing(id string) ([]FollowingOrganization, error) {
	var following []FollowingOrganization
	return following, DB.Where("student_id = ?", id).Find(&following).Error
}

//查询所有组织
func GetAllOrganizations(interface{}) ([]Organization, error) {
	var org []Organization
	return org, DB.Find(&org).Error
}

//查询指定组织的全部信息
func GetDetails(ID string, Name string) ([]Organization, error) {
	var org []Organization
	if Name == "" {
		return org, DB.Where("id = ?", ID).First(&org).Error
	} else {
		return org, DB.Where("organization_name = ?", Name).First(&org).Error
	}
}

//查看全部通知
func GetAnnouncements(interface{}) ([]Announcement, error) {
	var announce []Announcement

	return announce, DB.Find(&announce).Error
}

//更新组织信息
func (org *Organization) UpdateOrganization(orgID uint) error {
	return DB.Model(Organization{}).Where("id = ?").Updates(Organization{
		FounderID:         org.FounderID,
		OrganizationName:  org.OrganizationName,
		OrganizationIntro: org.OrganizationIntro,
		Avatar:            org.Avatar,
	}).Error
}

//验证身份是否为组织创建人
func JudgeFounder(PublisherId string, OrganizationId string) bool {
	var org Organization
	DB.First(&org)
	return PublisherId == org.FounderID
}

//验证身份是否为通知发布人
func JudgePublisher(id string, ID string) bool {
	var announcement Announcement
	DB.Where("id = ?", ID).First(&announcement)
	return id == announcement.PublisherID
}

//删除通告的分类
func DeleteAnnoucement(id string) error {
	var announcement Announcement
	return DB.Where("id = ?", id).Delete(&announcement).Error
}

//取消收藏
func CancelCollect(id string) error {
	var collect Collection
	return DB.Where("id = ?", id).Delete(&collect).Error
}
