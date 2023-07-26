package apis

type IPInfo struct {
	IP string `json:"ip"`
}

type TrackInfo struct {
	TrackID string `json:"track_id"`
}

type Response struct {
	Code int       `json:"code"`
	Msg  string    `json:"msg"`
	Data TrackInfo `json:"data"`
}

type IPGetResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data IPGetResult `json:"data"`
}

type IPGetResult struct {
	HttpList   []string `json:"http_list"`
	Socks5List []string `json:"socks5_list"`
	TrackID    string   `json:"track_id"`
}

type WhiteListResponse struct {
	Code int             `json:"code"`
	Msg  string          `json:"msg"`
	Data WhiteListResult `json:"data"`
}

type WhiteListResult struct {
	List    []string
	TrackID string `json:"track_id"`
}
