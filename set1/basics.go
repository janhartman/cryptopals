package set1

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"math"
	"strings"
)

func HexToBase64(input string) string {
	b := decodeHexString(input)

	return base64.StdEncoding.EncodeToString(b)
}

func FixedXor(s1, s2 string) string {
	b1 := decodeHexString(s1)
	b2 := decodeHexString(s2)

	return hex.EncodeToString(xorBytes(b1, b2))
}

func FindXorPlaintext(s string) (string, int) {
	b := decodeHexString(s)

	mostFreq := "ETAOIN SHRDLU"
	bestStr, bestScore := "", 0

	// try all bytes, score with most frequent characters in English
	for i := 0; i < 256; i++ {
		s := string(xorChar(b, byte(i)))
		score := 0

		for _, c := range mostFreq {
			score += strings.Count(strings.ToUpper(s), string(c))
		}

		if score > bestScore {
			bestScore = score
			bestStr = s
		}
	}

	return bestStr, bestScore
}

func DetectSingleCharXor(strs []string) string {
	bestStr, bestScore := "", 0

	for _, str := range strs {
		decryptedStr, score := FindXorPlaintext(str)

		if score > bestScore {
			bestStr = decryptedStr
			bestScore = score
		}
	}

	return bestStr
}

func EncryptRepeatingKeyXor(s string, key string) string {
	return hex.EncodeToString(xorRepeated([]byte(s), []byte(key)))
}

func BreakRepeatingKeyXor(encrypted []byte) string {
	bestKeySizes := []int{10000, 10000, 10000}

	for keySize := 2; keySize <= 40; keySize++ {
		b1 := encrypted[:keySize]
		b2 := encrypted[keySize : 2*keySize]

		dist := hammingDistance(b1, b2) / keySize

		for i, ks := range bestKeySizes {
			if dist < ks {
				bestKeySizes = append(bestKeySizes[:i], dist)
				bestKeySizes = append(bestKeySizes, bestKeySizes[i:2]...)
				break
			}
		}
	}

	for _, keySize := range bestKeySizes {
		// split into blocks of size keySize
		bytes.SplitN(encrypted, encrypted, keySize)
		// transpose

		// solve for single-character XOR

		// concatenate the best keys
	}

	// decode

	return ""
}

func decodeHexString(s string) []byte {
	b, err := hex.DecodeString(s)
	if err != nil {
		panic("Error decoding hex string")
	}

	return b
}

func xorBytes(b1, b2 []byte) []byte {
	out := make([]byte, len(b1))

	for i := 0; i < len(b1); i++ {
		out[i] = b1[i] ^ b2[i]
	}

	return out
}

func xorChar(buf []byte, b byte) []byte {
	return xorBytes(buf, bytes.Repeat([]byte{b}, len(buf)))
}

func xorRepeated(b, key []byte) []byte {
	repeats := int(math.Ceil(float64(len(b)) / float64(len(key))))
	repeatedKey := bytes.Repeat(key, repeats)[:len(b)]

	return xorBytes(b, repeatedKey)
}

func hammingDistance(b1, b2 []byte) int {
	diff := xorBytes(b1, b2)
	sum := 0
	for _, b := range diff {
		for b > 0 {
			sum += int(b & 1)
			b = b >> 1
		}

	}
	return sum
}

func _() {
	fmt.Printf("")
}
