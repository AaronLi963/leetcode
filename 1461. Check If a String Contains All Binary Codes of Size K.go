package main

import (
	"fmt"
)

func hasAllCodes(s string, k int) bool {
    hash := map[string]bool{}
    idx := 0
    tries := len(s)-k +1
    for counter := 1<<k ;  counter <= tries && idx < len(s); {
        exist := hash[s[idx : idx+k]]
        if !exist{
            counter--
            if counter == 0{
                return true
            }
            hash[s[idx : idx+k]] = true
        }
        idx++
        tries--
    }    
    
    return false

}

func main(){
	s := "00110110"
	k := 2
	fmt.Println(hasAllCodes(s,k))
}