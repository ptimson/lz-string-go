package main

import (
	"fmt"
	"math"
)

func main() {
	CompressToBase64(`During tattooing, ink is injected into the skin, initiating an immune response, and cells called "macrophages" move into the area and "eat up" the ink. The macrophages carry some of the ink to the body's lymph nodes, but some that are filled with ink stay put, embedded in the skin. That's what makes the tattoo visible under the skin. Dalhousie Uiversity's Alec Falkenham is developing a topical cream that works by targeting the macrophages that have remained at the site of the tattoo. New macrophages move in to consume the previously pigment-filled macrophages and then migrate to the lymph nodes, eventually taking all the dye with them. "When comparing it to laser-based tattoo removal, in which you see the burns, the scarring, the blisters, in this case, we've designed a drug that doesn't really have much off-target effect," he said. "We're not targeting any of the normal skin cells, so you won't see a lot of inflammation. In fact, based on the process that we're actually using, we don't think there will be any inflammation at all and it would actually be anti-inflammatory.`)
}

const keyStrBase64 = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/="
const keyStrUriSafe = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+-$"

var baseReverseDic = make(map[string]map[rune]int)

func getBaseValue(alphabet string, char rune) (int, bool) {
	_, ok := baseReverseDic[alphabet]
	if !ok {
		baseReverseDic[alphabet] = make(map[rune]int, len(alphabet))
		for i, c := range alphabet {
			baseReverseDic[alphabet][c] = i
		}
	}
	value, ok := baseReverseDic[alphabet][char]
	return value, ok
}

func CompressToBase64(input string) string {
	if input == "" {
		return ""
	}

	res := compress(input, 6, func(i int) rune {
		return rune(keyStrBase64[i])
	})

	switch len(res) % 4 {
	case 1:
		return res + "==="
	case 2:
		return res + "=="
	case 3:
		return res + "="
	}

	return res
}

func decompressFromBase64(input string) string {
	if input == "" {
		return ""
	}

	return decompress(len(input), 32, func(i int) (int, bool) {
		return getBaseValue(keyStrBase64, rune(keyStrBase64[i]))
	})
}

func decompress(length int, A int, f func(i int) (int, bool)) string {
	return ""
}

