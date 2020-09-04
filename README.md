# harassing-handle
用骚扰电话打败骚扰电话
提交一个骚扰电话, 推送到配置内容到其他骚扰电话的请求接口
被出卖的信息烦人的电话，把骚电话提交到 【买房买车】【男科医院】【线上赌博】【期货交易】的接口中，一次拒接二次贩卖，骚扰电话可也可以被骚扰。
## 使用
go run main.go submit -n 张三 -p 16601139225 -c "./config.json"
## 配置
```js
// "method": "http:post:form" |  "http:get" // post 支持 form 或 json 
{
  "list": [
    {
      "name": "测试",
      "method": "http:post:form",
      "requestUrl": "http://www.baidu.com",
      "requestData": {
        "header": {
          "accept": "*/*",
          "cookie": "ss"
        },
        "data": {
          "name": "${name}",
          "phone": "${phone}"
        }
      }
    }
]}
```
## todo:
  支持 http-proxy
  接口返回验证
 
