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
			case 'À', 'Á', 'Â', 'Ä', 'Ã', 'Å', 'Ā', 'Ą':
				result[i] = 'A'
			case 'Æ':
				result[i] = 'A' // make it "Ae" later
			case 'à', 'á', 'â', 'ä', 'ã', 'å', 'ā', 'ª', 'ą':
				result[i] = 'a'
			case 'æ':
				result[i] = 'a' // make it "ae" later
			case 'ß', 'Ɓ':
				result[i] = 'B'
			case 'Ç', 'Ć', 'Č':
				result[i] = 'C'
			case 'ç', 'ć', 'č':
				result[i] = 'c'
			case 'Ð', 'Ď':
				result[i] = 'D'
			case 'ð', 'ď', 'đ':
				result[i] = 'd'
			case 'È', 'É', 'Ê', 'Ë', 'Ē', 'Ė', 'Ę':
				result[i] = 'E'
			case 'è', 'é', 'ê', 'ë', 'ē', 'ė', 'ę':
				result[i] = 'e'
			case 'ƒ':
				result[i] = 'f'
			case 'Ĝ', 'Ğ', 'Ġ', 'Ģ':
				result[i] = 'G'
			case 'ĝ', 'ğ', 'ġ', 'ģ':
				result[i] = 'g'
			case 'Ĥ', 'Ħ':
				result[i] = 'H'
			case 'ĥ', 'ħ':
				result[i] = 'h'
			case 'Î', 'Ï', 'Í', 'Ī', 'Į', 'Ì':
				result[i] = 'I'
			case 'î', 'ï', 'í', 'ī', 'į', 'ì', '¡', 'ı', 'ĺ', 'ļ':
				result[i] = 'i'
			case 'Ł', 'Ĺ', 'Ļ', 'Ľ', 'Ŀ':
				result[i] = 'L'
			case 'ł', 'ľ', 'ŀ':
				result[i] = 'l'
			case 'Ñ', 'Ń', 'Ņ', 'Ň', 'Ŋ':
				result[i] = 'N'
			case 'ñ', 'ń', 'ņ', 'ň', 'ŉ', 'ŋ':
				result[i] = 'n'
			case 'ô', 'ö', 'ò', 'ó', 'ø', 'ō', 'õ', 'º', 'ő':
				result[i] = 'o'
			case 'Þ', 'þ':
				result[i] = 'p'
			case 'Ŕ', 'Ŗ', 'Ř':
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
			case 'Ŵ':
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
			case 'Ĳ', 'ĳ', 'Œ', 'œ':
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
