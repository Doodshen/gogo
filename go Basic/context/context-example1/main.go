/*
 * @Descripttion:
 * @version:
 * @Author: wzs
 * @Date: 2023-07-22 10:12:43
 * @LastEditors: Andy
 * @LastEditTime: 2023-07-22 10:21:54
 */
package main

import (
	"context"
	"fmt"
)

func f3(ctx context.Context, req any) {
	fmt.Println(ctx.Value("key0"))
	fmt.Println(ctx.Value("key1"))
	fmt.Println(ctx.Value("key2"))
}

func f2(ctx context.Context, req any) {
	ctx2 := context.WithValue(ctx, "key2", "value2")
	f3(ctx2, req)
}

func f1(ctx context.Context, req any) {
	ctx1 := context.WithValue(ctx, "key1", "value1")
	f2(ctx1, req)
}

func handle(ctx context.Context, req any) {
	ctx0 := context.WithValue(ctx, "key0", "value0")
	f1(ctx0, req)
}

func main() {
	rootCtx := context.Background()
	handle(rootCtx, "hello")
}
