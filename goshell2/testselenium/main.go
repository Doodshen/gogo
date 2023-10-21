package main

import (
	"fmt"
	"log"
	"time"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	//chromedriver路径 我们设置成绝对路径
	chromeDriver = "C://Program Files//Google//Chrome//Application//chromedriver.exe"

	//chromeDriver运行端口
	port = 8080
)

type Job struct {
	Name   string `json:"name" gorm:"column:name"`
	Area   string `json:"area" gorm:"column:area"`
	Pays   string `json:"pays" gorm:"column:pays"`
	Exp    string `json:"exp" gorm:"column:exp"`
	Tags    string `json:"tags" gorm:"column:tags"`
	Desc   string `json:"desc" gorm:"column:desc"`
	Publis string `json:"publis" gorm:"column:publis"`
	Cmp    string `json:"cmp" gorm:"column:cmp"`
	Scale  string `json:"scale" gorm:"column:scale"`
}

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



//获取当前页数的所有职位信息
func get_jobs(wd selenium.WebDriver) []Job{
	var jobs []Job
	jf,_ := wd.FindElements(selenium.ByClassName,"job-primary")
	for _,v := range jf{
		j := Job{}
		//获取职位名称
		name, _ := v.FindElement(selenium.ByClassName,"job-name")
		j.Name,_ = name.Text()


		//获取地点 
		area,_ := v.FindElement(selenium.ByClassName,"job-area")
		j.Area,_ = area.Text()

		//获取薪资
		pays,_ := v.FindElement(selenium.ByClassName,"red")
		j.Pays,_ = pays.Text()

		//获取经验学历
		exp,_ := v.FindElement(selenium.ByCSSSelector,`[class]="job-limit clearfix"]>p`)
		j.Exp,_ = exp.Text()
		//获取职位标签
		tags,_ := v.FindElement(selenium.ByClassName,"tags")
		j.Tags,_ = tags.Text()
		//获取工公司福利
		desc,_ := v.FindElement(selenium.ByClassName,"info-desc")
		j.Desc,_ = desc.Text()
		//获取公司人事信息
		publis,_ := v.FindElement(selenium.ByClassName,"info-publis")
		j.Publis,_ = publis.Text()
		//获取公司名称
		cmp,_ := v.FindElement(selenium.ByCSSSelector,`[class]="company-text"]>h3`)
		j.Cmp,_ = cmp.Text()
		//获取公式行业和规模
		scale,_ := v.FindElement(selenium.ByCSSSelector,`[class]="company-text"]>p`)
		j.Scale,_ = scale.Text()

		jobs = append(jobs, j)
	}
	return jobs
}

var (
    db     *gorm.DB
    job []*Job
)


func init() {
    dsn := "root:abc123@tcp(127.0.0.1:3306)/jobs?charset=utf8mb4&parseTime=True&loc=Local"
    d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
            SkipDefaultTransaction: true, // 关闭gorm默认开启的全局事务
            PrepareStmt:            true, // 开启每次执行SQL会预处理SQL
    })
    if err != nil {
            log.Println("连接数据库失败")
            return
    }
    db = d
    db.AutoMigrate(&Job{}) // 自动同步表
}



//保存数据 
func sava_data(jobs []Job){
	// 数据入库操作
    if err := db.Create(jobs).Error; err != nil {
        log.Println("插入数据失败", err.Error())
     	return
    }
    log.Println("插入数据成功")
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
	search,_:= wd.FindElement(selenium.ByCSSSelector,`[class="btn btn-search"]`)
	time.Sleep(5*time.Second)
	r1,_:= search.IsEnabled()
	fmt.Printf("判断网页元素是否可编辑或可点击状态：%v\n", r1)


	err := search.Click()
	fmt.Println(err)

	time.Sleep(2*time.Second)  

	//获取第一页的职位信息
	jobs := get_jobs(wd)
	

	//使用死循环1实现翻页
	for{
		np,err := wd.FindElement(selenium.ByCSSSelector,`[class="page"]>[class="next"]`)

		//err不等于nil说明无法单击下一页，终止死循环
		if err != nil{
			break
		}else{
			//单击下一页
			np.Click()
			time.Sleep(2*time.Second)
			
			//获取当前页职位信息
			//将当前页所有职位合并到切片jobs
			jobs = append(jobs, get_jobs(wd)...)
		}

		//保存数据
		sava_data(jobs)
	}
}