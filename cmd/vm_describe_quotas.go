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
	"github.com/jdcloud-api/jdcloud-sdk-go/services/common/models"
	"encoding/json"
	"fmt"
)

// describeQuotasCmd represents the describeQuotas command
var describeQuotasCmd = &cobra.Command{
	Use:   "describe-quotas",
	Short: "查询配额，支持：云主机、镜像、密钥、模板、镜像共享",
	Long: `

            查询配额，支持：云主机、镜像、密钥、模板、镜像共享; 。

            示例: jdc vm describe-quotas
`,
	Run: func(cmd *cobra.Command, args []string) {
		regionId, _ := cmd.Flags().GetString(config.REGION_ID)
		imageId, _ := cmd.Flags().GetString(config.IMAGE_ID)
		filters := []models.Filter{}
		//inputJson, _ := cmd.Flags().GetString(config.INPUT_JSON)
		//headers, _ := cmd.Flags().GetString(config.HEADERS)

		// 创建业务client
		vmclient := controller.VmClient()
		// 设置请求参数
		req := apis.NewDescribeQuotasRequestWithAllParams(regionId, filters, &imageId)
		// 执行请求得到响应
		resp, err := vmclient.DescribeQuotas(req)
		if err != nil {
			return
		}
		resultJsonByte, _ := json.MarshalIndent(resp, "", "    ")
		fmt.Println(string(resultJsonByte))
	},
}

func init() {
	vmCmd.AddCommand(describeQuotasCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// describeQuotasCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// describeQuotasCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	describeQuotasCmd.Flags().StringP(config.REGION_ID, "", controller.GetRegionId(), "地域ID")
	describeQuotasCmd.Flags().StringP(config.IMAGE_ID, "", "", "私有镜像Id，查询镜像共享(imageShare)配额时，此参数必传")
	describeQuotasCmd.Flags().StringP(config.INPUT_JSON, "", "", `以json字符串或文件绝对路径形式作为输入参数。
字符串方式举例：--input-json '{"field":"value"}';
文件格式举例：--input-json file:///xxxx.json`)
	describeQuotasCmd.Flags().StringP(config.HEADERS, "", "", `用户自定义Header，举例：'{"x-jdcloud-security-token":"abc","test":"123"}'`)
	describeQuotasCmd.Flags().StringP(config.FILTERS, "", "", `instanceId - 云主机ID，精确匹配，支持多个; privateIpAddress describeInstanceStatusCmd个; az - 可用区，精确匹配，支持多个; vpcId - 私有网络ID，精确匹配，支持多个; status - describeInstanceStatusCmdf="http://docs.jdcloud.com/virtual-machines/api/vm_status">参考云主机状态</a>; name - describeInstanceStatusCmdd - 镜像ID，精确匹配，支持多个; networkInterfaceId - 弹性网卡ID，精确匹配，支持多个; subnetId - describeInstanceStatusCmd用可用组id，支持单个; faultDomain - 错误域，支持多个`)
}