func compress(uncompressed string, bitsPerChar int, getCharFromInt func(i int) rune) string {
	if uncompressed == "" {
		return ""
	}

	w := ""
	value := 0
	dictSize := 3
	numBits := 2
	dataVal := 0
	dataPosition := 0
	enlargeIn := float64(2) // Compensate for the first entry which should not count
	var data []rune

	dictionary := make(map[string]int)
	dictionaryToCreate := make(map[string]struct{})

	for _, c := range uncompressed {
		char := fmt.Sprintf("%c", c)
		_, ok := dictionary[char]
		if !ok {
			dictionary[char] = dictSize
			dictSize++
			dictionaryToCreate[char] = struct{}{}
		}

		wc := fmt.Sprintf("%s%s", w, char)
		_, ok = dictionary[wc]
		if ok {
			w = wc
			continue
		}

		_, ok = dictionaryToCreate[w]
		if !ok {
			value, ok = dictionary[w]
			if !ok {
				value = 0
			}
			for i := 0; i < numBits; i++ {
				dataVal = (dataVal << 1) | (value & 1)
				if dataPosition == bitsPerChar-1 {
					data = append(data, getCharFromInt(dataVal))
					dataPosition = 0
					dataVal = 0
				} else {
					dataPosition++
				}
				value = value >> 1
			}
		} else {

			var max255 uint8 = 255
			if len(w) > 0 && w[0] <= max255 {
				for i := 0; i < numBits; i++ {
					dataVal = dataVal << 1
					if dataPosition == bitsPerChar-1 {
						data = append(data, getCharFromInt(dataVal))
						dataPosition = 0
						dataVal = 0
					} else {
						dataPosition++
					}
				}
				value = int(w[0])
				for i := 0; i < 8; i++ {
					dataVal = dataVal<<1 | (value & 1)
					if dataPosition == bitsPerChar-1 {
						data = append(data, getCharFromInt(dataVal))
						dataPosition = 0
						dataVal = 0
					} else {
						dataPosition++
					}
					value = value >> 1
				}

			} else {

				value = 1
				for i := 0; i < numBits; i++ {
					dataVal = (dataVal << 1) | value
					if dataPosition == bitsPerChar-1 {
						data = append(data, getCharFromInt(dataVal))
						dataPosition = 0
						dataVal = 0
					} else {
						dataPosition++
					}
					value = 0
				}

				value = int(w[0])

				for i := 0; i < 16; i++ {
					dataVal = dataVal<<1 | (value & 1)
					if dataPosition == bitsPerChar-1 {
						data = append(data, getCharFromInt(dataVal))
						dataPosition = 0
						dataVal = 0
					} else {
						dataPosition++
					}
					value = value >> 1
				}
			}
			enlargeIn--
			if enlargeIn == 0 {
				enlargeIn = math.Pow(2, float64(numBits))
				numBits++
			}
			delete(dictionaryToCreate, w)
		}

		enlargeIn--
		if enlargeIn == 0 {
			enlargeIn = math.Pow(2, float64(numBits))
			numBits++
		}
		// Add wc to the dictionary.
		dictionary[wc] = dictSize
		dictSize++
		w = char
	}

	/////////
	// Output the code for w.
	/////////
	_, ok := dictionaryToCreate[w]
	if w != "" && ok {
		var max255 uint8 = 255
		if len(w) > 0 && w[0] <= max255 {
			for i := 0; i < numBits; i++ {
				dataVal = dataVal << 1
				if dataPosition == bitsPerChar-1 {
					data = append(data, getCharFromInt(dataVal))
					dataPosition = 0
					dataVal = 0
				} else {
					dataPosition++
				}
			}
			value = int(w[0])
			for i := 0; i < 8; i++ {
				dataVal = dataVal<<1 | (value & 1)
				if dataPosition == bitsPerChar-1 {
					data = append(data, getCharFromInt(dataVal))
					dataPosition = 0
					dataVal = 0
				} else {
					dataPosition++
				}
				value = value >> 1
			}

		} else {

			value = 1
			for i := 0; i < numBits; i++ {
				dataVal = (dataVal << 1) | value
				if dataPosition == bitsPerChar-1 {
					data = append(data, getCharFromInt(dataVal))
					dataPosition = 0
					dataVal = 0
				} else {
					dataPosition++
				}
				value = 0
			}

			value = int(w[0])

			for i := 0; i < 16; i++ {
				dataVal = dataVal<<1 | (value & 1)
				if dataPosition == bitsPerChar-1 {
					data = append(data, getCharFromInt(dataVal))
					dataPosition = 0
					dataVal = 0
				} else {
					dataPosition++
				}
				value = value >> 1
			}
		}

		enlargeIn--
		if enlargeIn == 0 {
			enlargeIn = math.Pow(2, float64(numBits))
			numBits++
		}
		delete(dictionaryToCreate, w)
	}

	_, ok = dictionaryToCreate[w]
	if w != "" && !ok {
		value, ok = dictionary[w]
		if !ok {
			value = 0
		}
		for i := 0; i < numBits; i++ {
			dataVal = dataVal<<1 | (value & 1)
			if dataPosition == bitsPerChar-1 {
				data = append(data, getCharFromInt(dataVal))
				dataPosition = 0
				dataVal = 0
			} else {
				dataPosition++
			}
			value = value >> 1
		}
	}

	if w != "" {
		enlargeIn--
		if enlargeIn == 0 {
			enlargeIn = math.Pow(2, float64(numBits))
			numBits++
		}
	}

	// Mark the end of the stream
	value = 2
	for i := 0; i < numBits; i++ {
		dataVal = (dataVal << 1) | (value & 1)
		if dataPosition == bitsPerChar-1 {
			data = append(data, getCharFromInt(dataVal))
			dataPosition = 0
			dataVal = 0
		} else {
			dataPosition++
		}
		value = value >> 1
	}

	// Flush the last char
	for {
		dataVal = dataVal << 1
		if dataPosition == bitsPerChar-1 {
			data = append(data, getCharFromInt(dataVal))
			break
		} else {
			dataPosition++
		}
	}

	return string(data)
}
