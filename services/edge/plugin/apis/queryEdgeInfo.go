package apis

import (
	"github.com/nyuncloud/ncloud-go/core"
	"github.com/nyuncloud/ncloud-go/core/net/consts"
)

type QueryEdgeInfoRequest struct {
	core.NCloudRequest
	PageInfo
}

type PageInfo struct {
	/* 第几页*/
	Page int64 `json:"page"`
	/* 每页数量*/
	Size int64 `json:"size"`
}

func NewQueryEdgeInfoRequestWithParams(page, size int64) *QueryEdgeInfoRequest {
	return &QueryEdgeInfoRequest{
		NCloudRequest: core.NCloudRequest{
			Path:   "/edge/info/list",
			Method: consts.MethodPost,
		},
		PageInfo: PageInfo{
			Page: page,
			Size: size,
		},
	}
}

type QueryEdgeInfoRecordRequest struct {
	core.NCloudRequest
	QueryEdgeInfoConditions
}
type QueryEdgeInfoConditions struct {
	/* 边缘设备唯一标识*/
	UID string `json:"uid"`
	/* 整小时 时间戳*/
	StartTimestamp int64 `json:"start_timestamp"`
	/* 时长 最小3600秒*/
	Duration int64 `json:"duration"`
}

func NewQueryEdgeInfoRecordRequestAllParams(uid string,
	startTimestamp, duration int64) *QueryEdgeInfoRecordRequest {
	return &QueryEdgeInfoRecordRequest{
		NCloudRequest: core.NCloudRequest{
			Path:   "/edge/info/record",
			Method: consts.MethodPost,
		},
		QueryEdgeInfoConditions: QueryEdgeInfoConditions{
			UID:            uid,
			StartTimestamp: startTimestamp,
			Duration:       duration,
		},
	}
}

type QueryEdgeInfoResponse struct {
	Code int64          `json:"code"`
	Msg  string         `json:"msg"`
	Data EdgeInfoResult `json:"data"`
}

type EdgeInfoResult struct {
	Total int64          `json:"total"`
	Page  int64          `json:"page"`
	Size  int64          `json:"size"`
	List  []EdgeInfoItem `json:"list"`
}

type EdgeInfoItem struct {
	//设备唯一标识
	UID string `json:"uid"`
	// 厂商
	Refer string `json:"refer"`
	//总磁盘 单位MB
	TotalDisk int64 `json:"total_disk"`
	//总内存 单位MB
	TotalMem int64 `json:"total_mem"`
	//空闲磁盘 单位MB
	FreeDisk int64 `json:"free_disk"`
	//空闲内存 单位MB
	FreeMem int64 `json:"free_mem"`
	//插件固件版本
	PluginVersion int64 `json:"plugin_version"`
}
type EdgeInfoRecordResponse struct {
	Code int64                `json:"code"`
	Msg  string               `json:"msg"`
	Data EdgeInfoRecordResult `json:"data"`
}

type EdgeInfoRecordResult struct {
	List []*EdgeInfoRecordItem `json:"list"`
}

type EdgeInfoRecordItem struct {
	//设备唯一标识
	UID string `json:"uid"`
	//更新时间戳
	Timestamp int64 `json:"timestamp"`
	//空闲cpu 占比
	FreeCup float64 `json:"free_cup"`
	//空闲IO 占比
	FreeIO float64 `json:"free_io"`
	//空闲磁盘 占比
	FreeDisk float64 `json:"free_disk"`
	//空闲内存 占比
	FreeMem float64 `json:"free_mem"`
}
