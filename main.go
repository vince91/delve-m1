package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/core/types" // no issue without this import
)

func main() {
	l := types.Log{}

	str := "hello world"

	fmt.Println(str)
	_ = l
}
