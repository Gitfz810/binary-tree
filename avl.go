package main

import "fmt"

type AVLNode struct {
	Val    int
	Height int        // 节点深度
	Left   *AVLNode
	Right  *AVLNode
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (avl *AVLNode) insert(key int) *AVLNode {
	if avl == nil {
		return &AVLNode{
			Val:    key,
			Height: 1,
			Left:   nil,
			Right:  nil,
		}
	}
    // 递归调用找到合适的插入位置
	if key < avl.Val {
		avl.Left = avl.Left.insert(key)
		// determine adjust binary tree or not
		avl = avl.handleNode()
	} else if key > avl.Val {
		avl.Right = avl.Right.insert(key)
		avl = avl.handleNode()
	} else {
		fmt.Println("key is exists")
		return avl
	}
	avl.Height = max(avl.Left.getHeight(), avl.Right.getHeight()) + 1
	return avl
}

func (avl *AVLNode) delete(key int) *AVLNode {
	if avl == nil {
		return avl
	}
	if key < avl.Val {
		avl.Left = avl.Left.delete(key)
	} else if key > avl.Val {
		avl.Right = avl.Right.delete(key)
	} else {
		if avl.Left != nil && avl.Right != nil {
			// 把删除节点的值改为其后继节点的值
			avl.Val = avl.Right.getMin()
			// 删除后继节点
			avl.Right = avl.Right.delete(avl.Val)
		} else if avl.Left != nil {
			avl = avl.Left
		} else {
			avl = avl.Right
		}
	}
	if avl != nil {
		avl.Height = max(avl.Left.getHeight(), avl.Right.getHeight()) + 1
		avl = avl.handleNode()
	}
	return avl
}

// 获取删除节点的后继节点，即删除节点的右子树的最小值
func (avl *AVLNode) getMin() int {
	if avl == nil {
		return -1
	}
	if avl.Left == nil {
		return avl.Val
	} else {
		return avl.Left.getMin()
	}
}

func (avl *AVLNode) handleNode() *AVLNode {
	if avl.Left.getHeight() - avl.Right.getHeight() == 2 {
		// 当节点平衡因子大于0时，只需右转
		if avl.Left.Left.getHeight() - avl.Left.Right.getHeight() > 0 {
			avl = avl.rightRoute()
		} else {
			// 当节点平衡因子小于0时，需先左转在右转
			avl= avl.leftThenRightRoute()
		}
	} else if avl.Left.getHeight() - avl.Right.getHeight() == int(-2) {
		if avl.Right.Left.getHeight() - avl.Right.Right.getHeight() < 0 {
			avl = avl.leftRoute()
		} else {
			avl = avl.rightThenLeftRoute()
		}
	}
	return avl
}

// 获取节点深度
func (avl *AVLNode) getHeight() int {
	if avl == nil {
		return 0
	}
	return avl.Height
}

/*
 *       P                L
 *     /   \            /   \
 *    L    Pr   --->   Ll   P
 *   / \               |   / \
 *  Ll Lr              N  Lr Pr
 *  |
 *  N
 */
func (avl *AVLNode) rightRoute() *AVLNode {
	L := avl.Left
	avl.Left = L.Right
	L.Right = avl
	// 因为插入新的节点，所以更新节点深度
	avl.Height = max(avl.Left.getHeight(), avl.Right.getHeight()) + 1
	L.Height = max(L.Left.getHeight(), L.Right.getHeight()) + 1
	return L
}

/*
 *       P                L
 *     /   \            /   \
 *    Pl    L   --->   P    Lr
 *         / \        / \   |
 *        Ll Lr      Pl Ll  N
 *           |
 *           N
 */
func (avl *AVLNode) leftRoute() *AVLNode {
	L := avl.Right
	avl.Right = L.Left
	L.Left = avl
	// 更新节点深度
	avl.Height = max(avl.Left.getHeight(), avl.Right.getHeight()) + 1
	L.Height = max(L.Left.getHeight(), L.Right.getHeight()) + 1
	return L
}

/*
 *       P                 P                 Lr
 *     /   \    L左旋    /   \    P右旋     /    \
 *    L    Pr   --->   Lr    Pr   --->    L     P
 *   / \              / \               /  \   /  \
 * Ll  Lr            L  Lrr            Ll Lrl Lrr Pr
 *    /  \          / \                    |
 *   Lrl Lrr       Ll Lrl                  N
 *    |                |
 *    N                N
 */
func (avl *AVLNode) leftThenRightRoute() *AVLNode {
	Lr := avl.Left.leftRoute()
	avl.Left = Lr
	return avl.rightRoute()
}

/*
 *       P                 P                  Ll
 *     /   \    L右旋    /   \    P左旋      /    \
 *    Pl    L   --->   Pl    Ll   --->     P      L
 *         / \              /  \         /  \    /  \
 *       Ll  Lr           Lll  L        Pl Lll Llr Lr
 *      /  \                  / \               |
 *     Lll Llr              Llr  Lr             N
 *         |                 |
 *         N                 N
 */
func (avl *AVLNode) rightThenLeftRoute() *AVLNode {
	Ll := avl.Right.rightRoute()
	avl.Right = Ll
	return avl.leftRoute()
}
// 判断是否是二叉排序树
func inOrderRecu(node *AVLNode) {
	if node == nil {
		return
	}

	inOrderRecu(node.Left)
	fmt.Println(node.Val)
	inOrderRecu(node.Right)
}
// 判断是否是平衡二叉树
func (avl *AVLNode) isBalanced() bool {
	if avl == nil {
		return true
	}
	if avl.Left.getHeight() - avl.Right.getHeight() > 1 || avl.Left.getHeight() - avl.Right.getHeight() < int(-1) {
		return false
	}
	return avl.Left.isBalanced() && avl.Right.isBalanced()
}

func main() {
	var (
		avl *AVLNode
		a = []int{3, 2, 1, 4, 5, 6, 7, 10, 9, 8}
	)

	for _, i := range a {
		avl = avl.insert(i)
	}
	inOrderRecu(avl)
	fmt.Println(avl.isBalanced())

	avl = avl.delete(7)
	inOrderRecu(avl)
	fmt.Println(avl.isBalanced())
}
