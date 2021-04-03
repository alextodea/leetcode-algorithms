package trees 


type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Constructor(val int) *TreeNode {
    return &TreeNode{Val: val, Left: nil, Right: nil}
}

var preorderListFirstElementIndex int
var inorderListNodesMappings = make(map[int]int)
var thisPreorder []int

func buildTree(preorder []int, inorder []int) *TreeNode {
    thisPreorder = preorder
    
    for i,val := range inorder {
        inorderListNodesMappings[val] = i
    }
    
    return buildTreeRecursively(0,len(thisPreorder)-1)
}

func buildTreeRecursively(left_limit,right_limit int) *TreeNode {
        if left_limit > right_limit {
            return nil
        }
        
        rootValue := thisPreorder[preorderListFirstElementIndex]
        preorderListFirstElementIndex++
        
        index := inorderListNodesMappings[rootValue]
        root := Constructor(rootValue)
        root.Left = buildTreeRecursively(left_limit, index-1)
        root.Right = buildTreeRecursively(index+1,right_limit)
        return root 
    }