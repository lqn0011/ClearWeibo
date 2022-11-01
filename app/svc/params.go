package svc

const (
	MyBlogUrl      = "https://weibo.com/ajax/statuses/mymblog"
	BlogDestroyUrl = "https://weibo.com/ajax/statuses/destroy"
)

type MymblogResp struct {
	Data MymblogData `json:"data"`
	Ok   int         `json:"ok"`
}

type MymblogList struct {
	OriMid  string `json:"ori_mid,omitempty"`
	Mid     string `json:"mid"`
	ID      int64  `json:"id"`
	TextRaw string `json:"text_raw"`
}

type MymblogData struct {
	List  []MymblogList `json:"list"`
	Total int           `json:"total"`
}

type DelErrResp struct {
	Ok      int    `json:"ok"`
	Message string `json:"message"`
}
