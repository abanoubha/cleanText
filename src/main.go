package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

func main() {
	if contains(os.Args, "--help") || contains(os.Args, "-h") {
		printHelpScreen()
		os.Exit(0)
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		data := scanner.Text()

		if len(os.Args) < 2 {
			// default befavior
			data = normalizeSpecialChars(data)
			data = normalizeLigatures(data)
			data = convertSimilarChars(data)
		}

		if contains(os.Args, "--special-chars") || contains(os.Args, "--special") {
			data = normalizeSpecialChars(data)
		}

		if contains(os.Args, "--similar") {
			data = convertSimilarChars(data)
			// alt func
			data = convertSimilarCharsAlt(data)
		}

		if contains(os.Args, "--liga") || contains(os.Args, "--ligatures") {
			data = normalizeLigatures(data)
			// alt func
			data = normalizeLigaturesAlt(data)
		}

		fmt.Println(data)
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

		'à': "a",
		'á': "a",
		'â': "a",
		'ä': "a",
		'ã': "a",
		'å': "a",
		'ā': "a",
		'ª': "a",
		'ą': "a",
		'Ə': "a",
		'ǝ': "a",
		'ǎ': "a",
		'ǟ': "a",
		'ǡ': "a",
		'ǻ': "a",
		'ȁ': "a",
		'ȃ': "a",
		'ȧ': "a",
		'ɐ': "a",
		'ɑ': "a",
		'ɒ': "a",

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

func normalizeLigatures(input string) string {
	replacements := map[rune]string{
		'Æ': "AE",
		'æ': "ae",
		'Ĳ': "IJ",
		'ĳ': "ij",
		'Œ': "OE",
		'œ': "oe",
		'ƕ': "hv",
		'Ƣ': "OI",
		'ƣ': "oi",
		'Ǆ': "DZ",
		'ǅ': "Dz",
		'ǆ': "dz",
		'Ǉ': "LJ",
		'ǈ': "Lj",
		'ǉ': "lj",
		'Ǌ': "NJ",
		'ǋ': "Nj",
		'ǌ': "nj",
		'Ǣ': "AE",
		'ǣ': "ae",
		'Ǳ': "DZ",
		'ǲ': "Dz",
		'ǳ': "dz",
		'Ǽ': "AE",
		'ǽ': "ae",
		'Ȣ': "OU",
		'ȣ': "ou",
		'ȸ': "db",
		'ȹ': "qp",
		'ɮ': "lj",
		'ɶ': "oe",
		'ɷ': "w",
		'ʣ': "dz",
		'ʤ': "dj",
		'ʥ': "dz",
		'ʦ': "ts",
		'ʧ': "tf",
		'ʨ': "tc",
		'ʩ': "Feng",
		'ʪ': "ls",
		'ʫ': "lz",
		'ʬ': "w",
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

func normalizeLigaturesAlt(input string) string {
	output := make([]string, len(input))
	for i, c := range input {
		switch c {
		case 'Æ':
			output[i] = "AE"
		case 'æ':
			output[i] = "ae"
		case 'Ĳ':
			output[i] = "IJ"
		case 'ĳ':
			output[i] = "ij"
		case 'Œ':
			output[i] = "OE"
		case 'œ':
			output[i] = "oe"
		case 'ƕ':
			output[i] = "hv"
		case 'Ƣ':
			output[i] = "OI"
		case 'ƣ':
			output[i] = "oi"
		case 'Ǆ':
			output[i] = "DZ"
		case 'ǅ':
			output[i] = "Dz"
		case 'ǆ':
			output[i] = "dz"
		case 'Ǉ':
			output[i] = "LJ"
		case 'ǈ':
			output[i] = "Lj"
		case 'ǉ':
			output[i] = "lj"
		case 'Ǌ':
			output[i] = "NJ"
		case 'ǋ':
			output[i] = "Nj"
		case 'ǌ':
			output[i] = "nj"
		case 'Ǣ':
			output[i] = "AE"
		case 'ǣ':
			output[i] = "ae"
		case 'Ǳ':
			output[i] = "DZ"
		case 'ǲ':
			output[i] = "Dz"
		case 'ǳ':
			output[i] = "dz"
		case 'Ǽ':
			output[i] = "AE"
		case 'ǽ':
			output[i] = "ae"
		case 'Ȣ':
			output[i] = "OU"
		case 'ȣ':
			output[i] = "ou"
		case 'ȸ':
			output[i] = "db"
		case 'ȹ':
			output[i] = "qp"
		case 'ɮ':
			output[i] = "lj"
		case 'ɶ':
			output[i] = "oe"
		case 'ɷ':
			output[i] = "w"
		case 'ʣ':
			output[i] = "dz"
		case 'ʤ':
			output[i] = "dj"
		case 'ʥ':
			output[i] = "dz"
		case 'ʦ':
			output[i] = "ts"
		case 'ʧ':
			output[i] = "tf"
		case 'ʨ':
			output[i] = "tc"
		case 'ʩ':
			output[i] = "feng"
		case 'ʪ':
			output[i] = "ls"
		case 'ʫ':
			output[i] = "lz"
		case 'ʬ':
			output[i] = "w"
		default:
			output[i] = string(c)
		}
	}
	return strings.Join(output, "")
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}
