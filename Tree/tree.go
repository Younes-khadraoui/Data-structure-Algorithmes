package main

import "fmt"

type treeNode struct {
	value int
	left  *treeNode
	right *treeNode
}

func main() {
	values := []int{1, 2, 3, 4, 5, 6, 7}
	tree := createTree(values)

	inOrderTraversal(tree)
	fmt.Print("nil")
	fmt.Println()
	fmt.Println()
	levelOrderTraversal(tree)
	
}

func createTree(values []int) *treeNode {
	if len(values) == 0 {
		return nil
	}

	middle := len(values) / 2
	root := &treeNode{value: values[middle]}

	root.left = createTree(values[:middle])
	root.right = createTree(values[middle+1:])

	return root
}

func inOrderTraversal(tree *treeNode) {
	if tree == nil {
		return
	}

	inOrderTraversal(tree.left)
	fmt.Print(tree.value, " -> ")
	inOrderTraversal(tree.right)
}

func levelOrderTraversal(tree *treeNode) {
	if tree == nil {
		return
	}

	queue := []*treeNode{tree}

	for len(queue) > 0 {
		levelSize := len(queue)

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			fmt.Print(node.value, " ")

			if node.left != nil {
				queue = append(queue, node.left)
			}
			if node.right != nil {
				queue = append(queue, node.right)
			}
		}

		fmt.Println()
	}
}
