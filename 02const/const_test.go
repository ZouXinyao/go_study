package _2const

import "testing"

const a = iota

const (
	b = iota
)
// 从上到下顺着数下来，iota每行+1（空行除外）
// 没有初始化的常量，复用上一行的初始化操作。
// iota也复用，但是iota的值会+1，导致当前值与上一行不一致。
// 这变态的题目出自：https://studygolang.com/articles/22468?fr=sidebar
const (
	c = 0				// c = 0		iota = 0
	d = iota			// d = iota		iota = 1
	e					// e = iota		iota = 2
	f = "hello"			// f = "hello" 	iota = 3
	// nothing
	g					// g = "hello" 	iota = 4
	h = iota			// h = iota		iota = 5
	i					// i = iota		iota = 6
	j = 0				// j = 0		iota = 7
	k					// k = 0		iota = 8
	l, m = iota, iota	// l = iota, m = iota	iota = 9	iota的值是按行计算
	n, o				// l = iota, m = iota	iota = 10

	p = iota + 1		// p = iota + 1			iota = 11
	q					// q = iota + 1			iota = 12
	_					// _ = iota + 1			iota = 13
	r = iota * iota		// r = iota * iota		iota = 14
	s					// s = iota * iota		iota = 15
	t = r				// t = r				iota = 16
	u					// u = r				iota = 17
	v = 1 << iota		// v = 1 << iota		iota = 18
	w					// w = 1 << iota		iota = 19
	x = iota * 0.01		// x = iota * 0.01		iota = 20
	y float32 = iota * 0.01 // y = iota * 0.01	iota = 21
	z						// z = iota * 0.01	iota = 22
)

func TestConst(tst *testing.T) {
	tst.Logf("a : %T = %v\n", a, a)
	tst.Logf("b : %T = %v\n", b, b)
	tst.Logf("c : %T = %v\n", c, c)
	tst.Logf("d : %T = %v\n", d, d)
	tst.Logf("e : %T = %v\n", e, e)
	tst.Logf("f : %T = %v\n", f, f)
	tst.Logf("g : %T = %v\n", g, g)
	tst.Logf("h : %T = %v\n", h, h)
	tst.Logf("i : %T = %v\n", i, i)
	tst.Logf("j : %T = %v\n", j, j)
	tst.Logf("k : %T = %v\n", k, k)
	tst.Logf("l : %T = %v\n", l, l)
	tst.Logf("m : %T = %v\n", m, m)
	tst.Logf("n : %T = %v\n", n, n)
	tst.Logf("o : %T = %v\n", o, o)
	tst.Logf("p : %T = %v\n", p, p)
	tst.Logf("q : %T = %v\n", q, q)
	tst.Logf("r : %T = %v\n", r, r)
	tst.Logf("s : %T = %v\n", s, s)
	tst.Logf("t : %T = %v\n", t, t)
	tst.Logf("u : %T = %v\n", u, u)
	tst.Logf("v : %T = %v\n", v, v)
	tst.Logf("w : %T = %v\n", w, w)
	tst.Logf("x : %T = %v\n", x, x)
	tst.Logf("y : %T = %v\n", y, y)
	tst.Logf("z : %T = %v\n", z, z)
}