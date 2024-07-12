package trees

type MinHeap[T TreeNodeValue] struct {
	root *TreeNode[T]
	size int
}

func NewMinHeap[T TreeNodeValue]() *MinHeap[T] {
	return &MinHeap[T]{nil, 0}
}

func (t *MinHeap[T]) Insert(value T) *TreeNode[T] {

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

func (t *MinHeap[T]) RemoveNode(removeNode *TreeNode[T]) {
	if t.root == nil || removeNode == nil {
		return
	}

	if removeNode == t.root {
		t.Pop()
	} else {
		detachNodeFromTree(t.root, removeNode)
	}
	removeNode = nil

	t.size--
}

func (t *MinHeap[T]) Pop() *TreeNode[T] {
	oldRoot := t.root
	if t.root.left == nil && t.root.right == nil {
		t.root = nil
		return oldRoot
	}
	minLeft := oldRoot.left
	minRight := oldRoot.right
	minHelper(t.root.left, &minLeft) // not needed bc if root is always minimum then there will never be a node to the left of root but might leave it in just in case we want to implement a heap without the binary tree rules
	minHelper(t.root.right, &minRight)

	var newRoot *TreeNode[T]

	if minLeft == nil { //will Always be this case for min heap following binary tree insert rule
		newRoot = minRight
	} else if minRight == nil {
		newRoot = minLeft
	} else if minLeft != nil && minRight != nil { // will probably be this case if we are following the binary tree rules for insert after the root node but will leave in to make implementation more versatile
		if minLeft.Value < minRight.Value {
			newRoot = minLeft
		} else {
			newRoot = minRight
		}
	}

	detachNodeFromTree(t.root, newRoot)

	newRoot.left = t.root.left
	newRoot.right = t.root.right
	t.root = newRoot
	oldRoot.left = nil
	oldRoot.right = nil
	return oldRoot
}

func (t *MinHeap[T]) RemoveValue(value T) {
	var nodes []*TreeNode[T]
	treeFindAllHelper(t.root, value, &nodes)
	for _, node := range nodes {
		t.RemoveNode(node)
	}
}

func (t *MinHeap[T]) Contains(value T) bool {
	return treeContainsHelper(t.root, value)
}

func (t *MinHeap[T]) FindFirst(value T) *TreeNode[T] {
	var foundFirstNode *TreeNode[T]
	findFirstHelper(t.root, &foundFirstNode, value)
	return foundFirstNode
}

func (t *MinHeap[T]) Find(value T) []*TreeNode[T] {
	var nodes []*TreeNode[T]
	treeFindAllHelper(t.root, value, &nodes)
	return nodes
}

func (t *MinHeap[T]) Size() int {
	return treeSizeHelper(t.root)
}

func (t *MinHeap[T]) Min() any {
	currMin := t.root
	minHelper(t.root, &currMin)
	return currMin.Value
}

func (t *MinHeap[T]) Max() any {
	currMax := t.root
	maxHelper(t.root, &currMax)
	return currMax.Value
}

func (t *MinHeap[T]) Clear() bool {

	nodes := t.PreOrderValues()
	for i, _ := range nodes {
		nodes[i] = nil
	}
	t.root = nil
	t.size = 0
	return true
}

func (t *MinHeap[T]) PreOrderValues() []*TreeNode[T] {
	var res []*TreeNode[T]
	treePreOrderValueHelper(t.root, &res)
	return res
}

func (t *MinHeap[T]) InOrderValues() []*TreeNode[T] {
	var res []*TreeNode[T]
	treeInOrderValueHelper(t.root, &res)
	return res
}

func (t *MinHeap[T]) PostOrderValues() []*TreeNode[T] {
	var res []*TreeNode[T]
	treePostOrderValueHelper(t.root, &res)
	return res
}

func (t *MinHeap[T]) Root() *TreeNode[T] {
	return t.root
}

func (t *MinHeap[T]) GetHeapType() string {
	return "min"
}
