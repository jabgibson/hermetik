package hermetik

func Encrypt(cipher []byte, subject []byte) ([]byte, error) {
	return shift(encrypt, cipher, subject), nil
}

func Decrypt(cipher []byte, subject []byte) ([]byte, error) {
	return shift(decrypt, cipher, subject), nil
}

func encrypt(b byte) byte {
	return b - 128
}

func decrypt(b byte) byte {
	return 128 - b
}

func shift(mode Shift, cipher []byte, subject []byte) []byte {
	for i, ix := 0, 0; i < len(subject); i, ix = i+1, ix+1 {
		if ix == len(cipher) {
			ix = 0
		}
		subject[i] = shiftBytes(subject[i], cipher[ix], mode)
	}
	return subject
}

func shiftBytes(subject, b byte, direction Shift) byte {
	res := subject + direction(b)
	if res > 255 {
		return byte(int(res) - 256)
	} else if res < 0 {
		return byte(int(res) + 256)
	}
	return res
}
