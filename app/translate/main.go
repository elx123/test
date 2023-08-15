package main

import (
	"fmt"
	"log"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/alimt"
	"golang.design/x/clipboard"
)

func main() {
	err := clipboard.Init()
	if err != nil {
		panic(err)
	}
	result := clipboard.Read(clipboard.FmtText)
	log.Println(string(result))

	// 创建ecsClient实例
	alimtClient, err := alimt.NewClientWithAccessKey(
		"mt.cn-hangzhou.aliyuncs.com",    // 地域ID
		"LTAI5tDh4za5UkS3cM3Yc4Cp",       // 您的Access Key ID
		"9rxhHSvdLjMuQO09W8OrsSx94WZxUX") // 您的Access Key Secret
	if err != nil {
		// 异常处理
		panic(err)
	}
	// 创建API请求并设置参数
	request := alimt.CreateTranslateECommerceRequest()
	// 等价于 request.PageSize = "10"
	request.Method = "POST"             //设置请求
	request.FormatType = "text"         //翻译文本的格式
	request.SourceLanguage = "en"       //源语言
	request.SourceText = string(result) //原文
	request.TargetLanguage = "zh"       //目标语言
	request.Scene = "title"             //目标语言
	// 发起请求并处理异常
	response, err := alimtClient.TranslateECommerce(request)
	if err != nil {
		// 异常处理
		log.Println(err)
		panic(err)
	}
	fmt.Println(response.Data.Translated)
}
