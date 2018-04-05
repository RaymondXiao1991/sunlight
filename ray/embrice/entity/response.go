package entity

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Response 返回响应结构体
type Response struct {
	Code    int64       `json:"code"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message"`
}

// HanderSuccess 接口请求成功处理
func (rsp *Response) HanderSuccess(w http.ResponseWriter, obj interface{}) bool {
	var data []byte
	var err error
	rsp.Code = 0
	rsp.Message = "success"
	rsp.Data = obj

	if data, err = json.Marshal(&rsp); err != nil {
		fmt.Println("json marshal failed:", err)
		return false
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write([]byte(data))
	return true
}

// HandlerFailed 接口请求失败处理
func (rsp *Response) HandlerFailed(w http.ResponseWriter, message string) {
	var (
		data []byte
		err  error
	)
	rsp.Code = 1
	rsp.Message = message
	if data, err = json.Marshal(&rsp); err != nil {
		fmt.Println("json marshal failed:", err)
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write([]byte(data))
}
