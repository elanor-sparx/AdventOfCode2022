package AdventOfCode2022

// nicked from https://www.educative.io/answers/how-to-implement-a-stack-in-golang

type Stack []rune

// IsEmpty: check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *Stack) Push(r rune) {
	*s = append(*s, r) // Simply append the new value to the end of the stack
}

// Remove and return top element of stack. Return false if stack is empty.
func (s *Stack) Pop() (rune, bool) {
	if s.IsEmpty() {
		return -1, false
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the stack by slicing it off.
		return element, true
	}
}

// Remove and return top elements of stack. Return false if stack is empty.
func (s *Stack) PopMulti(n int) ([]rune, bool) {
	if len(*s) < n {
		return nil, false
	} else {
		index := len(*s) - n     // Get the index of the top most element.
		elements := (*s)[index:] // Index into the slice and obtain the element.
		*s = (*s)[:index]        // Remove it from the stack by slicing it off.
		return elements, true
	}
}

func (s *Stack) Peek() rune {
	if s.IsEmpty() {
		return -1
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		return element
	}
}
