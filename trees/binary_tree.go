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

func (t *BinaryTree[T]) Remove(value T) {

}

func (t *BinaryTree[T]) Contains(value T) {

}

func (t *BinaryTree[T]) Find(value T) []*TreeNode[T] {

}

func (t *BinaryTree[T]) Size() int {

}

func (t *BinaryTree[T]) Min() any {

}

func (t *BinaryTree[T]) Max() any {

}

func (t *BinaryTree[T]) Clear() bool {

}

func (t *BinaryTree[T]) Values() []T {

}

func (t *BinaryTree[T]) PreOrderTraverse() {

}

func (t *BinaryTree[T]) InOrderTraverse() {

}

func (t *BinaryTree[T]) PostOrderTraverse() {

}
