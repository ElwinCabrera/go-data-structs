package trees

type TreeNodeValue interface {
	~int | ~float32 | ~float64 | ~string
}

type TreeNode[T any] struct {
	Value  T
	weight int
	left   *TreeNode[T]
	right  *TreeNode[T]
}

type Tree[T any] interface {
	Insert(value T) *TreeNode[T]

	RemoveNode(n *TreeNode[T])
	RemoveValue(value T)
	Contains(value T) bool
	Find(value T) []*TreeNode[T]
	//Values() []T
	Size() int
	Clear() bool

	PreOrderValues() []*TreeNode[T]
	InOrderValues() []*TreeNode[T]
	PostOrderValues() []*TreeNode[T]
	Root() *TreeNode[T]
}

type MinMaxTree[T any] interface {
	Tree[T]
	Pop() *TreeNode[T]
}
