package harass

import (
	"runtime"
	"testing"
)

var jsonString string = `
{ "list" :[{
	"name":"测试",
    "method":"http:post:form",
    "requestUrl":"http://www.baidu.com",
    "requestData": {
		"header": {
			"accept": "*/*",
			"cookie":"ss"
		},
		"data":{
			"name":"${name}",
			"phone": "${phone}"
		}
    }
},{
	"name":"测试2",
    "method":"http:post:form",
    "requestUrl":"http://www.baidu.com",
    "requestData": {
		"header": {
			"accept": "*/*",
			"cookie":"ss"
		},
		"data":{
			"name":"${name}",
			"phone": "${phone}"
		}
    }
}]}`

func TestHarass_Do(t *testing.T) {
	runtime.GOMAXPROCS(4)
	h := NewHarass(&HarParamConfig{Phone: "12345", ConfigPath: "../config.json"})
	h.Do()
}
