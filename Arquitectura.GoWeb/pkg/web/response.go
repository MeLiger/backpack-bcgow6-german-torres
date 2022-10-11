package web

import  (
 		"encoding/json"
		"strconv"
)

{
	"code": 401,
	"error": "Token inválido"
}

{
	"code": 200,
	"data": {...}
}

type Response struct {
	Code string `json:"code"`
	Data interface{} `json:"data,omitempty"`
	Error string 		`json:"error,omitempty"`
}

func NewResponse(code int, data interface{}, err string) Response {
	if code < 300{
		return Response{strconv.FormatInt(int64(code),10), data, ""}
	}
}