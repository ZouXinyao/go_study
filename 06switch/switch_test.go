package _6switch

import (
	"testing"
)

func TestSwitch(t *testing.T) {
	value1 := [...]int{0, 1, 2, 3, 4, 5, 6}
	switch 1 + 1 {
	case value1[0], value1[1]:
		t.Log("0 or 1")
	case value1[2], value1[3]:
		t.Log("2 or 3")
		fallthrough				// 直接跳到下一个case
	case value1[4], value1[5], value1[6]:
		t.Log("4 or 5 or 6")
	}
}