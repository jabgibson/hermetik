package h7k

import (
	"github.com/cespare/xxhash/v2"
)

type shift func(byte) byte

func Encrypt(cipher []byte, subject []byte) ([]byte, error) {
	return transform(encrypt, cipher, subject), nil
}

func Decrypt(cipher []byte, subject []byte) ([]byte, error) {
	return transform(decrypt, cipher, subject), nil
}

func encrypt(b byte) byte {
	return b - 128
}

func decrypt(b byte) byte {
	return 128 - b
}

func transform(mode shift, cipher []byte, subject []byte) []byte {
	for i, ix := 0, 0; i < len(subject); i, ix = i+1, ix+1 {
		if ix == len(cipher) {
			ix = 0
		}
		subject[i] = shiftOne(subject[i], cipher[ix], mode)
	}
	return subject
}

func shiftOne(subject, b byte, direction shift) byte {
	res := subject + direction(b)
	if res > 255 {
		return byte(int(res) - 256)
	} else if res < 0 {
		return byte(int(res) + 256)
	}
	return res
}

func HashSeedFromString(source string) int64 {
	return int64(xxhash.Sum64String(source))
}
