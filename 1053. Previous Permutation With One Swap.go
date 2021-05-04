package main

func prevPermOpt1(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	l := len(arr) - 1
	for ; l > 0; l-- {
		if arr[l-1] > arr[l] {
			l--
			break
		}
	}
	if l == -1 {
		return arr
	}

	r := len(arr) - 1
	for ; r > l; r-- {
		if arr[r] < arr[l] {
			break
		}
	}
	for ; r > l; r-- {
		if arr[r] != arr[r-1] {
			break
		}
	}
	temp := arr[r]
	arr[r] = arr[l]
	arr[l] = temp
	return arr
}
