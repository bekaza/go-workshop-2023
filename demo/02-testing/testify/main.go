package main

import (
	"example/testify/bank"
	"fmt"
)

func main() {
	account := bank.Banking{}
	fmt.Println(account.Balance().String())
}
