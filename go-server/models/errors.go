package models

import "fmt"

type ErrorResp struct {
	Message string `json:"message"`
}

type SZLError struct {
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Type    string `json:"type,omitempty"`
	Param   string `json:"param,omitempty"`
	Errors  error  `json:"error,omitempty"`
}

func (r *SZLError) Error() string {
	return fmt.Sprintf("%v", r.Message)
}
