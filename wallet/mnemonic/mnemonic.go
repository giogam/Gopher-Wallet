package mnemonic

import (
	"crypto/rand"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/binary"
	"math/big"
	"strings"

	"golang.org/x/crypto/pbkdf2"
)

// Checksum appends checksum to the given entropy bytes
func Checksum(entropy []byte) []byte {

	checksumLen := uint(len(entropy)) / uint(4)

	hash := sha256.Sum256(entropy)

	intEntropy := new(big.Int).SetBytes(entropy)
	intChecksun := new(big.Int).SetUint64(uint64(hash[0]))

	intChecksun.Rsh(intChecksun, (8 - checksumLen))
	intEntropy.Lsh(intEntropy, checksumLen)

	return intChecksun.Or(intEntropy, intChecksun).Bytes()
}

// NewEntropy generates new entropy
func NewEntropy(bitLen uint) []byte {

	var entropyLen = bitLen / 8

	entropy := make([]byte, entropyLen)
	rand.Read(entropy)

	return entropy
}

// NewWordlist generates new entropy
func NewWordlist(language string) []string {
	return strings.Split(EnglishWords, "\n")
}

// NewMnemonic generates new mnemonic
func NewMnemonic(entropy []byte, wordlist []string) []string {

	entropyLen := len(entropy) * 8
	checksumLen := entropyLen / 32
	mnemonicLen := (entropyLen + checksumLen) / 11

	bitMask := new(big.Int).SetInt64(2047)
	entropyInt := new(big.Int).SetBytes(Checksum(entropy))
	slice := new(big.Int).SetInt64(0)

	mnemonic := make([]string, mnemonicLen)

	for i := mnemonicLen - 1; i >= 0; i-- {
		slice.And(entropyInt, bitMask)

		pad := make([]byte, 2-len(slice.Bytes()))
		bytes := append(pad, slice.Bytes()...)
		mnemonic[i] = wordlist[binary.BigEndian.Uint16(bytes)]

		entropyInt.Rsh(entropyInt, uint(11))
	}

	return mnemonic
}

// NewSeed generates new seed given a mnemonic and a passphrase
func NewSeed(mnemonic []string, passphrase string) []byte {

	return pbkdf2.Key([]byte(strings.Join(mnemonic, " ")), []byte("mnemonic"+passphrase), 2048, 64, sha512.New)
}

// Bip39RandomSeed generates new mnemonic and the respective seed
func Bip39RandomSeed(bitLen uint, language string, passphrase string) ([]byte, []string) {

	mnemonic := NewMnemonic(NewEntropy(bitLen), NewWordlist(language))

	return NewSeed(mnemonic, passphrase), mnemonic
}
