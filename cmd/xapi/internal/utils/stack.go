package utils

import "go/ast"

type NodeStack struct {
	node  ast.Node
	stack []ast.Node
}

// Push 压栈
func (s *NodeStack) Push(node ast.Node) {
	s.stack = append(s.stack, s.node)
	s.node = node
}

// Pop 出栈
func (s *NodeStack) Pop() {
	s.node = s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
}

// Top 倒数第n个元素
func (s *NodeStack) Top(n int) ast.Node {
	return s.stack[len(s.stack)-n]
}
