package account

import (
	store_account "github.com/CMedrado/DesafioStone/store/account"
	"testing"
)

type CreateAccountTestInput struct {
	ID        int
	Name      string
	CPF       string
	Secret    string
	Balance   int
	CreatedAt string
}

func TestCreateAccount(t *testing.T) {
	//prepare
	testTable := []struct { // tt := ....
		name    string                 //Nome do teste
		in      CreateAccountTestInput //Entrada da Função
		wantErr bool                   //Pra dizer se espera ou não um err
		want    int
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
			name: "should successfully create an account with unformulated CPF",
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
			accountStorage := store_account.NewStoredAccount()
			usecase := UseCase{
				StoredAccount: accountStorage,
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
		in      int
		wantErr bool
		want    int
	}{
		{
			name:    "should successfully get balance with formatted CPF",
			in:      982,
			wantErr: false,
			want:    5000,
		},
		{
			name:    "should successfully get balance with unformulated CPF",
			in:      981,
			wantErr: false,
			want:    6000,
		},
		{
			name:    "should unsuccessfully get balance when CPF is invalid",
			in:      398 - 6,
			wantErr: true,
		},
		{
			name:    "should unsuccessfully get balance when dont exist account",
			in:      06237,
			wantErr: true,
		},
	}

	for _, testCase := range tt {
		t.Run(testCase.name, func(t *testing.T) {
			listAccount := store_account.Account{ID: 982, Name: "Lucas", CPF: "08131391043", Secret: "lixo", Balance: 5000, CreatedAt: "06/01/2020"}
			listAccounts := store_account.Account{ID: 981, Name: "Rafael", CPF: "38453162093", Secret: "call", Balance: 6000, CreatedAt: "06/01/2020"}

			accountStorage := store_account.NewStoredAccount()
			usecase := UseCase{
				StoredAccount: accountStorage,
			}
			usecase.StoredAccount.CreateAccount(listAccount)
			usecase.StoredAccount.CreateAccount(listAccounts)
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
