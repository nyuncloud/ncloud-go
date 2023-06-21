package client

import (
	"errors"
	"github.com/bytedance/sonic"
	"github.com/nyuncloud/ncloud-go/core"
	"github.com/nyuncloud/ncloud-go/core/net/consts"
	"github.com/nyuncloud/ncloud-go/services/edge/plugin/apis"
)

type PluginClient struct {
	core.NCloudClient
}

func NewPluginClient(credential *core.Credential) *PluginClient {
	if credential == nil {
		return nil
	}
	config := core.NewConfig()
	config.SetScheme(consts.SchemeHttp)
	config.SetEndpoint("36.139.119.111:20005")
	return &PluginClient{
		core.NCloudClient{
			Credential:  *credential,
			Config:      *config,
			ServiceName: "plugin",
			Revision:    core.Version,
			Logger:      core.NewDefaultLogger(core.LogInfo),
		}}
}
func (c *PluginClient) DisableLogger() {
	c.Logger = core.NewDummyLogger()
}

func (c *PluginClient) QueryEdgeInfo(request *apis.QueryEdgeInfoRequest) (*apis.QueryEdgeInfoResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	jsonDataReq, err := sonic.Marshal(request.PageInfo)
	if err != nil {
		return nil, err
	}
	resp, err := c.Send(request.Method, request.Path, string(jsonDataReq))
	if err != nil {
		return nil, err
	}
	if len(resp) == 0 {
		return nil, nil
	}
	jsonResp := &apis.QueryEdgeInfoResponse{}
	err = sonic.Unmarshal(resp, jsonResp)
	if err != nil {
		c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
		return nil, err
	}
	return jsonResp, err
}

func (c *PluginClient) QueryEdgeInfoRecord(request *apis.QueryEdgeInfoRecordRequest) (*apis.EdgeInfoRecordResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	jsonDataReq, err := sonic.Marshal(request.QueryEdgeInfoConditions)
	if err != nil {
		return nil, err
	}
	resp, err := c.Send(request.Method, request.Path, string(jsonDataReq))

	if err != nil {
		return nil, err
	}
	if len(resp) == 0 {
		return nil, nil
	}
	jsonResp := &apis.EdgeInfoRecordResponse{}
	err = sonic.Unmarshal(resp, jsonResp)
	if err != nil {
		c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
		return nil, err
	}
	return jsonResp, err
}
func (c *PluginClient) QueryEdgeFlowRecord(request *apis.QueryEdgeFlowRequest) (*apis.QueryEdgeFlowResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	jsonDataReq, err := sonic.Marshal(request.QueryEdgeFlowConditions)
	if err != nil {
		return nil, err
	}
	resp, err := c.Send(request.Method, request.Path, string(jsonDataReq))

	if err != nil {
		return nil, err
	}
	if len(resp) == 0 {
		return nil, nil
	}
	jsonResp := &apis.QueryEdgeFlowResponse{}
	err = sonic.Unmarshal(resp, jsonResp)
	if err != nil {
		c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
		return nil, err
	}
	return jsonResp, err
}
