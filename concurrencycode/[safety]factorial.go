package main

import (
	"fmt"
	"time"
)

var (
    testMap = make(map[int]int, 10)
)

func testNum(num int) {
    res := 1
    for i := 1; i <= num; i++ {
        res *= i
    }
    testMap[num] = res
}

func main() {
	
    start := time.Now()
    for i := 1; i < 200; i++ {
        go testNum(i)
    }
    // 协程需要在main之后完毕
    time.Sleep(time.Second * 5)
    for key, val := range testMap {
        fmt.Printf("数字%v 对应的阶乘是 %v\n", key, val)
    }
    end := time.Since(start)
    fmt.Println(end)
}