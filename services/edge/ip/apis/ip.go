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
	Http    string `json:"http"`
	Socks5  string `json:"socks5"`
	TrackID string `json:"track_id"`
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
