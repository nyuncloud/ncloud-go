package plugin

type EdgeInfoObject struct {
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

type EdgeInfoRecordObject struct {
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
