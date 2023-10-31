package apis

import (
	"github.com/nyuncloud/ncloud-go/core"
	"github.com/nyuncloud/ncloud-go/core/net/consts"
)

type QueryIPGetRequest struct {
	core.NCloudRequest
}

func NewQueryIPGetRequest() *QueryIPGetRequest {
	return &QueryIPGetRequest{
		NCloudRequest: core.NCloudRequest{
			Path:   "/eip/ip",
			Method: consts.MethodGet,
		},
	}
}

type QueryWhiteListRequest struct {
	core.NCloudRequest
}

func NewQueryWhiteListRequest() *QueryWhiteListRequest {
	return &QueryWhiteListRequest{
		NCloudRequest: core.NCloudRequest{
			Path:   "/eip/whitelist",
			Method: consts.MethodGet,
		},
	}
}

type WhiteAddRequest struct {
	core.NCloudRequest
	IPInfo
}

func NewWhiteAddRequest(ip string) *WhiteAddRequest {
	return &WhiteAddRequest{
		NCloudRequest: core.NCloudRequest{
			Path:   "/eip/whitelist",
			Method: consts.MethodPost,
		},
		IPInfo: IPInfo{IP: ip},
	}
}

type WhiteDelRequest struct {
	core.NCloudRequest
	IPInfo
}

func NewWhiteDelRequest(ip string) *WhiteDelRequest {
	return &WhiteDelRequest{
		NCloudRequest: core.NCloudRequest{
			Path:   "/eip/whitelist",
			Method: consts.MethodDelete,
		},
		IPInfo: IPInfo{IP: ip},
	}
}

type CountGetRequest struct {
	core.NCloudRequest
}

func NewCountGetRequest() *CountGetRequest {
	return &CountGetRequest{
		NCloudRequest: core.NCloudRequest{
			Path:   "/eip/count",
			Method: consts.MethodGet,
		},
	}
}

type CountPostRequest struct {
	core.NCloudRequest
	CountPostInfo
}

func NewCountPostRequest(startDate, endDate string) *CountPostRequest {
	return &CountPostRequest{
		NCloudRequest: core.NCloudRequest{
			Path:   "/eip/count",
			Method: consts.MethodPost,
		},
		CountPostInfo: CountPostInfo{
			StartDate: startDate,
			EndDate:   endDate,
		},
	}
}
