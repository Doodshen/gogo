package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    // 生成随机数种子
    rand.Seed(time.Now().UnixNano())

    // 定义一个map用于存储已抽取的数字
    nums := make(map[int]bool)

    // 抽取10个不重复随机数
    for i := 0; i < 10; {
        num := rand.Intn(24) + 2
        if !nums[num] {
            nums[num] = true
            i++
            fmt.Println("抽取到的号码",num)
        }
    }
}
