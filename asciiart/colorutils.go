package asciiart

import (
	"fmt"
	"strings"
)

var (
	Cyan       = "\033[36m"
	Reset      = "\033[0m"
	colorCodes = map[string]string{
		"black":   "\033[30m",
		"red":     "\033[31m",
		"green":   "\033[32m",
		"yellow":  "\033[33m",
		"blue":    "\033[34m",
		"magenta": "\033[35m",
		"cyan":    "\033[36m",
		"white":   "\033[37m",
		"orange":  "\033[38;5;208m",
		"reset":   "\033[0m",
	}
)

// printAsciiWithColor prints ASCII art with a specific color.
func printAsciiWithColor(str string, alphabet map[rune][]string, color string, target string) {
	words, index := colorWord(str, target)

	lineOutput := ""
	for i := 0; i < CHARACTER_HEIGHT; i++ {
		for idx, word := range words {
			for _, l := range word {
				switch l {
				case '\n':
					fmt.Println()
				default:
					line := alphabet[rune(l)][i]
					// whole word
					if shouldColor(rune(l), target) && idx == index {
						lineOutput += color + line + Reset
					} else if shouldColor(rune(l), target) && index == -1 {
						// individual letters
						lineOutput += color + line + Reset
					} else {
						lineOutput += line
					}
				}
			}
			// add a line only if it is the final word
			if idx == len(words)-1 {
				fmt.Println(lineOutput)
				lineOutput = ""
			} else {
				lineOutput += alphabet[rune(' ')][i]
			}
		}
	}
}

// shouldColor determines if a letter should be colored
func shouldColor(letter rune, target string) bool {
	if target == "" {
		return true
	}
	return strings.ContainsRune(target, letter)
}

// checks whether we need to color the whole word
func colorWord(s string, target string) ([]string, int) {
	words := strings.Split(s, " ")
	for index, word := range words {
		if word == target {
			return words, index
		}
	}
	return words, -1
}

// printAsciiWithColor prints ASCII art with a specific color.
// func printAsciiWithColor(word string, alphabet map[rune][]string, color string, target string) {

// 	for i := 0; i < CHARACTER_HEIGHT; i++ {
// 		lineOutput := ""
// 		for _, l := range word {
// 			switch l {
// 			case '\n':
// 				fmt.Println()
// 			default:
// 				line := alphabet[rune(l)][i]
// 				if shouldColor(rune(l), target) {
// 					lineOutput += color + line + Reset
// 				} else {
// 					lineOutput += line
// 				}
// 			}
// 		}
// 		fmt.Println(lineOutput)
// 	}
// }
