package blueSkyThinking

import "hash/crc32"

// 累加hash

func addHashFun(seed uint, value string, max uint64) uint64 {
	hash := uint64(seed)
	for i := 0; i < len(value); i++ {
		hash += 9527 + uint64(value[i])
	}
	return hash % max
}

// 旋转hash

func rotatingHash(key []byte, r uint32, c uint32) uint32 {
	var hash uint32 = 0

	for i := 0; i < len(key); i++ {
		hash = (hash << r) + uint32(key[i])
		hash += c
	}

	return hash
}

// One-at-a-Time Hash

func oneAtATimeHash(key []byte) uint32 {
	var hash uint32 = 0

	for i := 0; i < len(key); i++ {
		hash += uint32(key[i])
		hash += hash << 10
		hash ^= hash >> 6
	}

	hash += hash << 3
	hash ^= hash >> 11
	hash += hash << 15

	return hash
}

// 字节hash

func bernsteinHash(key []byte) uint32 {
	var hash uint32 = 5381

	for i := 0; i < len(key); i++ {
		hash = ((hash << 5) + hash) + uint32(key[i])
	}

	return hash
}

// crc

func crc(str string) uint32 {
	return crc32.ChecksumIEEE([]byte(str))
}
