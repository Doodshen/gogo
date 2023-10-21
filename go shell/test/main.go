/*
 * @Descripttion:
 * @version:
 * @Author: wzs
 * @Date: 2023-05-19 15:59:56
 * @LastEditors: Andy
 * @LastEditTime: 2023-05-19 16:02:45
 */
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	urls := "https://search.51job.com/list/030200,000000,0000,00,9,99,golang,2,1.html"

	// 定义请求对象NewRequest，参数method可以为GET和POST
	// 参数url为HTTP的请求链接，参数body为请求参数，GET请求设为nil即可
	req, _ := http.NewRequest("GET", urls, nil)
	// POS请求将请求参数以k1=v1&k2=v2的形式表示
	// 再由strings.NewReader()转换格式
	// b := strings.NewReader("name=cjb")
	// req, _ := http.NewRequest("POST",urls, b)

	// 为请求对象NewRequest设置请求头
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64;x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/94.0.4606.81 Safari/537.36")

	// 设置Cookie
	// Cookie以结构体Cookie形式表示
	cookie := http.Cookie{Name: "clientcookieid", Value: "121"}
	req.AddCookie(&cookie)

	// 设置代理IP，代理IP必须以匿名函数表示
	// 因为Transport的参数Proxy以匿名函数定义
	proxy := func(_ *http.Request) (*url.URL, error) {
		return url.Parse("http://xxx.xxx.xxx.xxx:xxxx")
	}
	transport := &http.Transport{Proxy: proxy}

	// 在Client中设置参数Transport即可实现代理IP
	client := &http.Client{Transport: transport}
	// 发送HTTP请求
	resp, _ := client.Do(req)
	// 获取网站响应内容
	body, _ := ioutil.ReadAll(resp.Body)
	// 将响应内容转换为字符串格式输出
	fmt.Printf("获取网站响应内容：%v\n", string(body))
}
