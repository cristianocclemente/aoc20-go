package part1_test

import (
	"aoc20-go/day2/part1"
	"fmt"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) []string {
	fileBytes, _ := ioutil.ReadFile(path)
	sliceData := strings.Split(string(fileBytes), "\n")
	return sliceData
}

func TestExampleInput(t *testing.T) {
	lines := readLines("../example_input.txt")
	fmt.Print("lines:", lines)
	numberValidPasswords := part1.Part1(lines)
	assert.Equal(t, 2, numberValidPasswords)
}

func TestMyInput(t *testing.T) {
	lines := readLines("../my_input.txt")
	fmt.Print("lines:", lines)
	result := part1.Part1(lines)
	assert.Equal(t, 805731, result)
}
