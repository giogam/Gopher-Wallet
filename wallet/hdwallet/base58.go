package hdwallet

import (
	"crypto/sha256"
	"math"
	"strings"
)

const (
	b58alphabet         = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"
	log256      float64 = 5.54
	log58       float64 = 4.06
)

func encode(data []byte) (string, []byte) {

	dataSz := len(data)
	i, j, high, carry, zCount := 0, 0, 0, 0, 0

	for zCount < dataSz && data[zCount] == 0 {
		zCount++
	}

	size := int(math.Ceil(float64(dataSz-zCount) * log256 / log58))
	buff := make([]byte, size)

	for i, j, high = zCount, 0, size-1; i < dataSz; i, high = i+1, j {

		for carry, j = int(data[i]), size-1; j > high || carry != 0; j-- {

			carry += 256 * int(buff[j])
			buff[j] = byte(carry % 58)
			carry /= 58
		}
	}

	for j = 0; j < size && buff[j] == 0; j++ {
	}

	b58 := make([]byte, (size-j)+zCount)

	if zCount != 0 {
		for i := 0; i < zCount; i++ {
			b58[i] = '1'
		}
	}

	for i := zCount; j < size; i, j = i+1, j+1 {
		b58[i] = b58alphabet[buff[j]]
	}

	return string(b58), b58
}

func decode(str string) (string, []byte) {

	strSz := len(str)
	i, j, high, carry, zCount := 0, 0, 0, 0, 0

	for zCount < strSz && str[zCount] == '1' {
		zCount++
	}

	size := int(math.Ceil(float64(strSz-zCount) * log58 / log256))
	buff := make([]byte, size)

	for i, j, high = zCount, 0, size-1; i < strSz; i, high = i+1, j {

		for carry, j = strings.Index(b58alphabet, string(str[i])), size-1; j > high || carry != 0; j-- {

			carry += 58 * int(buff[j])
			buff[j] = byte(carry % 256)
			carry /= 256
		}
	}

	for j = 0; j < size && buff[j] == 0; j++ {
	}

	b256 := make([]byte, (size-j)+zCount)

	if zCount != 0 {
		for i := 0; i < zCount; i++ {
			b256[i] = byte(0)
		}
	}

	for i := zCount; j < size; i, j = i+1, j+1 {
		b256[i] = buff[j]
	}

	return string(b256), b256
}

// B58CheckEncode encodes data in Base58Check format
func B58CheckEncode(version int, data []byte) (string, []byte) {

	data = append([]byte{byte(version)}, data...)
	hash := sha256.Sum256(data)
	hash = sha256.Sum256(hash[:])

	return encode(append(data, hash[:4]...))
}

// B58CheckDecode decodes data encoded in Base58Check format
func B58CheckDecode(data string) (int, []byte, []byte) {

	_, encoded := decode(data)

	return int(encoded[0]),
		encoded[1 : len(encoded)-4],
		encoded[len(encoded)-4:]
}
