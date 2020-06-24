package TestSomeFunc

import "testing"

func TestSomeFunc(t *testing.T) {
	got := Abs(-1)
	if got != 1 {
		t.Errorf("ABS(-1) = %v; wanted 1\n", got)
	}
}

func Abs(n float64) float64 {
	if n < 0 {
		return -n
	}
	return n
}