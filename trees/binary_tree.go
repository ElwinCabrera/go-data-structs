package trees

type BinaryTree[T TreeNodeValue] struct {
	root *TreeNode[T]
	size int
}

func NewBinaryTree[T TreeNodeValue]() *BinaryTree[T] {
	return &BinaryTree[T]{nil, 0}
}

func (t *BinaryTree[T]) Insert(value T) *TreeNode[T] {
	var newNode *TreeNode[T]
	if t.root == nil {
		newNode = &TreeNode[T]{Value: value}
		t.root = newNode
	} else {
		newNode = binaryTreeInsert(t.root, value)
	}
	t.size++
	return newNode

}

func (t *BinaryTree[T]) RemoveNode(removeNode *TreeNode[T]) {
	if t.root == nil || removeNode == nil {
		return
	}

	if removeNode == t.root {
		removeRootNode(&t.root)
	} else {
		removeNodeHelper(t.root, removeNode)
	}
	removeNode = nil

	t.size--
}

func (t *BinaryTree[T]) RemoveValue(value T) {
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
	var foundFirstNode *TreeNode[T]
	findFirstHelper(t.root, &foundFirstNode, value)
	return foundFirstNode
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
	currMin := t.root
	minHelper(t.root, &currMin)
	return currMin.Value
}

func (t *BinaryTree[T]) Max() any {
	currMax := t.root
	maxHelper(t.root, &currMax)
	return currMax.Value
}

func (t *BinaryTree[T]) Clear() bool {

	nodes := t.PreOrderValues()
	for i, _ := range nodes {
		nodes[i] = nil
	}
	t.root = nil
	t.size = 0
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

func (t *BinaryTree[T]) Root() *TreeNode[T] {
	return t.root
}
