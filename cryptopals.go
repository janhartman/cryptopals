package main

import "fmt"
import "cryptopals/set1"

func main() {
	fmt.Printf("HexToBase64: %t\n", set1.TestHexToBase64())
	fmt.Printf("FixedXor: %t\n", set1.TestFixedXor())
	fmt.Printf("Single-byte xor: %s\n", set1.TestFindXorPlaintext())
	fmt.Printf("Detect single-byte xor: %s", set1.TestDetectSingleCharXor())
	fmt.Printf("EncryptRepeatingKeyXor: %t\n", set1.TestEncryptRepeatingKeyXor())
	fmt.Printf("BreakRepeatingKeyXor: %s\n", set1.TestBreakRepeatingKeyXor())
	fmt.Printf("DecryptAESInECBMode: %s\n", set1.TestDecryptAESInECBMode())
	fmt.Printf("DetectAESInECBMode: %s\n", set1.TestDetectAESInECBMode())
}
