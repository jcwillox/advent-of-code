package main

import (
	"strconv"
	"strings"
)

var words = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func Day1Part1(source string) (interface{}, error) {
	total := 0
	for _, line := range strings.Split(source, "\n") {
		num1 := ""
		num2 := ""
		// first number
		for _, c := range line {
			if c >= '1' && c <= '9' {
				num1 = string(c)
				break
			}
		}
		// last number
		for i := len(line) - 1; i >= 0; i-- {
			c := line[i]
			if c >= '1' && c <= '9' {
				num2 = string(c)
				break
			}
		}
		value, _ := strconv.Atoi(num1 + num2)
		total += value
	}
	return total, nil
}

func Day1Part2(source string) (interface{}, error) {
	total := 0
	for _, line := range strings.Split(source, "\n") {
		num1 := ""
		num2 := ""
		// first number
		for i, c := range line {
			if c >= '1' && c <= '9' {
				num1 = string(c)
				break
			}
			for word := range words {
				if strings.HasPrefix(line[i:], word) {
					num1 += words[word]
					break
				}
			}
			if num1 != "" {
				break
			}
		}
		// last number
		for i := len(line) - 1; i >= 0; i-- {
			c := line[i]
			if c >= '1' && c <= '9' {
				num2 = string(c)
				break
			}
			for word := range words {
				if strings.HasSuffix(line[:i+1], word) {
					num2 += words[word]
					break
				}
			}
			if num2 != "" {
				break
			}
		}
		value, _ := strconv.Atoi(num1 + num2)
		total += value
	}
	return total, nil
}
