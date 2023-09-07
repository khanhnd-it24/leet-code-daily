package valid_parentheses

/*
Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:
 - Open brackets must be closed by the same type of brackets.
 - Open brackets must be closed in the correct order.
 - Every close bracket has a corresponding open bracket of the same type.

Example 1:
	Input: s = "()"
	Output: true

Example 2:
	Input: s = "()[]{}"
	Output: true

Example 3:
	Input: s = "(]"
	Output: false

*/

var closingToOpening = map[byte]byte{
	')': '(',
	'}': '{',
	']': '[',
}

// Method 1: takes a lot of memory

func isValid(s string) bool {
	str := []byte(s)
	stack := NewStack()
	for _, ch := range str {
		matchingOpening, isClosing := closingToOpening[ch]

		if !isClosing {
			stack.Push(ch)
			continue
		}
		lastOpening := stack.Pop()
		if lastOpening == nil || lastOpening != matchingOpening {
			return false
		}
	}

	return stack.Len() == 0
}

type (
	Stack struct {
		top    *node
		length int
	}
	node struct {
		value interface{}
		prev  *node
	}
)

func NewStack() *Stack {
	return &Stack{nil, 0}
}

func (s *Stack) Len() int {
	return s.length
}

func (s *Stack) Peek() interface{} {
	if s.length == 0 {
		return nil
	}
	return s.top.value
}

func (s *Stack) Pop() interface{} {
	if s.length == 0 {
		return nil
	}

	n := s.top
	s.top = n.prev
	s.length--
	return n.value
}

func (s *Stack) Push(value interface{}) {
	n := &node{value, s.top}
	s.top = n
	s.length++
}

// Method 2: takes less memory

func isValid2(str string) bool {
	s := []byte(str)
	stack := make([]byte, 0)

	for _, ch := range s {
		matchingOpening, isClosing := closingToOpening[ch]

		if !isClosing {
			stack = append(stack, ch)
			continue
		}

		if len(stack) == 0 {
			return false
		}

		lastOpening := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if lastOpening != matchingOpening {
			return false
		}
	}

	return len(stack) == 0
}
