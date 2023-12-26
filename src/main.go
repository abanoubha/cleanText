package main

import (
	"bufio"
	"fmt"
	"math"
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
			// alt func
			fmt.Println(convertSimilarCharsAlt(inp))
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

func convertSimilarCharsAlt(input string) string {
	result := make([]byte, len(input))
	for i, c := range input {
		switch c {
		case 'À', 'Á', 'Â', 'Ä', 'Ã', 'Å', 'Ā', 'Ą', 'ƛ', 'Ǎ', 'Ǟ', 'Ǡ', 'Ǻ', 'Ȁ', 'Ȃ', 'Ȧ', 'Ⱥ', 'Ʌ':
			result[i] = 'A'
		case 'à', 'á', 'â', 'ä', 'ã', 'å', 'ā', 'ª', 'ą', 'Ə', 'ǝ', 'ǎ', 'ǟ', 'ǡ', 'ǻ', 'ȁ', 'ȃ', 'ȧ', 'ɐ', 'ɑ', 'ɒ':
			result[i] = 'a'
		case 'ß', 'Ɓ', 'Ƀ', 'ɞ', 'ʙ':
			result[i] = 'B'
		case 'Ƃ', 'ƃ', 'Ƅ', 'ƅ', 'ɓ':
			result[i] = 'b'
		case 'Ç', 'Ć', 'Č', 'Ƈ', 'Ȼ', 'ʗ':
			result[i] = 'C'
		case 'ç', 'ć', 'č', 'Ɔ', 'ƈ', 'ȼ', 'ɔ', 'ɕ':
			result[i] = 'c'
		case 'Ð', 'Ď', 'Ɖ', 'Ɗ':
			result[i] = 'D'
		case 'ð', 'ď', 'đ', 'Ƌ', 'ƌ', 'ȡ', 'ɖ', 'ɗ':
			result[i] = 'd'
		case 'È', 'É', 'Ê', 'Ë', 'Ē', 'Ė', 'Ę', 'Ǝ', 'Ɛ', 'Ʃ', 'Ƹ', 'ƹ', 'Ȅ', 'Ȇ', 'Ȩ', 'Ɇ', 'ɛ':
			result[i] = 'E'
		case 'è', 'é', 'ê', 'ë', 'ē', 'ė', 'ę', 'ȅ', 'ȇ', 'ȩ', 'ɇ', 'ɘ', 'ə', 'ɚ':
			result[i] = 'e'
		case 'ƒ', 'Ƒ', 'ʃ', 'ʄ':
			result[i] = 'f'
		case 'Ĝ', 'Ğ', 'Ġ', 'Ģ', 'Ɠ', 'Ǥ', 'Ǧ', 'Ǵ', 'ɢ', 'ʛ':
			result[i] = 'G'
		case 'ĝ', 'ğ', 'ġ', 'ģ', 'ƍ', 'ǥ', 'ǧ', 'ǵ', 'ɠ', 'ɡ':
			result[i] = 'g'
		case 'Ĥ', 'Ħ', 'Ƕ', 'Ȟ', 'ʜ':
			result[i] = 'H'
		case 'ĥ', 'ħ', 'ȟ', 'ɦ', 'ɧ', 'ʮ', 'ʯ', 'ʰ', 'ʱ':
			result[i] = 'h'
		case 'Î', 'Ï', 'Í', 'Ī', 'Į', 'Ì', 'Ɨ', 'Ǐ', 'Ȉ', 'Ȋ', 'ɪ':
			result[i] = 'I'
		case 'î', 'ï', 'í', 'ī', 'į', 'ì', '¡', 'ı', 'ĺ', 'ļ', 'Ɩ', 'ǐ', 'ȉ', 'ȋ', 'ɨ', 'ɩ':
			result[i] = 'i'
		case 'Ɉ':
			result[i] = 'J'
		case 'ǰ', 'ȷ', 'ɉ', 'ɟ', 'ʝ', 'ʲ':
			result[i] = 'j'
		case 'Ƙ', 'Ǩ':
			result[i] = 'K'
		case 'ƙ', 'ǩ', 'ʞ':
			result[i] = 'k'
		case 'Ł', 'Ĺ', 'Ļ', 'Ľ', 'Ŀ', 'Ƚ', 'ʟ':
			result[i] = 'L'
		case 'ł', 'ľ', 'ŀ', 'ƚ', 'ȴ', 'ɬ', 'ɭ', 'ʅ', 'ɫ':
			result[i] = 'l'
		case 'ʍ':
			result[i] = 'M'
		case 'ɱ':
			result[i] = 'm'
		case 'Ñ', 'Ń', 'Ņ', 'Ň', 'Ŋ', 'Ɲ', 'Ǹ', 'ɴ':
			result[i] = 'N'
		case 'ñ', 'ń', 'ņ', 'ň', 'ŉ', 'ŋ', 'ƞ', 'ǹ', 'Ƞ', 'ȵ', 'ɲ', 'ɳ':
			result[i] = 'n'
		case 'Ɵ', 'Ơ', 'Ǒ', 'Ǫ', 'Ǭ', 'Ǿ', 'Ȍ', 'Ȏ', 'Ȫ', 'Ȭ', 'Ȯ', 'Ȱ', 'ɸ', 'ʘ':
			result[i] = 'O'
		case 'ô', 'ö', 'ò', 'ó', 'ø', 'ō', 'õ', 'º', 'ő', 'ơ', 'ǒ', 'ǫ', 'ǭ', 'ǿ', 'ȍ', 'ȏ', 'ȫ', 'ȭ', 'ȯ', 'ȱ', 'ɵ':
			result[i] = 'o'
		case 'ƿ':
			result[i] = 'P'
		case 'Þ', 'þ', 'Ƥ', 'ƥ':
			result[i] = 'p'
		case 'Ɋ':
			result[i] = 'Q'
		case 'ƪ', 'ɋ', 'ʠ':
			result[i] = 'q'
		case 'Ŕ', 'Ŗ', 'Ř', 'Ʀ', 'Ȑ', 'Ȓ', 'Ɍ', 'ʀ', 'ʁ', 'ʶ':
			result[i] = 'R'
		case 'ŕ', 'ŗ', 'ř', 'ȑ', 'ȓ', 'ɍ', 'ɹ', 'ɺ', 'ɻ', 'ɼ', 'ɽ', 'ɾ', 'ɿ', 'ʳ', 'ʴ', 'ʵ':
			result[i] = 'r'
		case 'Ś', 'Š', '§', 'Ŝ', 'Ş', 'Ș', 'ȿ', 'ʂ':
			result[i] = 'S'
		case 'ś', 'š', 'ŝ', 'ş', 'ș':
			result[i] = 's'
		case 'Ţ', 'Ť', 'Ŧ', 'Ƭ', 'Ʈ', 'Ț', 'Ⱦ':
			result[i] = 'T'
		case 'ţ', 'ť', 'ŧ', 'ƫ', 'ƭ', 'ț', 'ȶ', 'ʇ', 'ʈ':
			result[i] = 't'
		case 'Ũ', 'Ū', 'Ŭ', 'Ů', 'Ű', 'Ų', 'Ư', 'Ǔ', 'Ǖ', 'Ǘ', 'Ǚ', 'Ǜ', 'Ȕ', 'Ȗ', 'Ʉ':
			result[i] = 'U'
		case 'û', 'ü', 'ù', 'ú', 'ū', 'µ', 'ũ', 'ŭ', 'ů', 'ű', 'ų', 'ư', 'Ʊ', 'ǔ', 'ǖ', 'ǘ', 'ǚ', 'ǜ', 'ȕ', 'ȗ', 'ɥ', 'ʉ', 'ʊ':
			result[i] = 'u'
		case 'Ɣ', 'Ʋ', 'ʋ', 'ʌ', 'ˬ', '˯':
			result[i] = 'v'
		case 'Ŵ', 'Ɯ':
			result[i] = 'W'
		case 'ŵ', 'ɯ', 'ɰ':
			result[i] = 'w'
		case 'Ÿ', 'Ŷ', 'Ƴ', 'Ȳ', 'Ɏ', 'ʏ':
			result[i] = 'Y'
		case 'ÿ', 'ŷ', 'ƴ', 'ȳ', 'ɏ', 'ʸ', 'ʎ':
			result[i] = 'y'
		case 'Ž', 'Ź', 'Ż', 'Ƶ', 'Ȥ', 'ɀ', 'ʐ', 'ʑ':
			result[i] = 'Z'
		case 'ž', 'ź', 'ż', 'ƶ', 'ȥ':
			result[i] = 'z'
		case '¢':
			result[i] = 'c' // or '€' : will add a flag '--keep-currency-symbols'
		case '€':
			result[i] = 'C' // or '€' : will add a flag '--keep-currency-symbols'
		case '£':
			result[i] = 'L' // will add a flag '--keep-currency-symbols'
		case '¤':
			result[i] = 'o' // will add a flag '--keep-currency-symbols'
		case '¥':
			result[i] = 'Y' // will add a flag '--keep-currency-symbols'
		case 'Ƨ', 'ƨ', 'ƻ':
			result[i] = '2'
		case 'Ʒ', 'ƺ', 'Ǯ', 'ǯ', 'Ȝ', 'ȝ', 'ɜ', 'ɝ', 'ʒ', 'ʓ':
			result[i] = '3'
		case 'Ƽ', 'ƽ', 'ƾ':
			result[i] = '5'
		case '·':
			result[i] = '.'
		case 'Ɂ', 'ɂ', 'ʔ', 'ʡ':
			result[i] = '?'
		case 'ʹ', 'ʼ':
			result[i] = '\''
		case 'ʺ', '˝', 'ˮ', '˵', '˶':
			result[i] = '"'
		case '˟':
			result[i] = '×'
		case '˸':
			result[i] = ':'
		default:
			if c <= math.MaxUint8 {
				result[i] = byte(c)
			} else {
				// Handle the case where c is too large to fit in a byte.
				// use a space
				result[i] = byte(' ')
			}
		}
	}
	return string(result)
}
