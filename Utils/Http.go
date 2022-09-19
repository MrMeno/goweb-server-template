package Utils

import "reflect"

const (
	SuccessCode      = 0
	FailedCode       = -1
	CommonFailMsg    = "调用失败"
	CommonSuccessMsg = "调用成功"
)

type ResponseInfo struct {
	Code int         `json:"code"`
	Msg  string      `json:"mode"`
	Data interface{} `json:"data"`
}

func GetReflectFields(T interface{}) []string {
	reflectFields := reflect.TypeOf(T)
	mapFields := []string{}
	for i := 0; i < reflectFields.NumField(); i++ {
		mapFields = append(mapFields, reflectFields.Field(i).Name)
	}
	return mapFields
}
