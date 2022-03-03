package main

import (
	"flag"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const (
	maxItems      = 100
	minItems      = 2
	maxItemLength = 32
)

var (
	letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	numbers = []rune("0123456789")
)

type genFunc func() string

var generateFuncs = map[bool]genFunc{
	true:  generateValid,
	false: generateInvalid,
}

func generateValid() string {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(maxItems)
	if n < minItems {
		n = minItems
	}
	parts := make([]string, n)
	for i := range parts {
		arr := letters
		if i%2 == 0 {
			arr = numbers
		}
		itemLength := rand.Intn(maxItemLength)
		if itemLength == 0 {
			itemLength++
		}
		s := make([]rune, itemLength)
		for j := range s {
			s[j] = arr[rand.Intn(len(arr))]
		}
		parts[i] = string(s)
	}
	return strings.Join(parts, "-")
}

func generateInvalid() string {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(maxItems)
	parts := make([]string, n)
	for i := range parts {
		arr := append(letters, numbers...)
		itemLength := rand.Intn(maxItemLength)
		if itemLength == 0 {
			itemLength++
		}
		s := make([]rune, itemLength)
		for j := range s {
			s[j] = arr[rand.Intn(len(arr))]
		}
		parts[i] = string(s)
	}
	return strings.Join(parts, "/")
}

func generate(valid bool) string {
	return generateFuncs[valid]()
}

var (
	n  = flag.Int("n", 1, "number of entries to output")
	ok = flag.Bool("v", true, "output valid entries")
)

func main() {
	flag.Parse()
	for i := 0; i < *n; i++ {
		fmt.Println(generate(*ok))
	}
}
