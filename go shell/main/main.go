/*
 * @Descripttion:
 * @version:
 * @Author: wzs
 * @Date: 2023-05-19 18:50:29
 * @LastEditors: Andy
 * @LastEditTime: 2023-05-20 11:22:46
 */
package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 定义结构体Movie 映射数据表movies
type Movie struct {
	//定义数据表字段
	gorm.Model
	Name        string `gorm:"type:varchar(60)"`
	Star        string `gorm:"type:varchar(60)"`
	Releasetime string `gorm:"type:varchar(60)"`
	Score       string `gorm:"type:varchar(60)"`
}

func get_data() string {
	urls := "https://www.maoyan.com/board/1?"
	fmt.Print(urls)

	// 定义请求对象NewRequest
	req, _ := http.NewRequest("GET", urls, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0(Windows NT 10.0; Win64; x64) AppleWebKit/537.36(KHTML, like Gecko) Chrome/94.0.4606.81 Safari/537.36")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Host", "www.maoyan.com")

	// 在Client设置参数Transport即可实现代理IP
	transport := &http.Transport{}
	client := &http.Client{Transport: transport}

	// 发送HTTP请求
	resp, _ := client.Do(req)
	if resp.StatusCode != http.StatusOK {
		log.Printf("请求失败，状态码：%d", resp.StatusCode)
	}
	// 获取网站响应内容
	body, _ := ioutil.ReadAll(resp.Body)
	// 网页响应内容转码
	result := string(body)
	// 设置延时，请求过快会引发反爬
	time.Sleep(5 * time.Second)
	return result
}

func clean_data(data string) []map[string]string {
	// 使用goquery解析HTML代码
	dom, _ := goquery.NewDocumentFromReader(strings.NewReader(data))
	// 定义变量result和info
	result := []map[string]string{}
	var info map[string]string
	// 遍历网页所有电影信息
	selection := dom.Find(".board-item-content")
	selection.Each(func(i int, selection *goquery.Selection) {
		// 记录每部电影信息，每存储一部电影必须清空集合
		info = map[string]string{}
		name := selection.Find(".name").Text()
		star := selection.Find(".star").Text()
		releasetime := selection.Find(".releasetime").Text()
		score := selection.Find(".score").Text()
		info["name"] = strings.TrimSpace(name)
		info["star"] = strings.TrimSpace(star)
		info["releasetime"] = strings.TrimSpace(releasetime)
		info["score"] = strings.TrimSpace(score)
		// 将电影信息写入切片
		result = append(result, info)
	})
	return result
}

func clean(data string) {
	// 使用goquery解析HTML代码
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(data))
	// 获取样式为"board-item-main"的div元素
	doc.Find(".board-item-main").Each(func(i int, s *goquery.Selection) {
		// 查找样式为"name"的p元素并获取其文本内容
		name := s.Find("p.name").Text()
		// 查找样式为"star"的p元素并获取其文本内容
		star := s.Find("p.star").Text()
		// 查找样式为"releasetime"的p元素并获取其文本内容
		releasetime := s.Find("p.releasetime").Text()
		// 查找样式为"score"的p元素并获取其文本内容
		score := s.Find("p.score").Text()

		// 输出结果
		fmt.Printf("Name: %s\nStar: %s\nRelease Time: %s\nScore: %s\n", name, star, releasetime, score)
	})
}

// 获取数据库连接
func sava_data(data []map[string]string) {
	//连接数据库
	dsn := `root:abc123@tcp(127.0.0.1:3306)/movie?parseTime=true`

	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	sqlDb, _ := db.DB()

	//关闭数据库 释放资源
	defer sqlDb.Close()

	//执行数据迁移
	db.AutoMigrate(&Movie{})

	//遍历变量data,获取每部电影信息
	for _, k := range data {
		fmt.Printf("当前数据：%v\n", k)
		//查找电影书否已经在数据库
		var m []Movie
		db.Where("name = ?", k["name"]).First(&m)

		//len(m)等于0说明数据不存在数据库
		if len(m) == 0 {
			//新增加数据
			m1 := Movie{
				Name:        k["name"],
				Star:        k["star"],
				Releasetime: k["releasetime"],
				Score:       k["score"],
			}
			db.Create(&m1)
		} else {
			//更新数据
			db.Where("name = ?", k["name"]).Find(&m).Update("Score", k["score"])
		}
	}
}

func sava_data2(data []map[string]string) {

	db, err := sql.Open("mysql", "root:abc123@(127.0.0.1:3306)/movie?parseTime=true")
	if err != nil {
		log.Println("连接数据库异常")
		panic(err)
	}

	//最大空闲数连接数 默认不配置 是2个最大空闲连接
	db.SetMaxIdleConns(5)
	//最大连接数 默认不配置 是不限制最大连接数
	db.SetMaxOpenConns(100)
	//连接最大存活实践
	db.SetConnMaxIdleTime(time.Minute * 3)
	//空闲连接最大存活时间
	db.SetConnMaxIdleTime(time.Minute * 1)
	err = db.Ping()
	if err != nil {
		log.Println("数据库无法连接")
		_ = db.Close()
		panic(err)
	}

	for _, k := range data {
		fmt.Printf("当前数据：%v\n", k)
		db.Exec("insert into movie"+
			"(name,star,releasetime,score)"+
			"values(?,?,?,?)",
			k["name"],
			k["star"],
			k["releasetime"],
			k["score"],
		)

	}

}

func main() {
	//循环十次 每次遍历代表不同页得网页信息

	//函数调用
	//调用次序：发起http请求-》清洗数据->数据入库
	webdata := get_data()
	clean(webdata)
}
