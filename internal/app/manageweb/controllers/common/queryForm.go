package common

import (
	"errors"

	"github.com/it234/goapp/pkg/convert"

	"github.com/gin-gonic/gin"
)

// GetQueryToStrE
func GetQueryToStrE(c *gin.Context,key string) (string,error) {
	str,ok:=c.GetQuery(key)
	if !ok {
		return "",errors.New("没有这个值传入")
	}
	return str,nil
}

// GetQueryToStr
func GetQueryToStr(c *gin.Context,key string,defaultValues ...string) string {
	var defaultValue string
	if len(defaultValues)>0{
		defaultValue=defaultValues[0]
	}
	str,err:=GetQueryToStrE(c,key)
	if str=="" || err!=nil{
		return defaultValue
	}
	return str
}

// QueryToUintE
func GetQueryToUintE(c *gin.Context,key string) (uint,error) {
	str,err:=GetQueryToStrE(c,key)
	if err !=nil {
		return 0,err
	}
	return convert.ToUintE(str)
}

// QueryToUint
func GetQueryToUint(c *gin.Context,key string,defaultValues ...uint) uint {
	var defaultValue uint
	if len(defaultValues)>0{
		defaultValue=defaultValues[0]
	}
	val,err:=GetQueryToUintE(c,key)
	if err!=nil {
		return defaultValue
	}
	return val
}

// QueryToUintE
func GetQueryToUint64E(c *gin.Context,key string) (uint64,error) {
	str,err:=GetQueryToStrE(c,key)
	if err !=nil {
		return 0,err
	}
	return convert.ToUint64E(str)
}

// QueryToUint
func GetQueryToUint64(c *gin.Context,key string,defaultValues ...uint64) uint64 {
	var defaultValue uint64
	if len(defaultValues)>0{
		defaultValue=defaultValues[0]
	}
	val,err:=GetQueryToUint64E(c,key)
	if err!=nil {
		return defaultValue
	}
	return val
}