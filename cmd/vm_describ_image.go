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
	"houyi/config"
	"houyi/controller"
	"github.com/jdcloud-api/jdcloud-sdk-go/services/vm/apis"
	"fmt"
	"encoding/json"
)

// describeImageCmd represents the describeImage command
var describeImageCmd = &cobra.Command{
	Use:   "describe-image",
	Short: "查询镜像详情",
	Long: `

            查询镜像详情。

            示例: jdc vm describe-image  --image-id xxx
`,
	Run: func(cmd *cobra.Command, args []string) {
		regionId, _ := cmd.Flags().GetString(config.REGION_ID)
		imageId, _ := cmd.Flags().GetString(config.IMAGE_ID)
		// 创建业务client
		vmclient := controller.VmClient()
		// 设置请求参数
		req := apis.NewDescribeImageRequest(regionId, imageId)
		// 执行请求得到响应
		resp, err := vmclient.DescribeImage(req)
		if err != nil {
			return
		}
		resultJsonByte, _ := json.MarshalIndent(resp, "", "    ")
		fmt.Println(string(resultJsonByte))

	},
}

func init() {
	vmCmd.AddCommand(describeImageCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// describeImageCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// describeImageCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	describeImageCmd.Flags().StringP(config.REGION_ID, "", controller.GetRegionId(), "地域ID")
	describeImageCmd.Flags().StringP(config.IMAGE_ID, "", "", "镜像ID")
	describeImageCmd.Flags().StringP(config.INPUT_JSON, "", "", `以json字符串或文件绝对路径形式作为输入参数。
字符串方式举例：--input-json '{"field":"value"}';
文件格式举例：--input-json file:///xxxx.json`)
	describeImageCmd.Flags().StringP(config.HEADERS, "", "", `用户自定义Header，举例：'{"x-jdcloud-security-token":"abc","test":"123"}'`)
}
