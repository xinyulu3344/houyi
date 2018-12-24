// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"github.com/spf13/cobra"
	"houyi/controller"
	"houyi/config"
)



// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "配置CLI运行环境，包括 access key, secret key 和区域等信息",
	Long: `

        配置CLI运行环境，包括 access key, secret key 和区域等信息。新增后即设置为当前配置。
        其中 access key 和 secret key 为必选参数。
        示例：jdc configure add --profile test --access-key xxx --secret-key xxx
        
        同时支持密码输入方式。
        示例：jdc configure add --profile test[Enter]
        Please input your access-key:
        Please input your secret-key:
`,
	Run: func(cmd *cobra.Command, args []string) {
		ak, _ := cmd.Flags().GetString(config.ACCESS_KEY)
		sk, _ := cmd.Flags().GetString(config.SECRET_KEY)
		rid, _ := cmd.Flags().GetString(config.REGION_ID)
		ep, _ := cmd.Flags().GetString(config.ENDPOINT)
		s, _ := cmd.Flags().GetString(config.SCHEME)
		t, _ := cmd.Flags().GetInt(config.TIMEOUT)
		p, _ := cmd.Flags().GetString(config.PROFILE)
		// 获取转换成map的json
		jsonMap := controller.ParseJsonMap(controller.ReadFile(config.CONFIG_PATH))
		// 向map中增加新项
		jsonMap = controller.AddJsonItem(jsonMap, ak, sk, rid, ep, s, t, p)
		// 将map转为json
		jsonStr := controller.Map2Json(jsonMap)
		// 将json写入config配置文件
		controller.WriteFile(config.CONFIG_PATH, jsonStr)
	},
}

func init() {
	ConfigureCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	addCmd.Flags().StringP(config.ACCESS_KEY, "", "", "access key")
	addCmd.Flags().StringP(config.SECRET_KEY, "", "", "secret key")
	addCmd.Flags().StringP(config.REGION_ID, "", "cn-north-1", "region id")
	addCmd.Flags().StringP(config.ENDPOINT, "", "www.jdcloud-api.com", "api gateway endpoint")
	addCmd.Flags().StringP(config.SCHEME, "", "https", "http scheme")
	addCmd.Flags().IntP(config.TIMEOUT, "", 20, "request timeout")
	addCmd.Flags().StringP(config.PROFILE, "", "", "request timeout")
}
