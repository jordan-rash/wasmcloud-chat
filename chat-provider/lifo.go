package main

import (
	"github.com/charmbracelet/lipgloss"
	chat "github.com/jordan-rash/wasmcloud-chat/interface"
)

// NewStack returns a new stack.
func NewStack(size int) *Stack {
	return &Stack{size: size}
}

// Stack is a basic LIFO stack that resizes as needed.
type Stack struct {
	size  int
	nodes []*chat.Msg
	count int
}

func (s Stack) Read() []*chat.Msg {
	return s.nodes
}

// Push adds a node to the stack.
func (s *Stack) Push(n *chat.Msg) {
	if !s.Find(n.Id) {
		if s.count >= s.size {
			s.Pop()
		}
		s.nodes = append(s.nodes[:s.count], n)
		s.count++
	}
}

func (s Stack) Find(id string) bool {
	for _, x := range s.nodes {
		if id == x.Id {
			return true
		}
	}
	return false
}

// Pop removes and returns a node from the stack in last to first order.
func (s *Stack) Pop() *chat.Msg {
	if s.count == 0 {
		return nil
	}
	s.count--
	return s.nodes[s.count]
}

func (s *Stack) Delete(m *chat.Msg) {
	for i, u := range s.nodes {
		if u.Owner.Name == m.Owner.Name {
			s.nodes = append(s.nodes[:i], s.nodes[i+1:]...)
			return
		}
	}
}

func (s *Stack) View() string {
	list := ""
	for _, n := range s.Read() {
		list += localUserStyle.
			Foreground(lipgloss.Color(n.Owner.Color)).
			Render(n.Owner.Name) + "\n"
	}
	return list
}
