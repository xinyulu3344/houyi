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
	"github.com/jdcloud-api/jdcloud-sdk-go/services/vm/apis"
	"github.com/jdcloud-api/jdcloud-sdk-go/services/common/models"
	"encoding/json"
)

// describeInstanceTypesCmd represents the describeInstanceTypes command
var describeInstanceTypesCmd = &cobra.Command{
	Use:   "describe-instance-types",
	Short: "查询实例规格信息列表",
	Long: `

            查询实例规格信息列表; 。

            示例: jdc vm describe-instance-types
`,
	Run: func(cmd *cobra.Command, args []string) {
		regionId, _ := cmd.Flags().GetString(config.REGION_ID)
		//filters, _ := cmd.Flags().GetString(config.FILTERS)
		filters := []models.Filter{}
		//inputJson, _ := cmd.Flags().GetString(config.INPUT_JSON)
		//headers, _ := cmd.Flags().GetString(config.HEADERS)

		// 获取业务client
		vmclient := controller.VmClient()
		// 设置请求参数
		req := apis.NewDescribeInstanceTypesRequestWithAllParams(regionId, filters)
		// 执行请求获得响应
		resp, err := vmclient.DescribeInstanceTypes(req)
		if err != nil {
			return
		}
		resultJsonByte, _ := json.MarshalIndent(resp, "", "    ")
		fmt.Println(string(resultJsonByte))
	},
}

func init() {
	vmCmd.AddCommand(describeInstanceTypesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// describeInstanceTypesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// describeInstanceTypesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	describeInstanceTypesCmd.Flags().StringP(config.REGION_ID, "", controller.GetRegionId(), "地域ID")
	describeInstanceTypesCmd.Flags().StringP(config.FILTERS, "", "", `instanceId - 云主机ID，精确匹配，支持多个; privateIpAddress describeInstanceStatusCmd个; az - 可用区，精确匹配，支持多个; vpcId - 私有网络ID，精确匹配，支持多个; status - describeInstanceStatusCmdf="http://docs.jdcloud.com/virtual-machines/api/vm_status">参考云主机状态</a>; name - describeInstanceStatusCmdd - 镜像ID，精确匹配，支持多个; networkInterfaceId - 弹性网卡ID，精确匹配，支持多个; subnetId - describeInstanceStatusCmd用可用组id，支持单个; faultDomain - 错误域，支持多个`)
	describeInstanceTypesCmd.Flags().StringP(config.INPUT_JSON, "", "", `(json) 以json字符串或文件绝对路径形式作为输入参数。`)
	describeInstanceTypesCmd.Flags().StringP(config.HEADERS, "", "", `用户自定义Header，举例：'{"x-jdcloud-security-token":"abc","test":"123"}'`)
}
