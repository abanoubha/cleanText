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
			case 'À', 'Á', 'Â', 'Ä', 'Ã', 'Å', 'Ā', 'Ą', 'ƛ', 'Ǎ', 'Ǟ', 'Ǡ', 'Ǻ', 'Ȁ', 'Ȃ', 'Ȧ', 'Ⱥ', 'Ʌ':
				result[i] = 'A'
			case 'Æ':
				result[i] = 'A' // make it "Ae" later
			case 'à', 'á', 'â', 'ä', 'ã', 'å', 'ā', 'ª', 'ą', 'Ə', 'ǝ', 'ǎ', 'ǟ', 'ǡ', 'ǻ', 'ȁ', 'ȃ', 'ȧ', 'ɐ', 'ɑ', 'ɒ':
				result[i] = 'a'
			case 'æ':
				result[i] = 'a' // make it "ae" later
			case 'ß', 'Ɓ', 'Ƀ', 'ɞ':
				result[i] = 'B'
			case 'Ƃ', 'ƃ', 'Ƅ', 'ƅ', 'ɓ':
				result[i] = 'b'
			case 'Ç', 'Ć', 'Č', 'Ƈ', 'Ȼ':
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
			case 'ʅ':
				// cant decide
			case 'Ĝ', 'Ğ', 'Ġ', 'Ģ', 'Ɠ', 'Ǥ', 'Ǧ', 'Ǵ', 'ɢ':
				result[i] = 'G'
			case 'ĝ', 'ğ', 'ġ', 'ģ', 'ƍ', 'ǥ', 'ǧ', 'ǵ', 'ɠ', 'ɡ':
				result[i] = 'g'
			case 'Ĥ', 'Ħ', 'Ƕ', 'Ȟ':
				result[i] = 'H'
			case 'ĥ', 'ħ', 'ȟ', 'ɦ', 'ɧ':
				result[i] = 'h'
			case 'Î', 'Ï', 'Í', 'Ī', 'Į', 'Ì', 'Ɨ', 'Ǐ', 'Ȉ', 'Ȋ', 'ɪ':
				result[i] = 'I'
			case 'î', 'ï', 'í', 'ī', 'į', 'ì', '¡', 'ı', 'ĺ', 'ļ', 'Ɩ', 'ǐ', 'ȉ', 'ȋ', 'ɨ', 'ɩ':
				result[i] = 'i'
			case 'Ɉ':
				result[i] = 'J'
			case 'ǰ', 'ȷ', 'ɉ', 'ɟ':
				result[i] = 'j'
			case 'Ƙ', 'Ǩ':
				result[i] = 'K'
			case 'ƙ', 'ǩ':
				result[i] = 'k'
			case 'Ł', 'Ĺ', 'Ļ', 'Ľ', 'Ŀ', 'Ƚ':
				result[i] = 'L'
			case 'ł', 'ľ', 'ŀ', 'ƚ', 'ȴ', 'ɬ', 'ɭ':
				result[i] = 'l'
			case 'ɫ':
				// supposed to be l
			case 'ɱ':
				result[i] = 'm'
			case 'Ñ', 'Ń', 'Ņ', 'Ň', 'Ŋ', 'Ɲ', 'Ǹ', 'ɴ':
				result[i] = 'N'
			case 'ñ', 'ń', 'ņ', 'ň', 'ŉ', 'ŋ', 'ƞ', 'ǹ', 'Ƞ', 'ȵ', 'ɲ', 'ɳ':
				result[i] = 'n'
			case 'Ɵ', 'Ơ', 'Ǒ', 'Ǫ', 'Ǭ', 'Ǿ', 'Ȍ', 'Ȏ', 'Ȫ', 'Ȭ', 'Ȯ', 'Ȱ', 'ɸ':
				result[i] = 'O'
			case 'ô', 'ö', 'ò', 'ó', 'ø', 'ō', 'õ', 'º', 'ő', 'ơ', 'ǒ', 'ǫ', 'ǭ', 'ǿ', 'ȍ', 'ȏ', 'ȫ', 'ȭ', 'ȯ', 'ȱ', 'ɵ':
				result[i] = 'o'
			case 'ƿ':
				result[i] = 'P'
			case 'Þ', 'þ', 'Ƥ', 'ƥ':
				result[i] = 'p'
			case 'Ɋ':
				result[i] = 'Q'
			case 'ƪ', 'ɋ':
				result[i] = 'q'
			case 'Ŕ', 'Ŗ', 'Ř', 'Ʀ', 'Ȑ', 'Ȓ', 'Ɍ', 'ʀ', 'ʁ':
				result[i] = 'R'
			case 'ŕ', 'ŗ', 'ř', 'ȑ', 'ȓ', 'ɍ', 'ɹ', 'ɺ', 'ɻ', 'ɼ', 'ɽ', 'ɾ', 'ɿ':
				result[i] = 'r'
			case 'Ś', 'Š', '§', 'Ŝ', 'Ş', 'Ș', 'ȿ', 'ʂ':
				result[i] = 'S'
			case 'ś', 'š', 'ŝ', 'ş', 'ș':
				result[i] = 's'
			case 'Ţ', 'Ť', 'Ŧ', 'Ƭ', 'Ʈ', 'Ț', 'Ⱦ':
				result[i] = 'T'
			case 'ţ', 'ť', 'ŧ', 'ƫ', 'ƭ', 'ț', 'ȶ':
				result[i] = 't'
			case 'Ũ', 'Ū', 'Ŭ', 'Ů', 'Ű', 'Ų', 'Ư', 'Ǔ', 'Ǖ', 'Ǘ', 'Ǚ', 'Ǜ', 'Ȕ', 'Ȗ', 'Ʉ':
				result[i] = 'U'
			case 'û', 'ü', 'ù', 'ú', 'ū', 'µ', 'ũ', 'ŭ', 'ů', 'ű', 'ų', 'ư', 'Ʊ', 'ǔ', 'ǖ', 'ǘ', 'ǚ', 'ǜ', 'ȕ', 'ȗ', 'ɥ':
				result[i] = 'u'
			case 'Ɣ', 'Ʋ':
				result[i] = 'V'
			case 'Ŵ', 'Ɯ':
				result[i] = 'W'
			case 'ŵ', 'ɯ', 'ɰ':
				result[i] = 'w'
			case 'Ÿ', 'Ŷ', 'Ƴ', 'Ȳ', 'Ɏ':
				result[i] = 'Y'
			case 'ÿ', 'ŷ', 'ƴ', 'ȳ', 'ɏ':
				result[i] = 'y'
			case 'Ž', 'Ź', 'Ż', 'Ƶ', 'Ȥ', 'ɀ':
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
			case '³', 'Ʒ', 'ƺ', 'Ǯ', 'ǯ', 'Ȝ', 'ȝ', 'ɜ', 'ɝ':
				result[i] = '3'
			case 'Ƽ', 'ƽ', 'ƾ':
				result[i] = '5'
			case '·':
				result[i] = '.'
			case 'Ɂ', 'ɂ':
				result[i] = '?'
			case 'Ĳ', 'ĳ', 'Œ', 'œ', 'ƕ', 'Ƣ', 'ƣ', 'Ǆ', 'ǅ', 'ǆ', 'Ǉ', 'ǈ', 'ǉ', 'Ǌ', 'ǋ', 'ǌ', 'Ǣ', 'ǣ', 'Ǳ', 'ǲ', 'ǳ', 'Ǽ', 'ǽ', 'Ȣ', 'ȣ', 'ȸ', 'ȹ', 'ɮ', 'ɶ', 'ɷ':
				// handle ligatures later
			case 'ɣ', 'ɤ':
				// later
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
