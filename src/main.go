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
			case 'À', 'Á', 'Â', 'Ä', 'Æ', 'Ã', 'Å', 'Ā':
				result[i] = 'A'
			case 'à', 'á', 'â', 'ä', 'æ', 'ã', 'å', 'ā':
				result[i] = 'a'
			case 'ß':
				result[i] = 'B'
			case 'Ç', 'Ć', 'Č':
				result[i] = 'C'
			case 'ç', 'ć', 'č':
				result[i] = 'c'
			case 'È', 'É', 'Ê', 'Ë', 'Ē', 'Ė', 'Ę':
				result[i] = 'E'
			case 'è', 'é', 'ê', 'ë', 'ē', 'ė', 'ę':
				result[i] = 'e'
			case 'ƒ':
				result[i] = 'f'
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
			case 'ô', 'ö', 'ò', 'ó', 'œ', 'ø', 'ō', 'õ':
				result[i] = 'o'
			case 'Ś', 'Š':
				result[i] = 'S'
			case 'ś', 'š':
				result[i] = 's'
			case 'û', 'ü', 'ù', 'ú', 'ū':
				result[i] = 'u'
			case 'Ÿ':
				result[i] = 'Y'
			case 'ÿ':
				result[i] = 'y'
			case 'Ž', 'Ź', 'Ż':
				result[i] = 'Z'
			case 'ž', 'ź', 'ż':
				result[i] = 'z'
			case '€', '¢':
				result[i] = 'C' // or '€' : will add a flag '--keep-currency-symbols'
			case '†', '‡', 'ˆ', '‹', '›', '•', '™':
				// skip these symbols
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
