package _4type

import "testing"

// 类型别名
type MyInt int64

func TestImplicit(t *testing.T) {
	var a int32 = 1
	var b int64
	// 不同类型不能隐式转换，不用的别名也不能隐式转换
	b = int64(a)
	var c MyInt
	c = MyInt(b)
	t.Log(a, b, c)
}

func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a
	//aPtr = aPtr + 1	不支持指针的+-
	t.Log(a, aPtr)
	t.Logf("%T %T", a, aPtr)
}

func TestString(t *testing.T) {
	// 字符串是一个值类型，
	var s string
	sPtr := &s
	t.Log("*" + s + "*") //初始化零值是“”
	t.Log(len(s))
	t.Log(sPtr)

}