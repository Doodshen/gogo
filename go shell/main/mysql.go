/*
 * @Descripttion:
 * @version:
 * @Author: wzs
 * @Date: 2023-05-19 20:39:53
 * @LastEditors: Andy
 * @LastEditTime: 2023-10-18 17:10:09
 */
package main

import (
	"database/sql"
	"log"
	"time"
)

var DB *sql.DB //通过DB对象来操作数据库

func init() {

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
	DB = db
}
