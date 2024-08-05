package m

import "unsafe"

// SizeOf return the size of T
func SizeOf[T any]() uint {
	var o T
	return uint(unsafe.Sizeof(o))
}

// Alloc allocates one of T and returns a pointer to T
func Alloc[T any]() *T {
	return (*T)(malloc(SizeOf[T]()))
}

// AllocSlice allocates the required num of T and return a slice of T
// NOTE: Do not append to the slice
func AllocSlice[T any](size uint) []T {
	return unsafe.Slice(
		(*T)(malloc(SizeOf[T]()*size)),
		size,
	)
}

// CallocSlice is the same with AllocSlice expect clear all memories
// NOTE: Do not append to the slice
func CallocSlice[T any](size uint) []T {
	return unsafe.Slice(
		(*T)(calloc(size, SizeOf[T]())),
		size,
	)
}

// ReallocSlice allocates new memory without change the data
// NOTE: Do not append to the slice
func ReallocSlice[T any](slice []T, newSize uint) []T {
	return unsafe.Slice(
		(*T)(realloc(unsafe.Pointer(&slice[0]), SizeOf[T]()*newSize)),
		newSize,
	)
}

// Free will free the memories allocated by Alloc
func Free[T any](t *T) {
	free(unsafe.Pointer(t))
}

// FreeSlice will free the slice allocated by AllocSlice, CallocSlice or ReallocSlice
func FreeSlice[T any](slice []T) {
	free(unsafe.Pointer(&slice[0]))
}
