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


/**
 *深度优先搜索
***/
func Dfs()  {

}