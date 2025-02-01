package hermetik

import (
	"fmt"
	"math/rand"
)

func New(secret string, length int) (*Service, error) {
	seed, err := FNVSeedFromString(secret)
	if err != nil {
		return nil, fmt.Errorf("could not create new shift: %v", err)
	}

	return &Service{
		key: buildShiftKey(rand.New(rand.NewSource(seed)), length),
	}, nil
}

type Service struct {
	key  []uint8
	seed int64
}

func (s *Service) Encrypt(subject []byte) []byte {
	for i := range subject {
		subject[i] = s.transform(subject[i], s.key[i], s.directionEncrypt)
	}
	return subject
}

func (s *Service) Decrypt(subject []byte) []byte {
	for i := range subject {
		subject[i] = s.transform(subject[i], s.key[i], s.directionDecrypt)
	}
	return subject
}

func (s *Service) Key() []byte {
	return s.key
}

func (s *Service) transform(subject, inc byte, direction func(byte) byte) byte {
	res := subject + direction(inc)
	if res > 255 {
		return byte(int(res) - 256)
	} else if res < 0 {
		return byte(int(res) + 256)
	}
	return res
}

func (s *Service) directionEncrypt(inc byte) byte {
	return inc - 128
}

func (s *Service) directionDecrypt(inc byte) byte {
	return 128 - inc
}

func buildShiftKey(generator *rand.Rand, length int) []uint8 {
	xs := make([]uint8, length)
	for i := 0; i < length; i++ {
		xs[i] = uint8(generator.Intn(255))
	}
	return xs
}
