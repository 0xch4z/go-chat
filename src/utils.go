package main

import (
	"crypto/rand"
	"fmt"
)

func randomKey() string {
	buf := make([]byte, 16)
	_, err := rand.Read(buf)

	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("%x", buf)
}
