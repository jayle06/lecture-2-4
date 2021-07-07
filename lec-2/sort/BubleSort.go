package sort

func Swap(x *int, y *int) {
	temp := *x
	*x = *y
	*y = temp
}

func BubbleSort(arr []int, n int) []int {
	var i, j int
	for i = 0; i < n-1; i++ {
		for j = 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				Swap(&arr[j], &arr[j+1])
			}
		}
	}
	return arr
}
