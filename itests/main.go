package main

import (
	"github.com/riyadennis/aes-encryption/ex"
)

func main() {
	cl := ex.NewClient("0.0.0.0:5052")
	StoreTest(cl)
}
