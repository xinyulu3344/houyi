package controller

import (
	"os"
	"fmt"
	"io/ioutil"
	"io"
	"github.com/bitly/go-simplejson"
	"encoding/json"
	"sort"
	"houyi/config"
)

type Config struct {
	AccessKey string
	SecretKey string
	RegionId  string
	Endpoint  string
	Scheme    string
	Timeout   json.Number
}

// 判断文件是否存在
func CheckFileIsExist(filename string) bool {
	exist := true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

// 获取文件所有内容
func ReadFile(filename string) string {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0766)
	if err != nil {
		fmt.Println("ReadFile os.OpenFile err: ", err)
	}
	content, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("ioutil.ReadAll err: ", err)
	}
	return string(content)
}

// 文件不存在，新建文件，并向文件中写入传入的内容
func WriteFile(filename string, content string) {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0766)
	if err != nil {
		fmt.Println("os.OpenFile err: ", err)
		fmt.Println("正在创建: ", filename)
	}
	fmt.Println("正在写入: \n", content)
	io.WriteString(file, content)
	fmt.Println("文件写入成功")
}

// 文件不存在，新建文件
func CreateFile(filename string) {
	_, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0766)
	if err != nil {
		fmt.Println("os.OpenFile err: ", err)
		fmt.Println("正在创建: ", filename)
	}
}

// 解析json
func ParseJsonConfig(jsonStr string, k string) Config {

	var c Config
	jsonObj, err := simplejson.NewJson([]byte(jsonStr))
	if err != nil {
		fmt.Println("simplejson.NewJson err: ", err)
		return c
	}
	jsonAll, err := jsonObj.Map()
	if err != nil {
		fmt.Println("jsonObj.Map() err: ", err)
	}
	for key, _ := range jsonAll {
		if key == k {
			jsonItem, err := jsonObj.Get(key).Map()
			if err != nil {
				fmt.Println("err")
			}
			for key, value := range jsonItem {
				switch key {
				case config.ACCESS_KEY:
					c.AccessKey = value.(string)
				case config.SECRET_KEY:
					c.SecretKey = value.(string)
				case config.REGION_ID:
					c.RegionId = value.(string)
				case config.ENDPOINT:
					c.Endpoint = value.(string)
				case config.SCHEME:
					c.Scheme = value.(string)
				case config.TIMEOUT:
					c.Timeout = value.(json.Number)
				}
			}
		}
	}
	return c
}

// 解析json，返回map
func ParseJsonMap(jsonStr string) (jsonAll map[string]interface{}) {
	jsonObj, err := simplejson.NewJson([]byte(jsonStr))
	jsonAll = make(map[string]interface{})
	if err != nil {
		fmt.Println("simplejson.NewJson err: ", err)
		return
	}
	jsonAll, err = jsonObj.Map()
	if err != nil {
		fmt.Println("jsonObj.Map() err: ", err)
	}
	return
}

// map转json
func Map2Json(m map[string]interface{}) string {
	jsonByte, err := json.MarshalIndent(m,"","    ")
	if err != nil {
		fmt.Println("json.Marshal err: ", err)
	}
	return string(jsonByte)
}

// 向map中增加新项
func AddJsonItem(jsonAll map[string]interface{}, ak string, sk string, rid string, ep string, s string, t int, p string) map[string]interface{}{
	//fmt.Println(jsonAll, ak, sk, rid, ep, s, t, p)
	//for k, v := range jsonAll{
	//	fmt.Printf("%s, %T, %v\n", k, v, v)
	//}
	m := make(map[string]interface{})
	m[config.ACCESS_KEY] = ak
	m[config.SECRET_KEY] = sk
	m[config.REGION_ID] = rid
	m[config.ENDPOINT] = ep
	m[config.SCHEME] = s
	m[config.TIMEOUT] = t
	jsonAll[p] = m
	return jsonAll
}

// 获取配置文件中Json的key
func GetJsonKey(jsonStr string) (keys []string) {
	jsonObj, err := simplejson.NewJson([]byte(jsonStr))
	if err != nil {
		fmt.Println("simplejson.NewJson err: ", err)
		fmt.Println("请使用add命令添加认证信息")
		return
	}
	jsonAll, err := jsonObj.Map()
	if err != nil {
		fmt.Println("jsonObj.Map() err: ", err)
	}
	for key, _ := range jsonAll {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return
}

// 展示内容
func ShowConfig(c Config, contentCurrent string) {
	fmt.Printf("================= %s =================\n", contentCurrent)
	fmt.Println("access_key:  	", c.AccessKey)
	fmt.Println("secret_key:  	", c.SecretKey)
	fmt.Println("region_id:  	", c.RegionId)
	fmt.Println("endpoint:  	", c.Endpoint)
	fmt.Println("scheme:  	", c.Scheme)
	fmt.Println("timeout:  	", c.Timeout)
}

// 获取regionId
func GetRegionId() (regionId string){

	return ParseJsonConfig(ReadFile(config.CONFIG_PATH),ReadFile(config.CONFIG_CURRENT_PATH)).RegionId
}
