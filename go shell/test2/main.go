package main
import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/axgle/mahonia"
	"io/ioutil"
	"net/http"
	"strings"
)






// 使用第三方mahonia实现网页内容的转码
func ConvertToString(src, srcCode, tagCode string) string {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult),true)
	result := string(cdata)
	return result
}


func main() {
	urls := "https://search.51job.com/list/030200,000000,0000,00,9,99,golang,2,1.html"


	// 定义请求对象NewRequest
	req, _ := http.NewRequest("GET",urls, nil)
	transport := &http.Transport{}


	// 在Client中设置参数Transport即可实现代理IP
	client := &http.Client{Transport: transport}

	// 发送HTTP请求
	resp, _ := client.Do(req)

	// 获取网站响应内容
	body, _ := ioutil.ReadAll(resp.Body)

	// 网页响应内容转码
	result := ConvertToString(string(body), "gbk", "utf-8")
	// 使用第三方包goquery读取HTML代码，读取方式有多种
	// NewDocumentFromReader：读取字符串的HTML代码
	// NewDocumentFromResponse：读取HTML对象，即net/http的resp.Body
	// NewDocument：从网址中直接读取HTML代码
	dom,_:=goquery.NewDocumentFromReader(strings.NewReader(result))


	// Find()是查找HTML里面所有符合要求的标签
	// 如果查找Class="ht"的标签，则使用Find(".ht")
	// 如果查找id="ht"的标签，则使用Find("#ht")
	// 多个标签使用同一个Class，如div和p标签使用Class="ht"
	// 若只需div标签，则使用Find("div[class=ht]")
	dom.Find(".ht").Each(func(i int,selection *goquery.Selection){
		 v := strings.TrimSpace(selection.Text())
		 fmt.Printf("查找Class=ht的标签：%v\n", v)
	})

	
	// 通过多层HTML标签查找，只需在Find里面设置多层标签的Class属性即可
	// 首先查找Class="rlk"的标签
	// 然后在Class="rlk"的标签中查找a标签
	// 因此查找方式为Find(".rlk a")，每个标签之间使用空格隔开
	dom.Find(".rlk a").Each(func(i int,selection *goquery.Selection){
		 // 获取数据
		 v1 := strings.TrimSpace(selection.Text())
		 fmt.Printf("当前数据：%v\n", v1)
		 // 获取数据所在的HTML代码
		 v2, _ := selection.Html()
		 fmt.Printf("获取数据所在HTML代码：%v\n", v2)
		 // 使用Attr获取标签的href属性
		 v3, _ := selection.Attr("href")
		 fmt.Printf("Attr()获取标签的href属性：%v\n", v3)
		 // 使用AttrOr获取标签的href属性
		 v4 := selection.AttrOr("href", "")
		 fmt.Printf("AttrOr()获取标签的href属性：%v\n", v4)
	})
}