package main

import (
	"fmt"
	"iter"
	"slices"
)

// 定义了一个列表
type List[T any] struct{
	head,tail *element[T]
}

// 定义了列表的元素
type element[T any] struct{
	next *element[T]
	val T
}

// 定义了链表的方法，如果空就接在首节点后面，如果不空，就把
func(lst *List[T]) Push(v T){
	if lst.tail == nil{
		lst.head = &element[T]{val:v}
		lst.tail = lst.head
	}else{
		lst.tail.next = &element[T]{val:v}
		lst.tail = lst.tail.next
	}
}
//
func (lst *List[T]) All() iter.Seq[T]{
	return func(yield func(T) bool){
		for e:= lst.head;e!=nil;e=e.next{
			if !yield(e.val){
				return
			}
		}
	}
}

func genFib() iter.Seq[int]{
	return func(yield func(int)bool){
		a,b := 1,1
		for{
			if !yield(a){
				return
			}
			a,b = b,a+b
		}
	}
}

func main(){
	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)

	for e := range lst.All(){
		fmt.Println(e)
	}

	all := slices.Collect(lst.All())
	fmt.Println("all:",all)

	for n:= range genFib(){
		if n>=10{
			break
		}
		fmt.Println(n)
	}
}

