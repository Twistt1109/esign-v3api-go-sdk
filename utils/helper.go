package tools

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"os"
)

// Base64编码
func Base64Encode(dataString string) string {
	encodeString := base64.StdEncoding.EncodeToString([]byte(dataString))
	return encodeString
}

// Base64解码
func Base64Decode(encodeString string) []byte {
	decodeBytes, err := base64.StdEncoding.DecodeString(encodeString)
	if err != nil {
		fmt.Println(err)
	}
	return decodeBytes
}

// 将文件进行Base64编码
func Base64EncodeByFile(filePath string) string {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	}
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	} else {
		file.Close()
	}
	encodeString := base64.StdEncoding.EncodeToString(fileBytes)
	return encodeString
}

// 保存文件
func SaveFileByBase64(base64String, outFilePath string) {
	// 将Base64字符串解码为[]byte
	var fileBytes = Base64Decode(base64String)

	saveFileErr := ioutil.WriteFile(outFilePath, fileBytes, 0666)

	if saveFileErr != nil {
		fmt.Println("文件保存失败:" + saveFileErr.Error())
		panic(saveFileErr)
	} else {
		fmt.Println("文件保存成功:" + outFilePath)
	}
}

// 摘要md5
func DohashMd5(body string) (md5Str string) {
	hash := md5.New()
	hash.Write([]byte(body))
	md5Data := hash.Sum(nil)
	return base64.StdEncoding.EncodeToString(md5Data)
}

// 计算文件content-md5
const fileChunk = 8192 // 8KB
func CountFileMd5(filePath string) (string, int64) {
	file, err := os.Open(filePath)
	if err != nil {
		return err.Error(), 0
	}
	defer file.Close()

	info, _ := file.Stat()
	fileSize := info.Size()

	blocks := uint64(math.Ceil(float64(fileSize) / float64(fileChunk)))
	hash := md5.New()

	for i := uint64(0); i < blocks; i++ {
		blockSize := int(math.Min(fileChunk, float64(fileSize-int64(i*fileChunk))))
		buf := make([]byte, blockSize)

		file.Read(buf)
		io.WriteString(hash, string(buf))
	}

	return base64.StdEncoding.EncodeToString(hash.Sum(nil)), fileSize
}

// sha256摘要签名
func DoSignatureBase64(message string, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(message))
	buf := h.Sum(nil)
	return base64.StdEncoding.EncodeToString(buf)
}

// 拼接请求参数
func AppendSignDataString(method string, accept string, contentMD5 string, contentType string, date string, headers string, url string) string {
	var buffer bytes.Buffer
	buffer.WriteString(method)
	buffer.WriteString("\n")
	buffer.WriteString(accept)
	buffer.WriteString("\n")
	buffer.WriteString(contentMD5)
	buffer.WriteString("\n")
	buffer.WriteString(contentType)
	buffer.WriteString("\n")
	buffer.WriteString(date)
	buffer.WriteString("\n")
	if len(headers) == 0 {
		buffer.WriteString(headers)
		buffer.WriteString(url)
	} else {
		buffer.WriteString(headers)
		buffer.WriteString("\n")
		buffer.WriteString(url)
	}
	return buffer.String()
}

// byte转json
func BytetoJson(initResult []byte) map[string]interface{} {
	var initResultJson interface{}
	json.Unmarshal(initResult, &initResultJson)
	jsonMap, err := initResultJson.(map[string]interface{})
	if !err {
		fmt.Println("DO SOMETHING!")
		return nil
	}
	return jsonMap
}
