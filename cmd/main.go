package main

import (
    "fmt"
    "most_frequent_finder/internal/finder"
    "strings"
)

func main() {
    // Test cases
    testCases := [][]string{
        {"apple", "pie", "apple", "red", "red", "red"},
        {"dog", "cat", "dog", "bird", "cat", "dog"},
        {"a", "b", "c", "a"},
        {},
    }

    // Process each test case
    for i, test := range testCases {
        result := finder.FindMostFrequent(test)
        fmt.Printf("\nTest Case %d:\n", i+1)
        fmt.Printf("Input: [%s]\n", strings.Join(test, ", "))
        fmt.Printf("Most frequent item: %s\n", result)
    }
}