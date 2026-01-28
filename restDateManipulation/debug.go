package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Item struct {
	Date  string `json:"date"`
	Value int    `json:"value"`
}

func main() {
	resp, _ := http.Get("https://coderbyte.com/api/challenges/json/date-list")
	defer resp.Body.Close()

	var items []Item
	json.NewDecoder(resp.Body).Decode(&items)

	fmt.Printf("Dados originais da API:\n")
	for _, item := range items {
		fmt.Printf("  %s -> %d\n", item.Date, item.Value)
	}

	fmt.Printf("\nProcessando datas:\n")
	for _, item := range items {
		t, _ := time.Parse(time.RFC3339, item.Date)
		year, month, day := t.Date()
		normalizedDay := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)

		fmt.Printf("  Original: %s -> Normalizada: %s -> Formatada: %s\n",
			item.Date,
			normalizedDay.String(),
			normalizedDay.Format("2006-01-02")+"T05:00:00.000Z")
	}
}

