package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

func main() {
	if len(os.Args) < 2 {
		// choose a default behavior
		printHelpScreen()
		return
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		inp := scanner.Text()

		switch os.Args[1] {
		case "--special-chars", "--special":
			output := normalizeSpecialChars(inp)
			fmt.Println(output)
		case "-h", "--help":
			printHelpScreen()
		default:
			printHelpScreen()
		}

	}
} // main

func printHelpScreen() {
	fmt.Println("cleanText is a small program to normalize and clean text")
	fmt.Println("cleanText <args>")
	fmt.Println("commands:")
	fmt.Println("cleanText -h (--help) : print the help screen (this screen)")
	fmt.Println("cleanText -e (--english) : normalize English letters (e.g convert ã into a) ")
}

func normalizeSpecialChars(input string) string {
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	result, _, _ := transform.String(t, input)
	return result
}
