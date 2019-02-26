package main

import (
	"fmt"

	arrstack "zachrey.win/demo/stack/arrStack"
	broswer "zachrey.win/demo/stack/broswer"
	linkstack "zachrey.win/demo/stack/linkStack"
)

func main() {
	var s arrstack.Stack
	s.Push(1)
	s.Push(2)
	s.Push(3)
	n := s.Push(4)

	fmt.Println(n)

	s.Pop()
	s.Pop()
	s.Pop()
	s.Pop()

	val, ok := s.Pop()
	fmt.Println(val, ok)

	fmt.Println(s.GetStackVal())

	var s1 linkstack.LinkStack
	s1.Push(1)
	s1.Push(2)
	s1.Push(3)
	n = s1.Push(4)
	fmt.Println(n)

	s1.Pop()
	s1.Pop()
	s1.Pop()
	s1.Pop()

	val, ok = s1.Pop()
	fmt.Println(val, ok)

	fmt.Println(s1.GetStackAll())

	var b broswer.Browser
	b.Push("A")
	b.Push("B")
	b.Push("C")

	pageName := b.Goback()
	fmt.Println(pageName)

	pageName = b.Goback()
	fmt.Println(pageName)

	pageName = b.Forward()
	fmt.Println(pageName)

	// pageName = b.Forward()
	// fmt.Println(pageName)

	b.Push("D")
	pageName = b.Forward()
	fmt.Println(pageName)
}
