package trees

type BinaryTree[T TreeNodeValue] struct {
	root *TreeNode[T]
	size int
}

func (t *BinaryTree[T]) Insert(value T) *TreeNode[T] {
	var newNode *TreeNode[T]
	if t.root == nil {
		newNode = &TreeNode[T]{Value: value}
		t.root = newNode
	} else {
		newNode = binaryTreeInsert(t.root, value)
	}

	return newNode

}

func (t *BinaryTree[T]) RemoveNode(node *TreeNode[T]) {
	treeRemoveNodeHelper(t.root, node)
}

func (t *BinaryTree[T]) Remove(value T) {
	var nodes []*TreeNode[T]
	treeFindAllHelper(t.root, value, &nodes)
	for _, node := range nodes {
		t.RemoveNode(node)
	}
}

func (t *BinaryTree[T]) Contains(value T) bool {
	return treeContainsHelper(t.root, value)
}

func (t *BinaryTree[T]) FindFirst(value T) *TreeNode[T] {
	return treeFindHelper(t.root, value)
}

func (t *BinaryTree[T]) Find(value T) []*TreeNode[T] {
	var nodes []*TreeNode[T]
	treeFindAllHelper(t.root, value, &nodes)
	return nodes
}

func (t *BinaryTree[T]) Size() int {
	return treeSizeHelper(t.root)
}

func (t *BinaryTree[T]) Min() any {
	currMin := t.root.Value
	treeMinHelper(t.root, &currMin)
	return currMin
}

func (t *BinaryTree[T]) Max() any {
	currMax := t.root.Value
	treeMaxHelper(t.root, &currMax)
	return currMax
}

func (t *BinaryTree[T]) Clear() bool {

	nodes := t.PreOrderValues()
	for i, _ := range nodes {
		nodes[i] = nil
	}
	return true
}

func (t *BinaryTree[T]) PreOrderValues() []*TreeNode[T] {
	var res []*TreeNode[T]
	treePreOrderValueHelper(t.root, &res)
	return res
}

func (t *BinaryTree[T]) InOrderValues() []*TreeNode[T] {
	var res []*TreeNode[T]
	treeInOrderValueHelper(t.root, &res)
	return res
}

func (t *BinaryTree[T]) PostOrderValues() []*TreeNode[T] {
	var res []*TreeNode[T]
	treePostOrderValueHelper(t.root, &res)
	return res
}
