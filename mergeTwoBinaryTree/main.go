package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root1 := TreeNode{Val: 2}
	root2 := TreeNode{Val: 3}
	fmt.Printf("mergeTrees(&root1, &root2): %v\n", mergeTrees(&root1, &root2))
}

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil && root2 == nil {
		return nil
	}
	mergeRoot := &TreeNode{}
	mergeChild(mergeRoot, root1, root2)
	return mergeRoot
}
func mergeChild(mergedRoot, root1, root2 *TreeNode) {

	if root1 == nil && root2 == nil {
		return
	}
	if root1 == nil {
		mergedRoot.Val = root2.Val

		if root2.Right != nil {
			mergedRoot.Right = &TreeNode{}
		}
		if root2.Left != nil {
			mergedRoot.Left = &TreeNode{}
		}

		mergeChild(mergedRoot.Right, nil, root2.Right)
		mergeChild(mergedRoot.Left, nil, root2.Left)
	} else if root2 == nil {
		mergedRoot.Val = root1.Val

		if root1.Right != nil {
			mergedRoot.Right = &TreeNode{}
		}
		if root1.Left != nil {
			mergedRoot.Left = &TreeNode{}
		}
		mergeChild(mergedRoot.Right, root1.Right, nil)
		mergeChild(mergedRoot.Left, root1.Left, nil)
	} else {

		mergedRoot.Val = root1.Val + root2.Val

		if root2.Right != nil || root1.Right != nil {
			mergedRoot.Right = &TreeNode{}
		}
		if root2.Left != nil || root1.Left != nil {
			mergedRoot.Left = &TreeNode{}
		}

		mergeChild(mergedRoot.Right, root1.Right, root2.Right)
		mergeChild(mergedRoot.Left, root1.Left, root2.Left)
	}

}
