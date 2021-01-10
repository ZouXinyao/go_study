package _8arrayAndSlice

import "testing"

func TestArray(t *testing.T) {
	// 所有元素都是0
	ary01 := [5]int{}
	t.Log("len:", len(ary01), ", cap:", cap(ary01), ", vals:", ary01)

	// 后面的元素为0
	ary02 := [5]int{1, 2}
	t.Log("len:", len(ary02), ", cap:", cap(ary02), ", vals:", ary02)

	// 这样arya03就变成切片了
	ary03 := ary02[1:3]
	t.Log("len:", len(ary03), ", cap:", cap(ary03), ", vals:", ary03)

	ary04 := ary03[:cap(ary03)]
	t.Log("len:", len(ary04), ", cap:", cap(ary04), ", vals:", ary04)

	ary05 := append(ary04, 1, 3, 5)
	t.Log("len:", len(ary05), ", cap:", cap(ary05), ", vals:", ary05)
}

func TestTraversal(t *testing.T) {
	ary01 := [5]int{1,2,3,4,5}
	t.Log(ary01)

	for i := 0; i < len(ary01); i++ {
		t.Log("index: ", i, "visit with index: ", ary01[i])
	}

	// 不使用 i 或者 val 时，可以用 _ 代替
	for i, val := range ary01 {
		t.Log("index: ", i, "value: ", val, "visit with index: ", ary01[i])
	}
}

func TestSlice(t *testing.T) {
	// make创建切片时候，make(类型，len，cap)
	s01 := make([]int, 5, 8)
	t.Log("s01", "len:", len(s01), ", cap:", cap(s01), ", vals:", s01)

	// 第三个值可以省略，这时len和cap相等
	s02 := make([]int, 5)
	t.Log("s02", "len:", len(s02), ", cap:", cap(s02), ", vals:", s02)

	// 也能像array那样定义，区别是[]内没有东西
	s03 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	t.Log("s03", "len:", len(s03), ", cap:", cap(s03), ", vals:", s03)

	// 用[i: j]取一部分，[i, j)左闭右开
	s04 := s03[2:6]
	t.Log("s04", "len:", len(s04), ", cap:", cap(s04), ", vals:", s04)

	// i前面的丢了，j后面的隐藏了，j后面的可以拿出来。
	s05 := s04[:cap(s04)]
	t.Log("s05", "len:", len(s05), ", cap:", cap(s05), ", vals:", s05)

	// 可以使用append(s, x)，x可以是想要添加的元素(1个或多个都行)
	s06 := append(s05, 1)
	// 可以是切片
	s07 := append(s05, make([]int, 3)...)

	tmp01 := []int{10, 20, 30}
	s08 := append(s05, tmp01...)

	// 不可以是数组
	//tmp02 := [3]int{10, 20, 30}
	//s09 := append(s05, tmp02...)

	t.Log("s06", "len:", len(s06), ", cap:", cap(s06), ", vals:", s06)
	t.Log("s07", "len:", len(s07), ", cap:", cap(s07), ", vals:", s07)
	t.Log("s08", "len:", len(s08), ", cap:", cap(s08), ", vals:", s08)

	/*
	append可以自动扩容，slice较短就是扩容2倍；
	如果不超过原切片的cap，底层数组不换地址；
	超过后，需要扩容，整体搬移到新地址。
	 */
}