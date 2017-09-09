package hdwallet

const b58alphabet = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"

// Encode encodes the given slice of data in Base58
// returns the encoded bytes and a string representation of them
func Encode(data []byte) (string, []byte) {

	dataSz := len(data)
	i, j, high, carry, zCount := 0, 0, 0, 0, 0

	// Count zero bytes
	for zCount < dataSz && data[zCount] == 0 {
		zCount++
	}

	size := (dataSz-zCount)*138/100 + 1

	buff := make([]byte, size)

	for i, j, high = zCount, 0, size-1; i < dataSz; i, high = i+1, j {

		for carry, j = int(data[i]), size-1; j > high || carry != 0; j-- {

			carry += 256 * int(buff[j])
			buff[j] = byte(carry % 58)
			carry /= 58
		}
	}

	j = 0
	for j < size && buff[j] == 0 {
		j++
	}

	b58 := make([]byte, (size-j)+zCount)

	if zCount != 0 {

		for i := 0; i < zCount; i++ {
			b58[i] = '1'
		}
	}
	for i := zCount; j < size; i++ {
		b58[i] = b58alphabet[buff[j]]
		j++
	}

	return string(b58), b58

}
