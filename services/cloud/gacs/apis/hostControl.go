package apis

import (
	"github.com/nyuncloud/ncloud-go/core"
	"github.com/nyuncloud/ncloud-go/core/net/consts"
)

type HostInfo struct {
	ID     string `json:"id"`
	Region string `json:"region"`
	Area   string `json:"area"`
}

type HostStartRequest struct {
	core.NCloudRequest
	HostInfo
}

func NewHostStartRequestWithParams(id, region, area string) *HostStartRequest {
	return &HostStartRequest{
		NCloudRequest: core.NCloudRequest{
			Path:   "/start",
			Method: consts.MethodPost,
		},
		HostInfo: HostInfo{
			ID:     id,
			Region: region,
			Area:   area,
		},
	}
}

type HostRestartRequest struct {
	core.NCloudRequest
	HostInfo
}

func NewHostRestartRequestWithParams(id, region, area string) *HostRestartRequest {
	return &HostRestartRequest{
		NCloudRequest: core.NCloudRequest{
			Path:   "/restart",
			Method: consts.MethodPost,
		},
		HostInfo: HostInfo{
			ID:     id,
			Region: region,
			Area:   area,
		},
	}
}

type HostStopRequest struct {
	core.NCloudRequest
	HostInfo
}

func NewHostStopRequestWithParams(id, region, area string) *HostStopRequest {
	return &HostStopRequest{
		NCloudRequest: core.NCloudRequest{
			Path:   "/stop",
			Method: consts.MethodPost,
		},
		HostInfo: HostInfo{
			ID:     id,
			Region: region,
			Area:   area,
		},
	}
}

type HostFindImageRequest struct {
	core.NCloudRequest
	HostInfo
}

func NewHostFindImageRequestWithParams(id, region, area string) *HostFindImageRequest {
	return &HostFindImageRequest{
		NCloudRequest: core.NCloudRequest{
			Path:   "/update/system/findImage",
			Method: consts.MethodPost,
		},
		HostInfo: HostInfo{
			ID:     id,
			Region: region,
			Area:   area,
		},
	}
}

type HostInfoWithPassword struct {
	HostInfo
	Password string `json:"admin_pass"`
}

type HostUpdatePasswordRequest struct {
	core.NCloudRequest
	HostInfoWithPassword
}

func NewHostUpdatePasswordRequestWithParams(id, region, area, password string) *HostUpdatePasswordRequest {
	return &HostUpdatePasswordRequest{
		NCloudRequest: core.NCloudRequest{
			Path:   "/update/password",
			Method: consts.MethodPost,
		},
		HostInfoWithPassword: HostInfoWithPassword{
			HostInfo: HostInfo{
				ID:     id,
				Region: region,
				Area:   area,
			},
			Password: password,
		},
	}
}

type HostInfoWithImage struct {
	HostInfo
	Password string `json:"admin_pass"`
	ImageID  string `json:"image_id"`
}

type HostUpdateSystemRequest struct {
	core.NCloudRequest
	HostInfoWithImage
}

func NewHostUpdateSystemRequestWithParams(id, region, area, password, imageID string) *HostUpdateSystemRequest {
	return &HostUpdateSystemRequest{
		NCloudRequest: core.NCloudRequest{
			Path:   "/update/system",
			Method: consts.MethodPost,
		},
		HostInfoWithImage: HostInfoWithImage{
			HostInfo: HostInfo{
				ID:     id,
				Region: region,
				Area:   area,
			},
			Password: password,
			ImageID:  imageID,
		},
	}
}
