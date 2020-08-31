package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type WebConfigItem struct {
	Name        string      `json:"name"`
	Method      string      `json:"method"`
	RequestURL  string      `json:"requestUrl"`
	RequestData requestData `json:"requestData"`
}
type requestData struct {
	Header map[string]string `json:"header"`
	Data   map[string]string `json:"data"`
}

func ParseWebConfigItem(item WebConfigItem) (req *http.Request) {
	// 暂时只有http
	methodArg := strings.Split(item.Method, ":")
	if len(methodArg) < 3 {
		panic(errors.New("WebConfigItem method error"))
	}
	var subtype = methodArg[2]
	var requestUrl = item.RequestURL
	var err error
	log.Println(subtype)
	switch subtype {
	case "json":
		bytesData, _ := json.Marshal(item.RequestData.Data)
		req, err = http.NewRequest("POST", requestUrl, bytes.NewReader(bytesData))
		if err != nil {
			fmt.Println(err, "==============")
		}
	case "form":
		urlValues := url.Values{}
		for k, v := range item.RequestData.Data {
			urlValues.Add(k, v)
		}
		req, err = http.NewRequest("POST", requestUrl, strings.NewReader(urlValues.Encode()))
		if err != nil {
			fmt.Println(err, "==============")
		}
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	default:
		params := url.Values{}
		Url, err := url.Parse(requestUrl)
		if err != nil {
			panic(err)
		}
		for k, v := range item.RequestData.Data {
			params.Set(k, v)
		}
		Url.RawQuery = params.Encode()
		urlPath := Url.String()
		req, err = http.NewRequest("GET", urlPath, strings.NewReader(params.Encode()))
		if err != nil {
			fmt.Println(err, "==============")
		}
	}
	for k, v := range item.RequestData.Header {
		req.Header.Add(k, v)
	}
	return req
}
