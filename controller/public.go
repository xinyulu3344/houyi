package controller

import (
	"houyi/config"
)


// 获取AccessKey
func GetAccessKey() string {
	contentCurrent := ReadFile(config.CONFIG_CURRENT_PATH)
	contentConfig := ReadFile(config.CONFIG_PATH)
	item := ParseJsonConfig(contentConfig, contentCurrent)
	return item.AccessKey
}

func GetSecretKey() string {
	contentCurrent := ReadFile(config.CONFIG_CURRENT_PATH)
	contentConfig := ReadFile(config.CONFIG_PATH)
	item := ParseJsonConfig(contentConfig, contentCurrent)
	return item.SecretKey
}

