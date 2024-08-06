package m

import (
	"testing"
)

func TestAlloc(t *testing.T) {
	type Point struct {
		x, y int
	}
	p := Alloc[Point]()
	if p == nil {
		t.Fatal("alloc failed")
	}
}

func TestAllocSlice(t *testing.T) {
	nums := AllocSlice[int](10)
	if nums == nil || len(nums) != 10 {
		t.Fatal("alloc slice failed")
	}
	FreeSlice(nums)
}

func TestReallocSlice(t *testing.T) {
	nums := AllocSlice[int](10)
	for i := 0; i < len(nums); i++ {
		nums[i] = i + 1
	}
	nums = ReallocSlice(nums, 12)
	if nums == nil || len(nums) != 12 {
		t.Fatal("realloc failed")
	}
	nums = ReallocSlice(nums, 8)
	if nums == nil || len(nums) != 8 {
		t.Fatal("realloc failed")
	}
	FreeSlice(nums)
}

func TestCallocSlice(t *testing.T) {
	nums := CallocSlice[int](10)
	if nums == nil || len(nums) != 10 {
		t.Fatal("calloc slice failed")
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			t.Fatal("memory not clear")
		}
	}
	FreeSlice(nums)
}

func TestRelease(t *testing.T) {
	type Point struct {
		x, y int
	}
	p := Alloc[Point]()
	Release(&p)
	if p != nil {
		t.Fatal("release failed")
	}
}

func TestReleaseSlice(t *testing.T) {
	nums := AllocSlice[int](10)
	ReleaseSlice(&nums)
	if nums != nil || len(nums) != 0 {
		t.Fatal("release slice failed")
	}
}
