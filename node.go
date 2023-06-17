package leetcode

import (
	"math"
)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 归并排序？
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow, fast := head, head
	if fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	firstTail := slow
	slow = slow.Next
	firstTail.Next = nil
	first, second := sortList(head), sortList(slow)
	return mergeList(first, second)
}

func mergeList(l1, l2 *ListNode) *ListNode {
	dummy := new(ListNode)
	cur := dummy
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}
	if l1 != nil {
		cur.Next = l1
	}
	if l2 != nil {
		cur.Next = l2
	}
	return dummy.Next
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	first := head
	second := head.Next
	others := head.Next.Next
	second.Next = first
	first.Next = swapPairs(others)
	return second
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			break
		}
	}
	if fast == nil || slow != fast {
		return nil
	}
	slow = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	l1, l2 := headA, headB
	for l1 != l2 {
		if l1 == nil {
			l1 = headB
			continue
		}
		if l2 == nil {
			l2 = headA
			continue
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	return l2
}

func kthSmallest(root *TreeNode, k int) int {
	res := 0
	var inOrder func(root *TreeNode) int
	inOrder = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := inOrder(root.Left)
		if left != 0 {
			return left
		}
		res++
		if res == k {
			return root.Val
		}
		right := inOrder(root.Right)
		if right != 0 {
			return right
		}
		return 0
	}
	return inOrder(root)
}

func kthSmallest1(root *TreeNode, k int) int {
	if root == nil {
		return 0
	}
	stack := []*TreeNode{}
	for p := root; p != nil && len(stack) > 0; p = p.Right {
		for p.Left != nil {
			stack = append(stack, p.Left)
		}
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		k--
		if k == 0 {
			return node.Val
		}
		p = node.Right
	}
	return -1
}

func isPalindrome(head *ListNode) bool {
	slow, fast := head, head
	var prev *ListNode
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		nxt := slow
		slow.Next = prev
		prev = slow
		slow = nxt
	}
	if fast != nil {
		slow = slow.Next
	}
	for slow != nil {
		if slow.Val != prev.Val {
			return false
		}
		slow = slow.Next
		prev = prev.Next
	}
	return true
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		nxt := head.Next
		head.Next = prev
		prev = head
		head = nxt
	}
	return prev
}

func diameterOfBinaryTree(root *TreeNode) int {
	var res int
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	var diameter func(root *TreeNode) int
	diameter = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := diameter(root.Left)
		right := diameter(root.Right)
		res = max(res, left+right)
		return max(left, right) + 1
	}
	diameter(root)
	return res
}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	dummy := new(Node)
	list := dummy
	m := make(map[*Node]*Node)
	for head != nil {
		node := &Node{
			Val: head.Val,
		}
		m[head] = node
		list.Next = node
		list = list.Next
		head = head.Next
	}
	list = dummy.Next
	for head != nil {
		list.Random = m[head.Random]
		list = list.Next
		head = head.Next
	}
	return dummy.Next
}

func pathSum(root *TreeNode, sum int) int {
	var res int
	var path func(root *TreeNode, notRoot bool, sum int)
	path = func(root *TreeNode, notRoot bool, sum int) {
		if root == nil {
			return
		}
		if notRoot {
			path(root.Left, false, sum)
			path(root.Right, false, sum)
		}
		path(root.Left, true, sum-root.Val)
		path(root.Right, true, sum-root.Val)
	}
	return res
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root == p || root == q {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left == nil && right == nil {
		return nil
	}
	if left != nil {
		return left
	}
	return right
}

func maxPathSum(root *TreeNode) int {
	var ret int
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	var pathSum func(root *TreeNode) int
	pathSum = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := max(pathSum(root.Left), 0)
		right := max(pathSum(root.Right), 0)
		ret = max(ret, left+right+root.Val)
		return root.Val + max(left, right)
	}
	pathSum(root)
	return ret
}

func isValidBST(root *TreeNode) bool {
	minV, maxV := math.MinInt64, math.MaxInt64
	var valid func(root *TreeNode, minV, maxV int) bool
	valid = func(root *TreeNode, minV, maxV int) bool {
		if root == nil {
			return true
		}
		if root.Val <= minV || root.Val >= maxV {
			return false
		}
		return valid(root.Left, minV, root.Val) && valid(root.Right, root.Val, maxV)
	}
	return valid(root, minV, maxV)
}

func rightSideView(root *TreeNode) []int {
	var ret []int
	var bt func(root *TreeNode, level int)
	bt = func(root *TreeNode, level int) {
		if root == nil {
			return
		}
		if level == len(ret) {
			ret = append(ret, root.Val)
		}
		bt(root.Right, level+1)
		bt(root.Left, level+1)
	}
	bt(root, 0)
	return ret
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	second := head
	for i := 1; i < k; i++ {
		if second == nil {
			return head
		}
		second = second.Next
	}
	others := second.Next
	second.Next = nil
	newHead := reverse(head)
	newHead.Next = reverseKGroup(others, k)
	return newHead
}

