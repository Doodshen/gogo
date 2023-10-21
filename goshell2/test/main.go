package main

import (
    "fmt"
    "io"
    "log"
    "net/http"
    "regexp"
    "strconv"
    "time"

    "gorm.io/driver/mysql"
    "github.com/PuerkitoBio/goquery"
    "gorm.io/gorm"
)

type Movie struct {
    gorm.Model
    Title        string  `json:"title" gorm:"column:title"`
    PublicYear   string  `json:"public_year" gorm:"column:public_year"`
    Score        float64 `json:"score" gorm:"column:score"`
    CommentCount int64   `json:"comment_count" gorm:"column:comment_count"`
    Quote        string  `json:"quote" gorm:"column:quote"`
}

func NewMovie(title, publicYear, quote string, score float64, commentCount int64) *Movie {
    return &Movie{
        Title:        title,
        PublicYear:   publicYear,
        Score:        score,
        CommentCount: commentCount,
        Quote:        quote,
    }
}

func (m *Movie) TableName() string {
    return "movie"
}

var (
    db     *gorm.DB
    movies []*Movie
)
 //初始化数据库 使用gorm 
func init() {
    dsn := "root:abc123@tcp(127.0.0.1:3306)/movies?charset=utf8mb4&parseTime=True&loc=Local"
    d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
            SkipDefaultTransaction: true, // 关闭gorm默认开启的全局事务
            PrepareStmt:            true, // 开启每次执行SQL会预处理SQL
    })
    if err != nil {
            log.Println("连接数据库失败")
            return
    }
    db = d
    db.AutoMigrate(&Movie{}) // 自动同步表
}

//正则表达式 
func ClearPlain(str string) string {
    reg := regexp.MustCompile(`\s`)
    return reg.ReplaceAllString(str, "")
}

func GetNumber(str string) string {
    reg := regexp.MustCompile(`\d+`)
    return reg.FindString(str)
}

//启动程序 
func Run(method, url string, body io.Reader, client *http.Client) {
    //发送请求 获取请求对象 
    req, err := http.NewRequest(method, url, body)
    if err != nil {
        log.Println("获取请求对象失败")
        return
    }
    //设置请求对象 
    req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36")
    req.Header.Set("host", "movie.douban.com")

    //发送请求 
    resp, err := client.Do(req)
    if err != nil {
        log.Println("发起请求失败")
        return
    }
    //校对请求码
    if resp.StatusCode != http.StatusOK {
        log.Printf("请求失败，状态码：%d", resp.StatusCode)
        return
    }

    // 关闭响应对象中的body
    defer resp.Body.Close() 


    //清洗数据 
    query, err := goquery.NewDocumentFromReader(resp.Body)
    if err != nil {
        log.Println("生成goQuery对象失败")
        return
    }
    query.Find("ol.grid_view li").Each(func(i int, s *goquery.Selection) {
        title := ClearPlain(s.Find("span.title").Text())
        year := ClearPlain(GetNumber(s.Find("div.bd>p").Text()))
        commentCountStr := ClearPlain(GetNumber(s.Find(".star>span").Eq(3).Text()))
        scoreStr := ClearPlain(s.Find("span.rating_num").Text())
        quote := ClearPlain(s.Find(".inq").Text())
        commentCount, _ := strconv.ParseInt(commentCountStr, 10, 64)
        score, _ := strconv.ParseFloat(scoreStr, 64) // 评分可能是小数，所以这里用的是ParseFloat方法
        fmt.Println(title)
        fmt.Println(year)
        fmt.Println(commentCount)
        fmt.Println(score)
        fmt.Println(quote)
        movies = append(movies, NewMovie(title, year, quote, score, commentCount))
        //time.Sleep(time.Second)
        fmt.Println("-------------------------")
    })
}

func main() {
    // 需求：
    // 1. 爬取片名、爬取年份、爬取评价人数、爬取评分、爬取描述。其他的信息大家自己解析
    // 2. 爬取到的数据入mysql数据库
    // 3. 起一个web服务，暴露接口直接获取到电影数据【分页处理】
    client := &http.Client{}
    url := "https://movie.douban.com/top250?start=%d&filter="
    method := "GET"
    // 数据爬取操作
    for i := 1; i <= 10; i++ {
        Run(method, fmt.Sprintf(url, i*25), nil, client)
        time.Sleep(time.Second * 2) // 主动等待下吧，别被ban了
    }

    // 数据入库操作
    if err := db.Create(movies).Error; err != nil {
        log.Println("插入数据失败", err.Error())
     	return
    }
    log.Println("插入数据成功")
}

