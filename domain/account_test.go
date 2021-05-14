package domain

import (
	"github.com/CMedrado/DesafioStone/store"
	"testing"
)

type CreateAccountTestInput struct {
	ID        int
	Name      string
	CPF       string
	Secret    string
	Balance   uint
	CreatedAt string
}

func TestCreateAccount(t *testing.T) {
	//prepare
	testTable := []struct { // tt := ....
		name    string                 //Nome do teste
		in      CreateAccountTestInput //Entrada da Função
		wantErr bool                   //Pra dizer se espera ou não um err
		want    uint
	}{
		{
			name: "should successfully create an account with formatted CPF",
			in: CreateAccountTestInput{
				Name:    "Rafael",
				CPF:     "081.313.910-43",
				Secret:  "lucas",
				Balance: 50000,
			},
			wantErr: false,
		},
		{
			name: "should successfully create an account with unformatted CPF",
			in: CreateAccountTestInput{
				Name:    "Lucas",
				CPF:     "38453162093",
				Secret:  "teo90",
				Balance: 60000,
			},
			wantErr: false,
		},
		{
			name: "should unsuccessfully create an account when CPF is invalid",
			in: CreateAccountTestInput{
				Name:    "Marcos",
				CPF:     "398.176200-26",
				Secret:  "marcos35",
				Balance: 7000,
			},
			wantErr: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			accountStorage := store.NewStoredAccount()
			usecase := AccountUsecase{
				Store: accountStorage,
			}

			//test
			gotID, gotErr := usecase.CreateAccount(testCase.in.Name, testCase.in.CPF, testCase.in.Secret, testCase.in.Balance)

			//assert
			if !testCase.wantErr && gotErr != nil { // O teste falhará pois não queremos erro e obtivemos um
				t.Errorf("unexpected error, wantErr=%v; gotErr=%s", testCase.wantErr, gotErr)
			}

			if testCase.wantErr && gotErr == nil { // O teste falhará pois queremos erro e não obtivemos um
				t.Error("wanted err but got nil")
			}

			if gotID == 0 && !testCase.wantErr && gotErr != nil {
				t.Errorf("expected an ID but got %d", gotID)
			}
		})
	}
}

func TestGetBalance(t *testing.T) {
	tt := []struct {
		name    string
		in      string
		wantErr bool
		want    uint
	}{
		{
			name:    "should successfully get balance with formatted CPF",
			in:      "081.313.910-43",
			wantErr: false,
			want:    5000,
		},
		{
			name:    "should successfully get balance with unformatted CPF",
			in:      "384.531.620-93",
			wantErr: false,
			want:    6000,
		},
		{
			name:    "should unsuccessfully get balance when CPF is invalid",
			in:      "398.176200-26",
			wantErr: true,
		},
		{
			name:    "should unsuccessfully get balance when dont exist account",
			in:      "062.136.280-37",
			wantErr: true,
		},
	}

	for _, testCase := range tt {
		t.Run(testCase.name, func(t *testing.T) {
			listAccount := store.Account{982, "Lucas", "08131391043", "lixo", 5000, "06/01/2020"}
			listAccounts := store.Account{981, "Rafael", "38453162093", "call", 6000, "06/01/2020"}

			accountStorage := store.NewStoredAccount()
			usecase := AccountUsecase{
				Store: accountStorage,
			}
			usecase.Store.TransferredAccount(listAccount)
			usecase.Store.TransferredAccount(listAccounts)
			//test
			gotBalance, gotErr := usecase.GetBalance(testCase.in)

			//assert
			if !testCase.wantErr && gotErr != nil { // O teste falhará pois não queremos erro e obtivemos um
				t.Errorf("unexpected error, wantErr=%v; gotErr=%s", testCase.wantErr, gotErr)
			}

			if testCase.wantErr && gotErr == nil { // O teste falhará pois queremos erro e não obtivemos um
				t.Error("wanted err but got nil")
			}

			if gotBalance != testCase.want {
				t.Errorf("expected an ID but got %d", gotBalance)
			}
		})
	}
}

//func TestGetAccounts (t *testing.T) {
//	tt := []struct{
//		name    string
//		wantErr bool
//		want	CreateAccountTestInput
//	}{
//		{
//			name:    "should successfully get accounts",
//			wantErr: false,
//			want:    CreateAccountTestInput{982, "Lucas", "08131391043", Hash("lixo"), 5000, "06/01/2020"},
//		},
//		{
//			name:    "should successfully get balance with unformatted CPF",
//			wantErr: false,
//			want:    CreateAccountTestInput{981,"Rafael","38453162093",Hash("call"),6000, "06/01/2020"},
//		},
//	}
//}
