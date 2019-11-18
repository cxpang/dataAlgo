package main

import (
	"fmt"
	"sync"
)

//tree 的深度优先与广度优先搜索
func main() {
	tree:= BinaryTree{nil,sync.RWMutex{}}
	tree.Insert(8,"a")
	tree.Insert(4,"b")
	tree.Insert(2,"c")
	tree.Insert(1,"d")
	tree.Insert(3,"e")
	tree.Insert(6,"f")
	tree.Insert(5,"g")
	tree.Insert(7,"h")
	tree.Insert(10,"i")
	tree.Insert(9,"j")
	tree.Insert(11,"k")
	tree.PreOrderTraverse()

	//进行广度优先搜索
	tree.BfsTree()
}

func (tree *BinaryTree) BfsTree()  {
	tree.lock.Lock()
	defer tree.lock.Unlock()
	Bfs(tree.root)
	fmt.Println("层序遍历 广度")
	PrintLevel(tree.root)
	fmt.Println(result)

	Dfslevel(tree.root)
	fmt.Println("层序遍历 深度")
	fmt.Println(Dresult)
}



/**
 *广度优先搜索
***/
func Bfs(node *Node){
	if node ==nil{
		return 
	}
	var queue []*Node
	queue=append(queue, node)

	for len(queue) > 0 {
		node:=queue[0]
		queue = queue[1:]
		fmt.Println(node.key)
		if node.left!=nil{
			queue = append(queue, node.left)
		}
		if node.right!=nil{
			queue = append(queue,node.right)
		}

	}

}
/*
*打印二叉树层次
**/
var result [][]interface{}
func PrintLevel(node *Node)[][]interface{}  {

	if node==nil{
		return  result
	}

	bfsHelper(node,0)
	return result
}
func bfsHelper(node *Node,level int)  {
	if node==nil{
		return
	}
	if len(result) < level+1{
		result=append(result,make([]interface{},0))
	}
	result[level] = append(result[level],node.val)
	if node.left!=nil{
		bfsHelper(node.left,level+1)
	}
	if node.right!=nil{
		bfsHelper(node.right,level+1)
	}
}

/**
 *深度优先搜索
***/
var Dresult [][]interface{}
func Dfslevel(node *Node)[][]interface{}  {
	if node==nil{
		return Dresult
	}
	Dfs(node,0)
	return Dresult
}
func Dfs(node *Node,level int)  {
	if node ==nil {
		return
	}
	if len(Dresult)<level+1{
		Dresult = append(Dresult,[]interface{}{node.val})
	}else {
		Dresult[level] = append(Dresult[level],node.val)
	}
	Dfs(node.left,level+1)
	Dfs(node.right,level+1)
}