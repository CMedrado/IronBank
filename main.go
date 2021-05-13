package main

import (
	"fmt"
	"github.com/CMedrado/DesafioStone/domain"
	"github.com/CMedrado/DesafioStone/store"
)

func main() {
	accountStorage := store.NewStoredAccount()
	accountUsecase := domain.AccountUsecase{Store: accountStorage}

	x, err := accountUsecase.CreateAccount("Rafael", "081.313.910-43", "lucas")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(x)
	}

	y, err2 := accountUsecase.CreateAccount("man", "398.176200-26", "jax")
	if err != nil {
		fmt.Println(err2)
	} else {
		fmt.Println(y)
	}
	l, err := domain.CreatedAccount("nil", "384.531.620-93", "man")
	if err != nil {
		fmt.Println(err3)
	} else {
		fmt.Println(l)
	}

	balance, err := accountUsecase.GetBalance("384.531.620-93")

	fmt.Println(accountUsecase.GetAccounts())
	fmt.Println(balance)

	//x, err := domain.CreatedAccount("Rafael", "081.313.910-43", "lucas")
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Println(x)
	//}
	//y, err := domain.CreatedAccount("man", "398.176200-26", "jax")
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Println(y)
	//}
	//l, err := domain.CreatedAccount("nil", "384.531.620-93", "man")
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Println(l)
	//}
	//f, err := domain.GetBalance("384.531.620-93")
	//fmt.Println(domain.GetAccounts())
	//fmt.Println(f)
}
