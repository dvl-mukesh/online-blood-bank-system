package model

import "encoding/json"

type Response struct {
	Success    bool        `json:"success"`
	SuccessMsg string      `json:"success_msg"`
	Data       interface{} `json:"data,omitempty"`
}

func (r *Response) ToJson() []byte {
	json_data, _ := json.Marshal(r)
	return json_data
}
