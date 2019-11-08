package main

import (
	"fmt"
	"sync"
)

type Node struct {
	key int
	val interface{}
	left ,right *Node
}

type BinaryTree struct {
	root *Node
	lock sync.RWMutex
}

func NewNode(key int,data interface{}) *Node {
	return &Node{key,data,nil,nil}
}
func (tree *BinaryTree)Insert(key int ,data interface{})  {
	node:=NewNode(key,data)
	tree.lock.Lock()
	defer tree.lock.Unlock()
	if tree.root == nil {
		tree.root = node
	} else {
		//调用下面的插入函数，另起一个方法👇
		insertNode(tree.root, node)
	}
}
func insertNode(node,newNode *Node)  {
    if node.key>newNode.key{
    	//插入左子树
		if node.left == nil {
			node.left = newNode
		}else {
			insertNode(node.left,newNode)
		}
	}else {
		//插入左子树
		if node.right == nil {
			node.right = newNode
		}else {
			insertNode(node.right,newNode)
		}
	}
}

func (tree *BinaryTree)searchTree (key int) bool {
	tree.lock.Lock()
	defer tree.lock.Unlock()
	if tree.root ==nil {
		return false
	}
	return  search(tree.root,key)
}
func search(node *Node,key int) bool {
	if node == nil {
		return false
	}
	//如果key的值小于节点的值，那么应该插入到左子树
	if key < node.key {
		//将左子树作为新节点，继续查询
		return search(node.left, key)
	}

	// 如果key的值大于节点的值，那么应该插入右子树
	if key > node.key {
		//将右子树作为新节点，继续查询
		return search(node.right, key)
	}

	//如果当前key的值==node.key 返回true
	return true
}
func (tree *BinaryTree) deleteTree(key int) *Node  {
	tree.lock.Lock()
	defer tree.lock.Unlock()
	return delete(tree.root,key)
}
func delete(node *Node,key int) *Node  {
    if node==nil{
    	return  nil
	}

	if key < node.key {

		// 将左节点作为新节点递归继续寻找

		node.left = delete(node.left, key)
		return node
	}

	// 如果key> node.key 则向右寻找

	if key > node.key {

		// 将右节点作为新节点递归继续寻找

		node.right = delete(node.right, key)
		return node
	}

    //删除该节点
    if node.left==nil && node.right==nil{
    	fmt.Println(node.key,":",node.val,"will delete")
    	node = nil
    	return  node

	}
	if node.left == nil {
		node = node.right
		return node
	}
	//如果key==node.key 判断node有没有左右子树，如果有右子树，则将右子树直接赋值给当前节点，完成覆盖删除

	if node.right == nil {
		node = node.left
		return node
	}
	mostLeftNode := node.right
	// 要删除的节点有2个字节点，找到右子树的最左节点，替换当前节点􏱚􏱐􏲝􏰈􏰖􏰗􏰛
	for {

		//一直遍历找到最左节点
		if mostLeftNode != nil && mostLeftNode.left != nil {
			mostLeftNode = mostLeftNode.left
		} else {
			break
		}

	}
	fmt.Println("find tihuan node",mostLeftNode.key,":",mostLeftNode.val)

	// 使用右子树的最左节点替换当前的节点，即删除当前节点
	node.key, node.val = mostLeftNode.key,mostLeftNode.val

	node.right = delete(node.right, node.key)

	return node


}
// 先序遍历：根节点 -> 左子树 ->右子树
func (tree *BinaryTree) PreOrderTraverse() {

	tree.lock.Lock()
	defer tree.lock.Unlock()

	preOrderTraverse(tree.root)

}

func preOrderTraverse(node *Node) {

	if node != nil {
		//先打印根节点
		fmt.Println(node.key,":",node.val)

		//然后递归调用自己，将左节点作为新节点，打印
		preOrderTraverse(node.left)
		//然后递归调用自己，将右节点作为新节点，打印
		preOrderTraverse(node.right)

	}

}





func main() {
	tree:= BinaryTree{nil,sync.RWMutex{}}
	tree.Insert(8,"pangchenxu")
    tree.Insert(4,"lichen")
	tree.Insert(2,"chenyang")
	tree.Insert(1,"chenyang")
	tree.Insert(3,"chenyang")
	tree.Insert(6,"chenyang")
	tree.Insert(5,"chenyang")
	tree.Insert(7,"chenyang")
	tree.Insert(10,"chenyang")
	tree.Insert(9,"chenyang")
	tree.Insert(11,"chenyang")
	tree.PreOrderTraverse()
    tree.deleteTree(4)
	fmt.Println("delete root node")
	tree.PreOrderTraverse()
}