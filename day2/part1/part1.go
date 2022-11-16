package part1

import (
	"fmt"
	"strings"
)

func Part1(lines []string) (numberValidPasswords int) {
	numberValidPasswords = 0

	for _, line := range lines {
		var minTimesLetterInPassword, maxTimesLetterInPassword int
		var letter, password string
		fmt.Sscanf(line, "%d-%d %s: %s", &minTimesLetterInPassword, &maxTimesLetterInPassword, &letter, &password)

		// fmt.Print(minTimesLetterInPassword, maxTimesLetterInPassword, letter, password)
		numberTimesLetterInPassword := strings.Count(password, letter)
		if numberTimesLetterInPassword >= minTimesLetterInPassword && numberTimesLetterInPassword >= maxTimesLetterInPassword {
			numberValidPasswords++
		}
	}

	return numberValidPasswords
}
