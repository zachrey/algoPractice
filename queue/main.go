package main

import (
	"fmt"

	"zachrey.win/demo/queue/arrqueue"
	"zachrey.win/demo/queue/linkqueue"
)

func main() {
	var aq arrqueue.ArrQueue
	aq.Push(1)
	aq.Push(2)
	aq.Push(3)
	n := aq.Push(4)
	fmt.Println(n)

	val, n := aq.Pop()
	fmt.Println(val, n)

	fmt.Println(aq.GetAll())
	fmt.Println("==============================")

	var lq linkqueue.LinkQueue
	lq.Push(1)
	lq.Push(2)
	lq.Push(3)
	n = lq.Push(4)
	fmt.Println(n)

	val, n, _ = lq.Pop()
	fmt.Println(val, n)

	fmt.Println(lq.GetAll())
	fmt.Println("==============================")

}
