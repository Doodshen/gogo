package main
import (
	"encoding/json"
	"fmt"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
	"os"
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
	urlPrefix := fmt.Sprintf("http://127.0.0.1:%d/wd/hub",port)
	wd, _ := selenium.NewRemote(caps, urlPrefix)
	// 关闭浏览器对象
	defer wd.Quit()


	// 访问网址
	wd.Get("https://www.baidu.com/s?wd=go")
	time.Sleep(3 * time.Second)


	// 获取所有Cookie
	// Cookies以结构体Cookie表示，可以转化为json存放在json文件中
	gc, _ := wd.GetCookies()
	for _, k := range gc {
		 fmt.Printf("Cookie信息：%v\n", k)
	}


	// 将Cookie写入JSON文件
	model := os.O_RDWR|os.O_CREATE|os.O_TRUNC
	f2, _ := os.OpenFile("output.json", model, 0755)
	encoder := json.NewEncoder(f2)
	// 将变量gc的数据写入JSON文件
	// 数据写入必须使用文件内容覆盖，即设置os.O_TRUNC模式，否则导致内容错乱
	encoder.Encode(gc)
	f2.Close()


	// 删除所有Cookies
	wd.DeleteAllCookies()


	// 添加Cookies信息
	// 如果Cookies存放在json文件中
	// 读取json文件并转换为结构体Cookie
	// 最后将结构体Cookie添加到浏览器对象中
	f1,_:=os.OpenFile("output.json",os.O_RDWR|os.O_CREATE,0755)
	// 定义结构体类型的切片
	var myCookies []selenium.Cookie
	// 实例化结构体Decoder，实现数据读取
	data := json.NewDecoder(f1)
	// 将已读取的数据加载到切片myCookies中
	data.Decode(&myCookies)
	f1.Close()


	// 读取json的Cookies信息
	fmt.Printf("读取json的Cookie信息：%v\n", myCookies)
	for _, k1 := range myCookies {
		 wd.AddCookie(&k1)
	}
}