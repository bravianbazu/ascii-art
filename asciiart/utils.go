package asciiart

import (
	"fmt"
	"strings"
)

// isValid checks if the input string contains only acceptable characters. ([32,126] & [8,10])
// The special characters newline (10), carriage return (13), and backspace (8).
// - bool: true if contains only acceptable, false otherwise.
func isValidInput(args string) (bool, rune) {
	for _, ch := range args {
		if !((ch >= 32 && ch <= 126) || (ch >= 8 && ch <= 10)) {
			return false, ch
		}
	}
	return true, ' '
}

/*
* Replace all special characters with characters that can be recognized and processed by golang
* Make them fit inside a rune
 */
func replaceSpecialcharacters(s string) string {
	replacer := strings.NewReplacer(
		"\\v", "\\n\\n\\n\\n",
		"\n", "\\n",
		"\\t", "    ",
		"\\b", "\b",
		"\\r", "\r",
		"\\a", "\a",
		"\\f", "\f",
	)
	return replacer.Replace(s)
}

// printAscii prints the ascii art representation of the given word using the given alphabet
func printAscii(word string, alphabet map[rune][]string) {
	for i := 0; i < CHARACTER_HEIGHT; i++ {
		lineOutput := ""
		for _, l := range word {
			switch l {
			case '\n':
				fmt.Println()
			default:
				lineOutput += alphabet[rune(l)][i]
			}
		}
		fmt.Println(lineOutput)
	}
}

/*
* Processes the input array and extracts whatever the program needs
* It is also responsible of dealing with optional parameters like the file names
* Exits if input does not meet the expected format
 */
func processInput(arr []string) (bool, string) {
	if len(arr) < 2 || len(arr) > 4 {
		return false, ""
	}

	filename := "standard.txt"
	if len(arr) == 3 {
		switch arr[2] {
		case "shadow", "shadow.txt":
			filename = "shadow.txt"
		case "thinkertoy", "thinkertoy.txt":
			filename = "thinkertoy.txt"
		}
	}
	return true, filename
}

func processBackspace(s string) string {
	if len(s) > 1 {
		temp := ""
		for index, ch := range s {
			if ch != '\b' {
				temp += string(ch)
			} else {
				temp = temp[:index-1]
			}
		}
		return temp
	} else {
		return ""
	}
}
