package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		inp := scanner.Text()
		result := make([]byte, len(inp))
		for i, c := range inp {
			switch c {
			case 'À', 'Á', 'Â', 'Ä', 'Ã', 'Å', 'Ā', 'Ą', 'ƛ', 'Ǎ', 'Ǟ', 'Ǡ':
				result[i] = 'A'
			case 'Æ':
				result[i] = 'A' // make it "Ae" later
			case 'à', 'á', 'â', 'ä', 'ã', 'å', 'ā', 'ª', 'ą', 'Ə', 'ǝ', 'ǎ', 'ǟ', 'ǡ':
				result[i] = 'a'
			case 'æ':
				result[i] = 'a' // make it "ae" later
			case 'ß', 'Ɓ':
				result[i] = 'B'
			case 'Ƃ', 'ƃ', 'Ƅ', 'ƅ':
				result[i] = 'b'
			case 'Ç', 'Ć', 'Č', 'Ƈ':
				result[i] = 'C'
			case 'ç', 'ć', 'č', 'Ɔ', 'ƈ':
				result[i] = 'c'
			case 'Ð', 'Ď', 'Ɖ', 'Ɗ':
				result[i] = 'D'
			case 'ð', 'ď', 'đ', 'Ƌ', 'ƌ':
				result[i] = 'd'
			case 'È', 'É', 'Ê', 'Ë', 'Ē', 'Ė', 'Ę', 'Ǝ', 'Ɛ', 'Ʃ', 'Ƹ', 'ƹ':
				result[i] = 'E'
			case 'è', 'é', 'ê', 'ë', 'ē', 'ė', 'ę':
				result[i] = 'e'
			case 'ƒ', 'Ƒ':
				result[i] = 'f'
			case 'Ĝ', 'Ğ', 'Ġ', 'Ģ', 'Ɠ', 'Ǥ', 'Ǧ':
				result[i] = 'G'
			case 'ĝ', 'ğ', 'ġ', 'ģ', 'ƍ', 'ǥ', 'ǧ':
				result[i] = 'g'
			case 'Ĥ', 'Ħ':
				result[i] = 'H'
			case 'ĥ', 'ħ':
				result[i] = 'h'
			case 'Î', 'Ï', 'Í', 'Ī', 'Į', 'Ì', 'Ɨ', 'Ǐ':
				result[i] = 'I'
			case 'î', 'ï', 'í', 'ī', 'į', 'ì', '¡', 'ı', 'ĺ', 'ļ', 'Ɩ', 'ǐ':
				result[i] = 'i'
			case 'ǰ':
				result[i] = 'j'
			case 'Ƙ', 'Ǩ':
				result[i] = 'K'
			case 'ƙ', 'ǩ':
				result[i] = 'k'
			case 'Ł', 'Ĺ', 'Ļ', 'Ľ', 'Ŀ':
				result[i] = 'L'
			case 'ł', 'ľ', 'ŀ', 'ƚ':
				result[i] = 'l'
			case 'Ñ', 'Ń', 'Ņ', 'Ň', 'Ŋ', 'Ɲ':
				result[i] = 'N'
			case 'ñ', 'ń', 'ņ', 'ň', 'ŉ', 'ŋ', 'ƞ':
				result[i] = 'n'
			case 'Ɵ', 'Ơ', 'Ǒ', 'Ǫ', 'Ǭ':
				result[i] = 'O'
			case 'ô', 'ö', 'ò', 'ó', 'ø', 'ō', 'õ', 'º', 'ő', 'ơ', 'ǒ', 'ǫ', 'ǭ':
				result[i] = 'o'
			case 'ƿ':
				result[i] = 'P'
			case 'Þ', 'þ', 'Ƥ', 'ƥ':
				result[i] = 'p'
			case 'ƪ':
				result[i] = 'q'
			case 'Ŕ', 'Ŗ', 'Ř', 'Ʀ':
				result[i] = 'R'
			case 'ŕ', 'ŗ', 'ř':
				result[i] = 'r'
			case 'Ś', 'Š', '§', 'Ŝ', 'Ş':
				result[i] = 'S'
			case 'ś', 'š', 'ŝ', 'ş':
				result[i] = 's'
			case 'Ţ', 'Ť', 'Ŧ', 'Ƭ', 'Ʈ':
				result[i] = 'T'
			case 'ţ', 'ť', 'ŧ', 'ƫ', 'ƭ':
				result[i] = 't'
			case 'Ũ', 'Ū', 'Ŭ', 'Ů', 'Ű', 'Ų', 'Ư', 'Ǔ', 'Ǖ', 'Ǘ', 'Ǚ', 'Ǜ':
				result[i] = 'U'
			case 'û', 'ü', 'ù', 'ú', 'ū', 'µ', 'ũ', 'ŭ', 'ů', 'ű', 'ų', 'ư', 'Ʊ', 'ǔ', 'ǖ', 'ǘ', 'ǚ', 'ǜ':
				result[i] = 'u'
			case 'Ɣ', 'Ʋ':
				result[i] = 'V'
			case 'Ŵ', 'Ɯ':
				result[i] = 'W'
			case 'ŵ':
				result[i] = 'w'
			case 'Ÿ', 'Ŷ', 'Ƴ':
				result[i] = 'Y'
			case 'ÿ', 'ŷ', 'ƴ':
				result[i] = 'y'
			case 'Ž', 'Ź', 'Ż', 'Ƶ':
				result[i] = 'Z'
			case 'ž', 'ź', 'ż', 'ƶ':
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
			case '†', '‡', 'ˆ', '‹', '«', '›', '»', '•', '™', '¦', '¨', '©', '®', '¯', '¬', '´', '¶', '¸', '¿', 'ſ', 'ǀ', 'ǁ', 'ǂ', 'ǃ':
				// skip these symbols
			case '¼', '½', '¾', '÷':
				// skipped math
			case '×': // multiplication sign
				result[i] = 'x'
			case '¹':
				result[i] = '1'
			case '²', 'Ƨ', 'ƨ', 'ƻ':
				result[i] = '2'
			case '³', 'Ʒ', 'ƺ', 'Ǯ', 'ǯ':
				result[i] = '3'
			case 'Ƽ', 'ƽ', 'ƾ':
				result[i] = '5'
			case '·':
				result[i] = '.'
			case 'Ĳ', 'ĳ', 'Œ', 'œ', 'ƕ', 'Ƣ', 'ƣ', 'Ǆ', 'ǅ', 'ǆ', 'Ǉ', 'ǈ', 'ǉ', 'Ǌ', 'ǋ', 'ǌ', 'Ǣ', 'ǣ', 'Ǳ', 'ǲ', 'ǳ':
				// handle ligatures later
			default:
				if c <= math.MaxUint8 {
					result[i] = byte(c)
				} else {
					// Handle the case where c is too large to fit in a byte.
				}
			}
		}
		fmt.Println(string(result))
	}
} // main
