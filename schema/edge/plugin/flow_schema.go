package plugin

type EdgeFlowObject struct {
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
