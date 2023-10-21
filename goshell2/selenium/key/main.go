package main

import (
	"fmt"
	"time"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
	
	
)

const (
	//chromedriver路径 我们设置成绝对路径
	chromeDriver = "C:\\Program Files\\Google\\Chrome\\Application\\chromedriver.exe"

	//chromeDriver运行端口
	port = 8080
)


// 创建浏览器对象
func get_wd() (selenium.WebDriver,*selenium.Service){
	//开启selenium服务
	s,_ := selenium.NewChromeDriverService(chromeDriver,port)

	//连接webDriver服务
	//设置浏览器功能
	caps := selenium.Capabilities{}
	//设置chrome特定功能 
	chromeCaps := chrome.Capabilities{
		//使用开发者调试模式
		ExcludeSwitches : []string{"enable-automatin"},
		//基本功能
		Args: []string{
			"--no-sandbox",
			// 设置请求头
			"--user-agent=Mozilla/5.0 (Windows NT 10.0; Win64; "+
				 "x64) AppleWebKit/537.36 (KHTML, like Gecko) "+
				 "Chrome/94.0.4606.61 Safari/537.36",
		},
		
	}


	//将谷歌浏览器的特定功能chromecaps添加到caps
	caps.AddChrome(chromeCaps)
	
	//根据浏览器功能连接
	urlPrefix:=fmt.Sprintf("http://localhost:%d/wd/hub",port)
	wd, _ := selenium.NewRemote(caps, urlPrefix)
	return wd,s
}



func main() {
	//获取浏览器对象
	wd,s := get_wd()
	//关闭服务
	defer s.Stop()
	//关闭浏览器对象 
	defer wd.Quit()

	//访问网页
	wd.Get("https://www.zhipin.com/")

	//最大话窗口
	

	//输入查询职位
	query,_ := wd.FindElement(selenium.ByName,"query")
	query.SendKeys("go语言")
	time.Sleep(2*time.Second)

	//单击输入按钮
	search,_ := wd.FindElement(selenium.ByCSSSelector,`[class="btn btn-search"]`)
	r1,_:= search.IsEnabled()
	fmt.Printf("判断网页元素是否可编辑或可点击状态：%v\n", r1)


	search.Click()
	time.Sleep(2*time.Second)  


}