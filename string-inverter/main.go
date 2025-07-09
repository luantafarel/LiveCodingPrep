package main

import (
	"fmt"
	"strings"
)

func stringInverter(s string) string {
	var response string
	sArray := strings.Split(s, "")
	for i := len(sArray); i > 0; i-- {
		response += sArray[i-1]
	}
	return response
}

func main() {
	fmt.Println(stringInverter("golang"))
	return
}
