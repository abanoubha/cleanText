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
			case 'ß':
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
			case 'ĝ', 'ğ', 'ġ':
				result[i] = 'g'
			case 'Î', 'Ï', 'Í', 'Ī', 'Į', 'Ì':
				result[i] = 'I'
			case 'î', 'ï', 'í', 'ī', 'į', 'ì', '¡':
				result[i] = 'i'
			case 'Ł':
				result[i] = 'L'
			case 'ł':
				result[i] = 'l'
			case 'Ñ', 'Ń':
				result[i] = 'N'
			case 'ñ', 'ń':
				result[i] = 'n'
			case 'ô', 'ö', 'ò', 'ó', 'œ', 'ø', 'ō', 'õ', 'º':
				result[i] = 'o'
			case 'Þ', 'þ':
				result[i] = 'p'
			case 'Ś', 'Š', '§':
				result[i] = 'S'
			case 'ś', 'š':
				result[i] = 's'
			case 'û', 'ü', 'ù', 'ú', 'ū', 'µ':
				result[i] = 'u'
			case 'Ÿ':
				result[i] = 'Y'
			case 'ÿ':
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
			case '†', '‡', 'ˆ', '‹', '«', '›', '»', '•', '™', '¦', '¨', '©', '®', '¯', '¬', '´', '¶', '¸', '¿':
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
