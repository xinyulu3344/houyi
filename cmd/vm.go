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
	"fmt"

	"github.com/spf13/cobra"
)


// vmCmd represents the vm command
var vmCmd = &cobra.Command{
	Use:   "vm",
	Short: "云主机",
	Long: `
              {associate-elastic-ip,attach-disk,attach-network-interface,copy-images,create-image,create-instances,create-keypair,delete-image,delete-instance,delete-keypair,describe-image,describe-image-constraints,describe-image-constraints-batch,describe-image-members,describe-images,describe-instance,describe-instance-private-ip-address,describe-instance-status,describe-instance-types,describe-instance-vnc-url,describe-instances,describe-keypairs,describe-quotas,detach-disk,detach-network-interface,disassociate-elastic-ip,generate-skeleton,import-keypair,modify-image-attribute,modify-instance-attribute,modify-instance-disk-attribute,modify-instance-network-attribute,modify-instance-password,reboot-instance,rebuild-instance,resize-instance,share-image,start-instance,stop-instance,un-share-image}
              ...

        vm cli 子命令，云主机实例、镜像、实例规格、实例模板、配额相关的接口。
        OpenAPI文档地址为：https://www.jdcloud.com/help/detail/376/isCatalog/0
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("vm called")
	},
}

func init() {
	RootCmd.AddCommand(vmCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// vmCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// vmCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
