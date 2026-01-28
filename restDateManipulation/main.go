package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sort"
	"time"
)

type Item struct {
	Date  string `json:"date"`
	Value int    `json:"value"`
}

// firstReverse inverte a string passada como parâmetro
func firstReverse(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	resp, err := http.Get("https://coderbyte.com/api/challenges/json/date-list")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var items []Item
	json.Unmarshal(body, &items)

	var min, max time.Time
	const day = "2006-01-02"

	for i, it := range items {
		t, _ := time.Parse(time.RFC3339, it.Date)
		year, month, day := t.Date()
		normalizedDay := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
		if i == 0 || normalizedDay.Before(min) {
			min = normalizedDay
		}
		if i == 0 || normalizedDay.After(max) {
			max = normalizedDay
		}
	}
	var out []Item

	existingDates := map[string]bool{}
	for _, item := range items {
		t, _ := time.Parse(time.RFC3339, item.Date)
		year, month, day := t.Date()
		dk := fmt.Sprintf("%04d-%02d-%02d", year, month, day)
		out = append(out, Item{
			Date:  item.Date, // Manter horário original
			Value: item.Value,
		})
		existingDates[dk] = true
	}
	for d := min; !d.After(max); d = d.Add(24 * time.Hour) {
		k := d.Format(day)
		if !existingDates[k] {
			out = append(out, Item{
				Date:  k + "T05:00:00.000Z",
				Value: 0,
			})
		}
	}
	// Ordenar por data (earliest to latest)
	sort.Slice(out, func(i, j int) bool {
		return out[i].Date < out[j].Date
	})

	b, _ := json.Marshal(out)
	fmt.Print(string(b))
}
