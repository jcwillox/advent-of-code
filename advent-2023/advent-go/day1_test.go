package main

import (
	. "advent-go/utils"
	"testing"
)

func TestDay1Part1(t *testing.T) {
	RunTests(t, Day1Part1)
}

func TestDay1Part2(t *testing.T) {
	RunTests(t, Day1Part2)
}

func BenchmarkDay1Part1(b *testing.B) {
	RunBenchmarks(b, Day1Part1)
}

func BenchmarkDay1Part2(b *testing.B) {
	RunBenchmarks(b, Day1Part2)
}
