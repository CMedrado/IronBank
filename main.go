package main

import (
	"github.com/CMedrado/DesafioStone/domain"
	https "github.com/CMedrado/DesafioStone/https"
	"log"
	"net/http"
)

func main() {
	//accountStorage := store.NewStoredAccount()
	//accountUsecase := domain.AccountUsecase{Store: accountStorage}
	//
	//x, err := accountUsecase.CreateAccount("Rafael", "081.313.910-43", "lucas", 5000)
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Println(x)
	//}
	//
	//y, err2 := accountUsecase.CreateAccount("man", "398.176200-26", "jax", 4000)
	//if err != nil {
	//	fmt.Println(err2)
	//} else {
	//	fmt.Println(y)
	//}
	//l, err3 := accountUsecase.CreateAccount("nil", "384.531.620-93", "man", 3000)
	//if err != nil {
	//	fmt.Println(err3)
	//} else {
	//	fmt.Println(l)
	//}
	//
	//
	//
	//fmt.Println(accountUsecase.GetAccounts())
	//
	//_, tr := accountUsecase.AuthenticatedLogin("38453162093", "man")
	//_, vf := accountUsecase.AuthenticatedLogin("081.313.910-43", "lucas")
	//
	//fmt.Println(tr)
	//_, sl:= accountUsecase.MakeTransfers(19727887, tr, 98498081, 300)
	//fmt.Println(sl)
	//_, sa:= accountUsecase.MakeTransfers(19727887, tr, 98498081, 400)
	//fmt.Println(sa)
	//_, as:= accountUsecase.MakeTransfers(19727887, tr, 98498081, 100)
	//fmt.Println(as)
	//
	//tm, _ := accountUsecase.GetTransfers(98498081, vf)
	//ts, _ := accountUsecase.GetTransfers(19727887, tr)
	//fmt.Println(tm)
	//fmt.Println(ts)
	//balance, err := accountUsecase.GetBalance("384.531.620-93")
	//fmt.Println(balance)
	//fmt.Println(accountUsecase.GetAccounts())
	accounStorage := domain.AccountUsecase{}
	servidor := https.NewServerAccount(&accounStorage)

	if err := http.ListenAndServe(":5000", servidor); err != nil {
		log.Fatal("could not hear on port 5000 ")
	}
}
