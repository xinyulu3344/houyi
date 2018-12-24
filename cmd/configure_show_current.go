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
)


// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show-current",
	Short: "显示当前配置组",
	Long: `

			显示当前配置组。
			示例：jdc configure show-current
			`,
	Run: func(cmd *cobra.Command, args []string) {
		/**
		1. 判断配置文件是否存在，如果不存在新建，写入default
		2. 解析文件内容
		3. 打印
		*/

		// 获取current文件内容
		contentCurrent := controller.ReadFile(config.CONFIG_CURRENT_PATH)

		// 获取config内容
		contentConfig := controller.ReadFile(config.CONFIG_PATH)

		// 获取show-current命令要展示的内容
		item := controller.ParseJsonConfig(contentConfig, contentCurrent)

		// 展示内容
		controller.ShowConfig(item, contentCurrent)


	},
}

func init() {
	ConfigureCmd.AddCommand(showCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// showCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// showCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
