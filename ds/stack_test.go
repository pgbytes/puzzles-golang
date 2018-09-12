package ds

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStack_Push(t *testing.T) {
	t.Log("testing Stack.Push")
	stack := NewStack()
	stack.Push("element1")
	stack.Push("element2")
	stack.Push("element3")
	assert.Equal(t, 3, len(stack.data), "After stack push, stack length should be 3")
	assert.Equal(t, "element1", stack.data[0], "first element in stack should be 'element1'")
}

func TestStack_Peep(t *testing.T) {
	t.Log("testing Stack.Peep")
	stack := NewStack()
	stack.Push("element1")
	stack.Push("element2")
	stack.Push("element3")
	assert.Equal(t, "element3", stack.Peep(), "top element should be element3")
}

func TestStack_Pop(t *testing.T) {
	t.Log("testing Stack.Pop")
	stack := NewStack()
	stack.Push("element1")
	stack.Push("element2")
	stack.Push("element3")
	top, err := stack.Pop()
	assert.Nil(t, err, "stack should not have been empty")
	assert.Equal(t, "element3", top, "top element should be element3")
	assert.Equal(t, 2, len(stack.data))
}

func TestStack_IsEmpty(t *testing.T) {
	t.Log("testing Stack.IsEmpty")
	stack := NewStack()
	assert.Equal(t, true, stack.IsEmpty(), "stack should have been empty")
	stack.Push("element1")
	assert.Equal(t, false, stack.IsEmpty(), "stack should not have been empty")
}
