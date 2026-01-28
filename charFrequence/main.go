package main

import (
	"fmt"
	"strings"
)

func charFrequency(s string) map[string]int {
	strCount := make(map[string]int)
	strArr := strings.Split(s, "")
	for i := 0; i < len(strArr); i++ {
		strCount[strArr[i]]++
	}
	return strCount
}

func main() {
	fmt.Println(charFrequency("golang"))
	return
}
