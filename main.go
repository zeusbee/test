package main

import (
	"fmt"
	"hash/crc32"
	"time"
)

func main() {
	t := time.Unix(0, 0)
	fmt.Println(t)
	fmt.Sprintf("111111")
}

func hash(account string) int64 {
	hash := int64(crc32.ChecksumIEEE([]byte(account)))
	if hash >= 0 {
		return hash
	}
	if -hash >= 0 {
		return -hash
	}
	return 0
}
