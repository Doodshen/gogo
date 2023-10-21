/*
 * @Descripttion:
 * @version:
 * @Author: wzs
 * @Date: 2023-05-19 21:58:35
 * @LastEditors: Andy
 * @LastEditTime: 2023-05-19 22:29:36
 */
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	url := "https://movie.douban.com/chart"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Cookie", "bid=x2Sv-yIDV2k")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
