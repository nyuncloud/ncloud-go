package client

import (
	"fmt"
	"github.com/nyuncloud/ncloud-go/core"
	"github.com/nyuncloud/ncloud-go/services/edge/plugin/apis"
	"testing"
)

const (
	accessKey = "JgQxbyg0g1BwKyHp"
	secretKey = "mT6XJEgsNqTs6N0GKmqOIStqlZKMGXpK"
)

func TestPluginClient_QueryEdgeInfo(t *testing.T) {
	credentials := core.NewCredentials(accessKey, secretKey)
	pluginClient := NewPluginClient(credentials)
	dataRequest := apis.NewQueryEdgeInfoRequestWithParams(1, 10)
	//pluginClient.DisableLogger() //关闭日志输出
	dataResponse, err := pluginClient.QueryEdgeInfo(dataRequest)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("dataResponse:", dataResponse)
}

func TestPluginClient_QueryEdgeInfoRecord(t *testing.T) {
	credentials := core.NewCredentials(accessKey, secretKey)
	pluginClient := NewPluginClient(credentials)
	dataRequest := apis.NewQueryEdgeInfoRecordRequestAllParams("NC_000C292F729F", 1687258800, 3600*2)
	dataResponse, err := pluginClient.QueryEdgeInfoRecord(dataRequest)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("dataResponse:", dataResponse)

}

func TestPluginClient_QueryEdgeFlowRecord(t *testing.T) {
	credentials := core.NewCredentials(accessKey, secretKey)
	pluginClient := NewPluginClient(credentials)
	dataRequest := apis.NewQueryEdgeFlowRequest()
	dataRequest.SetReferList([]string{"NC"})
	dataRequest.SetStartTimestamp(1687231440)
	dataRequest.SetDuration(300)
	dataResponse, err := pluginClient.QueryEdgeFlowRecord(dataRequest)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("dataResponse:", dataResponse)
}
