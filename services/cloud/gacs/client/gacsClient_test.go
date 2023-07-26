package client

import (
	"fmt"
	"github.com/nyuncloud/ncloud-go/core"
	"github.com/nyuncloud/ncloud-go/services/cloud/gacs/apis"
	"testing"
)

const (
	accessKey = "JgQxbyg0g1BwKyHp"
	secretKey = "mT6XJEgsNqTs6N0GKmqOIStqlZKMGXpK"
)

func TestGACSClient_QueryHostList(t *testing.T) {
	credentials := core.NewCredentials(accessKey, secretKey)
	gacsClient := NewGACSClient(credentials)
	dataRequest := apis.NewQueryHostListRequestWithParams(1, 10)
	//pluginClient.DisableLogger() //关闭日志输出
	dataResponse, err := gacsClient.QueryHostList(dataRequest)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("dataResponse:", dataResponse)
}

func TestGACSClient_QueryHostAll(t *testing.T) {
	credentials := core.NewCredentials(accessKey, secretKey)
	gacsClient := NewGACSClient(credentials)
	dataRequest := apis.NewQueryHostAll()
	//pluginClient.DisableLogger() //关闭日志输出
	dataResponse, err := gacsClient.QueryHostAll(dataRequest)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("dataResponse:", dataResponse)
}

func TestGACSClient_StartHost(t *testing.T) {
	credentials := core.NewCredentials(accessKey, secretKey)
	gacsClient := NewGACSClient(credentials)
	dataRequest := apis.NewHostStartRequestWithParams("", "", "")
	//pluginClient.DisableLogger() //关闭日志输出
	dataResponse, err := gacsClient.StartHost(dataRequest)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("dataResponse:", dataResponse)
}

func TestGACSClient_ReStartHost(t *testing.T) {
	credentials := core.NewCredentials(accessKey, secretKey)
	gacsClient := NewGACSClient(credentials)
	dataRequest := apis.NewHostRestartRequestWithParams("", "", "")
	//pluginClient.DisableLogger() //关闭日志输出
	dataResponse, err := gacsClient.ReStartHost(dataRequest)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("dataResponse:", dataResponse)
}

func TestGACSClient_StopHost(t *testing.T) {
	credentials := core.NewCredentials(accessKey, secretKey)
	gacsClient := NewGACSClient(credentials)
	dataRequest := apis.NewHostStopRequestWithParams("", "", "")
	//pluginClient.DisableLogger() //关闭日志输出
	dataResponse, err := gacsClient.StopHost(dataRequest)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("dataResponse:", dataResponse)
}

func TestGACSClient_UpdateHostPassword(t *testing.T) {
	credentials := core.NewCredentials(accessKey, secretKey)
	gacsClient := NewGACSClient(credentials)
	dataRequest := apis.NewHostUpdatePasswordRequestWithParams("", "", "", "")
	//pluginClient.DisableLogger() //关闭日志输出
	dataResponse, err := gacsClient.UpdateHostPassword(dataRequest)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("dataResponse:", dataResponse)
}

func TestGACSClient_UpdateHostSystem(t *testing.T) {
	credentials := core.NewCredentials(accessKey, secretKey)
	gacsClient := NewGACSClient(credentials)
	dataRequest := apis.NewHostUpdateSystemRequestWithParams("", "", "", "", "")
	//pluginClient.DisableLogger() //关闭日志输出
	dataResponse, err := gacsClient.UpdateHostSystem(dataRequest)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("dataResponse:", dataResponse)
}

func TestGACSClient_FindHostImage(t *testing.T) {
	credentials := core.NewCredentials(accessKey, secretKey)
	gacsClient := NewGACSClient(credentials)
	dataRequest := apis.NewHostFindImageRequestWithParams("", "", "")
	//pluginClient.DisableLogger() //关闭日志输出
	dataResponse, err := gacsClient.FindHostImage(dataRequest)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("dataResponse:", dataResponse)
}
