package main

import (
	"fmt"
	"sync"
)

type Item interface {
}
type Queue struct {
	items []Item
	lock sync.RWMutex
}

func NewQueue() *Queue  {
	 q := &Queue{}
	 q.items = []Item{}
	 return  q
}
func (q *Queue) Enqueue(t Item) {
	q.lock.Lock()
	 q.items = append(q.items, t)
	defer q.lock.Unlock()
}

func (q *Queue) Dequeue() *Item {
	q.lock.Lock()
	if len(q.items)==0{
		return nil
	}
	item := q.items[0]
	q.items = q.items[1:len(q.items)]
	defer q.lock.Unlock()
	return &item
}


func main() {
	que:=NewQueue()
	fmt.Println(que.items)
	que.Enqueue("pangchnxu")
	que.Enqueue("lichen")
	a:=que.Dequeue()
	fmt.Println(*a)
	b:=que.Dequeue()
	fmt.Println(*b)
	c:=que.Dequeue()
	fmt.Println(c)
}
