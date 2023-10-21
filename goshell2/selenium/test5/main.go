/*
 * @Descripttion:
 * @version:
 * @Author: wzs
 * @Date: 2023-05-20 14:04:26
 * @LastEditors: Andy
 * @LastEditTime: 2023-05-20 14:05:42
 */
package main

import (
	"fmt"
	"time"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

func main() {
	// ChromeDriver路径信息
	chromeDriver := "C:\\Program Files\\Google\\Chrome\\Application\\chromedriver.exe"
	// ChromeDriver运行端口
	port := 8080

	/* 开启WebDriver服务 */
	s, _ := selenium.NewChromeDriverService(chromeDriver, port)
	// 关闭服务
	defer s.Stop()

	/* 连接WebDriver服务 */
	caps := selenium.Capabilities{}

	// 设置Chrome特定功能
	chromeCaps := chrome.Capabilities{
		// 使用开发者调试模式
		ExcludeSwitches: []string{"enable-automation"},
	}

	// 将谷歌浏览器特定功能chromeCaps添加到caps
	caps.AddChrome(chromeCaps)

	// 根据浏览器功能连接Selenium
	urlPrefix := fmt.Sprintf("http://127.0.0.1:%d/wd/hub", port)
	wd, _ := selenium.NewRemote(caps, urlPrefix)
	// 关闭浏览器对象
	defer wd.Quit()

	// 访问网址
	wd.Get("https://www.baidu.com/s?wd=go")
	time.Sleep(3 * time.Second)

	// 等待网页元素加载，参数condition是匿名函数，参数timeout是等待超时
	wd.WaitWithTimeout(func(wd selenium.WebDriver) (bool, error) {
		_, err := wd.FindElement(selenium.ByID, "su")
		if err == nil {
			return true, nil
		} else {
			return false, err
		}
	}, 60*time.Second)

	// 等待网页元素加载，参数condition是匿名函数，默认等待1分钟
	wd.Wait(func(wd selenium.WebDriver) (bool, error) {
		_, err := wd.FindElement(selenium.ByID, "su")
		if err == nil {
			return true, nil
		} else {
			return false, err
		}
	})
}
