package easy_dingding

import (
	"bytes"
	"csf/library/easy_config"
	"csf/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var DingDing dingDing

type dingDing struct {
}

func NewDingDing() dingDing {
	DingDing = dingDing{}
	return DingDing
}

func (s *dingDing) SendMsg(env string, data interface{}) (resp Response, err error) {
	if !s.isPass(env) {
		return
	}
	urlTmp := s.getUrl()
	var dataTpl interface{}
	dataTpl = s.getTemplate(data)

	dataByte, _ := json.Marshal(dataTpl)
	var respByte []byte
	respByte, err = s.send("POST", urlTmp, dataByte)
	if err != nil {
		return
	}
	err = json.Unmarshal(respByte, &resp)
	if err != nil {
		return
	}
	return
}

func (s *dingDing) send(method, url string, dataByte []byte) (respByte []byte, err error) {
	payload := bytes.NewReader(dataByte)
	var req *http.Request
	req, err = http.NewRequest(method, url, payload)
	if err != nil {
		return
	}

	req.Header.Add("Content-Type", "application/json; charset-utf8")

	client := &http.Client{}
	var resp *http.Response
	resp, err = client.Do(req)
	if err != nil {
		return
	}

	respByte, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	return
}

func (s *dingDing) getTemplate(data interface{}) (msgTpl interface{}) {
	switch val := data.(type) {
	case Text:
		msgTpl = s.getTextTpl(val)
	case MarkDown:
		msgTpl = s.getMarkdownTpl(val)
	case Link:
		msgTpl = s.getLinkTpl(val)
	case int:
		msgTpl = s.getTextTpl(Text{Content: strconv.Itoa(val)})
	case float64:
		msgTpl = s.getTextTpl(Text{Content: strconv.FormatFloat(val, 'f', -1, 64)})
	case string:
		msgTpl = s.getTextTpl(Text{Content: val})
	case []byte:
		var str string
		if utils.IsJsonWithInterface(string(val)) {
			str = string(val)
		} else {
			tmp, _ := json.Marshal(data)
			str = string(tmp)
		}
		msgTpl = s.getTextTpl(Text{Content: str})
	default:
		tmp, _ := json.Marshal(data)
		msgTpl = s.getTextTpl(Text{Content: string(tmp)})
	}
	return
}

func (s *dingDing) isPass(env string) (flag bool) {
	var envs = strings.Split(easy_config.Config.DingDing.Env, ",")
	for _, val := range envs {
		if val == env {
			flag = true
			break
		}
	}
	return
}

func (s *dingDing) getUrl() string {
	timestamp := time.Now().UnixMilli()
	signStr := fmt.Sprintf("%d\n%v", timestamp, easy_config.Config.DingDing.Secret)
	sign := utils.HmacSign(signStr)
	urlTmp := fmt.Sprintf("%v?access_token=%v&timestamp=%v&sign=%v",
		easy_config.Config.DingDing.Url,
		easy_config.Config.DingDing.AccessToken,
		timestamp,
		sign,
	)
	return urlTmp
}

func (s *dingDing) getTextTpl(msg interface{}) interface{} {
	return textTpl{
		MsgType: "text",
		Text:    msg,
	}
}

func (s *dingDing) getMarkdownTpl(msg interface{}) interface{} {
	return markdownTpl{
		MsgType:  "markdown",
		Markdown: msg,
	}
}

func (s *dingDing) getLinkTpl(msg interface{}) interface{} {
	return linkTpl{
		MsgType: "link",
		Link:    msg,
	}
}
