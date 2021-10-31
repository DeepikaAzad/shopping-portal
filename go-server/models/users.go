package models

type RegisterUserReq struct {
	UserName string `json:"user_name"`
	FullName string `json:"full_name"`
	Password string `json:"password"`
}

type RegisterUserResp struct {
	Message string `json:"message"`
}

type LoginUserReq struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

type LoginUserResp struct {
	Token string `json:"token"`
}

type UserList struct {
	User []User `json:"users"`
}

type User struct {
	ID       uint   `json:"id"`
	UserName string `json:"user_name"`
	FullName string `json:"full_name"`
}
