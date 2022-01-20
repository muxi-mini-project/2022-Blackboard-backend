package model

//用户信息
type User struct {
	ID          string `json:"id"`
	StudentID   string `json:"student_id"`
	PassWord    string `json:"password"`
	NickName    string `json:"nickname"`
	HeadPotrait string `json:"headpotrait"`
}

//用户收藏
type Collection struct {
	ID             string `json:"id"`
	StudentID      string `json:"student_id"`
	AnnouncementID string `json:"announcement_id"`
	Announcement   string `json:"announcement"`
}

//用户创建的组织
type OrganizationCreated struct {
	ID               string `json:"id"`
	StudentID        string `json:"student_id"`
	OrganizationID   string `json:"organization_id"`
	OrganizationName string `json:"organization_name"`
}

//用户关注的组织
type OrganizationFollowing struct {
	ID               string `json:"id"`
	StudentID        string `json:"student_id"`
	OrganizationID   string `json:"organization_id"`
	OrganizationName string `json:"organization_name"`
}

//组织信息
type Organization struct {
	OrganizationID    string `json:"organization_id"`
	OrganizationLogo  string `json:"organization_logo"`
	FounderID         string `json:"founder_id"`
	OrganizationName  string `json:"organization_name"`
	OrganizationIntro string `json;"intro"`
}

//组织公告分组
type Group struct {
	GroupID          string `json:"group_id"`
	OrganizationID   string `json:"organization_id"`
	OrganizationName string `json:"organization_name"`
	GroupName        string `json:"group_name"`
}

//通告
type Announcement struct {
	ID               string `json:"id"`
	OrganizationID   string `json:"organization_id"`
	OrganizationName string `json:"organization_name"`
	GroupID          string `json:"group_id"`
	GroupName        string `json:"group_name"`
	Contents         string `json:"contents"`
}
