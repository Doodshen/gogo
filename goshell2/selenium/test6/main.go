package main
import (
	"fmt"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
	"time"
)
func main() {
	// ChromeDriver路径信息
	chromeDriver := "C:\\Program Files\\Google\\Chrome\\Application\\chromedriver.exe"
	// ChromeDriver运行端口
	port := 8080


	/* 开启WebDriver服务 */
	s, _ := selenium.NewChromeDriverService(chromeDriver,port)
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
	url:="https://zhidao.baidu.com/question/496886588283871492.html"
	wd.Get(url)



	/* 切换iframe标签 */
	// 在iframe标签中嵌套HTML网页，不同HTML代码需要SwitchFrame切换
	// 通过iframe标签的id属性切换
	wd.SwitchFrame("ueditor_0")
	// 定位iframe标签中HTML网页的网页元素p
	e1, _ := wd.FindElement(selenium.ByTagName, "p")
	e1.SendKeys("aaa")

	// 若参数frame设为nil，则切换主HTML
	wd.SwitchFrame(nil)


	/* 切换浏览器多标签页 */
	// 获取浏览器所有标签页，以切片格式返回

	ss, _ := wd.WindowHandles()
	fmt.Printf("获取浏览器所有标签页：%v\n", ss)


	// 获取浏览器当前正在显示的标签页
	cs, _ := wd.CurrentWindowHandle()
	fmt.Printf("获取浏览器当前正在显示的标签页：%v\n", cs)


	// 切换标签页，ss[len(ss)-1]获取切片ss某个元素
	wd.SwitchWindow(ss[len(ss)-1])

	
	// 获取浏览器当前正在显示的标签页
	cs1, _ := wd.CurrentWindowHandle()
	fmt.Printf("获取浏览器当前正在显示的标签页：%v\n", cs1)
	time.Sleep(3 * time.Second)
}