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
	"fmt"
	"houyi/controller"
	"encoding/json"
	"github.com/jdcloud-api/jdcloud-sdk-go/services/vm/apis"
)

// describeImageConstraintsCmd represents the describeImageConstraints command
var describeImageConstraintsCmd = &cobra.Command{
	Use:   "describe-image-constraints",
	Short: "查询镜像的实例规格限制; 通过此接口可以查看镜像不支持的实例规格。只有官方镜像、第三方镜像有实例规格的限制，个人的私有镜像没有此限制。",
	Long: `

            查询镜像的实例规格限制。<br>; 通过此接口可以查看镜像不支持的实例规格。只有官方镜像、第三方镜像有实例规格的限制，个人的私有镜像没有此限制。; 。

            示例: jdc vm describe-image-constraints  --image-id xxx
`,
	Run: func(cmd *cobra.Command, args []string) {
		regionId, _ := cmd.Flags().GetString(config.REGION_ID)
		imageId, _ := cmd.Flags().GetString(config.IMAGE_ID)
		//inputJson, _ := cmd.Flags().GetString(config.INPUT_JSON)
		//headers, _ := cmd.Flags().GetString(config.HEADERS)

		// 创建业务client
		vmclient := controller.VmClient()
		// 设置请求参数
		req := apis.NewDescribeImageConstraintsRequestWithAllParams(regionId,imageId)
		// 执行请求得到响应
		resp, err := vmclient.DescribeImageConstraints(req)
		if err != nil {
			return
		}
		resultJsonByte, _ := json.MarshalIndent(resp, "", "    ")
		fmt.Println(string(resultJsonByte))
	},
}

func init() {
	vmCmd.AddCommand(describeImageConstraintsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// describeImageConstraintsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// describeImageConstraintsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	describeImageConstraintsCmd.Flags().StringP(config.REGION_ID, "", controller.GetRegionId(), "地域ID")
	describeImageConstraintsCmd.Flags().StringP(config.IMAGE_ID, "", "", "镜像ID")
	describeImageConstraintsCmd.Flags().StringP(config.INPUT_JSON, "", "", `以json字符串或文件绝对路径形式作为输入参数。
字符串方式举例：--input-json '{"field":"value"}';
文件格式举例：--input-json file:///xxxx.json`)
	describeImageConstraintsCmd.Flags().StringP(config.HEADERS, "", "", `用户自定义Header，举例：'{"x-jdcloud-security-token":"abc","test":"123"}'`)
}
