package client

import (
	"errors"
	"github.com/bytedance/sonic"
	"github.com/nyuncloud/ncloud-go/core"
	"github.com/nyuncloud/ncloud-go/core/net/consts"
	"github.com/nyuncloud/ncloud-go/services/edge/ip/apis"
)

type IPClient struct {
	core.NCloudClient
}

func NewIPClient(credential *core.Credential) *IPClient {
	if credential == nil {
		return nil
	}
	config := core.NewConfig()
	config.SetScheme(consts.SchemeHttp)
	config.SetEndpoint("114.67.174.156:60001")
	return &IPClient{
		core.NCloudClient{
			Credential:  *credential,
			Config:      *config,
			ServiceName: "public",
			Revision:    core.Version,
			Logger:      core.NewDefaultLogger(core.LogInfo),
		}}
}
func (c *IPClient) DisableLogger() {
	c.Logger = core.NewDummyLogger()
}

func (c *IPClient) QueryIPGet(request *apis.QueryIPGetRequest) (*apis.IPGetResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	resp, err := c.Send(request.Method, request.Path, "")
	if err != nil {
		return nil, err
	}
	if len(resp) == 0 {
		return nil, nil
	}
	jsonResp := &apis.IPGetResponse{}
	err = sonic.Unmarshal(resp, jsonResp)
	if err != nil {
		c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
		return nil, err
	}
	return jsonResp, err
}

func (c *IPClient) GetWhiteList(request *apis.QueryWhiteListRequest) (*apis.WhiteListResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	resp, err := c.Send(request.Method, request.Path, "")
	if err != nil {
		return nil, err
	}
	if len(resp) == 0 {
		return nil, nil
	}
	jsonResp := &apis.WhiteListResponse{}
	err = sonic.Unmarshal(resp, jsonResp)
	if err != nil {
		c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
		return nil, err
	}
	return jsonResp, err
}

func (c *IPClient) AddWhiteList(request *apis.WhiteAddRequest) (*apis.Response, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	jsonDataReq, err := sonic.Marshal(request.IPInfo)
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
	jsonResp := &apis.Response{}
	err = sonic.Unmarshal(resp, jsonResp)
	if err != nil {
		c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
		return nil, err
	}
	return jsonResp, err
}

func (c *IPClient) DelWhiteList(request *apis.WhiteDelRequest) (*apis.Response, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	jsonDataReq, err := sonic.Marshal(request.IPInfo)
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
	jsonResp := &apis.Response{}
	err = sonic.Unmarshal(resp, jsonResp)
	if err != nil {
		c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
		return nil, err
	}
	return jsonResp, err
}
