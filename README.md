# Most Frequent Item Finder

A Go application that finds the most frequently occurring item in a given array/list of strings.

## Table of Contents
- [Features](#features)
- [Project Structure](#project-structure)
- [Requirements](#requirements)
- [Installation](#installation)
- [Usage](#usage)
- [Development](#development)
- [Testing](#testing)
- [Algorithm](#algorithm)
- [Contributing](#contributing)

## Features
- Finds most frequent item in a string array
- Handles various edge cases (empty input, equal frequencies)
- Case-sensitive comparison
- Comprehensive test coverage
- Performance optimized using hash map
- Clean and maintainable code structure

## Requirements
- Go 1.20 or higher
- Git

## Installation
1. Clone the repository:
git clone https://github.com/sahnfatih/most_frequent_finder.git

2. Navigate to project directory:
cd most_frequent_finder

3. Build the project:
go build ./cmd/main.go

## Usage
Run the program:
go run cmd/main.go

Example output:
Test Case 1:
Input: [apple, pie, apple, red, red, red]
Most frequent item: red

Test Case 2:
Input: [dog, cat, dog, bird, cat, dog]
Most frequent item: dog

## Development
### Code Structure
- All function and variable names follow Go naming conventions
- Comments are in English and follow Go best practices
- Core logic is separated into the internal package
- Input handling and main logic in cmd package

### Naming Conventions
- Functions: PascalCase for exported functions (e.g., `FindMostFrequent`)
- Variables: camelCase for internal variables (e.g., `maxCount`)
- Packages: lowercase, single-word names (e.g., `finder`)

## Testing
Run all tests:
go test ./...

Run tests with coverage:
go test ./... -cover

Run benchmarks:
go test -bench=. ./...

## Algorithm
The program uses a hash map (map[string]int) to count occurrences of each item:
1. Create an empty hash map for frequency counting
2. Iterate through the input array once
3. For each item, increment its count in the map
4. Track the highest frequency and corresponding item
5. Return the item with highest frequency

Time Complexity: O(n) where n is the length of input array
Space Complexity: O(k) where k is the number of unique items

## Error Handling
- Empty input returns empty string
- Case-sensitive comparison
- Handles multiple items with same frequency (returns first occurrence)
- No panic conditions

### Commit Message Format
- feat: New feature
- fix: Bug fix
- docs: Documentation changes
- test: Adding or modifying tests
- refactor: Code refactoring

## Contributing
1. Fork the repository
2. Create your feature branch
3. Commit your changes
4. Push to the branch
5. Create a Pull Request

## Future Improvements
- Support for different data types (generic implementation)
- Command-line flags for case-sensitivity option
- File input/output support
- Concurrent processing for large datasets
- API endpoint option

## Test and Results
C:\Users\sahnf\most_frequent_finder>go test ./...
?       most_frequent_finder/cmd        [no test files]
ok      most_frequent_finder/internal/finder    0.157s

C:\Users\sahnf\most_frequent_finder>
C:\Users\sahnf\most_frequent_finder>go run cmd/main.go

Test Case 1:
Input: [apple, pie, apple, red, red, red]
Most frequent item: red

Test Case 2:
Input: [dog, cat, dog, bird, cat, dog]
Most frequent item: dog

Test Case 3:
Input: [a, b, c, a]
Most frequent item: a

Test Case 4:
Input: []
Most frequent item: