package main

import (
	"fmt"
	"strings"
)

func removeDuplicates(num []string) string {
	mapNum := make(map[string]string)

	for i := 0; i < len(num); i++ {
		mapNum[num[i]] = num[i]
	}
	output := make([]string, 0, len(mapNum))
	for _, val := range mapNum {
		output = append(output, val)
	}
	return strings.Join(output, "")
}
func main() {
	s := strings.Split("sajsndabdabdbajsnajdaijdbnajsbasjondajndjsanjoa", "")
	fmt.Println(removeDuplicates(s))
	return
}
