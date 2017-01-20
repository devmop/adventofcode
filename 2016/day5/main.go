package main

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
)

func main() {
	key := "wtnhxymk"

	var index int
	var count uint
	for count < 8 {
		h := hash(key, index)
		println(hash)
		if h[0:5] == "00000" {
			count++
			print(h[5])
		}
		index++
	}
}

func hash(key string, index int) string {
	input := key + strconv.Itoa(index)
	hasher := md5.New()
	println(input)
	hasher.Write([]byte(input))
	return hex.EncodeToString(hasher.Sum(nil))
}
