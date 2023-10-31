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
	List    []IPObject `json:"list"`
	TrackID string     `json:"track_id"`
}

type IPObject struct {
	IP     string `json:"ip"`
	Port   int64  `json:"port"`
	S5Port int64  `json:"s5port"`
}

type WhiteListResponse struct {
	Code int             `json:"code"`
	Msg  string          `json:"msg"`
	Data WhiteListResult `json:"data"`
}

type WhiteListResult struct {
	List    []string `json:"list"`
	TrackID string   `json:"track_id"`
}

type CountPostInfo struct {
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

type CountPostResult struct {
	TrackID string        `json:"track_id"`
	List    []CountObject `json:"list"`
}
type CountPostResponse struct {
	Code int             `json:"code"`
	Msg  string          `json:"msg"`
	Data CountPostResult `json:"data"`
}

type CountObject struct {
	Date  string `json:"date"`
	Count int64  `json:"count"`
}

type CountGetResult struct {
	TrackID string `json:"track_id"`
	Count   int64  `json:"count"`
}

type CountGetResponse struct {
	Code int            `json:"code"`
	Msg  string         `json:"msg"`
	Data CountGetResult `json:"data"`
}
