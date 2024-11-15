package http

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"time"

	"github.com/valyala/fastjson"
)

type ReqBody map[string]interface{}

func ParseJson(res string) (*fastjson.Value, error) {
	var p fastjson.Parser
	return p.Parse(res)
}

func PostForm(
	reqbody map[string]interface{},
	url string,
	actionHandler func(res http.ResponseWriter, req *http.Request),
) (respBody string, err error) {
	bodyBytes, err := json.Marshal(reqbody)
	if err != nil {
		return "", err
	}
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return "", err
	}

	//lint:ignore SA1029 按要求用 type contextKey string 代替，会在 resp.go 中取不出
	ctx := context.WithValue(req.Context(), "params", bodyBytes)
	req = req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	res := httptest.NewRecorder()
	actionHandler(res, req)
	time.Sleep(time.Second * 1)
	return res.Body.String(), nil
}

// PostJson 发送 json 请求
func PostJson(
	reqbody map[string]interface{},
	url string,
	actionHandler func(res http.ResponseWriter, req *http.Request),
) (respBody string, err error) {
	bodyBytes, err := json.Marshal(reqbody)
	if err != nil {
		return "", err
	}
	fmt.Println(string(bodyBytes))
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()
	actionHandler(res, req)
	time.Sleep(time.Second * 1)
	return res.Body.String(), nil
}
