package main

import "fmt"

type Node struct {
	value int
	next *Node
}
type List struct {
	size uint64 // 车辆数量
	head *Node  // 车头
	tail *Node  // 车尾
}

func (list *List)InitLize()  {
	list.size = 0
	list.head = nil
	list.tail=nil
}
func (list *List)Add(node *Node)  {
	if list.size == 0{
		list.head = node
		list.tail =node
		list.size = 1
	}else {
		old:=list.tail
		old.next = node
		list.tail = node
		list.size++
	}

}
func (list *List)Foreach()  {
	if list.size == 0 {
		fmt.Println("false")
	}
	if list.tail != nil{
		fmt.Println(list.tail.value)
	}

}
func main()  {
	var list List //不能直接声明为地址
	list.InitLize()
	fmt.Println(list)

	lichen:=Node{12,nil}
	list.Add(&lichen)
	list.Foreach()
}
