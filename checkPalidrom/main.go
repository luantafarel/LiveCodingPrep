package main

import (
	"fmt"
	"strings"
)

func checkPalindrome(s string) bool {
	sArray := strings.Split(strings.ToUpper(s), "")
	i := 0
	for j := len(sArray); i < j; j-- {
		if sArray[i] != sArray[j-1] {
			return false
		}
		i++
	}
	return true
}

func main() {
	fmt.Println(checkPalindrome("mam"))
	return
}
