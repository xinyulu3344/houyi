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

package main

import (
	"houyi/controller"
	"houyi/config"
	"houyi/cmd"
)

func main() {
	cmd.Execute()
}

func init() {
	// 判断config文件是否存在
	if !controller.CheckFileIsExist(config.CONFIG_PATH) {
		controller.WriteFile(config.CONFIG_PATH, "{}")
	}
	// 判断current文件是否存在
	if !controller.CheckFileIsExist(config.CONFIG_CURRENT_PATH) {
		controller.WriteFile(config.CONFIG_CURRENT_PATH, config.CONFIG_CURRENT_DEFAULT)
	}
}