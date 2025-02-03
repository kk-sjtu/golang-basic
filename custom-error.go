package main

import(
	"errors"
	"fmt"
)

type argError struct{
	arg int
	message string
}
//虽然没有直接调用Error()，但是在errors.As()中会调用Error()方法
func (e *argError) Error() string{
	return fmt.Sprintf("%d - %s",e.arg,e.message)
}

func f(arg int)(int,error){
	if arg == 42{
		return -1,&argError{arg,"can't work with it"}
	}
	return arg +3,nil
}

func main(){
	/*
	_,err := f(42)
	var ae *argError
	if errors.As(err,&ae){
		fmt.Println(ae.arg)
		fmt.Println(ae.message)
	}else{
		fmt.Println("err doesn't match argError")
	}
*/
	_,errc := f(45)
	var ace *argError
	if errors.As(errc,&ace){
		fmt.Println(ace.arg)
		fmt.Println(ace.message)
	}else{
		fmt.Println("err doesn't match argError")
	}
}