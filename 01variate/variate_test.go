package _1variate

import "testing"

func TestVariate(t *testing.T) {
	var a int = 1
	t.Log(a)

	var b = 2
	t.Log(b)

	c := 3
	t.Log(c)

	var (
		d = 4
		e = 5
	)
	t.Log(d, e)
}

func TestSwap(t *testing.T) {
	a, b := 1, 2
	t.Log("before swap(a, b):", a, b)
	a, b = b, a
	t.Log("after  swap(a, b):", a, b)
}
