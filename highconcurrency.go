package main


/* 1. 按顺序运行
func main(){
	count(5,"🐑")
	count(5,"🐂")
}

func count(n int,animal string){
	for i:=0;i<n;i++{
		fmt.Println(i+1,animal)
		time.Sleep(time.Millisecond*500)
	}
}
*/

/* 2. go
func main(){
	go count(5,"🐑")
	count(5,"🐂")
}

func count(n int,animal string){
	for i:=0;i<n;i++{
		fmt.Println(i+1,animal)
		time.Sleep(time.Millisecond*500)
	}
}
*/

/* 3. waitgroup 
import(
	"fmt"
	"time"
	"sync"
)
func main(){
	var wg sync.WaitGroup
	wg.Add(2)
	go func(){
		count(5,"🐑")
		wg.Done()
		}()
	func(){
		go count(5,"🐂")
		wg.Done()
	}()
	wg.Wait()
}

func count(n int,animal string){
	for i:=0;i<n;i++{
		fmt.Println(i+1,animal)
		time.Sleep(time.Millisecond*500)
	}
}
*/

/* 
channel
import(
	"fmt"
	"time"
)
func main(){

	c := make(chan string)

	count(5,"🐑",c)
	for{
		message,open := <-c
		if !open{
			break
		}
		fmt.Println(message)
	}
	
	for message := range c{
		fmt.Println(message)
	}
		
	
}

func count(n int,animal string,c chan string){
	for i:=0;i<n;i++{
		c<-animal
		time.Sleep(time.Millisecond*500)
	}
	close(c)
}
*/


import(
	"fmt"
	"time"
)

func main(){

	c1 := make(chan string)
	c2 := make(chan string)

	go func(){
		for{
			c1 <- "🐑"
			time.Sleep(time.Millisecond*500)
		}
	}()

	go func(){
		for{
			c2 <- "🐂"
			time.Sleep(time.Millisecond*2000)
		}
	}()

	for {
		select{
		case msg := <-c1:
			fmt.Println(msg)
		
		case msg := <-c2:
			fmt.Println(msg)
	}
	}}