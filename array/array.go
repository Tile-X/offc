package array

import (
	"fmt"
	"github.com/tile-x/offc/m"
	"strconv"
)

const (
	defaultArraySize = 10
)

// Array is just an off-heap array
type Array[T any] struct {
	data []T
	len  int
}

// Init will initialize the array with default size
func (a *Array[T]) Init() *Array[T] {
	a.data = m.AllocSlice[T](defaultArraySize)
	a.len = 0
	return a
}

// InitWithSize will initialize the array with initial size
func (a *Array[T]) InitWithSize(initialSize int) *Array[T] {
	if initialSize > 0 {
		a.data = m.AllocSlice[T](uint(initialSize))
		a.len = 0
	} else if initialSize == 0 {

	} else {
		panic("illegal initial size: " + strconv.Itoa(initialSize))
	}
	return a
}

// InitWithSlice will initialize the array and copy the data to array
func (a *Array[T]) InitWithSlice(data []T) *Array[T] {
	if len(data) != 0 {
		a.len = len(data)
		a.data = m.AllocSlice[T](uint(len(data)))
		copy(a.data, data)
	}
	return a
}

// New returns an initialized array
func New[T any]() *Array[T] { return (&Array[T]{}).Init() }

// Empty returns true if array is empty
func (a *Array[T]) Empty() bool { return a.len == 0 }

// Len returns the number of elements of array a
// The complexity is O(1).
func (a *Array[T]) Len() int { return a.len }

func (a *Array[T]) Swap(x, y int) {
	tmp := a.data[x]
	a.data[x] = a.data[y]
	a.data[y] = tmp
}

// Front returns the first element of array a or panic if the array is empty
func (a *Array[T]) Front() T {
	a.emptyCheck()
	return a.data[0]
}

// Back returns the last element of array a or panic if the array is empty
func (a *Array[T]) Back() T {
	a.emptyCheck()
	return a.data[a.len-1]
}

// At returns the i-th element of array a or panic if the index out of range
func (a *Array[T]) At(i int) T {
	a.rangeCheck(i)
	return a.data[i]
}

// Set will set the value of index i to v
func (a *Array[T]) Set(i int, v T) {
	a.rangeCheck(i)
	a.data[i] = v
}

// emptyCheck do empty check
func (a *Array[T]) emptyCheck() {
	if a.Empty() {
		panic("array is empty")
	}
}

// rangeCheck do empty check
func (a *Array[T]) rangeCheck(i int) {
	if i < 0 || i >= a.len {
		panic("index out of range with: index " + strconv.Itoa(i))
	}
}

// lazyInit initialize the array only if the array is not initialized
func (a *Array[T]) lazyInit() {
	if a.data == nil {
		a.Init()
	}
}

// cap returns the capacity of the array
func (a *Array[T]) cap() int {
	return cap(a.data)
}

// grow will enlarge the capacity
func (a *Array[T]) grow(minCap int) {
	oldCap := a.cap()
	newCap := oldCap + (oldCap >> 1)
	if newCap < 0 || newCap < minCap {
		newCap = minCap
	}
	a.data = m.ReallocSlice(a.data, uint(newCap))
}

// EnsureCapacity make sure the capacity >= minCap
func (a *Array[T]) EnsureCapacity(minCap int) {
	if minCap < 0 {
		panic("out of memory")
	}
	if a.cap()-minCap < 0 {
		a.grow(minCap)
	}
}

// PushFront add v at the front of the array
func (a *Array[T]) PushFront(v T) {
	a.PushBack(v)
	a.cycleShiftRight(0)
}

// PushBack add v at the end of the array
func (a *Array[T]) PushBack(v T) {
	a.lazyInit()
	a.EnsureCapacity(a.len + 1)
	a.data[a.len] = v
	a.len++
}

// cycleShiftLeft will shift the array a[i:] to left once
func (a *Array[T]) cycleShiftLeft(i int) {
	if i == a.len {
		return
	}
	first := a.data[i]
	for ; i < a.len-1; i++ {
		a.data[i] = a.data[i+1]
	}
	a.data[i] = first
}

// cycleShiftRight will shift the array a[i:] to right once
func (a *Array[T]) cycleShiftRight(i int) {
	if i == a.len {
		return
	}
	last := a.data[a.len-1]
	for j := a.len - 1; j > i; j-- {
		a.data[j] = a.data[j-1]
	}
	a.data[i] = last
}

// InsertBefore insert v before index i
func (a *Array[T]) InsertBefore(i int, v T) {
	a.rangeCheck(i)
	a.PushBack(v)
	a.cycleShiftRight(i)
}

// InsertAfter insert v after index i
func (a *Array[T]) InsertAfter(i int, v T) {
	a.rangeCheck(i)
	a.PushBack(v)
	a.cycleShiftRight(i + 1)
}

// PopFront pop the first element of array
func (a *Array[T]) PopFront() {
	a.emptyCheck()
	a.cycleShiftLeft(0)
	a.len--
}

// PopBack pop the last element of array
func (a *Array[T]) PopBack() {
	a.emptyCheck()
	a.len--
}

// RemoveAt remove the element at index
func (a *Array[T]) RemoveAt(index int) {
	a.rangeCheck(index)
	a.cycleShiftLeft(index)
	a.len--
}

// RemoveAll remove all the elements in array
func (a *Array[T]) RemoveAll() {
	a.len = 0
}

// Free will free the array
func (a *Array[T]) Free() {
	if a.data == nil {
		return
	}
	m.FreeSlice(a.data)
	a.len = 0
	a.data = nil
}

func (a *Array[T]) String() string {
	return fmt.Sprintf("%v", a.data[:a.len])
}
