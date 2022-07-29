package part2

import "strconv"

func Part2(lines []string) (result int) {
	var entries []int
	for _, line := range lines {
		entry, _ := strconv.Atoi(line)
		entries = append(entries, entry)
	}

	for index1, entry1 := range entries {
		for index2, entry2 := range entries[index1:] {
			for _, entry3 := range entries[index2:] {
				if entry1+entry2+entry3 == 2020 {
					result = entry1 * entry2 * entry3
				}
			}
		}
	}

	return result
}
