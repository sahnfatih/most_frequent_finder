package finder

import "testing"

func TestFindMostFrequent(t *testing.T) {
    tests := []struct {
        name     string
        input    []string
        expected string
    }{
        {
            name:     "Example case",
            input:    []string{"apple", "pie", "apple", "red", "red", "red"},
            expected: "red",
        },
        {
            name:     "Empty slice",
            input:    []string{},
            expected: "",
        },
        {
            name:     "Single item",
            input:    []string{"apple"},
            expected: "apple",
        },
        {
            name:     "All items same frequency",
            input:    []string{"a", "b", "c"},
            expected: "a",
        },
        {
            name:     "Multiple items with same highest frequency",
            input:    []string{"a", "a", "b", "b"},
            expected: "a",
        },
        {
            name:     "Case sensitive",
            input:    []string{"Apple", "apple", "APPLE", "Apple"},
            expected: "Apple",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := FindMostFrequent(tt.input)
            if result != tt.expected {
                t.Errorf("FindMostFrequent() = %v, want %v", result, tt.expected)
            }
        })
    }
}

// Benchmark for performance testing
func BenchmarkFindMostFrequent(b *testing.B) {
    input := []string{"apple", "pie", "apple", "red", "red", "red"}
    for i := 0; i < b.N; i++ {
        FindMostFrequent(input)
    }
}