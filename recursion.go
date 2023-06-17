package leetcode

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	rootVal := preorder[0]
	findIdx := func(nums []int, val int) int {
		for i, v := range nums {
			if v == val {
				return i
			}
		}
		return -1
	}
	rootIdx := findIdx(inorder, rootVal)
	return &TreeNode{
		Val:   rootVal,
		Left:  buildTree(preorder[1:rootIdx+1], inorder[:rootIdx]),
		Right: buildTree(preorder[rootIdx+1:], inorder[rootIdx+1:]),
	}
}