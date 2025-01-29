package hermetik

import "hash/fnv"

func fnvSeedFromString(source string) (int64, error) {
	h := fnv.New64a()
	_, err := h.Write([]byte(source))
	if err != nil {
		return 0, err
	}
	return int64(h.Sum64()), nil
}
