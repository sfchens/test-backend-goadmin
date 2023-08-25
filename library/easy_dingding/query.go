package easy_dingding

type Text struct {
	Content string `json:"content"`
}
type textTpl struct {
	MsgType string      `json:"msgtype"`
	Text    interface{} `json:"text"`
}

type MarkDown struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}
type markdownTpl struct {
	MsgType  string      `json:"msgtype"`
	Markdown interface{} `json:"markdown"`
}

type Link struct {
	Title      string `json:"title"`
	Text       string `json:"text"`
	MessageUrl string `json:"messageUrl"`
	PicUrl     string `json:"picUrl"`
}
type linkTpl struct {
	MsgType string      `json:"msgtype"`
	Link    interface{} `json:"link"`
}

type Response struct {
	Errcode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}
