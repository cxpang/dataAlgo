package main

import (
	"fmt"
)

//双向链表

type ListNode struct {
	prev  *ListNode // 前一个节点
	next  *ListNode // 后一个节点
	value interface{}    // 数据
}

// 创建一个节点
func NewListNode(value interface{}) (listNode *ListNode) {
	listNode = &ListNode{
		value: value,
	}

	return
}
//获取节点值

// 获取节点的值
func (n *ListNode) GetValue() (value interface{}) {
	if n == nil {
		return
	}
	value = n.value
	return
}
// 链表
type List struct {
	head *ListNode // 表头节点
	tail *ListNode // 表尾节点
	len  int       // 链表的长度
}
// 创建一个空链表
func NewList() (list *List) {
	list = &List{
	}
	return
}
func (l *List)add(val interface{}) *List  {
	node:=NewListNode(val)
	if l.len == 0{
		l.head = node
		l.tail = node
	}else {
		//尾部插入
		tail:=l.tail
		tail.next = node
		node.prev = tail
		l.tail = node
	}
	l.len ++
    return  l
}

func (l *List)del() *ListNode {
	if l.len == 0{
		fmt.Println("empty node")
	}
	node:=l.head
	l.head = node.next
	l.len--
	fmt.Println(node.GetValue(),"出队列")
	return node
}
func (l *List)length()  {
	fmt.Println(l.len)
}
func main()  {
	list:=NewList()
	list.add("livhen")
	list.length()
	list.add("庞晨旭")
	list.length()
	list.del()
	list.del()
}