package hermetik

import (
	"github.com/cespare/xxhash"
	"hash/fnv"
)

func FNVSeedFromString(source string) (int64, error) {
	h := fnv.New64a()
	_, err := h.Write([]byte(source))
	if err != nil {
		return 0, err
	}
	return int64(h.Sum64()), nil
}

func HashSeedFromString(source string) int64 {
	return int64(xxhash.Sum64String(source))
}
