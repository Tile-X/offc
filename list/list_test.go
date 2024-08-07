package list

import "testing"

func TestList(t *testing.T) {
	l := New[int]()
	for i := 0; i < 10; i++ {
		l.PushBack(i)
	}
	for i := 0; i < 10; i++ {
		if l.Front().Value != i {
			t.Errorf("expect %d, got %d", i, l.Front().Value)
		}
		l.Remove(l.Front())
	}
	l.Free()
}
