package main

import (
	"fmt"
	"sort"
)

type Order struct {
	User     string
	Quantity int
}

func main() {
	orders := []Order{
		{"John", 10}, {"Mary", 25}, {"Peter", 15},
		{"Anna", 30}, {"John", 20}, {"Carlos", 40},
		{"Anna", 12}, {"Mary", 18}, {"Bruno", 22},
	}

	totals := make(map[string]int)
	for _, o := range orders {
		totals[o.User] += o.Quantity
	}

	type UserTotal struct {
		Name  string
		Total int
	}
	var users []UserTotal
	for name, total := range totals {
		users = append(users, UserTotal{name, total})
	}

	sort.Slice(users, func(i, j int) bool {
		return users[i].Total > users[j].Total
	})

	fmt.Println("Top 3 users:")
	for i := 0; i < 3 && i < len(users); i++ {
		fmt.Printf("%d. %s: %d units\n", i+1, users[i].Name, users[i].Total)
	}
}
