package main

import (
	"fmt"
	"slices"
	"sort"
	"strings"
	"sync"
	"time"
)

func anagramChecker(s1, s2 string) bool {
	arrS1 := strings.Split(s1, "")
	arrS2 := strings.Split(s2, "")
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		slices.Sort(arrS1)
	}()
	go func() {
		defer wg.Done()
		slices.Sort(arrS2)
	}()
	wg.Wait()
	return strings.Join(arrS1, "") == strings.Join(arrS2, "")
}

func bubbleSortAsc(a []int) []int {
	for i := 0; i < len(a)-1; i++ {
		if a[i] > a[i+1] {
			c := a[i]
			a[i] = a[i+1]
			a[i+1] = c
			i = -1
			continue
		}
	}
	return a
}

func charFrequency(s string) map[string]int {
	m := make(map[string]int)
	st := strings.Split(s, "")
	for _, c := range st {
		m[c]++
	}
	return m
}

func removeDuplicates(s string) string {
	m := make(map[string]int)
	var out []string
	st := strings.Split(s, "")
	for _, c := range st {
		m[c] = 1
	}
	for k := range m {
		out = append(out, k)
	}
	return strings.Join(out, "")
}

func checkPalindrome(s string) bool {
	t := true
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[(len(s)-1)-i] {
			t = false
		}
	}
	return t
}

func evenNumbers() {
	ch := make(chan int, 200)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 2; i < 100; i += 2 {
			ch <- i
			time.Sleep(time.Millisecond * 10) // changed to a time.Duration
		}
	}()
	go func() {
		defer wg.Done()
		for i := 100; i < 200; i += 2 {
			ch <- i
			time.Sleep(time.Millisecond * 10) // changed to a time.Duration
		}
	}()
	wg.Wait()
	for i := 0; i < 20; i++ {
		x := <-ch
		fmt.Println(x)

	}
}

// substitutions applies replacements described by subs to s.
// subs maps replacement string -> [start, end) indices on the original string s.
// Returns an error if any index is out of range or intervals overlap.
func substitutions(s string, subs map[string][]int) (string, error) {
	type ent struct {
		start int
		end   int
		repl  string
	}

	n := len(s)
	ents := make([]ent, 0, len(subs))
	for repl, idx := range subs {
		if len(idx) != 2 {
			return "", fmt.Errorf("invalid indices")
		}
		start, end := idx[0], idx[1]
		if start < 0 || end < start || end > n {
			return "", fmt.Errorf("index out of range")
		}
		ents = append(ents, ent{start: start, end: end, repl: repl})
	}

	// sort by start index so we can build the final string in one pass
	sort.Slice(ents, func(i, j int) bool { return ents[i].start < ents[j].start })

	// detect overlaps
	for i := 1; i < len(ents); i++ {
		if ents[i-1].end > ents[i].start {
			return "", fmt.Errorf("overlapping intervals")
		}
	}

	var b strings.Builder
	pos := 0
	for _, e := range ents {
		if pos < e.start {
			b.WriteString(s[pos:e.start])
		}
		b.WriteString(e.repl)
		pos = e.end
	}
	if pos < n {
		b.WriteString(s[pos:n])
	}
	return b.String(), nil
}

func main() {
	// fmt.Println(anagramChecker("golang", "galnog"))
	// fmt.Println(bubbleSortAsc([]int{2, 3, 5, 6, 1, 7, 8, 9}))
	// fmt.Println(removeDuplicates("gooeqeqdkndksannqwojqebjbajsbdijbqioasdbobnspinqlang"))
	// fmt.Println(checkPalindrome("aegoogea"))
	evenNumbers()
}
