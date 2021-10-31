package models

type RegisterUserReq struct {
	UserName string `json:"user_name,omitempty"`
	FullName string `json:"full_name,omitempty"`
	Password string `json:"password,omitempty"`
}

type LoginUserReq struct {
	UserName string `json:"user_name,omitempty"`
	Password string `json:"password,omitempty"`
}

type LoginUserResp struct {
	Token string `json:"token,omitempty"`
}
