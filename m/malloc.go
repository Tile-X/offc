package m

//#include <stdlib.h>
import "C"
import "unsafe"

/*
Just a simple wrapper for C function
*/

func malloc(size uint) unsafe.Pointer {
	return C.malloc(C.size_t(size))
}

func calloc(count, size uint) unsafe.Pointer {
	return C.calloc(C.size_t(count), C.size_t(size))
}

func free(ptr unsafe.Pointer) {
	C.free(ptr)
}

func realloc(ptr unsafe.Pointer, size uint) unsafe.Pointer {
	return C.realloc(ptr, C.size_t(size))
}
