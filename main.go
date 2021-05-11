package main

import (
	"fmt"
	"github.com/CMedrado/DesafioStone/domain"
)

func main() {

	x, err := domain.CreatedAccount("Rafael", "081.313.910-43", "lucas")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(x)
	}
	y, err := domain.CreatedAccount("man", "398.176200-26", "jax")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(y)
	}
	l, err := domain.CreatedAccount("nil", "384.531.620-93", "man")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(l)
	}
	f, err := domain.GetBalance("384.531.620-93")
	fmt.Println(domain.GetAccounts())
	fmt.Println(f)
}
