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
			case 'à', 'á', 'â', 'ä', 'æ', 'ã', 'å', 'ā':
				result[i] = 'a'
			case 'ô', 'ö', 'ò', 'ó', 'œ', 'ø', 'ō', 'õ':
				result[i] = 'o'
			case 'û', 'ü', 'ù', 'ú', 'ū':
				result[i] = 'u'
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
