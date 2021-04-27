package common

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func BuildTree(vals []int) *TreeNode {
	var root = new(TreeNode)
	root.Val = vals[0]
	for _, val := range vals[1:] {
		if val == 0 {
			continue
		}
		root.Insert(&TreeNode{Val: val})
	}
	return root
}

func (nd *TreeNode) Insert(newNode *TreeNode) {
	if newNode.Val == nd.Val {
		return
	}
	if newNode.Val > nd.Val {
		if nd.Right == nil {
			nd.Right = newNode
		} else {
			nd.Right.Insert(newNode)
		}
	} else { //3 小于 继续比较插入到 左孩子
		if nd.Left == nil {
			nd.Left = newNode
		} else {
			nd.Left.Insert(newNode)
		}
	}
}
