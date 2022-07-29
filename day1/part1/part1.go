package part1

import "strconv"

func Part1(lines []string) (result int) {
	var entries []int
	for _, line := range lines {
		entry, _ := strconv.Atoi(line)
		entries = append(entries, entry)
	}

	for index, entry1 := range entries {
		for _, entry2 := range entries[index:] {
			if entry1+entry2 == 2020 {
				result = entry1 * entry2
			}
		}
	}

	return result
}
