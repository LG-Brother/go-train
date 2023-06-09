package arithmetic

import (
	"math/rand"
	"time"
)

/**
定义数据结构
BFS：Breadth-First Search 广度优先搜索算法（用队列实现）
DFS：Depth-First Search 深度优先搜索算法（用递归、栈实现）
*/

// ListNode 链表结构
type ListNode struct {
	Val  int
	Next *ListNode
}

// TreeNode 二叉树结构
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// CreateLinkedList 创建指定个数的随机链表
func CreateLinkedList(count int) *ListNode {
	head := &ListNode{}
	tail := head
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < count; i++ {
		node := &ListNode{Val: rand.Intn(100)}
		tail.Next = node
		tail = node
	}
	return head
}

// PrintLinkedList 打印链表
func PrintLinkedList(head *ListNode) {
	if head == nil {
		return
	}
	p := head.Next
	for p != nil {
		print(p.Val, " ")
		p = p.Next
	}
	println()
}

// PreorderPrintTree 使用非递归方法先序遍历打印二叉树
func PreorderPrintTree(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	// 定义存放遍历结果
	var ans []int
	ans = append(ans, root.Val)
	// 定义临时栈
	var stack []*TreeNode
	stack = append(stack, root)
	// 定义遍历树的指针
	var p *TreeNode
	// 定义是否遍历左子树标识（在遍历右子树时开启标识）
	flag := true
	for len(stack) != 0 {
		p = stack[len(stack)-1]
		// 遍历左子树
		for p.Left != nil && flag {
			p = p.Left
			ans = append(ans, p.Val)
			stack = append(stack, p)
		}
		// 弹栈
		stack = stack[:len(stack)-1]
		// 遍历右子树
		if p.Right != nil {
			p = p.Right
			ans = append(ans, p.Val)
			stack = append(stack, p)
			flag = true
		} else {
			flag = false
		}
	}
	return ans
}

// PreorderPrintTreeByRecursion 使用递归方法先序遍历打印二叉树
func PreorderPrintTreeByRecursion(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var ans []int
	left := PreorderPrintTreeByRecursion(root.Left)
	right := PreorderPrintTreeByRecursion(root.Right)
	ans = append(ans, root.Val)
	ans = append(ans, left...)
	ans = append(ans, right...)
	return ans
}

// InorderPrintTreeByRecursion 使用递归方法中序遍历打印二叉树
func InorderPrintTreeByRecursion(root *TreeNode) []int {
	var ans []int
	var inorder func(*TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		ans = append(ans, node.Val)
		inorder(node.Right)
	}
	inorder(root)
	return ans
}

// InorderPrintTree 使用非递归方法中序遍历打印二叉树
func InorderPrintTree(root *TreeNode) (ans []int) {
	var stack []*TreeNode
	node := root
	for node != nil || len(stack) != 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		ans = append(ans, stack[len(stack)-1].Val)
		node = stack[len(stack)-1].Right
		stack = stack[:len(stack)-1]
	}
	return ans
}

// PostorderPrintTreeByRecursion 递归后序遍历打印二叉树
func PostorderPrintTreeByRecursion(root *TreeNode) (ans []int) {
	var postorder func(*TreeNode)
	postorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		postorder(node.Left)
		postorder(node.Right)
		ans = append(ans, node.Val)
	}
	postorder(root)
	return ans
}

// PostorderPrintTree 非递归后序遍历打印二叉树
func PostorderPrintTree(root *TreeNode) (ans []int) {
	var stack []*TreeNode
	var prev *TreeNode
	for root != nil || len(stack) > 0 {
		// 遍历左子树
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		// 弹栈
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		// root.Right == prev 是为了表示该节点的右子树已经遍历
		if root.Right == nil || root.Right == prev {
			// 取值
			ans = append(ans, root.Val)
			prev = root
			root = nil
		} else {
			// 遍历右子树
			stack = append(stack, root)
			root = root.Right
		}
	}
	return
}

// LevelOrder 二叉树进行层序遍历（BFS-广度优先搜索）
func LevelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return nil
	}
	// 定义一个队列实现
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		innerSize := len(queue)
		var innerVal []int
		for i := 0; i < innerSize; i++ {
			node := queue[i]
			innerVal = append(innerVal, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(res, innerVal)
		queue = queue[innerSize:]
	}
	return res
}
