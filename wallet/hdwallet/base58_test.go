package hdwallet

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

type test struct {
	input  string
	output string
}

func testVector() []test {
	return []test{
		{
			input:  "",
			output: "",
		},
		{
			input:  "00",
			output: "1",
		},
		{
			input:  "61",
			output: "2g",
		},
		{
			input:  "626262",
			output: "a3gV",
		},
		{
			input:  "636363",
			output: "aPEr",
		},
		{
			input:  "572e4794",
			output: "3EFU7m",
		},
		{
			input:  "10c8511e",
			output: "Rt5zm",
		},
		{
			input:  "516b6fcd0f",
			output: "ABnLTmg",
		},
		{
			input:  "bf4f89001e670274dd",
			output: "3SEo3LWLoPntC",
		},
		{
			input:  "00000000000000000000",
			output: "1111111111",
		},
		{
			input:  "ecac89cad93923c02321",
			output: "EJDM8drfXA6uyA",
		},
		{
			input:  "73696d706c792061206c6f6e6720737472696e67",
			output: "2cFupjhnEsSn59qHXstmK2ffpLv2",
		},
		{
			input:  "00eb15231dfceb60925886b67d065299925915aeb172c06647",
			output: "1NS17iag9jJgTHD1VXjvLCEnZuQ3rJDE9L",
		},
		{
			input:  "00010966776006953d5567439e5e39f86a0d273beed61967f6",
			output: "16UwLL9Risc3QfPqBUvKofHmBQ7wMtjvM",
		},
	}
}

var decodeMap [256]int8

func init() {
	for i := range decodeMap {
		decodeMap[i] = -1
	}
	for i, b := range b58alphabet {
		decodeMap[b] = int8(i)
	}
}

var (
	bn0  = big.NewInt(0)
	bn58 = big.NewInt(58)
)

func BigIntBase58Encoding(data []byte) (string, []byte) {
	idx := len(data)*138/100 + 1
	buf := make([]byte, idx)
	bn := new(big.Int).SetBytes(data)
	var mo *big.Int
	for bn.Cmp(bn0) != 0 {
		bn, mo = bn.DivMod(bn, bn58, new(big.Int))
		idx--
		buf[idx] = b58alphabet[mo.Int64()]
	}
	for i := range data {
		if data[i] != 0 {
			break
		}
		idx--
		buf[idx] = b58alphabet[0]
	}
	return string(buf[idx:]), buf
}

func BigIntBase58Decoding(str string) ([]byte, error) {
	var zcount int
	for i := 0; i < len(str) && str[i] == '1'; i++ {
		zcount++
	}
	leading := make([]byte, zcount)

	var padChar rune = -1
	src := []byte(str)
	j := 0
	for ; j < len(src) && src[j] == byte(padChar); j++ {
	}

	n := new(big.Int)
	for i := range src[j:] {
		c := decodeMap[src[i]]
		if c == -1 {
			return nil, fmt.Errorf("illegal base58 data at input index: %d", i)
		}
		n.Mul(n, bn58)
		n.Add(n, big.NewInt(int64(c)))
	}
	return append(leading, n.Bytes()...), nil
}

func TestBase58Enc(t *testing.T) {
	for _, test := range testVector() {
		input, err := hex.DecodeString(test.input)
		assert.NoError(t, err)

		output, _ := encode(input)
		assert.Equal(t, test.output, output)
	}
}

func TestBase58Dec(t *testing.T) {
	for _, test := range testVector() {
		_, output := decode(test.output)

		input, err := hex.DecodeString(test.input)
		assert.NoError(t, err)

		assert.Equal(t, input, output)
	}
}

func BenchmarkEncode(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, test := range testVector() {
			input, _ := hex.DecodeString(test.input)

			encode(input)
		}
	}
}

func BenchmarkDecode(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, test := range testVector() {

			decode(test.output)
		}
	}
}

func BenchmarkBigIntEncode(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, test := range testVector() {
			input, _ := hex.DecodeString(test.input)

			BigIntBase58Encoding(input)
		}
	}
}

func BenchmarkBigIntDecode(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, test := range testVector() {

			BigIntBase58Decoding(test.output)
		}
	}
}
