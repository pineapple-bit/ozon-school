package main

import (
	"fmt"
	"math/big"
)

func main() {
	var item0, item1 string
	fmt.Scan(&item0, &item1)
	bignum0, _ := new(big.Int).SetString(item0, 0)
	bignum1, _ := new(big.Int).SetString(item1, 0)
	result := new(big.Int).Add(bignum0, bignum1)
	fmt.Println(result.String())
}
