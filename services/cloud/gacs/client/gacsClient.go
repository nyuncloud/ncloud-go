package client

import (
	"errors"
	"github.com/bytedance/sonic"
	"github.com/nyuncloud/ncloud-go/core"
	"github.com/nyuncloud/ncloud-go/core/net/consts"
	"github.com/nyuncloud/ncloud-go/services/cloud/gacs/apis"
)

type GACSClient struct {
	core.NCloudClient
}

func NewGACSClient(credential *core.Credential) *GACSClient {
	if credential == nil {
		return nil
	}
	config := core.NewConfig()
	config.SetScheme(consts.SchemeHttp)
	config.SetEndpoint("114.67.174.156:20001")
	return &GACSClient{
		core.NCloudClient{
			Credential:  *credential,
			Config:      *config,
			ServiceName: "cloudlist",
			Revision:    "v2",
			Logger:      core.NewDefaultLogger(core.LogInfo),
		}}
}

func (c *GACSClient) DisableLogger() {
	c.Logger = core.NewDummyLogger()
}

func (c *GACSClient) QueryHostList(request *apis.QueryHostListRequest) (*apis.QueryHostListResponse, error) {
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
	jsonResp := &apis.QueryHostListResponse{}
	err = sonic.Unmarshal(resp, jsonResp)
	if err != nil {
		c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
		return nil, err
	}
	return jsonResp, err
}
func (c *GACSClient) QueryHostAll(request *apis.QueryHostAllRequest) (*apis.QueryHostAllResponse, error) {
	resp, err := c.Send(request.Method, request.Path, "")
	if err != nil {
		return nil, err
	}
	if len(resp) == 0 {
		return nil, nil
	}
	jsonResp := &apis.QueryHostAllResponse{}
	err = sonic.Unmarshal(resp, jsonResp)
	if err != nil {
		c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
		return nil, err
	}
	return jsonResp, err
}

func (c *GACSClient) StartHost(request *apis.HostStartRequest) (*apis.QueryResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	jsonDataReq, err := sonic.Marshal(request.HostInfo)
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
	jsonResp := &apis.QueryResponse{}
	err = sonic.Unmarshal(resp, jsonResp)
	if err != nil {
		c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
		return nil, err
	}
	return jsonResp, err
}

func (c *GACSClient) ReStartHost(request *apis.HostRestartRequest) (*apis.QueryResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	jsonDataReq, err := sonic.Marshal(request.HostInfo)
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
	jsonResp := &apis.QueryResponse{}
	err = sonic.Unmarshal(resp, jsonResp)
	if err != nil {
		c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
		return nil, err
	}
	return jsonResp, err
}
func (c *GACSClient) StopHost(request *apis.HostStopRequest) (*apis.QueryResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	jsonDataReq, err := sonic.Marshal(request.HostInfo)
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
	jsonResp := &apis.QueryResponse{}
	err = sonic.Unmarshal(resp, jsonResp)
	if err != nil {
		c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
		return nil, err
	}
	return jsonResp, err
}

func (c *GACSClient) UpdateHostPassword(request *apis.HostUpdatePasswordRequest) (*apis.QueryResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	jsonDataReq, err := sonic.Marshal(request.HostInfoWithPassword)
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
	jsonResp := &apis.QueryResponse{}
	err = sonic.Unmarshal(resp, jsonResp)
	if err != nil {
		c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
		return nil, err
	}
	return jsonResp, err
}

func (c *GACSClient) UpdateHostSystem(request *apis.HostUpdateSystemRequest) (*apis.QueryResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	jsonDataReq, err := sonic.Marshal(request.HostInfoWithImage)
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
	jsonResp := &apis.QueryResponse{}
	err = sonic.Unmarshal(resp, jsonResp)
	if err != nil {
		c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
		return nil, err
	}
	return jsonResp, err
}

func (c *GACSClient) FindHostImage(request *apis.HostFindImageRequest) (*apis.QueryResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	jsonDataReq, err := sonic.Marshal(request.HostInfo)
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
	jsonResp := &apis.QueryResponse{}
	err = sonic.Unmarshal(resp, jsonResp)
	if err != nil {
		c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
		return nil, err
	}
	return jsonResp, err
}
