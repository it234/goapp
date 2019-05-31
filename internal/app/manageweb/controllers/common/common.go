package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	SUCCESS_CODE          = 20000      //成功的状态码
	FAIL_CODE             = 30000      //失败的状态码
	MD5_PREFIX            = "jkfldfsf" //MD5加密前缀字符串
	TOKEN_KEY             = "X-Token"  //页面token键名
	USER_ID_Key           = "X-USERID" //页面用户ID键名
	USER_UUID_Key         = "X-UUID"   //页面UUID键名
	SUPER_ADMIN_ID uint64 = 956986 // 超级管理员账号ID
)

type ResponseModel struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ResponseModelBase struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// 响应成功
func ResSuccess(c *gin.Context, v interface{}) {
	ret := ResponseModel{Code: SUCCESS_CODE, Message: "ok", Data: v}
	ResJSON(c, http.StatusOK, &ret)
}

// 响应成功
func ResSuccessMsg(c *gin.Context) {
	ret := ResponseModelBase{Code: SUCCESS_CODE, Message: "ok"}
	ResJSON(c, http.StatusOK, &ret)
}

// 响应失败
func ResFail(c *gin.Context, msg string) {
	ret := ResponseModelBase{Code: FAIL_CODE, Message: msg}
	ResJSON(c, http.StatusOK, &ret)
}

// 响应失败
func ResFailCode(c *gin.Context, msg string, code int) {
	ret := ResponseModelBase{Code: code, Message: msg}
	ResJSON(c, http.StatusOK, &ret)
}

// 响应JSON数据
func ResJSON(c *gin.Context, status int, v interface{}) {
	c.JSON(status, v)
	c.Abort()
}

// 响应错误-服务端故障
func ResErrSrv(c *gin.Context, err error) {
	ret := ResponseModelBase{Code: FAIL_CODE, Message: "服务端故障"}
	ResJSON(c, http.StatusOK, &ret)
}

// 响应错误-用户端故障
func ResErrCli(c *gin.Context, err error) {
	ret := ResponseModelBase{Code: FAIL_CODE, Message: "err"}
	ResJSON(c, http.StatusOK, &ret)
}

type ResponsePageData struct {
	Total uint64      `json:"total"`
	Items interface{} `json:"items"`
}

type ResponsePage struct {
	Code    int              `json:"code"`
	Message string           `json:"message"`
	Data    ResponsePageData `json:"data"`
}

// 响应成功-分页数据
func ResSuccessPage(c *gin.Context, total uint64, list interface{}) {
	ret := ResponsePage{Code: SUCCESS_CODE, Message: "ok", Data: ResponsePageData{Total: total, Items: list}}
	ResJSON(c, http.StatusOK, &ret)
}

// 获取页码
func GetPageIndex(c *gin.Context) uint64 {
	return GetQueryToUint64(c, "page", 1)
}

// 获取每页记录数
func GetPageLimit(c *gin.Context) uint64 {
	limit := GetQueryToUint64(c, "limit", 20)
	if limit > 500 {
		limit = 20
	}
	return limit
}

// 获取排序信息
func GetPageSort(c *gin.Context) string {
	return GetQueryToStr(c, "sort")
}

// 获取搜索关键词信息
func GetPageKey(c *gin.Context) string {
	return GetQueryToStr(c, "key")
}
