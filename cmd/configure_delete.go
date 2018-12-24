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
	"fmt"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "删除某个配置组",
	Long: `

        删除某个配置组。
        示例：jdc configure delete --profile test
`,
	Run: func(cmd *cobra.Command, args []string) {
		p, _ := cmd.Flags().GetString(config.PROFILE)
		// 读取配置文件内容
		contentConfig := controller.ReadFile(config.CONFIG_PATH)

		// 将Json转换成map
		jsonMap := controller.ParseJsonMap(contentConfig)

		// 判断待删除的profile是否在配置文件中
		_, exists := jsonMap[p]
		if exists {
			delete(jsonMap, p)

			// 将map转为json
			jsonStr := controller.Map2Json(jsonMap)

			// 将json写入config配置文件
			controller.WriteFile(config.CONFIG_PATH, jsonStr)
		}else {
			fmt.Printf("No profile %s", p)
		}

	},
}

func init() {
	ConfigureCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	deleteCmd.Flags().StringP(config.PROFILE, "", "", "PROFILE  配置组名称")
}
