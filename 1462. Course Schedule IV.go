package main

import (
	"fmt"
)

func checkIfPrerequisite(n int, prerequisites [][]int, queries [][]int) []bool {
    preMap := map[string]bool{}
    for _,pre := range prerequisites{
        preMap[sliceToPair(pre)] = true
    }
    ans := []bool{}
    for _,q := range queries{
        ans = append(ans, preMap[sliceToPair(q)])
    }
    return ans
}

func sliceToPair(s []int) string {
    return fmt.Sprintf("%d->%d",s[0],s[1])
}

func main(){
	n := 2
	pre := [][]int{[]int{1,0}}
	query := [][]int{[]int{0,1},[]int{1,0}}
	fmt.Println(checkIfPrerequisite(n,pre,query))
}