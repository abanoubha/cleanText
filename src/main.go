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
			case 'À', 'Á', 'Â', 'Ä', 'Ã', 'Å', 'Ā', 'Ą', 'ƛ':
				result[i] = 'A'
			case 'Æ':
				result[i] = 'A' // make it "Ae" later
			case 'à', 'á', 'â', 'ä', 'ã', 'å', 'ā', 'ª', 'ą', 'Ə':
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
			case 'È', 'É', 'Ê', 'Ë', 'Ē', 'Ė', 'Ę', 'Ǝ', 'Ɛ':
				result[i] = 'E'
			case 'è', 'é', 'ê', 'ë', 'ē', 'ė', 'ę':
				result[i] = 'e'
			case 'ƒ', 'Ƒ':
				result[i] = 'f'
			case 'Ĝ', 'Ğ', 'Ġ', 'Ģ', 'Ɠ':
				result[i] = 'G'
			case 'ĝ', 'ğ', 'ġ', 'ģ', 'ƍ':
				result[i] = 'g'
			case 'Ĥ', 'Ħ':
				result[i] = 'H'
			case 'ĥ', 'ħ':
				result[i] = 'h'
			case 'Î', 'Ï', 'Í', 'Ī', 'Į', 'Ì', 'Ɨ':
				result[i] = 'I'
			case 'î', 'ï', 'í', 'ī', 'į', 'ì', '¡', 'ı', 'ĺ', 'ļ', 'Ɩ':
				result[i] = 'i'
			case 'Ƙ':
				result[i] = 'K'
			case 'ƙ':
				result[i] = 'k'
			case 'Ł', 'Ĺ', 'Ļ', 'Ľ', 'Ŀ':
				result[i] = 'L'
			case 'ł', 'ľ', 'ŀ', 'ƚ':
				result[i] = 'l'
			case 'Ñ', 'Ń', 'Ņ', 'Ň', 'Ŋ', 'Ɲ':
				result[i] = 'N'
			case 'ñ', 'ń', 'ņ', 'ň', 'ŉ', 'ŋ', 'ƞ':
				result[i] = 'n'
			case 'Ɵ', 'Ơ':
				result[i] = 'O'
			case 'ô', 'ö', 'ò', 'ó', 'ø', 'ō', 'õ', 'º', 'ő', 'ơ':
				result[i] = 'o'
			case 'Þ', 'þ', 'Ƥ', 'ƥ':
				result[i] = 'p'
			case 'Ŕ', 'Ŗ', 'Ř', 'Ʀ':
				result[i] = 'R'
			case 'ŕ', 'ŗ', 'ř':
				result[i] = 'r'
			case 'Ś', 'Š', '§', 'Ŝ', 'Ş':
				result[i] = 'S'
			case 'ś', 'š', 'ŝ', 'ş':
				result[i] = 's'
			case 'Ţ', 'Ť', 'Ŧ':
				result[i] = 'T'
			case 'ţ', 'ť', 'ŧ':
				result[i] = 't'
			case 'Ũ', 'Ū', 'Ŭ', 'Ů', 'Ű', 'Ų':
				result[i] = 'U'
			case 'û', 'ü', 'ù', 'ú', 'ū', 'µ', 'ũ', 'ŭ', 'ů', 'ű', 'ų':
				result[i] = 'u'
			case 'Ɣ':
				result[i] = 'V'
			case 'Ŵ', 'Ɯ':
				result[i] = 'W'
			case 'ŵ':
				result[i] = 'w'
			case 'Ÿ', 'Ŷ':
				result[i] = 'Y'
			case 'ÿ', 'ŷ':
				result[i] = 'y'
			case 'Ž', 'Ź', 'Ż':
				result[i] = 'Z'
			case 'ž', 'ź', 'ż':
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
			case '†', '‡', 'ˆ', '‹', '«', '›', '»', '•', '™', '¦', '¨', '©', '®', '¯', '¬', '´', '¶', '¸', '¿', 'ſ':
				// skip these symbols
			case '¼', '½', '¾', '÷':
				// skipped math
			case '×': // multiplication sign
				result[i] = 'x'
			case '¹':
				result[i] = '1'
			case '²':
				result[i] = '2'
			case '³':
				result[i] = '3'
			case '·':
				result[i] = '.'
			case 'Ĳ', 'ĳ', 'Œ', 'œ', 'ƕ', 'Ƣ', 'ƣ':
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
