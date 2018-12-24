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
	"github.com/jdcloud-api/jdcloud-sdk-go/services/common/models"
)

// vmDescribeInstancesCmd represents the vmDescribeInstances command
var describeInstancesCmd = &cobra.Command{
	Use:   "describe-instances",
	Short: "批量查询云主机的详细信息; 此接口支持分页查询，默认每页20条。",
	Long: `

            批量查询云主机的详细信息; 此接口支持分页查询，默认每页20条。

            示例: jdc vm describe-instances
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
		req := apis.NewDescribeInstancesRequestWithAllParams(regionId,&pageNumber,&pageSize,filters)
		// 执行请求得到响应
		resp, err := vmclient.DescribeInstances(req)
		if err != nil {
			return
		}
		resultJsonByte, _ := json.MarshalIndent(resp, "", "    ")
		fmt.Println(string(resultJsonByte))
	},
}

func init() {
	vmCmd.AddCommand(describeInstancesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// vmDescribeInstancesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// vmDescribeInstancesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	describeInstancesCmd.Flags().StringP(config.REGION_ID, "", controller.GetRegionId(), "地域ID")
	describeInstancesCmd.Flags().IntP(config.PAGE_NUMBER, "", 1, "页码；默认为1")
	describeInstancesCmd.Flags().IntP(config.PAGE_SIZE, "", 20, "分页大小；默认为20；取值范围[10, 100]")
	describeInstancesCmd.Flags().StringP(config.FILTERS, "", "", `instanceId - 云主机ID，精确匹配，支持多个; privateIpAddress - 主网卡内网主IP地址，模糊匹配，支持多个; az - 可用区，精确匹配，支持多个; vpcId - 私有网络ID，精确匹配，支持多个; status - 云主机状态，精确匹配，支持多个，<a href="http://docs.jdcloud.com/virtual-machines/api/vm_status">参考云主机状态</a>; name - 云主机名称，模糊匹配，支持单个; imageId - 镜像ID，精确匹配，支持多个; networkInterfaceId - 弹性网卡ID，精确匹配，支持多个; subnetId - 子网ID，精确匹配，支持多个; agId - 使用可用组id，支持单个; faultDomain - 错误域，支持多个`)
	describeInstancesCmd.Flags().StringP(config.INPUT_JSON, "", "", `(json) 以json字符串或文件绝对路径形式作为输入参数。`)
	describeInstancesCmd.Flags().StringP(config.HEADERS, "", "", `用户自定义Header，举例：'{"x-jdcloud-security-token":"abc","test":"123"}'`)
}
