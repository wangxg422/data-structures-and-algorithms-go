package list

type Node struct {
	Data any
	Next *Node
}

type List struct {
	Data *Node
}

func NewNode(data any) *Node {
	return &Node{Data: data}
}

func NewList() *List {
	return &List{Data: NewNode(nil)}
}

type INode interface {
	Len() int
	First() *Node
	Last() *Node
	AppendStart(data any)
	AppendEnd(data any)
}

func (s *List) Len() int {
	if s.Data == nil {
		return 0
	}

	len := 0
	p := s.Data
	for p != nil {
		p = p.Next
		len++
	}

	return len
}

func (s *List) First() *Node {
	if s.Data == nil {
		return nil
	}

	return s.Data
}

func (s *List) Last() *Node {
	p := s.Data
	for p != nil {
		p = p.Next
	}

	return p
}

func (s *List) AppendFirst(data any) {
	node := NewNode(data)
	if s.Data == nil {
		s.Data = node
		return
	}

	node.Next = s.Data
	s.Data = node
}

func (s *List) AppendEnd(data any) {
	node := NewNode(data)
	if s.Data == nil {
		s.Data = node
		return
	}

	p := s.Data
	for p.Next != nil {
		p = p.Next
	}

	p.Next = node
}

func main() {

}
