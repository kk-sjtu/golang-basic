package main

import (
	"fmt"
	"time"
)
func hello(){
	fmt.Println("Hello from goroutine")
	time.Sleep(1 * time.Second)
}

func main(){
	for i := 0; i < 10; i++ {
		time.Sleep(2 * time.Second)
		go hello()
	}
}