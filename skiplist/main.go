package main

import (
	"fmt"

	"github.com/zachrey/algoPractice/skiplist/skiplist"
)

func main() {
	sl := skiplist.NewList()
	sl.Add(1, "111")
	sl.Add(2, "222")
	sl.Add(3, "333")
	sl.Add(5, "555")

	fmt.Println(sl.GetListData())

	for _, v := range sl.Header.ForWard {
		fmt.Println(v)
	}
}
