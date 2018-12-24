package controller

import (
	"github.com/jdcloud-api/jdcloud-sdk-go/core"
	"github.com/jdcloud-api/jdcloud-sdk-go/services/vm/client"
	"regexp"
	"fmt"
	"strings"
)

// 返回VM client
func VmClient () *client.VmClient{
	ak := GetAccessKey()
	sk := GetSecretKey()
	credentials := core.NewCredentials(ak, sk)
	vmclient := client.NewVmClient(credentials)
	return vmclient
}

// 字符串转换成字符串数组
func Str2Arr(s string) (strArr []string){
	// 判断字符串是否被[]包裹
	// 去掉首尾字符
	// 逗号分隔成字符串数组
	// 返回数组
	// [aaa, bbb,... ]
	if ok, _ := regexp.MatchString("^[[].*[]]$", s); ok {
		s = strings.TrimLeft(s, "[")
		s = strings.TrimRight(s, "]")
	}else {
		fmt.Println("Parameter ids is invalid!")
	}
	strArr = strings.Split(s, ",")
	return
}
