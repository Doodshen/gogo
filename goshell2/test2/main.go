/*
 * @Descripttion: 
 * @version: 
 * @Author: wzs
 * @Date: 2023-05-20 15:40:52
 * @LastEditors: Andy
 * @LastEditTime: 2023-05-20 15:40:57
 */
package main

import (
    "github.com/PuerkitoBio/goquery"
    "log"
    "fmt"
    "strconv"
    "time"
)


func main() {
    url := "https://www.zhipin.com/c101010100/?query=Go&page="
    t := time.Now()
    fmt.Println("============== 千锋教育Go语言开发教学部 职位信息分析 ================")
    for offset := 0; offset < 10; offset++ {
        time.Sleep(1 * time.Second)
        doc, err := goquery.NewDocument(url + strconv.Itoa(offset))
        handleErr(err)
        fmt.Printf("第 %d 页的数据：\n", offset)
        doc.Find(".job-primary").Each(func(i int, selection *goquery.Selection) {
            item := Item{}
            fmt.Printf("职位序号：第%d个职位\n", (i + 1))
            item.position_name = selection.Find("div .job-title").Text()
            fmt.Printf("职位名称：%s\n", item.position_name)
            item.position_salary = selection.Find("div .red").Text()
            fmt.Printf("职位薪酬：%s\n", item.position_salary)
            item.work_address = selection.Find(".info-primary p").Children().Nodes[0].PrevSibling.Data
            fmt.Printf("工作地点：%s\n", item.work_address)
            item.work_experience = selection.Find(".info-primary p").Children().Nodes[0].NextSibling.Data
            fmt.Printf("职位所需工作经历：%s\n", item.work_experience)
            item.education = selection.Find(".info-primary p").Children().Nodes[1].NextSibling.Data
            fmt.Printf("学历要求：%s\n", item.education)
            item.company_name = selection.Find(".company-text .name").Children().First().Text()
            fmt.Printf("公司名称：%s\n", item.position_name)
            item.company_type = selection.Find(".company-text p").Children().Nodes[0].PrevSibling.Data
            fmt.Printf("公司类型：%s\n", item.company_type )
            if selection.Find(".company-text p").Children().Size() == 2 {
                item.company_development_stage = selection.Find(".company-text p").Children().Nodes[0].NextSibling.Data
                fmt.Printf("公司发展阶段：%s\n", item.company_development_stage)
                item.company_size = selection.Find(".company-text p").Children().Nodes[1].NextSibling.Data
                fmt.Printf("公司规模：%s\n", item.company_size )
            } else if selection.Find(".company-text p").Children().Size() == 1 {
                item.company_size = selection.Find(".company-text p").Children().Nodes[0].NextSibling.Data
                fmt.Printf("公司规模：%s\n", item.company_size)
            }
            fmt.Println("================================================================\n")
        })
    }
    elapsed := time.Since(t)
    fmt.Println("app elapsed:", elapsed)
}
type Item struct {
    // 职位名称
    position_name string
    // 职位薪酬
    position_salary string
    //工作地点
    work_address string
    // 职位所需工作经历
    work_experience string
    // 学历要求
    education string
    // 公司名称
    company_name string
    // 公司类型
    company_type string
    // 公司发展阶段
    company_development_stage string
    //公司规模
    company_size string
}
func handleErr(err error) {
    if err != nil {
        log.Fatal(err)
    }
}