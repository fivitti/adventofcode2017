package main

import (
	"errors"
	"strings"
	"../utils/intutils"
	"../utils/stringutils"
	"../utils/argparse"
	"fmt"
)

type node struct {
	name string
	weight int
	children []*node
	parent *node
}

func (n *node) addChild(child *node) {
	if n.children == nil {
		n.children = make([]*node, 0)
	}
	n.children = append(n.children, child)
}

func (n *node) getCumulativeWeight() int {
	cumulativeWeight := n.weight

	for _, child := range n.children {
		cumulativeWeight += child.getCumulativeWeight()
	}

	return cumulativeWeight
}

func (n *node) hasChildren() bool {
	return n.children != nil && len(n.children) != 0
}

func (n *node) isBalanced() bool {
	if !n.hasChildren() {
		return true
	}

	expectedWeight := n.children[0].getCumulativeWeight()

	for _, child := range n.children {
		if child.getCumulativeWeight() != expectedWeight {
			return false
		}
	}
	return true
}

func (n *node) getUnbalancedChild() *node {
	if !n.hasChildren() {
		return nil
	}

	expectedWeight := n.getDominateChildCumulativeWeight()

	for _, child := range n.children {
		if child.getCumulativeWeight() != expectedWeight {
			return child
		}
	}

	return nil
}

func (n *node) getDominateChildCumulativeWeight() int {
	if !n.hasChildren() {
		return 0
	}

	if len(n.children) < 3 {
		return n.children[0].getCumulativeWeight()
	}

	n0Weight := n.children[0].getCumulativeWeight()
	n1Weight := n.children[1].getCumulativeWeight()
	n2Weight := n.children[2].getCumulativeWeight()

	if n0Weight == n1Weight {
		return  n0Weight
	} else if n0Weight == n2Weight {
		return n0Weight
	} else {
		return n1Weight
	}
}

func main() {
	if err := execute(); err != nil {
		fmt.Println("Day 7: Recursive Circus. Parameters: /path/to/file(string)")
		fmt.Println("Error: ", err)
	}
}

func execute() error {
	if err := argparse.ValidateLength(2); err != nil {
		return err
	}

	inputMatrix, err := argparse.ReadStringMatrix(1, " ")
	if err != nil {
		return err
	}
	root, err := buildTree(inputMatrix)
	if err != nil {
		return err
	}

	fmt.Println("Root name: ", root.name)

	unbalancedNode := findUnbalancedNode(root)

	if unbalancedNode == nil {
		fmt.Println("Tree is balanced")
		return nil
	}

	fmt.Println("Unbalanced node:", unbalancedNode.name, "balance weight:", getBalanceWeight(unbalancedNode))
	return nil
}

func buildTree(input [][]string) (*node, error) {
	nodes := make(map[string]*node)

	// Create nodes
	for _, row := range input {
		name := row[0]
		weightRow := strings.Trim(row[1], "()")
		weight, err := intutils.ParseInt(weightRow)
		if err != nil {
			return nil, err
		}
		nodes[name] = &node{name:name, weight:weight}
	}

	// Append children
	for _, row := range input {
		if len(row) < 4 {
			continue
		}
		parentName := row[0]
		childrenNames := stringutils.MapStrings(row[3:], func (name string) string {
			return strings.TrimRight(name, ",")
		})

		parentNode := nodes[parentName]

		for _, childName := range childrenNames {
			childNode := nodes[childName]
			parentNode.addChild(childNode)
			childNode.parent = parentNode
		}
	}

	for _, node := range nodes {
		if node.parent == nil {
			return node, nil
		}
	}

	return nil, errors.New("invalid flow")
}

func findUnbalancedNode(node *node) *node {
	if node.isBalanced() {
		return node
	}

	unbalancedChild := node.getUnbalancedChild()
	return findUnbalancedNode(unbalancedChild)
}

func getBalanceWeight(unbalanced *node) int {
	expectedWeight := unbalanced.parent.getDominateChildCumulativeWeight()
	unbalancedWeight := unbalanced.getCumulativeWeight()
	diff := unbalanced.weight - (unbalancedWeight - expectedWeight)
	return diff
}