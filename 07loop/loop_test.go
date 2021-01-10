package _7loop

import (
	"testing"
)

func TestRange(t *testing.T)  {
	numbers1 := []int{1, 2, 3, 4, 5, 6}
	// 只有一个值时候，i是索引
	for i := range numbers1 {
		if i == 3 {
			numbers1[i] |= i
		}
	}
	t.Log("nums01", numbers1)

	// range表达式只会在for语句开始执行时被求值一次；
	// range表达式的求值结果会被复制。
	numbers2 := []int{1, 2, 3, 4, 5, 6}
	maxIndex2 := len(numbers2) - 1
	// 这里每次循环时候的e，在第一次开始执行for循环时候就已经确定了，不是动态确定的。
	for i, e := range numbers2 {
		if i == maxIndex2 {
			numbers2[0] += e
		} else {
			numbers2[i+1] += e
		}
	}
	t.Log("nums02", numbers2)
}

func TestLoop(t *testing.T)  {
	numbers3 := []int{1, 2, 3, 4, 5, 6}
	for i := 0; i < len(numbers3); i++ {
		t.Logf("index %d is %d", i, numbers3[i])
	}
	i := 0
	// 也可以像c中的while一样，只有判断条件
	for i < len(numbers3) {
		t.Logf("index %d is %d", i, numbers3[i])
		i++
	}
}