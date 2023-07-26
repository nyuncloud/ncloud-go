package client

import (
	"fmt"
	"github.com/nyuncloud/ncloud-go/core"
	"github.com/nyuncloud/ncloud-go/services/edge/ip/apis"
	"testing"
)

const (
	accessKey = "JgQxbyg0g1BwKyHp"
	secretKey = "mT6XJEgsNqTs6N0GKmqOIStqlZKMGXpK"
)

func TestIPClient_QueryIPGet(t *testing.T) {
	credentials := core.NewCredentials(accessKey, secretKey)
	pluginClient := NewIPClient(credentials)
	dataRequest := apis.NewQueryIPGetRequest()
	//pluginClient.DisableLogger() //关闭日志输出
	dataResponse, err := pluginClient.QueryIPGet(dataRequest)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("dataResponse:", dataResponse)
}

func TestIPClient_GetWhiteList(t *testing.T) {
	credentials := core.NewCredentials(accessKey, secretKey)
	pluginClient := NewIPClient(credentials)
	dataRequest := apis.NewQueryWhiteListRequest()
	//pluginClient.DisableLogger() //关闭日志输出
	dataResponse, err := pluginClient.GetWhiteList(dataRequest)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("dataResponse:", dataResponse)
}

func TestIPClient_AddWhiteList(t *testing.T) {
	credentials := core.NewCredentials(accessKey, secretKey)
	pluginClient := NewIPClient(credentials)
	dataRequest := apis.NewWhiteAddRequest("223.74.107.179")
	//pluginClient.DisableLogger() //关闭日志输出
	dataResponse, err := pluginClient.AddWhiteList(dataRequest)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("dataResponse:", dataResponse)
}

func TestIPClient_DelWhiteList(t *testing.T) {
	credentials := core.NewCredentials(accessKey, secretKey)
	pluginClient := NewIPClient(credentials)
	dataRequest := apis.NewWhiteDelRequest("223.74.107.179")
	//pluginClient.DisableLogger() //关闭日志输出
	dataResponse, err := pluginClient.DelWhiteList(dataRequest)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("dataResponse:", dataResponse)
}
