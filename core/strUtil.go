package core

import (
	"math/rand"
	"time"
)

// UUID 生成随机字符串
func UUID(size int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	bytes := make([]byte, size)
	for i := 0; i < size/2; i++ {
		b := r.Intn(25) + 65
		bytes[i] = byte(b)
	}
	for i := size / 2; i < size; i++ {
		b := r.Intn(25) + 97
		bytes[i] = byte(b)
	}
	rand.Shuffle(size, func(i, j int) {
		t := bytes[i]
		bytes[i] = bytes[j]
		bytes[j] = t
	})
	return string(bytes)
}
