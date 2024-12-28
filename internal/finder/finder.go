package finder

// FindMostFrequent returns the most frequently occurring element in the given slice.
// If there are multiple elements with the same highest frequency, it returns the first one.
func FindMostFrequent(items []string) string {
    // Handle empty input
    if len(items) == 0 {
        return ""
    }

    // Create frequency map
    frequency := make(map[string]int)
    
    // Count occurrences of each item
    var maxCount int
    var mostFrequent string
    
    for _, item := range items {
        frequency[item]++
        
        // Update most frequent if current item has higher count
        if frequency[item] > maxCount {
            maxCount = frequency[item]
            mostFrequent = item
        }
    }
    
    return mostFrequent
}