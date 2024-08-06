package array

import (
	"fmt"
	"testing"
)

func TestArray_InitWithSlice(t *testing.T) {
	var arr Array[int]
	nums := []int{1, 2, 3}
	arr.InitWithSlice(nums)
	for i := 0; i < len(nums); i++ {
		if arr.At(i) != nums[i] {
			t.Fatal("init with slice error")
		}
	}
}

func TestArray_PushBack(t *testing.T) {
	var a Array[int]
	for i := 0; i < 8; i++ {
		a.PushBack(i + 1)
	}
	a.PushBack(100)
	fmt.Println(a.data)
}

func TestArray_PushFront(t *testing.T) {
	var a Array[int]
	for i := 0; i < 8; i++ {
		a.PushBack(i + 1)
	}
	a.PushFront(100)
	fmt.Println(a.data)
}

func TestArray_cycleShiftLeft(t *testing.T) {
	var a Array[int]
	a.InitWithSize(7)
	for i := 0; i < 7; i++ {
		a.PushBack(i + 1)
	}
	fmt.Println(a.Len())
	fmt.Println(a.data)
	a.cycleShiftLeft(1)
	fmt.Println(a.data)
}

func TestArray_cycleShiftRight(t *testing.T) {
	var a Array[int]
	a.InitWithSize(7)
	for i := 0; i < 7; i++ {
		a.PushBack(i + 1)
	}
	fmt.Println(a.Len())
	fmt.Println(a.data)
	a.cycleShiftRight(0)
	fmt.Println(a.data)
}

func TestArray_Pop(t *testing.T) {
	var a Array[int]
	a.InitWithSlice([]int{1, 2, 3})
	a.PopBack()
	if a.Back() != 2 {
		t.Fatal("pop back error")
	}
	a.PopFront()
	if a.Front() != 2 {
		t.Fatal("pop front error")
	}
}

func TestArray_Remove(t *testing.T) {
	var a Array[int]
	a.InitWithSlice([]int{1, 2, 3, 4})
	a.RemoveAt(1)
	fmt.Println(a.String())
	a.RemoveAll()
	fmt.Println(a.String())
}
