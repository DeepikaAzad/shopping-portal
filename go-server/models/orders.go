package models

type OrderListReq struct {
	UsersID uint `json:"users_id"`
}

type OrderListResp struct {
	Order []Order `json:"orders"`
}

type Order struct {
	CartID    uint       `json:"cart_id"`
	ItemNames []ItemName `json:"items"`
}

type ItemName struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
