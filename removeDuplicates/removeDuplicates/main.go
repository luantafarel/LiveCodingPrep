package main

import (
	"fmt"
	"strings"
)

func mergeMaps(map1, map2 map[string]int) string {
	for key, val := range map2 {
		map1[key] = val
	}
	var i []string
	for key, _ := range map1 {
		i = append(i, key)
	}
	return strings.Join(i, "")
}

func removeDuplicates(s []string, c chan map[string]int) {
	m := make(map[string]int)
	for _, val := range s {
		m[val] = 1
	}
	c <- m
}

func main() {
	c := make(chan map[string]int)
	s := "sajsndabdabdbajsnajdaijdbnajsbasjondajndjsanjoa"
	arrS := strings.Split(s, "")
	go removeDuplicates(arrS[len(arrS)/2:], c)
	go removeDuplicates(arrS[:len(arrS)/2], c)
	x, y := <-c, <-c
	fmt.Println(mergeMaps(x, y))
	return
}
