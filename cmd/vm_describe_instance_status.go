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
	"github.com/jdcloud-api/jdcloud-sdk-go/services/common/models"
	"github.com/jdcloud-api/jdcloud-sdk-go/services/vm/apis"
	"encoding/json"
)

// describeInstanceStatusCmd represents the describeInstanceStatus command
var describeInstanceStatusCmd = &cobra.Command{
	Use:   "describe-instance-status",
	Short: "批量查询云主机状态",
	Long: `

            批量查询云主机状态。

            示例: jdc vm describe-instance-status 
`,
	Run: func(cmd *cobra.Command, args []string) {
		regionId, _ := cmd.Flags().GetString(config.REGION_ID)
		pageNumber, _ := cmd.Flags().GetInt(config.PAGE_NUMBER)
		pageSize, _ := cmd.Flags().GetInt(config.PAGE_SIZE)
		//filters, _ := cmd.Flags().GetString(config.FILTERS)
		filters := []models.Filter{}
		//inputJson, _ := cmd.Flags().GetString(config.INPUT_JSON)
		//headers, _ := cmd.Flags().GetString(config.HEADERS)

		// 获取业务client
		vmclient := controller.VmClient()
		// 设置请求参数
		req := apis.NewDescribeInstanceStatusRequestWithAllParams(regionId, &pageNumber, &pageSize,filters)
		// 执行请求获得响应
		resp, err := vmclient.DescribeInstanceStatus(req)
		if err != nil {
			return
		}
		resultJsonByte, _ := json.MarshalIndent(resp, "", "    ")
		fmt.Println(string(resultJsonByte))
	},
}

func init() {
	vmCmd.AddCommand(describeInstanceStatusCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// describeInstanceStatusCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// describeInstanceStatusCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	describeInstanceStatusCmd.Flags().StringP(config.REGION_ID, "", controller.GetRegionId(), "地域ID")
	describeInstanceStatusCmd.Flags().IntP(config.PAGE_NUMBER, "", 1, "页码；默认为1")
	describeInstanceStatusCmd.Flags().IntP(config.PAGE_SIZE, "", 20, "分页大小；默认为20；取值范围[10, 100]")
	describeInstanceStatusCmd.Flags().StringP(config.FILTERS, "", "", `instanceId - 云主机ID，精确匹配，支持多个; privateIpAddress describeInstanceStatusCmd个; az - 可用区，精确匹配，支持多个; vpcId - 私有网络ID，精确匹配，支持多个; status - describeInstanceStatusCmdf="http://docs.jdcloud.com/virtual-machines/api/vm_status">参考云主机状态</a>; name - describeInstanceStatusCmdd - 镜像ID，精确匹配，支持多个; networkInterfaceId - 弹性网卡ID，精确匹配，支持多个; subnetId - describeInstanceStatusCmd用可用组id，支持单个; faultDomain - 错误域，支持多个`)
	describeInstanceStatusCmd.Flags().StringP(config.INPUT_JSON, "", "", `(json) 以json字符串或文件绝对路径形式作为输入参数。`)
	describeInstanceStatusCmd.Flags().StringP(config.HEADERS, "", "", `用户自定义Header，举例：'{"x-jdcloud-security-token":"abc","test":"123"}'`)
}
