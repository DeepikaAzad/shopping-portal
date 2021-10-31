package models

type AddItemToCartReq struct {
	ItemID   uint   `json:"item_id,omitempty"`
	ItemName string `json:"item_name,omitempty"`
}

type AddItemToCartResp struct {
	CartID  uint   `json:"cart_id,omitempty"`
	ItemsID string `json:"items_id,omitempty"`
	Message string `json:"message,omitempty"`
}

type PlaceOrderReq struct {
	CartID uint `json:"cart_id,omitempty"`
	UserID uint `json:"user_id,omitempty"`
}

type PlaceOrderResp struct {
	CartID  uint   `json:"cart_id,omitempty"`
	Message string `json:"message,omitempty"`
}
