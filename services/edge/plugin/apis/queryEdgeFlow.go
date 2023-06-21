package apis

import (
	"github.com/nyuncloud/ncloud-go/core"
	"github.com/nyuncloud/ncloud-go/core/net/consts"
)

type QueryEdgeFlowConditions struct {
	//边缘设备唯一标识列表
	UIDList []string `json:"uid_list"`
	// 厂商列表
	ReferList []string `json:"refer_list"`
	// 项目名称列表
	ProjectList []string `json:"project_list"`
	//开始时间戳
	StartTimestamp int64 `json:"start_timestamp"`
	//持续时间
	Duration int64 `json:"duration"`
}

type QueryEdgeFlowRequest struct {
	core.NCloudRequest
	QueryEdgeFlowConditions
}

func NewQueryEdgeFlowRequest() *QueryEdgeFlowRequest {
	return &QueryEdgeFlowRequest{
		NCloudRequest: core.NCloudRequest{
			Path:   "/edge/flow/record",
			Method: consts.MethodPost,
		},
	}
}
func NewQueryEdgeFlowRequestWithAllParams(referList []string,
	projectList []string, uidList []string,
	startTimestamp, duration int64) *QueryEdgeFlowRequest {
	return &QueryEdgeFlowRequest{
		NCloudRequest: core.NCloudRequest{
			Path:   "/edge/flow/record",
			Method: consts.MethodPost,
		},
		QueryEdgeFlowConditions: QueryEdgeFlowConditions{
			UIDList:        uidList,
			ReferList:      referList,
			ProjectList:    projectList,
			StartTimestamp: startTimestamp,
			Duration:       duration,
		},
	}
}

func (r *QueryEdgeFlowRequest) SetUidList(uidList []string) {
	r.UIDList = uidList
}

func (r *QueryEdgeFlowRequest) SetReferList(referList []string) {
	r.ReferList = referList
}

func (r *QueryEdgeFlowRequest) SetProjectList(projectList []string) {
	r.ProjectList = projectList
}

func (r *QueryEdgeFlowRequest) SetStartTimestamp(startTimestamp int64) {
	r.StartTimestamp = startTimestamp
}

func (r *QueryEdgeFlowRequest) SetDuration(duration int64) {
	r.Duration = duration
}

type QueryEdgeFlowResponse struct {
	Code int64          `json:"code"`
	Msg  string         `json:"msg"`
	Data EdgeFlowResult `json:"data"`
}

type EdgeFlowResult struct {
	List []EdgeFlowItem `json:"list"`
}

type EdgeFlowItem struct {
	//设备唯一标识
	UID string `json:"uid"`
	// 所属厂商
	Refer string `json:"refer"`
	//所属项目
	Project string `json:"project"`
	//地区
	Region string `json:"region"`
	//运营商
	Isp string `json:"isp"`
	//开始时间戳
	StartTimestamp int64 `json:"start_timestamp"`
	//结束时间戳
	EndTimestamp int64 `json:"end_timestamp"`
	//上行流量
	UpFlow int64 `json:"up_flow"`
	//下行流量
	DownFlow int64 `json:"down_flow"`
}
