package linkedlist_test

import (
	"go-dsa/datastructures/linkedlist"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingly_InsertTail(t *testing.T) {
	list := linkedlist.NewSingly[int]()
	list.InsertTail(1)
	list.InsertTail(2)
	list.InsertTail(3)
	list.InsertTail(4)

	want := []int{1, 2, 3, 4}
	got := []int{}
	cur := list.Head
	got = append(got, cur.Data)

	for cur.Next != nil {
		cur = cur.Next
		got = append(got, cur.Data)
	}

	assert.Equal(t, want, got)
	assert.Equal(t, 4, list.Len())
}

func TestSingly_Insert(t *testing.T) {
	list := linkedlist.NewSingly[int]()

	err := list.Insert(3, -2)
	assert.Equal(t, linkedlist.ErrPosOutOfRange, err)

	list.Insert(1, 0)
	list.Insert(2, 1)
	list.Insert(3, 2)
	list.Insert(4, 3)
	list.Insert(12, 2)

	err = list.Insert(5, 7)
	assert.Equal(t, linkedlist.ErrPosOutOfRange, err)

	want := []int{1, 2, 12, 3, 4}
	got := []int{}
	cur := list.Head
	got = append(got, cur.Data)

	for cur.Next != nil {
		cur = cur.Next
		got = append(got, cur.Data)
	}

	assert.Equal(t, want, got)
	assert.Equal(t, 5, list.Len())
}
