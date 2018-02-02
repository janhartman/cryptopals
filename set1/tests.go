package set1

import "io/ioutil"
import "strings"

func TestHexToBase64() bool {
	input := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	output := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"

	return HexToBase64(input) == output
}

func TestFixedXor() bool {
	in1 := "1c0111001f010100061a024b53535009181c"
	in2 := "686974207468652062756c6c277320657965"
	out := "746865206b696420646f6e277420706c6179"

	return FixedXor(in1, in2) == out
}

func TestFindXorPlaintext() string {
	input := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"

	str, _, _ := FindXorPlaintext(input)
	return str
}

func TestDetectSingleCharXor() string {
	contents, err := ioutil.ReadFile("set1/input4.txt")
	if err != nil {
		panic("Could not read file")
	}

	strs := strings.Split(string(contents), "\n")
	return DetectSingleCharXor(strs)
}

func TestEncryptRepeatingKeyXor() bool {
	input := "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"
	key := "ICE"

	output := "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f"

	return EncryptRepeatingKeyXor(input, key) == output
}

func TestBreakRepeatingKeyXor() string {
	contents, err := ioutil.ReadFile("set1/input6.txt")
	if err != nil {
		panic("Could not read file")
	}

	key, _ := BreakRepeatingKeyXor(contents)
	return key
}

func TestDecryptAESInECBMode() string {
	contents, err := ioutil.ReadFile("set1/input7.txt")
	if err != nil {
		panic("Could not read file")
	}

	key := []byte("YELLOW SUBMARINE")
	return DecryptAESInECBMode(contents, key)
}
