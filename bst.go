package main

import (
	"fmt"
	"sync"
)
// 构建二叉树
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
// 构建二叉排序树
type BST struct {
	Root *TreeNode
	Lock sync.RWMutex
}
// 查询节点是否存在
func search(node *TreeNode, key int) bool {
	if node == nil {
		return false
	}

	if key == node.Val {
		return true
	} else if key > node.Val {
		return search(node.Right, key)
	} else {
		return search(node.Left, key)
	}
}
// 插入节点数据
func insert(node, addNode *TreeNode) {
	if addNode.Val < node.Val {
		if node.Left == nil {
			node.Left = addNode
		} else {
			insert(node.Left, addNode)
		}
	} else {
		if node.Right == nil {
			node.Right = addNode
		} else {
			insert(node.Right, addNode)
		}
	}
}
// 插入节点
func (b *BST) insertNode(data int) {
	b.Lock.Lock()
	defer b.Lock.Unlock()
	// 构建数据结构
	treeNode := TreeNode{
		Val:   data,
		Left:  nil,
		Right: nil,
	}

	if b.Root == nil {
		b.Root = &treeNode
	} else {
		if !search(b.Root, data) {
			insert(b.Root, &treeNode)
		}
	}
}
// 删除节点数据
func del(node *TreeNode, key int) *TreeNode {
	if node == nil {
		return nil
	}

	if key < node.Val {
		node.Left = del(node.Left, key)
		return node
	} else if key > node.Val {
		node.Right = del(node.Right, key)
		return node
	} else {
		if node.Right == nil && node.Left == nil {
			node = nil
			return nil
		}

		if node.Right == nil {
			node = node.Left
			return node
		} else if node.Left == nil {
			node = node.Right
			return node
		} else {
			// 找删除节点的后继节点
			nodeRight := node.Right
			for {
				if nodeRight != nil && nodeRight.Left != nil {
					nodeRight = nodeRight.Left
				} else {
					break
				}
			}
			// nodeRight就是删除节点的后继节点
			node.Val = nodeRight.Val
			node.Right = del(node.Right, node.Val)
			return node
		}
	}
}
// 删除节点
func (b *BST) deleteNode(key int) {
	b.Lock.Lock()
	defer b.Lock.Unlock()

	del(b.Root, key)
}

func main() {
	var (
		binarySearchTree BST
		a = [10]int{66, 88, 58, 47, 35, 73, 51, 99, 37, 93}
	)
	for _, i := range a {
		binarySearchTree.insertNode(i)
	}

	fmt.Println(search(binarySearchTree.Root, 93))
	binarySearchTree.deleteNode(93)
	fmt.Println(search(binarySearchTree.Root, 93))
}
