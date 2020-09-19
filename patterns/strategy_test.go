package patterns

import (
	"fmt"
	"testing"
)

func TestStrategy(t *testing.T) {
	var sort StrategySort

	arr1 := []int{31, 15, 10, 2, 4, 2, 14, 23, 12, 66}
	sort = InsertionSort{}
	context := Context{sort, arr1}

	context.Sort()
	fmt.Printf("%v\n", arr1)

	arr2 := []int{1, 5, 10, 2, 4, 12, 14, 23, 12, 66}
	sort = InsertionSort{}
	context = Context{sort, arr2}

	context.Sort()
	fmt.Printf("%v\n", arr1)

	arr3 := []int{31, 15, 10, 2, 4, 2, 14, 23, 12, 66}
	sort = BubbleSort{}
	context = Context{sort, arr3}

	context.Sort()
	fmt.Printf("%v\n", arr3)
}
