package main

import (
	"encoding/binary"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestZipList(t *testing.T) {
	list := New()
	list.Insert(NewStringEntry("foo"))
	list.Insert(NewStringEntry("bar"))
	list.Insert(NewIntEntry(10086))
	assert.Equal(t, "foo", string(list.Get(0)))
	assert.Equal(t, "bar", string(list.Get(1)))

	assert.Equal(t, uint32(10086), binary.BigEndian.Uint32(list.Get(2)))
}
