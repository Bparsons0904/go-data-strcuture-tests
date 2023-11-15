package main

import "time"

type BinaryTreeNode struct {
	Value int
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

type BinaryTree struct {
	Root *BinaryTreeNode
}

func (tree *BinaryTree) Insert(value int) {
	newNode := &BinaryTreeNode{Value: value}
	if tree.Root == nil {
		tree.Root = newNode
		return
	}
	tree.Root.insert(newNode)
}

func testBinaryTree() {
	var createTimes [numberOfTests]time.Duration
	var removeTimes [numberOfTests]time.Duration

	for i := 0; i < numberOfTests; i++ {
		createTime, tree := testBinaryTreeCreate()
		createTimes[i] = createTime
		removeTime := testBinaryTreeRemove(tree)
		removeTimes[i] = removeTime
	}

	binaryTreeCreateStats := getStatistics(createTimes[:])
	binaryTreeRemoveStats := getStatistics(removeTimes[:])
	binaryTreeCombinedStats := combineStatistics(createTimes[:], removeTimes[:])
	printStatistics("Binary Tree", binaryTreeCreateStats, binaryTreeRemoveStats, binaryTreeCombinedStats)
}

func testBinaryTreeCreate() (time.Duration, *BinaryTree) {
	startTime := time.Now()
	tree := &BinaryTree{}
	for _, value := range testCreateOrder {
		tree.Insert(value)
	}
	return time.Since(startTime), tree
}

func testBinaryTreeRemove(tree *BinaryTree) time.Duration {
	startTime := time.Now()
	for _, value := range testRemoveOrder {
		tree.Delete(value)
	}
	return time.Since(startTime)
}

func (node *BinaryTreeNode) insert(newNode *BinaryTreeNode) {
	if newNode.Value < node.Value {
		if node.Left == nil {
			node.Left = newNode
		} else {
			node.Left.insert(newNode)
		}
	} else {
		if node.Right == nil {
			node.Right = newNode
		} else {
			node.Right.insert(newNode)
		}
	}
}

func (tree *BinaryTree) Delete(value int) {
	tree.Root = deleteNode(tree.Root, value)
}

func deleteNode(node *BinaryTreeNode, value int) *BinaryTreeNode {
	if node == nil {
		return nil
	}
	if value < node.Value {
		node.Left = deleteNode(node.Left, value)
	} else if value > node.Value {
		node.Right = deleteNode(node.Right, value)
	} else {
		if node.Left == nil {
			return node.Right
		} else if node.Right == nil {
			return node.Left
		}

		node.Value = minValue(node.Right)
		node.Right = deleteNode(node.Right, node.Value)
	}
	return node
}

func minValue(node *BinaryTreeNode) int {
	current := node
	for current.Left != nil {
		current = current.Left
	}
	return current.Value
}
