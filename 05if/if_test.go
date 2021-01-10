package _5if

import "testing"

func TestIf(t *testing.T) {
	// go中没有else if
	for i := 0; i < 5; i++ {
		// 左大括号必须放在if所在行
		if i % 2 == 0 {
			t.Log("i is even")
		} else {
			// else必须放在if的右大括号后面，
			t.Log("i is odd")
		}
	}
}
