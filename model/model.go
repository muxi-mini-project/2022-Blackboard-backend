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
	return DB.Model(&user).Where("student_id = ?", user.StudentID).Update(user.NickName).Error
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
func GetDetails(ID string) (Organization, error) {
	var org Organization
	return org, DB.Where("organization_id = ?", ID).First(&org).Error
}
