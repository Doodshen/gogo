package main

import(
	"net/http"
	"fmt"
)

func main(){
	http.HandleFunc("/",hello)
	http.ListenAndServe("8080",nil)
}

func  hello(w http.ResponseWriter,r *http.Request)  {
	fmt.Println("hello")
}