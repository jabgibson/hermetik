package h7k

import "math/rand"

func BuildCipher(seed int64, size int) []byte {
	generator := rand.New(rand.NewSource(seed))
	xs := make([]uint8, size)
	for i := 0; i < size; i++ {
		xs[i] = uint8(generator.Intn(255))
	}
	return xs
}
