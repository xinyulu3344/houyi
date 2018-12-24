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
	"encoding/json"
	"fmt"
)

// describeInstanceVncUrlCmd represents the describeInstanceVncUrl command
var describeInstanceVncUrlCmd = &cobra.Command{
	Use:   "describe-instance-vnc-url",
	Short: "获取云主机vnc，用于连接管理云主机; vnc地址的有效期为1个小时，调用接口获取vnc地址后如果1个小时内没有使用，vnc地址自动失效，再次使用需要重新获取。",
	Long: `
usage: jdc vm describe-instance-vnc-url [-h] [--region-id REGIONID] --instance-id INSTANCEID [--input-json INPUT_JSON] [--headers HEADERS]

            获取云主机vnc，用于连接管理云主机; vnc地址的有效期为1个小时，调用接口获取vnc地址后如果1个小时内没有使用，vnc地址自动失效，再次使用需要重新获取。

            示例: jdc vm describe-instance-vnc-url  --instance-id xxx
`,
	Run: func(cmd *cobra.Command, args []string) {
		regionId, _ := cmd.Flags().GetString(config.REGION_ID)
		instanceId, _ := cmd.Flags().GetString(config.INSTANCE_ID)
		//inputJson, _ := cmd.Flags().GetString(config.INPUT_JSON)
		//headers, _ := cmd.Flags().GetString(config.HEADERS)

		// 创建业务client
		vmclient := controller.VmClient()
		// 设置请求参数
		req := apis.NewDescribeInstanceVncUrlRequestWithAllParams(regionId,instanceId)
		// 执行请求得到响应
		resp, err := vmclient.DescribeInstanceVncUrl(req)
		if err != nil {
			return
		}
		resultJsonByte, _ := json.MarshalIndent(resp, "", "    ")
		fmt.Println(string(resultJsonByte))
	},
}

func init() {
	vmCmd.AddCommand(describeInstanceVncUrlCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// describeInstanceVncUrlCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// describeInstanceVncUrlCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	describeInstanceVncUrlCmd.Flags().StringP(config.REGION_ID, "", controller.GetRegionId(), "地域ID")
	describeInstanceVncUrlCmd.Flags().StringP(config.INSTANCE_ID, "", "", "云主机ID")
	describeInstanceVncUrlCmd.Flags().StringP(config.INPUT_JSON, "", "", `以json字符串或文件绝对路径形式作为输入参数。
字符串方式举例：--input-json '{"field":"value"}';
文件格式举例：--input-json file:///xxxx.json`)
	describeInstanceVncUrlCmd.Flags().StringP(config.HEADERS, "", "", `用户自定义Header，举例：'{"x-jdcloud-security-token":"abc","test":"123"}'`)
}
