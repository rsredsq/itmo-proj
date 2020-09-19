package pract4

type StrategySort interface {
	Title() string
	Sort(arr []int)
}

type InsertionSort struct {
}

func (s InsertionSort) Title() string {
	return "Сортировка вставками"
}

func (s InsertionSort) Sort(arr []int) {
	for i := 1; i < len(arr); i++ {
		j := 0
		buffer := arr[i]
		for j = i - 1; j >= 0; j-- {
			if arr[j] < buffer {
				break
			}
			arr[j+1] = arr[j]
		}
		arr[j+1] = buffer
	}
}

type Context struct {
	strategy StrategySort
	arr      []int
}

func (c Context) Sort() {
	c.strategy.Sort(c.arr)
}

type BubbleSort struct {
}

func (s BubbleSort) Title() string {
	return "Сортировка пузырьком"
}

func (s BubbleSort) Sort(arr []int) {
	swapped := true
	for swapped {
		swapped = false
		for i := 1; i < len(arr); i++ {
			if arr[i-1] > arr[i] {
				arr[i], arr[i-1] = arr[i-1], arr[i]
				swapped = true
			}
		}
	}
}
