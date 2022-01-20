package model

type User struct {
	UserID      string `json:"user_id"`
	NickName    string `json:"nickname"`
	HeadPotrait string `json:"headpotrait"`
}

type Collection struct {
}
