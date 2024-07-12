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

func binaryTreeInsertNode[T TreeNodeValue](current, insertNode *TreeNode[T]) {
	if current == nil || insertNode == nil {
		return
	}

	insertNode.left = nil
	insertNode.right = nil

	if insertNode.Value <= current.Value {
		if current.left == nil {
			current.left = insertNode
			return
		}
		binaryTreeInsertNode(current.left, insertNode)

	} else {
		if current.right == nil {
			current.right = insertNode
			return
		}
		binaryTreeInsertNode(current.right, insertNode)
	}

}

// removes any node that in within currents sub trees else is current == remove then this will not work
func detachNodeFromTree[T TreeNodeValue](current *TreeNode[T], remove *TreeNode[T]) {
	if current == nil || remove == nil {
		return
	}

	var replacementNode *TreeNode[T]
	if current.left == remove || current.right == remove {

		replacementNode = findAndDetachReplacementNode(remove)

		//point the current node (left or right ) to the replacement node
		if current.left == remove {
			current.left = replacementNode
		} else if current.right == remove {
			current.right = replacementNode
		}
		//else if current == remove {
		//	//only time is safe to do this is when current == remove and current does not have a parent
		//	(this is not guaranteed since there is no way of knowing if current is a root node )
		//	replacementNode.left = current.left
		//	replacementNode.right = current.right
		//}
		remove.left = nil
		remove.right = nil
	} else {
		detachNodeFromTree(current.left, remove)
		detachNodeFromTree(current.right, remove)
	}

}

// Be careful if root is not actually a root node then the tree will be messed up and cause issues when doing other operations
func detachRootNode[T TreeNodeValue](root **TreeNode[T]) {
	if root == nil {
		return
	}
	replacementNode := findAndDetachReplacementNode(*root)
	//this is redundant because the previous function call does this, but keeping it somehow makes me feel better so I will keep it for now
	if replacementNode != nil {
		replacementNode.left = (*root).left
		replacementNode.right = (*root).right
	}
	(*root) = replacementNode
}

// this function will find the correct node to use to replace whatever was passed in as argument and will detach it from the tree
// if the replacementNode is having only one subtree or if it's a leaf node. lastly, the newly found and detached nodes children will point
// to the children of whatever was passed in as argument
func findAndDetachReplacementNode[T TreeNodeValue](replace *TreeNode[T]) *TreeNode[T] {
	var replacementNode *TreeNode[T]
	replacementNode = nil
	if replace.left != nil {
		replacementNode = replace.left
		getMaxLeafOrSingleSubTreeNode(replace.left, &replacementNode)
	} else if replace.right != nil {
		replacementNode = replace.right
		getMinLeafOrSingleSubTreeNode(replace.right, &replacementNode)
	} // else it means 'replace' has no children and is a leaf node thus 'replacementNode' will be nil

	if isSingleSubTreeNode(replacementNode) {
		//replacement node is not a leaf node, but it only has one subtree( only one left or right node but not both)
		// in this case just point the current node (left or right ) to the only subtree in replacement node and free replacementNode
		detachSingleSubtreeNodeAndFixTree(replace, replacementNode)
	}
	// at this point if replacementNode is not nil it means it either was a leaf node or it had one subtree which in the
	//previous block of code we freed so replacementNode's children are nil and ready to point to something new

	if replacementNode != nil {
		replacementNode.left = replace.left
		replacementNode.right = replace.right
	}

	return replacementNode // if replacementNode is nil it means that 'replace' was a leaf node
}

// this function assumes that 'singleSubTreeNode' has only one left or right node and from a starting position 'current'
// it will recursively find that node (stopping one node before it hits 'singleSubTreeNode'). Then, it will
// take 'singleSubTreeNode's parent and point it to 'singleSubTreeNode's subtree.
// at the end 'singleSubTreeNode's left and right pointers should be free to safely point to something else
func detachSingleSubtreeNodeAndFixTree[T any](current *TreeNode[T], singleSubTreeNode *TreeNode[T]) {
	if current == nil || singleSubTreeNode == nil || !isSingleSubTreeNode(singleSubTreeNode) {
		return
	}

	if current.left == singleSubTreeNode || current.right == singleSubTreeNode {
		var subTree *TreeNode[T]

		//find and assign new subtre root
		if singleSubTreeNode.left != nil {
			subTree = singleSubTreeNode.left
		} else {
			subTree = singleSubTreeNode.right

		}

		//fix tree
		if current.left == singleSubTreeNode {
			current.left = subTree
		}
		if current.right == singleSubTreeNode {
			current.right = subTree
		}
		//detach
		singleSubTreeNode.left = nil
		singleSubTreeNode.right = nil

	} else {
		detachSingleSubtreeNodeAndFixTree(current.left, singleSubTreeNode)
		detachSingleSubtreeNodeAndFixTree(current.right, singleSubTreeNode)
	}
}

func findFirstHelper[T TreeNodeValue](current *TreeNode[T], foundNode **TreeNode[T], value T) {
	if current == nil || *foundNode != nil {
		return
	}

	if value == current.Value {
		*foundNode = current
	} else {
		findFirstHelper(current.left, foundNode, value)
		findFirstHelper(current.right, foundNode, value)
	}
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

func minHelper[T TreeNodeValue](current *TreeNode[T], currMin **TreeNode[T]) {
	if current == nil {
		return
	}
	if current.Value <= (*currMin).Value {
		*currMin = current
	}
	minHelper(current.left, currMin)
	minHelper(current.right, currMin)
}

func maxHelper[T TreeNodeValue](current *TreeNode[T], currMax **TreeNode[T]) {
	if current == nil {
		return
	}
	if current.Value >= (*currMax).Value {
		*currMax = current
	}
	maxHelper(current.left, currMax)
	maxHelper(current.right, currMax)
}

func getMinLeafOrSingleSubTreeNode[T TreeNodeValue](current *TreeNode[T], minSoFar **TreeNode[T]) {
	if current == nil {
		return
	}
	if current.Value <= (*minSoFar).Value {
		if isLeafNode(current) || (current.left == nil && current.right != nil) {
			*minSoFar = current
		}

	}
	getMinLeafOrSingleSubTreeNode(current.left, minSoFar)
	getMinLeafOrSingleSubTreeNode(current.right, minSoFar)
}

func getMaxLeafOrSingleSubTreeNode[T TreeNodeValue](current *TreeNode[T], maxSoFar **TreeNode[T]) {
	if current == nil {
		return
	}

	if current.Value >= (*maxSoFar).Value {
		if isLeafNode(current) || (current.left != nil && current.right == nil) {
			*maxSoFar = current
		}
	}
	getMaxLeafOrSingleSubTreeNode(current.right, maxSoFar)
	getMaxLeafOrSingleSubTreeNode(current.left, maxSoFar)

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

func isSingleSubTreeNode[T any](node *TreeNode[T]) bool {
	if node == nil {
		return false
	}
	return !(node.left != nil && node.right != nil)
}

func isLeafNode[T any](node *TreeNode[T]) bool {
	if node == nil {
		return false
	}
	return node.left == nil && node.right == nil
}
