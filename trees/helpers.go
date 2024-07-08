package trees

func binaryTreeInsert[T TreeNodeValue](current *TreeNode[T], value T) *TreeNode[T] {
	if current == nil {
		return nil
	}

	var newNode *TreeNode[T]

	if value <= current.Value {
		if current.left == nil {
			newNode = &TreeNode[T]{Value: value}
			current.left = newNode
			return newNode
		}
		newNode = binaryTreeInsert(current.left, value)

	} else {
		if current.right == nil {
			newNode = &TreeNode[T]{Value: value}
			current.right = newNode
			return newNode
		}
		newNode = binaryTreeInsert(current.right, value)

	}
	return newNode

}

func treeRemoveNodeHelper[T any](current *TreeNode[T], remove *TreeNode[T]) {
	if current == nil || remove == nil {
		return
	}

	removeLeft := remove.left
	removeRight := remove.right

	if current.left == remove {
		current.left = removeLeft
		remove = nil
	} else if current.right == remove {
		current.right = removeRight
		remove = nil
	} else {
		treeRemoveNodeHelper(current.left, remove)
		treeRemoveNodeHelper(current.right, remove)
	}
}

func treeFindHelper[T TreeNodeValue](current *TreeNode[T], value T) *TreeNode[T] {
	if current == nil {
		return nil
	}
	var foundNode *TreeNode[T]
	if value == current.Value {
		foundNode = current
	} else {
		treeFindHelper(current.left, value)
		treeFindHelper(current.right, value)
	}
	return foundNode
}

func treeFindAllHelper[T TreeNodeValue](current *TreeNode[T], value T, foundValues *[]*TreeNode[T]) {
	if current == nil {
		return
	}
	if value == current.Value {
		*foundValues = append(*foundValues, current)
	}
	treeFindAllHelper(current.left, value, foundValues)
	treeFindAllHelper(current.right, value, foundValues)
}

func treeContainsHelper[T TreeNodeValue](current *TreeNode[T], value T) bool {
	if current == nil {
		return false
	}

	if value == current.Value {
		return true
	}
	return treeContainsHelper(current.left, value) || treeContainsHelper(current.right, value)
}

func treeSizeHelper[T any](current *TreeNode[T]) int {
	if current == nil {
		return 0
	}
	return treeSizeHelper(current.left) + treeSizeHelper(current.right) + 1
}

func treeMinHelper[T TreeNodeValue](current *TreeNode[T], currMin *T) {
	if current == nil {
		return
	}
	if current.Value < *currMin {
		*currMin = current.Value
	}
	treeMinHelper(current.left, currMin)
	treeMinHelper(current.right, currMin)
}

func treeMaxHelper[T TreeNodeValue](current *TreeNode[T], currMax *T) {
	if current == nil {
		return
	}
	if current.Value < *currMax {
		*currMax = current.Value
	}
	treeMaxHelper(current.left, currMax)
	treeMaxHelper(current.right, currMax)
}

func treeClearHelper[T any](current *TreeNode[T]) {

}

func treePreOrderValueHelper[T TreeNodeValue](current *TreeNode[T], nodes *[]*TreeNode[T]) {
	if current == nil {
		return
	}
	*nodes = append(*nodes, current)
	treePreOrderValueHelper(current.left, nodes)
	treePreOrderValueHelper(current.right, nodes)
}

func treeInOrderValueHelper[T TreeNodeValue](current *TreeNode[T], nodes *[]*TreeNode[T]) {
	if current == nil {
		return
	}
	treeInOrderValueHelper(current.left, nodes)
	*nodes = append(*nodes, current)
	treeInOrderValueHelper(current.right, nodes)
}

func treePostOrderValueHelper[T TreeNodeValue](current *TreeNode[T], nodes *[]*TreeNode[T]) {
	if current == nil {
		return
	}
	treePostOrderValueHelper(current.left, nodes)
	treePostOrderValueHelper(current.right, nodes)
	*nodes = append(*nodes, current)
}
