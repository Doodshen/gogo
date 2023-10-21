/*
 * @Descripttion:
 * @version:
 * @Author: wzs
 * @Date: 2023-05-20 13:40:51
 * @LastEditors: Andy
 * @LastEditTime: 2023-05-20 16:40:29
 */
package main

import (
	"fmt"
	"time"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

// 设置常量
const (
	// ChromeDriver路径信息
	chromeDriver = "chromedriver.exe"

	// ChromeDriver运行端口
	port = 8080
)

func main() {
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

	// 通过class属性定位元素
	ele1, _ := wd.FindElement(selenium.ByClassName, "s_ipt")
	ele1.SendKeys("Golang")
	time.Sleep(3 * time.Second)

	// 通过id属性定位元素
	ele2, _ := wd.FindElement(selenium.ByID, "kw")
	ele2.SendKeys("good")
	time.Sleep(3 * time.Second)

	// 通过HTML标签定位元素，先定位局部范围，再定位标签
	ele3, _ := wd.FindElement(selenium.ByClassName, "quickdelete-wrap")
	ele31, _ := ele3.FindElement(selenium.ByTagName, "input")
	ele31.SendKeys("nice")
	time.Sleep(3 * time.Second)

	// 通过name属性定位元素
	ele4, _ := wd.FindElement(selenium.ByName, "wd")
	ele4.SendKeys("very")
	time.Sleep(3 * time.Second)

	// 通过CssSelector定位元素
	ele5, _ := wd.FindElement(selenium.ByCSSSelector, ".s_ipt")
	ele5.SendKeys("go语言")
	time.Sleep(3 * time.Second)

	// 通过链接文本定位元素
	ele6, _ := wd.FindElement(selenium.ByLinkText, "资讯")
	t6, _ := ele6.Text()
	fmt.Printf("链接文本：%v\n", t6)

	// 通过链接文本定位元素，支持模糊匹配
	ele7, _ := wd.FindElement(selenium.ByPartialLinkText, "视")
	t7, _ := ele7.Text()
	fmt.Printf("模糊匹配链接文本：%v\n", t7)

	// 通过Xpath语法定位元素
	ele8, _ := wd.FindElement(selenium.ByXPATH, `//*[@id="s_tab"]/div/b`)
	t8, _ := ele8.Text()
	fmt.Printf("Xpath语法定位元素：%v\n", t8)
}
