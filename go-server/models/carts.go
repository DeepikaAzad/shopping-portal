package models

type AddItemToCartReq struct {
	ItemID   uint   `json:"item_id"`
	ItemName string `json:"item_name"`
}

type AddItemToCartResp struct {
	Message string `json:"message"`
}

type PlaceOrderReq struct {
	CartID uint `json:"cart_id"`
	UserID uint `json:"user_id"`
}

type PlaceOrderResp struct {
	CartID  uint   `json:"cart_id"`
	Message string `json:"message"`
}

type CartListReq struct {
	UserID uint `json:"user_id"`
}

type CartListResp struct {
	CartID uint   `json:"cart_id"`
	Items  []Item `json:"item_list"`
}
