package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
)

func main() {
	modePtr := flag.String("mode", "sha256", "can be sha256, ")
	flag.Parse()

	var value string
	fmt.Scanf("%s", &value)

	switch *modePtr {
	case "sha256":
		fmt.Printf("value %s has 256 hash %X", value, sha256.Sum256([]byte(value)))
	case "sha384":
		fmt.Printf("value %s has 384 hash %X", value, sha512.Sum384([]byte(value)))
	}

}
