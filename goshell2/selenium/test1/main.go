package main
import (
	"fmt"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)
//设置常量
const (
	//chromedriver路径 我们设置成绝对路径
	chromeDriver  = "C:\\Program Files\\Google\\Chrome\\Application\\chromedriver.exe"

	//chromeDriver运行端口
	port = 8080
)

func main() {
	/* 开启WebDriver服务 */
	// 设置Selenium的服务配置
	opts := []selenium.ServiceOption{
		 // 开启Selenium的执行记录
		 //selenium.Output(os.Stderr),
	}


	// 设置Selenium的调试模式
	selenium.SetDebug(false)
	// 开启Selenium服务
	// 参数path是ChromeDriver的路径信息
	// 参数port是ChromeDriver的运行端口
	// 参数opts是ChromeDriver的服务配置
	s,_:=selenium.NewChromeDriverService(chromeDriver,port,opts...)



	// 关闭服务
	defer s.Stop()



	/* 连接WebDriver服务 */
	// 设置浏览器功能，变量caps通用于Chrome和Firefox
	caps := selenium.Capabilities{}
	// chrome.Capabilities{}来自tebeka/selenium/chrome
	// 火狐使用tebeka/selenium/firefox的firefox.Capabilities{}
	// 设置Chrome的特定功能
	chromeCaps := chrome.Capabilities{
		// 禁止加载图片，加快渲染速度
		Prefs: map[string]interface{}{
			 "profile.managed_default_content_settings.images":2},
		// 使用开发者调试模式
		ExcludeSwitches: []string{"enable-automation"},
		// 基本功能
		Args: []string{
			 // 无界面模式，不会打开浏览器，程序后台运行
			 //"--headless",
			 // 浏览器窗口全屏模式
			 //"--start-maximized",
			 // 浏览器窗口大小设置
			 //"--window-size=1200x600",
			 // 取消沙盒模式
			 "--no-sandbox",
			 // 设置请求头
			 "--user-agent=Mozilla/5.0 (Windows NT 10.0; Win64; "+
				  "x64) AppleWebKit/537.36 (KHTML, like Gecko) "+
				  "Chrome/94.0.4606.61 Safari/537.36",
			 // 禁止扩展功能
			 "--disable-extensions",
			 // 禁用沙盒模式
			 "--disable-setuid-sandbox",
			 // 禁止使用shm
			 "--disable-dev-shm-usage",
			 // 禁用GPU加速
			 "--disable-gpu",
			 // 关闭安全策略
			 "--disable-web-security",
			 // 允许运行不安全的内容
			 "--allow-running-insecure-content",
		},
	}


	// 将谷歌浏览器特定功能chromeCaps添加到caps
	caps.AddChrome(chromeCaps)



	// 设置浏览器的代理IP
	http := "http://xxx.xxx.xxx.xxx:xxxx"
	caps.AddProxy(selenium.Proxy{Type: selenium.Manual, HTTP: http, HTTPPort: 0000})


	// 根据浏览器功能连接
	urlPrefix:=fmt.Sprintf("http://localhost:%d/wd/hub",port)
	wd, _ := selenium.NewRemote(caps, urlPrefix)
	// 关闭浏览器对象
	defer wd.Quit()
	// 访问网址
	wd.Get("http://httpbin.org/get")
	// 获取网页内容
	pg, _ := wd.PageSource()
	// 输出网页内容
	fmt.Println(pg)
}