package main

import (
	"fmt"
)

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

func printTree(node *TreeNode, level int) {
	if node == nil {
		return
	}

	// Imprime a indentação baseada no nível atual
	for i := 0; i < level; i++ {
		fmt.Print("  ")
	}

	// Imprime o valor do nó
	fmt.Println(node.Value)

	// Imprime os nós filhos recursivamente
	printTree(node.Left, level+1)
	printTree(node.Right, level+1)
}

func main() {
	// Criando uma árvore binária de exemplo
	root := &TreeNode{
		Value: 1,
		Left: &TreeNode{
			Value: 2,
			Left: &TreeNode{
				Value: 4,
			},
			Right: &TreeNode{
				Value: 5,
			},
		},
		Right: &TreeNode{
			Value: 3,
			Left: &TreeNode{
				Value: 6,
			},
			Right: &TreeNode{
				Value: 7,
			},
		},
	}

	// Imprime a árvore
	fmt.Println("Árvore Binária:")
	printTree(root, 0)
}
