package aepapi

import (
	aepsdkcore "github.com/sacker1225/ctwing-aep-sdk/core"
	"net/http"
)

// 参数MasterKey: 类型String, 参数不可以为空
//
//	描述:MasterKey在该设备所属产品的概况中可以查看
//
// 参数body: 类型json, 参数不可以为空
//
//	描述:body,具体参考平台api说明
func CreateCommand(appKey string, appSecret string, MasterKey string, body string) (*http.Response, error) {
	path := "/aep_device_command/command"

	var headers map[string]string = make(map[string]string)
	headers["MasterKey"] = MasterKey

	var param map[string]string = nil
	version := "20190712225145"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, body, version, application, key, "POST")
}

// 参数MasterKey: 类型String, 参数不可以为空
//
//	描述:MasterKey在该设备所属产品的概况中可以查看
//
// 参数commandId: 类型String, 参数不可以为空
//
//	描述:创建指令成功响应中返回的id，
//
// 参数productId: 类型long, 参数不可以为空
//
//	描述:
//
// 参数deviceId: 类型String, 参数不可以为空
//
//	描述:设备ID
func QueryCommand(appKey string, appSecret string, MasterKey string, commandId string, productId string, deviceId string) (*http.Response, error) {
	path := "/aep_device_command/command"

	var headers map[string]string = make(map[string]string)
	headers["MasterKey"] = MasterKey

	var param map[string]string = make(map[string]string)
	param["commandId"] = commandId
	param["productId"] = productId
	param["deviceId"] = deviceId

	version := "20190712225241"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, "", version, application, key, "GET")
}

// 参数MasterKey: 类型String, 参数不可以为空
//
//	描述:
//
// 参数body: 类型json, 参数不可以为空
//
//	描述:body,具体参考平台api说明
func CancelCommand(appKey string, appSecret string, MasterKey string, body string) (*http.Response, error) {
	path := "/aep_device_command/cancelCommand"

	var headers map[string]string = make(map[string]string)
	headers["MasterKey"] = MasterKey

	var param map[string]string = nil
	version := "20190615023142"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, body, version, application, key, "PUT")
}

// 参数MasterKey: 类型String, 参数不可以为空
//
//	描述:产品MasterKey，在该设备所属产品的概况中可以查看
//
// 参数productId: 类型long, 参数不可以为空
//
//	描述:产品ID
//
// 参数deviceId: 类型String, 参数不可以为空
//
//	描述:设备ID
//
// 参数startTime: 类型long, 参数不可以为空
//
//	描述:起始时间戳，精确到毫秒, 例如：1539565440000
//
// 参数endTime: 类型long, 参数不可以为空
//
//	描述:结束时间戳，精确到毫秒, 例如：1539565460000
//
// 参数pageTime: 类型long, 参数可以为空
//
//	描述:分页时间戳，精确到毫秒, 例如：1539565460000；查询第一页时pageTime可不传 或 填endTime，第二页开始需使用返回值中的pageTime，若返回值中pageTime为空，则说明无下一页数据
//
// 参数pageSize: 类型long, 参数可以为空
//
//	描述:每页记录数
func QueryCommandDetailList(appKey string, appSecret string, MasterKey string, productId string, deviceId string, startTime string, endTime string, pageTime string, pageSize string) (*http.Response, error) {
	path := "/aep_device_command/commandDetailList"

	var headers map[string]string = make(map[string]string)
	headers["MasterKey"] = MasterKey

	var param map[string]string = make(map[string]string)
	param["productId"] = productId
	param["deviceId"] = deviceId
	param["startTime"] = startTime
	param["endTime"] = endTime
	param["pageTime"] = pageTime
	param["pageSize"] = pageSize

	version := "20241218142112"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, "", version, application, key, "GET")
}
