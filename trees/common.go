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
	Remove(n *TreeNode[T]) bool
	Contains(value T) bool
	Find(value T) []*TreeNode[T]
	Values() []T
	Size() int
	Clear() bool
}
