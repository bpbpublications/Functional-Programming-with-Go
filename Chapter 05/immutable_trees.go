type TreeNode struct {
    Value int
    Left, Right *TreeNode
}

func NewTreeNode(value int, left, right *TreeNode) *TreeNode {
    return &TreeNode{Value: value, Left: left, Right: right}
}

func (n *TreeNode) Add(value int) *TreeNode {
    if n == nil {
        return NewTreeNode(value, nil, nil)
    }
    if value < n.Value {
        return NewTreeNode(n.Value, n.Left.Add(value), n.Right)
    }
    return NewTreeNode(n.Value, n.Left, n.Right.Add(value))
}
