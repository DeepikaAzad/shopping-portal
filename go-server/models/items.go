package models

type AddItemReq struct {
	Name string `json:"name"`
}

type AddItemResp struct {
	Message string `json:"message"`
}

type ItemList struct {
	Item []Item `json:"items"`
}

type Item struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
