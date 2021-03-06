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

// useCmd represents the use command
var useCmd = &cobra.Command{
	Use:   "use",
	Short: "切换到某个配置组",
	Long: `

        切换到某个配置组。
        示例：jdc configure use --profile default
`,
	Run: func(cmd *cobra.Command, args []string) {
		p, _ := cmd.Flags().GetString(config.PROFILE)
		// 读取配置文件内容
		contentConfig := controller.ReadFile(config.CONFIG_PATH)
		// 获取配置文件中json的key，返回key数组
		keys := controller.GetJsonKey(contentConfig)
		// 判断命令行传入的值是否存在于配置文件
		for _, key := range keys {
			if p == key {
				controller.WriteFile(config.CONFIG_CURRENT_PATH, p)
				return
			}
		}
		fmt.Printf("Profile %s do not exist! Configure failed!", p)
	},
}

func init() {
	ConfigureCmd.AddCommand(useCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// useCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// useCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	useCmd.Flags().StringP(config.PROFILE, "", "default", "PROFILE  配置组名称")
}
