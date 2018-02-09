package set1

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func getState() [][]byte {
	var state = [][]byte{{0, 1, 2, 3}, {4, 5, 6, 7}, {8, 9, 10, 11}, {12, 13, 14, 15}}
	return state
}

func TestInvShiftRows(t *testing.T) {
	state := getState()
	var rightShifted = [][]byte{{0, 1, 2, 3}, {7, 4, 5, 6}, {10, 11, 8, 9}, {13, 14, 15, 12}}

	shifted := invShiftRows(state)

	//fmt.Println(rightShifted, shifted)

	assert.True(t, reflect.DeepEqual(rightShifted, shifted), "shiftRows equal")
}

func TestShiftRows(t *testing.T) {
	state := getState()
	var leftShifted = [][]byte{{0, 1, 2, 3}, {5, 6, 7, 4}, {10, 11, 8, 9}, {15, 12, 13, 14}}

	shifted := shiftRows(state)

	//fmt.Println(leftShifted, shifted)

	assert.True(t, reflect.DeepEqual(leftShifted, shifted), "invShiftRows equal")
}

func TestTo2Darray(t *testing.T) {
	state := getState()
	var b = []byte{0, 4, 8, 12, 1, 5, 9, 13, 2, 6, 10, 14, 3, 7, 11, 15}

	assert.True(t, reflect.DeepEqual(state, to2Darray(b)), "to2Darray equal")
}

func TestTo1Darray(t *testing.T) {
	state := getState()
	var b = []byte{0, 4, 8, 12, 1, 5, 9, 13, 2, 6, 10, 14, 3, 7, 11, 15}

	assert.True(t, reflect.DeepEqual(b, to1Darray(state)), "to1Darray equal")
}

func TestGenerateKeys(t *testing.T) {
	var keys = [][]byte{{0x69, 0x20, 0xe2, 0x99, 0xa5, 0x20, 0x2a, 0x6d, 0x65, 0x6e, 0x63, 0x68, 0x69, 0x74, 0x6f, 0x2a}, {0xfa, 0x88, 0x07, 0x60, 0x5f, 0xa8, 0x2d, 0x0d, 0x3a, 0xc6, 0x4e, 0x65, 0x53, 0xb2, 0x21, 0x4f}, {0xcf, 0x75, 0x83, 0x8d, 0x90, 0xdd, 0xae, 0x80, 0xaa, 0x1b, 0xe0, 0xe5, 0xf9, 0xa9, 0xc1, 0xaa}, {0x18, 0x0d, 0x2f, 0x14, 0x88, 0xd0, 0x81, 0x94, 0x22, 0xcb, 0x61, 0x71, 0xdb, 0x62, 0xa0, 0xdb}, {0xba, 0xed, 0x96, 0xad, 0x32, 0x3d, 0x17, 0x39, 0x10, 0xf6, 0x76, 0x48, 0xcb, 0x94, 0xd6, 0x93}, {0x88, 0x1b, 0x4a, 0xb2, 0xba, 0x26, 0x5d, 0x8b, 0xaa, 0xd0, 0x2b, 0xc3, 0x61, 0x44, 0xfd, 0x50}, {0xb3, 0x4f, 0x19, 0x5d, 0x09, 0x69, 0x44, 0xd6, 0xa3, 0xb9, 0x6f, 0x15, 0xc2, 0xfd, 0x92, 0x45}, {0xa7, 0x00, 0x77, 0x78, 0xae, 0x69, 0x33, 0xae, 0x0d, 0xd0, 0x5c, 0xbb, 0xcf, 0x2d, 0xce, 0xfe}, {0xff, 0x8b, 0xcc, 0xf2, 0x51, 0xe2, 0xff, 0x5c, 0x5c, 0x32, 0xa3, 0xe7, 0x93, 0x1f, 0x6d, 0x19}, {0x24, 0xb7, 0x18, 0x2e, 0x75, 0x55, 0xe7, 0x72, 0x29, 0x67, 0x44, 0x95, 0xba, 0x78, 0x29, 0x8c}, {0xae, 0x12, 0x7c, 0xda, 0xdb, 0x47, 0x9b, 0xa8, 0xf2, 0x20, 0xdf, 0x3d, 0x48, 0x58, 0xf6, 0xb1}}

	key := keys[0]

	keys2d := make([][][]byte, 11)
	for i := 0; i < 11; i++ {
		keys2d[i] = to2Darray(keys[i])
	}

	generatedKeys := generateKeys(key)

	for k := 0; k < 11; k++ {
		for i := 0; i < 4; i++ {
			for j := 0; j < 4; j++ {
				if keys2d[k][i][j] != generatedKeys[k][i][j] {
					fmt.Println(k, i, j)
				}
			}
		}
	}

	assert.True(t, reflect.DeepEqual(keys2d, generatedKeys), "generateKeys equal")
}

func TestMixColumns(t *testing.T) {
	var state = [][]byte{{219, 242, 1, 198}, {19, 10, 1, 198}, {83, 34, 1, 198}, {69, 92, 1, 198}}
	var mixedState = [][]byte{{142, 159, 1, 198}, {77, 220, 1, 198}, {161, 88, 1, 198}, {188, 157, 1, 198}}

	mixedColumns := mixColumns(state)

	assert.True(t, reflect.DeepEqual(mixedState, mixedColumns), "mixColumns equal")
}

func TestInvMixColumns(t *testing.T) {
	var state = [][]byte{{219, 242, 1, 198}, {19, 10, 1, 198}, {83, 34, 1, 198}, {69, 92, 1, 198}}
	var mixedState = [][]byte{{142, 159, 1, 198}, {77, 220, 1, 198}, {161, 88, 1, 198}, {188, 157, 1, 198}}

	mixedColumns := invMixColumns(mixedState)

	assert.True(t, reflect.DeepEqual(state, mixedColumns), "invMixColumns equal")
}

func TestAES(t *testing.T) {
	plaintext := []byte("hello darkness my old friend, I'")
	key := []byte("yellow submarine")
	encrypted := encryptAESInECBMode(plaintext, key)
	decrypted := decryptAESInECBMode(encrypted, key)

	assert.True(t, reflect.DeepEqual(plaintext, decrypted), "AES equal")
}