func reverse(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		nxt := head.Next
		head.Next = prev
		prev = head
		head = nxt
	}
	return prev
}

func flatten(root *TreeNode) {
	var preorder func(root *TreeNode) *TreeNode
	preorder = func(root *TreeNode) *TreeNode {
		if root == nil {
			return nil
		}
		left := preorder(root.Left)
		right := preorder(root.Right)
		root.Left = nil
		root.Right = left
		tmp := root
		for tmp.Right != nil {
			tmp = tmp.Right
		}
		tmp.Right = right
		return root
	}
	preorder(root)
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
	return root
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{
		Next: head,
	}
	for i := 0; i < n; i++ {
		head = head.Next
	}
	prev := dummy
	for head != nil {
		prev = prev.Next
		head = head.Next
	}
	prev.Next = prev.Next.Next
	return dummy.Next
}

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var res []int
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, node.Val)
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
	}
	return res
}

func inorderTraversal(root *TreeNode) []int {
	var ret []int
	stack := []*TreeNode{}
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ret = append(ret, node.Val)
		root = node.Right
	}
	return ret
}

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var ret []int
	stack := []*TreeNode{}
	var prev *TreeNode
	for root != nil && len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		node := stack[len(stack)-1]
		if node.Right == nil || node.Right == prev {
			stack = stack[:len(stack)-1]
			ret = append(ret, node.Val)
			prev = node
		} else {
			root = node.Right
		}
	}
	return ret
}

func levelOrder(root *TreeNode) [][]int {
	var ret [][]int
	deque := []*TreeNode{root}
	for len(deque) > 0 {
		level := make([]int, 0)
		for _, node := range deque {
			deque = deque[1:]
			level = append(level, node.Val)
			if node.Left != nil {
				deque = append(deque, node.Left)
			}
			if node.Right != nil {
				deque = append(deque, node.Right)
			}
		}
		ret = append(ret, level)
	}
	return ret
}

func largestValues(root *TreeNode) []int {
	var ret []int
	var bfs func(root *TreeNode, level int)
	bfs = func(root *TreeNode, level int) {
		if root == nil {
			return
		}
		if level == len(ret) {
			ret = append(ret, root.Val)
		} else if ret[level] < root.Val {
			ret[level] = root.Val
		}
		bfs(root.Left, level+1)
		bfs(root.Right, level+1)
	}
	bfs(root, 0)
	return ret
}

func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	l := &ListNode{
		Next: head,
	}
	for prev, i := head, head.Next; i != nil; i = i.Next {
		if prev.Val <= i.Val {
			prev = i
			continue
		}
		j := l
		for j.Next.Val <= i.Val {
			j = j.Next
		}
		prev.Next = i.Next
		i.Next = j.Next
		j.Next = i
		i = prev
	}
	return l.Next
}

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var ret [][]int
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		one := []int{}
		for _, root := range stack {
			stack = stack[1:]
			one = append(one, root.Val)
			if root.Left != nil {
				stack = append(stack, root.Left)
			}
			if root.Right != nil {
				stack = append(stack, root.Right)
			}
		}
		ret = append([][]int{one}, ret...)
	}
	return ret
}

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}
	return build(1, n)
}

func build(lo, hi int) []*TreeNode {
	res := []*TreeNode{}
	// base case
	if lo > hi {
		// nil 需要返回,构造出子树为nil的节点
		res = append(res, nil)
		return res
	}
	// 穷举root节点的所有可能
	for i := lo; i <= hi; i++ {
		// bst定义left<root<right
		// 递归到最小子树时, 根据值可以构造出左右子树所有合法bst
		// 只有一个值时, 例如2, 返回是 nil,2,nil
		left := build(lo, i-1)
		right := build(i+1, hi)
		//.排列组合
		for _, l := range left {
			for _, r := range right {
				root := &TreeNode{Val: i}
				root.Left = l
				root.Right = r
				res = append(res, root)
			}
		}
	}
	return res
}

func averageOfLevels(root *TreeNode) []float64 {
	var res []float64
	var count []int
	var sumOfLevels func(root *TreeNode, level int)
	sumOfLevels = func(root *TreeNode, level int) {
		if root == nil {
			return
		}
		if level == len(res) {
			res = append(res, float64(root.Val))
			count = append(count, 1)
		} else {
			res[level] += float64(root.Val)
			count[level]++
		}
		sumOfLevels(root.Left, level+1)
		sumOfLevels(root.Right, level+1)
	}
	sumOfLevels(root, 0)
	for i := range res {
		res[i] /= float64(count[i])
	}
	return res
}
