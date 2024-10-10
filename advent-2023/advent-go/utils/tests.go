package utils

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func getDayPart(tb testing.TB) (string, string) {
	name := strings.TrimPrefix(tb.Name(), "TestDay")
	name = strings.TrimPrefix(name, "BenchmarkDay")
	parts := strings.SplitN(name, "Part", 2)

	day := parts[0]
	part := parts[1]

	return day, part
}

func loadSample(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		return ""
	}
	data = bytes.ReplaceAll(data, []byte("\r\n"), []byte("\n"))
	data = bytes.TrimSpace(data)
	return string(data)
}

func loadCases(tb testing.TB) ([]string, []string) {
	day, part := getDayPart(tb)

	outputs, err := filepath.Glob(fmt.Sprintf("../samples/day%s/p%s.*.output.txt", day, part))
	if err != nil {
		tb.Fatalf("failed to find samples: %v", err)
	}
	for i, output := range outputs {
		outputs[i] = loadSample(output)
	}

	inputs := make([]string, len(outputs))

	for i := range outputs {
		inputs[i] = loadSample(fmt.Sprintf("../samples/day%s/p%s.s%d.input.txt", day, part, i+1))
		if inputs[i] == "" {
			inputs[i] = loadSample(fmt.Sprintf("../samples/day%s/s%d.input.txt", day, i+1))
		}
	}

	return inputs, outputs
}

func RunTests(t *testing.T, f func(string) (interface{}, error)) {
	inputs, outputs := loadCases(t)

	for i, input := range inputs {
		t.Run(fmt.Sprint("Sample", i+1), func(t *testing.T) {
			result, err := f(input)
			if err != nil {
				t.Errorf("failed running test: %v", err)
			}
			assert.Equal(t, outputs[i], fmt.Sprintf("%v", result))
		})
	}
}

func RunBenchmarks(b *testing.B, f func(string) (interface{}, error)) {
	inputs, _ := loadCases(b)

	b.ResetTimer()

	for i, input := range inputs {
		b.Run(fmt.Sprint("Sample", i+1), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _ = f(input)
			}
		})
	}
}
