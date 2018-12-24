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
	"encoding/json"
	"fmt"
	"github.com/jdcloud-api/jdcloud-sdk-go/services/vm/apis"
)

// describeImageConstraintsBatchCmd represents the describeImageConstraintsBatch command
var describeImageConstraintsBatchCmd = &cobra.Command{
	Use:   "describe-image-constraints-batch",
	Short: "批量查询镜像的实例规格限制; 通过此接口可以查看镜像不支持的实例规格。只有官方镜像、第三方镜像有实例规格的限制，个人的私有镜像没有此限制。",
	Long: `

            批量查询镜像的实例规格限制。<br>; 通过此接口可以查看镜像不支持的实例规格。只有官方镜像、第三方镜像有实例规格的限制，个人的私有镜像没有此限制。; 。

            示例: jdc vm describe-image-constraints-batch 
`,
	Run: func(cmd *cobra.Command, args []string) {
		regionId, _ := cmd.Flags().GetString(config.REGION_ID)
		ids, _ := cmd.Flags().GetString(config.IDS)
		//inputJson, _ := cmd.Flags().GetString(config.INSTANCE_ID)
		//headers, _ := cmd.Flags().GetString(config.HEADERS)

		idsArr := controller.Str2Arr(ids)
		fmt.Println(idsArr)
		// 获取业务client
		vmclient := controller.VmClient()
		// 设置请求参数
		req := apis.NewDescribeImageConstraintsBatchRequestWithAllParams(regionId, idsArr)
		// 执行请求得到响应
		resp, err := vmclient.DescribeImageConstraintsBatch(req)
		if err != nil {
			return
		}
		resultJsonByte, _ := json.MarshalIndent(resp, "", "    ")
		fmt.Println(string(resultJsonByte))
	},
}

func init() {
	vmCmd.AddCommand(describeImageConstraintsBatchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// describeImageConstraintsBatchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// describeImageConstraintsBatchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	describeImageConstraintsBatchCmd.Flags().StringP(config.REGION_ID, "", controller.GetRegionId(), "地域ID")
	describeImageConstraintsBatchCmd.Flags().StringP(config.IDS,"","","镜像ID列表")
	describeImageConstraintsBatchCmd.Flags().StringP(config.INPUT_JSON, "", "", `(json) 以json字符串或文件绝对路径形式作为输入参数。`)
	describeImageConstraintsBatchCmd.Flags().StringP(config.HEADERS, "", "", `用户自定义Header，举例：'{"x-jdcloud-security-token":"abc","test":"123"}'`)
}
