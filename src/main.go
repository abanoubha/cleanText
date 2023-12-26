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
		case "--similar":
			fmt.Println(convertSimilarChars(inp))
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

func convertSimilarChars(input string) string {
	replacements := map[rune]string{
		'À': "A",
		'Á': "A",
		'Â': "A",
		'Ä': "A",
		'Ã': "A",
		'Å': "A",
		'Ā': "A",
		'Ą': "A",
		'ƛ': "A",
		'Ǎ': "A",
		'Ǟ': "A",
		'Ǡ': "A",
		'Ǻ': "A",
		'Ȁ': "A",
		'Ȃ': "A",
		'Ȧ': "A",
		'Ⱥ': "A",
		'Ʌ': "A",

		'ƙ': "k",

		'Ɖ': "D",
	}

	// var output []rune
	output := make([]rune, len(input)) // use less space

	for _, r := range input {
		if replacement, ok := replacements[r]; ok {
			output = append(output, []rune(replacement)...)
		} else {
			output = append(output, r)
		}
	}

	return string(output)
}
