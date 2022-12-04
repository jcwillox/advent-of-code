package main

import (
	. "advent-go/utils"
	"testing"
)

func TestDay{{ cookiecutter.day }}Part1(t *testing.T) {
	RunTests(t, Day{{ cookiecutter.day }}Part1)
}

func TestDay{{ cookiecutter.day }}Part2(t *testing.T) {
	RunTests(t, Day{{ cookiecutter.day }}Part2)
}

func BenchmarkDay{{ cookiecutter.day }}Part1(b *testing.B) {
	RunBenchmarks(b, Day{{ cookiecutter.day }}Part1)
}

func BenchmarkDay{{ cookiecutter.day }}Part2(b *testing.B) {
	RunBenchmarks(b, Day{{ cookiecutter.day }}Part2)
}
