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
	"encoding/json"
)

// describeImagesCmd represents the describeImages command
var describeImagesCmd = &cobra.Command{
	Use:   "describe-images",
	Short: "查询镜像信息列表; 通过此接口可以查询到京东云官方镜像、第三方镜像、私有镜像、或其他用户共享给您的镜像; 此接口支持分页查询，默认每页20条。",
	Long: `
查询镜像信息列表; 通过此接口可以查询到京东云官方镜像、第三方镜像、私有镜像、或其他用户共享给您的镜像; 此接口支持分页查询，默认每页20条。
`,
	Run: func(cmd *cobra.Command, args []string) {
		regionId, _ := cmd.Flags().GetString(config.REGION_ID)
		pageNumber, _ := cmd.Flags().GetInt(config.PAGE_NUMBER)
		pageSize, _ := cmd.Flags().GetInt(config.PAGE_SIZE)
		//filters, _ := cmd.Flags().GetString(config.FILTERS)
		//filters := []models.Filter{}
		//inputJson, _ := cmd.Flags().GetString(config.INPUT_JSON)
		//headers, _ := cmd.Flags().GetString(config.HEADERS)
		imageSource, _ := cmd.Flags().GetString(config.IMAGE_SOURCE)
		platform, _ := cmd.Flags().GetString(config.PLATFORM)
		rootDeviceType, _ := cmd.Flags().GetString(config.ROOT_DEVICE_TYPE)
		status, _ := cmd.Flags().GetString(config.STATUS)
		ids, _ := cmd.Flags().GetString(config.IDS)
		// 将ids参数转化成字符串数组

		//idsArr := controller.Str2Arr(ids)
		idsArr := controller.Str2Arr(ids)
		// 获取业务client
		vmclient := controller.VmClient()
		// 设置请求参数
		req := apis.NewDescribeImagesRequestWithAllParams(regionId,&imageSource,&platform,idsArr,&rootDeviceType,&status,&pageNumber,&pageSize)
		// 执行请求得到响应
		resp, err := vmclient.DescribeImages(req)
		if err != nil {
			return
		}
		resultJsonByte, _ := json.MarshalIndent(resp, "", "    ")
		fmt.Println(string(resultJsonByte))
	},
}

func init() {
	vmCmd.AddCommand(describeImagesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// describeImagesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// describeImagesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	describeImagesCmd.Flags().StringP(config.REGION_ID, "", controller.GetRegionId(), "地域ID")
	describeImagesCmd.Flags().IntP(config.PAGE_NUMBER, "", 1, "页码；默认为1")
	describeImagesCmd.Flags().IntP(config.PAGE_SIZE, "", 20, "分页大小；默认为20；取值范围[10, 100]")
	describeImagesCmd.Flags().StringP(config.INPUT_JSON, "", "", `(json) 以json字符串或文件绝对路径形式作为输入参数。`)
	describeImagesCmd.Flags().StringP(config.HEADERS, "", "", `用户自定义Header，举例：'{"x-jdcloud-security-token":"abc","test":"123"}'`)
	describeImagesCmd.Flags().StringP(config.IMAGE_SOURCE,"", "","镜像来源，如果没有指定ids参数，此参数必传；取值范围：public、shared、thirdparty、private")
	describeImagesCmd.Flags().StringP(config.PLATFORM, "", "", "操作系统平台，取值范围：Windows Server、CentOS、Ubuntu")
	describeImagesCmd.Flags().StringP(config.ROOT_DEVICE_TYPE,"","","镜像支持的系统盘类型，[localDisk,cloudDisk]")
	describeImagesCmd.Flags().StringP(config.STATUS, "","",`<a href="http://docs.jdcloud.com/virtual-machines/api/image_status">参考镜像状态</a> `)
	describeImagesCmd.Flags().StringP(config.IDS,"","","镜像ID列表，如果指定了此参数，其它参数可为空")
}
