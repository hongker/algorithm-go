package skiplist

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewSkipList(t *testing.T) {
	s := NewSkipList(100.00, 5)
	s.Insert(1, "a")
	s.Insert(2, "b")
	s.Insert(3, "c")

	assert.Equal(t, "a", s.Search(1))
}
