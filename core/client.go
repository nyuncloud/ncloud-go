package core

import (
	"encoding/base64"
	"github.com/nyuncloud/ncloud-go/core/net/consts"
	coreFasthttp "github.com/nyuncloud/ncloud-go/core/net/fasthttp"
	"github.com/valyala/fasthttp"
	"time"
)

type NCloudClient struct {
	Credential  Credential
	Config      Config
	ServiceName string
	Revision    string
	Logger      Logger
}

func (c NCloudClient) Send(method, path, data string) ([]byte, error) {
	signer := NewSigner(c.Credential)
	reqUrl := "/" + c.Revision + "/" + c.ServiceName + path
	sign, timestamp := signer.Sign(reqUrl)
	url := c.Config.Scheme + "://" + c.Config.Endpoint + reqUrl
	return coreFasthttp.RetryDoWithErr(func() ([]byte, error) {
		return c.doSend(method, url,
			data, sign, timestamp, c.Config.Timeout)
	}, c.Config.Retry, consts.DefaultSleep)
}

func (c NCloudClient) doSend(method, url, data, signed, t string,
	timeout time.Duration) ([]byte, error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	req.Header.SetMethod(method)
	c.setHeader(req, c.Credential.AccessKey, signed, t)
	req.SetRequestURI(url)
	if data != "" {
		req.SetBodyString(data)
	}
	rsp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(rsp)
	if err := fasthttp.DoTimeout(req, rsp, timeout); err != nil {
		return nil, err
	}
	if data != "" {
		c.Logger.Log(LogInfo, "req data:", data)
	}
	c.Logger.Log(LogInfo, "trace_id:", string(rsp.Header.Peek("X-UUID")))
	if rsp.Header.StatusCode() > fasthttp.StatusBadRequest {
		c.Logger.Log(LogError, "status msg:", string(rsp.Header.StatusMessage()))
	} else {
		c.Logger.Log(LogInfo, "rsp data:", string(rsp.Body()))
	}
	return rsp.Body(), nil

}

func (c NCloudClient) setHeader(req *fasthttp.Request, ak string, signed string, t string) {
	req.Header.Set("Content-Type", "application/json")
	stdEncoding := "ak=" + ak + "&sign=" + signed + "&t=" + t
	req.Header.Set("Authorization",
		base64.StdEncoding.EncodeToString([]byte(stdEncoding)))

}
