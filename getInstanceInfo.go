package main

import (
	"github.com/jdcloud-api/jdcloud-sdk-go/core"
	vmcli "github.com/jdcloud-api/jdcloud-sdk-go/services/vm/client"
	vmapis "github.com/jdcloud-api/jdcloud-sdk-go/services/vm/apis"
	"fmt"
)

func vmInit(accessKey string, secretKey string) *vmcli.VmClient {
	credentials := core.NewCredentials(accessKey, secretKey)
	vmClient := vmcli.NewVmClient(credentials)
	return vmClient
}

func getInstanceInfo(vmclient *vmcli.VmClient, region string) {
	req := vmapis.NewDescribeInstancesRequest(region)
	resp, err := vmclient.DescribeInstances(req)
	if err != nil {
		return
	}
	for _, v := range resp.Result.Instances {
		//fmt.Println(v)
		fmt.Println("主机Id: ", v.InstanceId)
		fmt.Println("云主机状态: ", v.Status)
		fmt.Println("云主机名称: ", v.InstanceName)
		fmt.Println("实例规格: ", v.InstanceType)
		fmt.Println("主网卡IP地址: ", v.PrivateIpAddress)
		fmt.Println("主网卡主IP绑定弹性IP的地址: ", v.ElasticIpAddress)
		fmt.Println("云主机描述: ", v.Description)
		fmt.Println("镜像ID: ", v.ImageId)
		fmt.Println("云主机创建时间: ", v.LaunchTime)
		fmt.Println("云主机所在可用区: ", v.Az)
		// 计费信息
		fmt.Println("支付模式: ", v.Charge.ChargeMode)
		fmt.Println("费用支付状态: ", v.Charge.ChargeStatus)
		fmt.Println("计费开始时间: ", v.Charge.ChargeStartTime)
		fmt.Println("过期时间: ", v.Charge.ChargeExpiredTime)
		fmt.Println("预期释放时间: ", v.Charge.ChargeRetireTime)
		// 磁盘相关
		fmt.Println("数据盘挂载状态: ", v.SystemDisk.Status)
		fmt.Println("随云主机一起删除: ", v.SystemDisk.AutoDelete)
		fmt.Println("数据盘逻辑挂载点: ", v.SystemDisk.DeviceName)
		fmt.Println("本地磁盘类型: ", v.SystemDisk.LocalDisk.DiskType)
		fmt.Println("本地磁盘大小: ", v.SystemDisk.LocalDisk.DiskSizeGB)
		fmt.Println("密钥对名称: ", v.KeyNames)
		// 高可用组
		fmt.Println("高可用组Id: ", v.Ag.Id)
		fmt.Println("高可用组名称: ", v.Ag.Name)
		fmt.Println("高可用组错误区域: ", v.FaultDomain)
		fmt.Println("\n\n")
	}
}

func getInstanceIpInfo(vmclient *vmcli.VmClient, region string) {
	req := vmapis.NewDescribeInstancePrivateIpAddressRequest(region)
	resp, err := vmclient.DescribeInstancePrivateIpAddress(req)
	if err != nil {
		return
	}
	for _, v := range resp.Result.InstancePrivateIpAddress {
		fmt.Println(v)
	}
}


func main1() {
	accessKey := "EE35CF02BF28533628A64594A8641E04"
	secretKey := "8E110D81823C9C5F1A773FCC54415DC8"
	vmClient := vmInit(accessKey, secretKey)
	getInstanceInfo(vmClient, "cn-north-1")
	//getInstanceIpInfo(vmClient, "cn-north-1")
}
