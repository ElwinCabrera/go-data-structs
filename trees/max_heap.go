package trees

type MaxHeap[T TreeNodeValue] struct {
	root *TreeNode[T]
	size int
}

func NewMaxHeap[T TreeNodeValue]() *MaxHeap[T] {
	return &MaxHeap[T]{nil, 0}
}

func (t *MaxHeap[T]) Insert(value T) *TreeNode[T] {

	var newNode *TreeNode[T]
	if t.root == nil {
		newNode = &TreeNode[T]{Value: value}
		t.root = newNode
	} else if value < t.root.Value {
		oldRoot := t.root
		newNode = &TreeNode[T]{Value: value}
		newNode.left = oldRoot.left
		newNode.right = oldRoot.right

		oldRoot.left = nil
		oldRoot.right = nil
		t.root = newNode
		binaryTreeInsertNode(t.root, oldRoot)
	} else {
		newNode = binaryTreeInsert(t.root, value)
	}
	t.size++
	return newNode

}

func (t *MaxHeap[T]) RemoveNode(removeNode *TreeNode[T]) {
	if t.root == nil || removeNode == nil {
		return
	}

	if removeNode == t.root {
		t.Pop()
	} else {
		removeNodeHelper(t.root, removeNode)
	}
	removeNode = nil

	t.size--
}

func (t *MaxHeap[T]) Pop() *TreeNode[T] {
	oldRoot := t.root
	maxLeft := oldRoot.left
	maxRight := oldRoot.right
	maxHelper(t.root.left, &maxLeft)
	maxHelper(t.root.right, &maxRight)

	var newRoot *TreeNode[T]
	if maxLeft.Value < maxRight.Value {
		newRoot = maxLeft
	} else {
		newRoot = maxRight
	}
	newRoot.left = t.root.left
	newRoot.right = t.root.right
	t.root = newRoot
	oldRoot.left = nil
	oldRoot.right = nil
	return oldRoot
}

func (t *MaxHeap[T]) RemoveValue(value T) {
	var nodes []*TreeNode[T]
	treeFindAllHelper(t.root, value, &nodes)
	for _, node := range nodes {
		t.RemoveNode(node)
	}
}

func (t *MaxHeap[T]) Contains(value T) bool {
	return treeContainsHelper(t.root, value)
}

func (t *MaxHeap[T]) FindFirst(value T) *TreeNode[T] {
	var foundFirstNode *TreeNode[T]
	findFirstHelper(t.root, &foundFirstNode, value)
	return foundFirstNode
}

func (t *MaxHeap[T]) Find(value T) []*TreeNode[T] {
	var nodes []*TreeNode[T]
	treeFindAllHelper(t.root, value, &nodes)
	return nodes
}

func (t *MaxHeap[T]) Size() int {
	return treeSizeHelper(t.root)
}

func (t *MaxHeap[T]) Min() any {
	currMin := t.root
	minHelper(t.root, &currMin)
	return currMin.Value
}

func (t *MaxHeap[T]) Max() any {
	currMax := t.root
	maxHelper(t.root, &currMax)
	return currMax.Value
}

func (t *MaxHeap[T]) Clear() bool {

	nodes := t.PreOrderValues()
	for i, _ := range nodes {
		nodes[i] = nil
	}
	t.root = nil
	t.size = 0
	return true
}

func (t *MaxHeap[T]) PreOrderValues() []*TreeNode[T] {
	var res []*TreeNode[T]
	treePreOrderValueHelper(t.root, &res)
	return res
}

func (t *MaxHeap[T]) InOrderValues() []*TreeNode[T] {
	var res []*TreeNode[T]
	treeInOrderValueHelper(t.root, &res)
	return res
}

func (t *MaxHeap[T]) PostOrderValues() []*TreeNode[T] {
	var res []*TreeNode[T]
	treePostOrderValueHelper(t.root, &res)
	return res
}

func (t *MaxHeap[T]) Root() *TreeNode[T] {
	return t.root
}
