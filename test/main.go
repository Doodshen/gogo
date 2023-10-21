/*
 * @Descripttion: 
 * @version: 
 * @Author: wzs
 * @Date: 2023-07-24 15:40:16
 * @LastEditors: Andy
 * @LastEditTime: 2023-07-24 15:40:35
 */

package main

import (
   "fmt"
   "net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
   fmt.Fprintf(w, "Hello World")
}

func main () {
   http.HandleFunc("/", HelloHandler)
   http.ListenAndServe(":8000", nil)
}