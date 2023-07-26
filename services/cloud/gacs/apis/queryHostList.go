package apis

import (
	"github.com/nyuncloud/ncloud-go/core"
	"github.com/nyuncloud/ncloud-go/core/net/consts"
)

type QueryHostListRequest struct {
	core.NCloudRequest
	PageInfo
}

func NewQueryHostListRequestWithParams(page, size int64) *QueryHostListRequest {
	return &QueryHostListRequest{
		NCloudRequest: core.NCloudRequest{
			Path:   "/list",
			Method: consts.MethodPost,
		},
		PageInfo: PageInfo{
			Page: page,
			Size: size,
		},
	}
}

type QueryHostAllRequest struct {
	core.NCloudRequest
}

func NewQueryHostAll() *QueryHostAllRequest {
	return &QueryHostAllRequest{
		NCloudRequest: core.NCloudRequest{
			Path:   "/all",
			Method: consts.MethodPost,
		},
	}
}

type QueryHostListResponse struct {
	Code int64      `json:"code"`
	Msg  string     `json:"msg"`
	Data HostResult `json:"data"`
}

type QueryHostAllResponse struct {
	Code int64    `json:"code"`
	Msg  string   `json:"msg"`
	Data HostItem `json:"data"`
}

type HostResult struct {
	Total int64      `json:"total"`
	Page  int64      `json:"page"`
	Size  int64      `json:"page_size"`
	Data  []HostItem `json:"data"`
}

type HostItem struct {
	SuperID          int64  `json:"super_id"`
	SuperName        string `json:"super_name"`
	Area             string `json:"area"`
	ZoneDesc         string `json:"zone_desc"`
	CloudID          string `json:"cloud_id"`
	Name             string `json:"name"`
	VCpu             int64  `json:"vcpu"`
	VMemory          int64  `json:"vmemory"`
	VDisk            int64  `json:"vdisk"`
	Status           int8   `json:"status"`
	ImageRef         string `json:"image_ref"`
	ImageName        string `json:"image_name"`
	ImageOSType      string `json:"image_os_type"`
	SystemDiskID     string `json:"system_disk_id"`
	CloudCreatedTime string `json:"cloud_created_time"`
	ServerType       string `json:"server_type"`
	ServerVmtype     string `json:"server_vmtype"`
	ProductType      string `json:"product_type"`
	OpStatus         string `json:"op_status"`
	ECStatus         string `json:"ec_status"`
	BootVolumeType   string `json:"boot_volume_type"`
	AdminPaused      bool   `json:"admin_paused"`
	UserName         string `json:"user_name"`
	Visible          bool   `json:"visible"`
	Region           string `json:"region"`
	SpecsName        string `json:"specs_name"`
	RecycleStatus    string `json:"recycle_status"`
	ReleaseProtect   bool   `json:"release_protect"`
	StoppedMode      bool   `json:"stopped_mode"`
	PrivateIP        string `json:"private_ip"`
	PublicIP         string `json:"public_ip"`
}
