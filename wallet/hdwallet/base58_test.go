package hdwallet

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

type b58test struct {
	input  string
	output string
}

func b58testVector() []b58test {
	return []b58test{
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

type b58Checktest struct {
	input   string
	version string
	output  string
}

func b58ChecktestVector() []b58Checktest {
	return []b58Checktest{
		{
			input:   "65a16059864a2fdbc7c99a4723a8395bc6f188eb",
			version: "0",
			output:  "1AGNa15ZQXAZUgFiqJ2i7Z2DPU2J6hW62i",
		},
		{
			input:   "74f209f6ea907e2ea48f74fae05782ae8a665257",
			version: "5",
			output:  "3CMNFxN1oHBc4R1EpboAL5yzHGgE611Xou",
		},
		{
			input:   "53c0307d6851aa0ce7825ba883c6bd9ad242b486",
			version: "111",
			output:  "mo9ncXisMeAoXwqcV5EWuyncbmCcQN4rVs",
		},
		{
			input:   "6349a418fc4578d10a372b54b45c280cc8c4382f",
			version: "196",
			output:  "2N2JD6wb56AfK4tfmM6PwdVmoYk2dCKf4Br",
		},
		{
			input:   "eddbdc1168f1daeadbd3e44c1e3f8f5a284c2029f78ad26af98583a499de5b19",
			version: "128",
			output:  "5Kd3NBUAdUnhyzenEwVLy9pBKxSwXvE9FMPyR4UKZvpe6E3AgLr",
		},
		{
			input:   "36cb93b9ab1bdabf7fb9f2c04f1b9cc879933530ae7842398eef5a63a56800c2",
			version: "239",
			output:  "9213qJab2HNEpMpYNBa7wHGFKKbkDn24jpANDs2huN3yi4J11ko",
		},
		{
			input:   "6d23156cbbdcc82a5a47eee4c2c7c583c18b6bf4",
			version: "0",
			output:  "1Ax4gZtb7gAit2TivwejZHYtNNLT18PUXJ",
		},
		{
			input:   "fcc5460dd6e2487c7d75b1963625da0e8f4c5975",
			version: "5",
			output:  "3QjYXhTkvuj8qPaXHTTWb5wjXhdsLAAWVy",
		},
		{
			input:   "f1d470f9b02370fdec2e6b708b08ac431bf7a5f7",
			version: "111",
			output:  "n3ZddxzLvAY9o7184TB4c6FJasAybsw4HZ",
		},
		{
			input:   "c579342c2c4c9220205e2cdc285617040c924a0a",
			version: "196",
			output:  "2NBFNJTktNa7GZusGbDbGKRZTxdK9VVez3n",
		},
		{
			input:   "a326b95ebae30164217d7a7f57d72ab2b54e3be64928a19da0210b9568d4015e",
			version: "128",
			output:  "5K494XZwps2bGyeL71pWid4noiSNA2cfCibrvRWqcHSptoFn7rc",
		},
		{
			input:   "d6bca256b5abc5602ec2e1c121a08b0da2556587430bcf7e1898af2224885203",
			version: "239",
			output:  "93DVKyFYwSN6wEo3E2fCrFPUp17FtrtNi2Lf7n4G3garFb16CRj",
		},
		{
			input:   "7987ccaa53d02c8873487ef919677cd3db7a6912",
			version: "0",
			output:  "1C5bSj1iEGUgSTbziymG7Cn18ENQuT36vv",
		},
		{
			input:   "63bcc565f9e68ee0189dd5cc67f1b0e5f02f45cb",
			version: "5",
			output:  "3AnNxabYGoTxYiTEZwFEnerUoeFXK2Zoks",
		},
		{
			input:   "ef66444b5b17f14e8fae6e7e19b045a78c54fd79",
			version: "111",
			output:  "n3LnJXCqbPjghuVs8ph9CYsAe4Sh4j97wk",
		},
		{
			input:   "c3e55fceceaa4391ed2a9677f4a4d34eacd021a0",
			version: "196",
			output:  "2NB72XtkjpnATMggui83aEtPawyyKvnbX2o",
		},
		{
			input:   "e75d936d56377f432f404aabb406601f892fd49da90eb6ac558a733c93b47252",
			version: "128",
			output:  "5KaBW9vNtWNhc3ZEDyNCiXLPdVPHCikRxSBWwV9NrpLLa4LsXi9",
		},
		{
			input:   "44c4f6a096eac5238291a94cc24c01e3b19b8d8cef72874a079e00a242237a52",
			version: "239",
			output:  "927CnUkUbasYtDwYwVn2j8GdTuACNnKkjZ1rpZd2yBB1CLcnXpo",
		},
		{
			input:   "adc1cc2081a27206fae25792f28bbc55b831549d",
			version: "0",
			output:  "1Gqk4Tv79P91Cc1STQtU3s1W6277M2CVWu",
		},
		{
			input:   "188f91a931947eddd7432d6e614387e32b244709",
			version: "5",
			output:  "33vt8ViH5jsr115AGkW6cEmEz9MpvJSwDk",
		},
		{
			input:   "1694f5bc1a7295b600f40018a618a6ea48eeb498",
			version: "111",
			output:  "mhaMcBxNh5cqXm4aTQ6EcVbKtfL6LGyK2H",
		},
		{
			input:   "3b9b3fd7a50d4f08d1a5b0f62f644fa7115ae2f3",
			version: "196",
			output:  "2MxgPqX1iThW3oZVk9KoFcE5M4JpiETssVN",
		},
		{
			input:   "091035445ef105fa1bb125eccfb1882f3fe69592265956ade751fd095033d8d0",
			version: "128",
			output:  "5HtH6GdcwCJA4ggWEL1B3jzBBUB8HPiBi9SBc5h9i4Wk4PSeApR",
		},
		{
			input:   "b4204389cef18bbe2b353623cbf93e8678fbc92a475b664ae98ed594e6cf0856",
			version: "239",
			output:  "92xFEve1Z9N8Z641KQQS7ByCSb8kGjsDzw6fAmjHN1LZGKQXyMq",
		},
		{
			input:   "c4c1b72491ede1eedaca00618407ee0b772cad0d",
			version: "0",
			output:  "1JwMWBVLtiqtscbaRHai4pqHokhFCbtoB4",
		},
		{
			input:   "f6fe69bcb548a829cce4c57bf6fff8af3a5981f9",
			version: "5",
			output:  "3QCzvfL4ZRvmJFiWWBVwxfdaNBT8EtxB5y",
		},
		{
			input:   "261f83568a098a8638844bd7aeca039d5f2352c0",
			version: "111",
			output:  "mizXiucXRCsEriQCHUkCqef9ph9qtPbZZ6",
		},
		{
			input:   "e930e1834a4d234702773951d627cce82fbb5d2e",
			version: "196",
			output:  "2NEWDzHWwY5ZZp8CQWbB7ouNMLqCia6YRda",
		},
		{
			input:   "d1fab7ab7385ad26872237f1eb9789aa25cc986bacc695e07ac571d6cdac8bc0",
			version: "128",
			output:  "5KQmDryMNDcisTzRp3zEq9e4awRmJrEVU1j5vFRTKpRNYPqYrMg",
		},
		{
			input:   "037f4192c630f399d9271e26c575269b1d15be553ea1a7217f0cb8513cef41cb",
			version: "239",
			output:  "91cTVUcgydqyZLgaANpf1fvL55FH53QMm4BsnCADVNYuWuqdVys",
		},
		{
			input:   "5eadaf9bb7121f0f192561a5a62f5e5f54210292",
			version: "0",
			output:  "19dcawoKcZdQz365WpXWMhX6QCUpR9SY4r",
		},
		{
			input:   "3f210e7277c899c3a155cc1c90f4106cbddeec6e",
			version: "5",
			output:  "37Sp6Rv3y4kVd1nQ1JV5pfqXccHNyZm1x3",
		},
		{
			input:   "c8a3c2a09a298592c3e180f02487cd91ba3400b5",
			version: "111",
			output:  "myoqcgYiehufrsnnkqdqbp69dddVDMopJu",
		},
		{
			input:   "99b31df7c9068d1481b596578ddbb4d3bd90baeb",
			version: "196",
			output:  "2N7FuwuUuoTBrDFdrAZ9KxBmtqMLxce9i1C",
		},
		{
			input:   "c7666842503db6dc6ea061f092cfb9c388448629a6fe868d068c42a488b478ae",
			version: "128",
			output:  "5KL6zEaMtPRXZKo1bbMq7JDjjo1bJuQcsgL33je3oY8uSJCR5b4",
		},
		{
			input:   "ea577acfb5d1d14d3b7b195c321566f12f87d2b77ea3a53f68df7ebf8604a801",
			version: "239",
			output:  "93N87D6uxSBzwXvpokpzg8FFmfQPmvX4xHoWQe3pLdYpbiwT5YV",
		},
		{
			input:   "1ed467017f043e91ed4c44b4e8dd674db211c4e6",
			version: "0",
			output:  "13p1ijLwsnrcuyqcTvJXkq2ASdXqcnEBLE",
		},
		{
			input:   "5ece0cadddc415b1980f001785947120acdb36fc",
			version: "5",
			output:  "3ALJH9Y951VCGcVZYAdpA3KchoP9McEj1G",
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

func TestBase58CheckEnc(t *testing.T) {
	for _, test := range b58ChecktestVector() {
		input, err := hex.DecodeString(test.input)
		assert.NoError(t, err)
		version, err := strconv.Atoi(test.version)
		assert.NoError(t, err)

		output, _ := B58CheckEncode(version, input)
		assert.Equal(t, test.output, output)
	}
}
func TestBase58CheckDec(t *testing.T) {
	for _, test := range b58ChecktestVector() {

		ver, data, check := B58CheckDecode(test.output)

		input, err := hex.DecodeString(test.input)
		assert.NoError(t, err)
		assert.Equal(t, input, data)

		version := strconv.Itoa(ver)
		assert.Equal(t, version, test.version)

		data = append([]byte{byte(ver)}, data...)
		hash := sha256.Sum256(data)
		hash = sha256.Sum256(hash[:])
		assert.Equal(t, check, hash[:4])
	}
}
func TestBase58Enc(t *testing.T) {
	for _, test := range b58testVector() {
		input, err := hex.DecodeString(test.input)
		assert.NoError(t, err)

		output, _ := encode(input)
		assert.Equal(t, test.output, output)
	}
}
func TestBase58Dec(t *testing.T) {
	for _, test := range b58testVector() {
		_, output := decode(test.output)

		input, err := hex.DecodeString(test.input)
		assert.NoError(t, err)

		assert.Equal(t, input, output)
	}
}
func BenchmarkEncode(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, test := range b58testVector() {
			input, _ := hex.DecodeString(test.input)

			encode(input)
		}
	}
}
func BenchmarkDecode(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, test := range b58testVector() {

			decode(test.output)
		}
	}
}
func BenchmarkBigIntEncode(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, test := range b58testVector() {
			input, _ := hex.DecodeString(test.input)

			BigIntBase58Encoding(input)
		}
	}
}
func BenchmarkBigIntDecode(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, test := range b58testVector() {

			BigIntBase58Decoding(test.output)
		}
	}
}
