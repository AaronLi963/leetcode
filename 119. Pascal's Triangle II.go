package main

func getRow(rowIndex int) []int {
	paskal := make([][]int, rowIndex+1)
	if rowIndex == 0 {
		return []int{1}
	}
	if rowIndex == 1 {
		return []int{1, 1}
	}
	paskal[0] = []int{1}
	paskal[1] = []int{1, 1}

	for i := 2; i <= rowIndex; i++ {
		row := make([]int, i+1)
		for j := range row {
			if j == 0 || j == i {
				row[j] = 1
				continue
			}
			row[j] = paskal[i-1][j-1] + paskal[i-1][j]
		}
		paskal[i] = row
	}
	return paskal[rowIndex]
}
