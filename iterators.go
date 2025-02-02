package main

import (
	"fmt"
	"iter"
	"slices"
)

type List[T any] struct{
	head,tail *element[T]
}

type element[T any] struct{
	next *element[T]
	val T
}

func(lst *List[T]) Push(v T){
	if lst.tail == nil{
		lst.head = &element[T]{val:v}
		lst.tail = lst.head
	}else{
		lst.tail.next = &element[T]{val:v}
		lst.tail = lst.tail.next
	}
}

func (lst *List[T]) All() iter.Seq[T]{
	return func()
}