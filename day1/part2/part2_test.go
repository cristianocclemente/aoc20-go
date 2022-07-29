package part2_test

import (
	"aoc20-go/day1/part2"
	"bufio"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func TestExampleInput(t *testing.T) {
	lines, _ := readLines("../example_input.txt")
	result := part2.Part2(lines)
	assert.Equal(t, 241861950, result)
}

func TestMyInput(t *testing.T) {
	lines, _ := readLines("../my_input.txt")
	result := part2.Part2(lines)
	assert.Equal(t, 192684960, result)
}
