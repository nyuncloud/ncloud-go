package apis

type QueryResponse struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
}

type PageInfo struct {
	/* 第几页*/
	Page int64 `json:"page"`
	/* 每页数量*/
	Size int64 `json:"pageSize"`
}
