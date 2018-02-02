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

func FindXorPlaintext(s string) (string, int, byte) {
	b := decodeHexString(s)

	mostFreq := "ETAOIN SHRDLU"
	bestStr, bestScore, key := "", 0, 0

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
			key = i
		}
	}

	return bestStr, bestScore, byte(key)
}

func DetectSingleCharXor(strs []string) string {
	bestStr, bestScore := "", 0

	for _, str := range strs {
		decryptedStr, score, _ := FindXorPlaintext(str)

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

func BreakRepeatingKeyXor(contents []byte) (string, string) {
	encrypted := make([]byte, base64.StdEncoding.DecodedLen(len(contents)))
	base64.StdEncoding.Decode(encrypted, contents)

	bestKeySizes := map[float64]int{10000.0: 0, 10001.0: 0, 10002.0: 0}

	for keySize := 2; keySize <= 40; keySize++ {

		pos := 0
		dist := 0.0
		for i := 0; i < len(encrypted)/(2*keySize); i++ {
			b1 := encrypted[pos : pos+keySize]
			b2 := encrypted[pos+keySize : pos+2*keySize]
			pos += 2 * keySize

			dist += float64(hammingDistance(b1, b2))
		}

		dist = dist / float64(len(encrypted)/(2*keySize)) / float64(keySize)
		maxDist := 0.0
		for d := range bestKeySizes {
			if d > maxDist {
				maxDist = d
			}
		}

		if dist < maxDist {
			delete(bestKeySizes, maxDist)
			bestKeySizes[dist] = keySize
		}

		//fmt.Printf("keySize %2d dist %1.3f\n", keySize, dist)
	}

	//fmt.Println("best key sizes:", bestKeySizes)

	bestScore, decrypted, bestKey := 0, "", ""
	for _, keySize := range bestKeySizes {
		key := make([]byte, keySize)
		score := 0

		// split into blocks of size keySize
		// transpose
		transposed := splitAndTransposeArray(encrypted, keySize)

		// solve for single-character XOR
		// concatenate the best keys
		for i, block := range transposed {
			_, s, bKey := FindXorPlaintext(hex.EncodeToString(block))
			key[i] = bKey
			score += s
		}

		// decode
		if score > bestScore {
			bestScore = score
			decrypted = string(xorRepeated(encrypted, key))
			bestKey = string(key)
		}

	}

	return bestKey, decrypted
}

func DecryptAESInECBMode(contents []byte, key []byte) string {
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

func splitAndTransposeArray(arr []byte, blockLen int) [][]byte {
	out := make([][]byte, blockLen)
	for i := 0; i < blockLen; i++ {
		out[i] = make([]byte, 0)
	}

	for i, b := range arr {
		out[i%blockLen] = append(out[i%blockLen], b)
	}

	return out
}

func _() {
	fmt.Printf("")
}
