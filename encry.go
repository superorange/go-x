// hash_sha1.go
package main

import (
	"crypto/sha1"
	"fmt"
	"log"
)

func main() {
	hasher := sha1.New()
	hasher.Write([]byte("test"))
	//io.WriteString(hasher, "test")
	b := []byte{}
	fmt.Printf("Result: %x\n", hasher.Sum(b))
	fmt.Printf("Result: %d\n", hasher.Sum(b))
	//
	hasher.Reset()
	data := []byte("We shall overcome!")
	n, err := hasher.Write(data)
	if n != len(data) || err != nil {
		log.Printf("Hash write error: %v / %v", n, err)
	}
	checksum := hasher.Sum(b)
	fmt.Printf("Result: %x\n", checksum)
}
