package request

import uuid "github.com/satori/go.uuid"

// User register structure
// type Register struct {
// 	Username    string `json:"userName"`
// 	Password    string `json:"passWord"`
// 	NickName    string `json:"nickName"`
// 	HeaderImg   string `json:"headerImg"`
// 	AuthorityId uint   `json:"authorityId" gorm:"default:2"`
// }
type Register struct {
	Username string `json:"userName"`
	Password string `json:"passWord"`
}

// User login structure
type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Modify password structure
type ChangePasswordStruct struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	NewPassword string `json:"newPassword"`
}

// Modify  user's auth structure
type SetUserAuth struct {
	UUID        uuid.UUID `json:"uuid"`
	AuthorityId string    `json:"authorityId"`
}
