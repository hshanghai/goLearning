package main

import (
	"./queue"
	"fmt"
)

func main()  {
	q := queue.Queue{1}
	q.Push(2)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
}