package kapp

import (
	"reflect"
	"fmt"
)

// Returns 所有请求都返回这个类型，其中 data 是返回的实际内容
type Returns struct {
	// 应答码
	Code    string `json:"code"`

	// 错误消息，如果请求成功，不返回这个字段
	Message string `json:"message,omitempty"`

	// 错误提示，用于客户端提醒用户
	Alert   string `json:"alert,omitempty"`

	// 分页时返回总条数
	Total   int64 `json:"total,omitempty"`

	// 返回数据，如果请求错误，不返回这个字段
	Data    interface{} `json:"data,omitempty"`
}

func (r *Returns) Error() string {
	return fmt.Sprintf(`code: %s, message: %s`, r.Code, r.Message)
}

func (r *Returns) String() string {
	return r.Error()
}


// NewReturns 创建 Returns
func NewReturns(code, message string, data interface{}) *Returns {
	return &Returns{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

// NewErrReturns 创建错误 Returns
func NewErrReturns(code, message string) *Returns {
	return &Returns{
		Code:    code,
		Message: message,
		Alert:   STATUS.Msg(code),
	}
}

func NewDataReturns(data interface{}) *Returns {
	return &Returns{Code: STATUS.Ok, Data: data}
}

func NewListReturns(data interface{}) *Returns {

	if data == nil {
		return &Returns{Code: STATUS.Ok, Data: []interface{}{}}
	}
	t := reflect.TypeOf(data)
	switch t.Kind() {
	case reflect.Array, reflect.Slice:
		// 如果 data 是 array/slice, 判断长度是否为 0
		// 返回 `{"code":"00000", "data": []}`
		if reflect.ValueOf(data).Len() == 0 {
			return &Returns{Code: STATUS.Ok, Data: []interface{}{}}
		}
	default:
	}
	return &Returns{Code: STATUS.Ok, Data: data}
}

func NewTotalListReturns(total int64, data interface{}) *Returns {
	if data == nil {
		return &Returns{Code: STATUS.Ok, Data: []interface{}{}}
	}
	t := reflect.TypeOf(data)
	switch t.Kind() {
	case reflect.Array, reflect.Slice:
		// 如果 data 是 array/slice, 判断长度是否为 0
		// 返回 `{"code":"00000", "data": []}`
		if reflect.ValueOf(data).Len() == 0 {
			return &Returns{Code: STATUS.Ok, Data: []interface{}{}}
		}
	default:
	}
	return &Returns{Code: STATUS.Ok, Total: total, Data: data}
}
