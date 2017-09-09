package hdwallet

import (
	"encoding/hex"
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

func TestBase58Enc(t *testing.T) {
	for _, test := range testVector() {
		input, err := hex.DecodeString(test.input)
		assert.NoError(t, err)

		output, _ := Encode(input)
		assert.Equal(t, test.output, output)
	}
}
