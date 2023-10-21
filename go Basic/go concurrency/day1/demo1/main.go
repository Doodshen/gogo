package main

import (
	"fmt"
	"sync"
)


func main(){
	//互斥锁保护计数器 
	var mu sync.Mutex

	var count = 0 
	//辅助变量 使用WaitGroup等待10个goroutine完成 
	var wg sync.WaitGroup
	wg.Add(10)

	//启动10个goroutine 
	for i:= 0;i<10;i++{
		go func(){
			defer wg.Done()
			//对变量count指向10次加一
			for j:= 0;j<100000;j++{
				mu.Lock()
				count++
				mu.Unlock()
			}
		}()
	}
	//等待10个goroutine完成 
	wg.Wait()
	fmt.Println(count)

}
