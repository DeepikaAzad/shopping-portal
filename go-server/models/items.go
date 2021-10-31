package models

type AddItemReq struct {
	Name string `json:"name,omitempty"`
}

type AddItemResp struct {
	Message string `json:"message,omitempty"`
}
