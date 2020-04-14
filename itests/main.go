package main

import (
	"github.com/riyadennis/aes-encryption/ex"
)

func main() {
	cl := ex.NewClient()
	StoreRetrieveTest(cl)
}
