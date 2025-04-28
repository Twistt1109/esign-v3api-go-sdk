package esign

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func SendHttp(apiUrl string, data string, method string, headers map[string]string) ([]byte, int) {
	// API接口返回值
	var apiResult []byte
	url := apiUrl
	var jsonStr = []byte(data)
	var req *http.Request
	var err error
	if method == "GET" || method == "DELETE" {
		req, err = http.NewRequest(method, url, nil)
	} else {
		req, err = http.NewRequest(method, url, bytes.NewBuffer(jsonStr))
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	} else {
		var httpStatus = resp.StatusCode
		if httpStatus != http.StatusOK {
			return apiResult, httpStatus
		}
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		apiResult = body
		return apiResult, httpStatus
	}

}

// 文件上传
func UpLoadFile(uploadUrl string, filePath string, contentMD5 string, contentType string) string {
	//创建一个缓冲区对象,后面的要上传的body都存在这个缓冲区里
	bodyBuf := &bytes.Buffer{}
	//要上传的文件
	//创建第一个需要上传的文件,filepath.Base获取文件的名称
	//打开文件
	fd1, _ := os.Open(filePath)
	defer fd1.Close()
	//把第一个文件流写入到缓冲区里去
	_, _ = io.Copy(bodyBuf, fd1)
	//获取请求Content-Type类型,后面有用
	//contentType := bodyWriter.FormDataContentType()
	//创建一个http客户端请求对象
	client := &http.Client{}
	//创建一个post请求
	req, _ := http.NewRequest("PUT", uploadUrl, nil)
	//设置请求头
	req.Header.Set("Content-MD5", contentMD5)
	//这里的Content-Type值就是上面contentType的值
	req.Header.Set("Content-Type", contentType)
	//转换类型
	req.Body = ioutil.NopCloser(bodyBuf)
	//发送数据
	data, _ := client.Do(req)
	//读取请求返回的数据
	bytes, _ := ioutil.ReadAll(data.Body)
	defer data.Body.Close()
	//返回数据
	return string(bytes)
}

func (o *OpenApiClient) SendCommHttp(apiUrl, dataJsonStr, method, token string) (initResult []byte, httpStatus int) {
	log.Println("请求参数JSON字符串：" + dataJsonStr)
	log.Println("发送地址: " + apiUrl)

	initResult, httpStatus = SendHttp(apiUrl, dataJsonStr, method, buildCommHeader(o.appId, token))
	return initResult, httpStatus
}

func buildCommHeader(appID, token string) (header map[string]string) {
	headers := map[string]string{}
	headers["X-Tsign-Open-App-Id"] = appID
	headers["X-Tsign-Open-Token"] = token
	headers["Content-Type"] = "application/json; charset=UTF-8"
	return headers
}
