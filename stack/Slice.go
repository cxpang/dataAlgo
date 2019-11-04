package main

import (
	"fmt"
	"sync"
)

type Item interface {
}

type  ItemStack struct {
	items []Item
	lock sync.RWMutex
}

func NewStack() *ItemStack  {
    s:=&ItemStack{}
    s.items = []Item{}
    return s
}

func (stack *ItemStack) print()  {
	fmt.Println(stack.items)
}
func (stack *ItemStack) Push(t Item) {
	stack.lock.Lock()
	defer stack.lock.Unlock()
	stack.items = append(stack.items, t)
}

func (stack *ItemStack) Pop() Item {
	stack.lock.Lock()
	defer stack.lock.Unlock()
	if len(stack.items) == 0 {
		return nil
	}
	item := stack.items[len(stack.items)-1]
	stack.items = stack.items[0 : len(stack.items)-1]
	return item
}



func main() {
	stack := NewStack()
	stack.Push("cxpang")
	stack.Push("李晨")
	stack.print()
	a:=stack.Pop()
	fmt.Println(a)
	stack.print()
	stack.Pop()
	b:=stack.Pop()
	fmt.Println(b)
}
