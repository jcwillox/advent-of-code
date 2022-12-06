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

func loadInputs(tb testing.TB) ([]string, []string) {
	name := strings.TrimPrefix(tb.Name(), "TestDay")
	name = strings.TrimPrefix(name, "BenchmarkDay")
	parts := strings.SplitN(name, "Part", 2)

	loadFiles := func(path string) []string {
		files, err := filepath.Glob(path)
		if err != nil {
			tb.Fatalf("failed to find samples: %v", err)
		}
		contents := make([]string, len(files))
		for i, file := range files {
			data, err := os.ReadFile(file)
			if err != nil {
				tb.Fatalf("failed to read samples: %v", err)
			}
			contents[i] = string(bytes.TrimSpace(bytes.ReplaceAll(data, []byte("\r\n"), []byte("\n"))))
		}
		return contents
	}

	return loadFiles("../samples/day" + parts[0] + "/*.input.txt"),
		loadFiles("../samples/day" + parts[0] + "/*." + parts[1] + ".output.txt")
}

func RunTests(t *testing.T, f func(string) (interface{}, error)) {
	inputs, outputs := loadInputs(t)

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
	inputs, _ := loadInputs(b)

	b.ResetTimer()

	for i, input := range inputs {
		b.Run(fmt.Sprint("Sample", i+1), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _ = f(input)
			}
		})
	}
}
