package domain

import (
	"github.com/CMedrado/DesafioStone/store"
	"testing"
)

type CreateLoginInput struct {
	CPF    string
	Secret string
}

func TestAuthenticatedLogin(t *testing.T) {
	tt := []struct {
		name    string
		in      CreateLoginInput
		wantErr bool
	}{
		{
			name: "should successfully authenticated login with formatted CPF",
			in: CreateLoginInput{
				CPF:    "081.313.910-43",
				Secret: "lixo",
			},
			wantErr: false,
		},
		{
			name: "should successfully authenticated login with unformatted CPF",
			in: CreateLoginInput{
				CPF:    "38453162093",
				Secret: "call",
			},
			wantErr: false,
		},
		{
			name: "should unsuccessfully authenticated login when CPF is invalid",
			in: CreateLoginInput{
				CPF:    "384531620.93",
				Secret: "call",
			},
			wantErr: true,
		},
		{
			name: "should unsuccessfully authenticated login when cpf is not registered",
			in: CreateLoginInput{
				CPF:    "38453162793",
				Secret: "call",
			},
			wantErr: true,
		},
		{
			name: "should unsuccessfully authenticated login when secret is not correct",
			in: CreateLoginInput{
				CPF:    "081.313.910-43",
				Secret: "call",
			},
			wantErr: true,
		},
	}

	for _, testCase := range tt {
		t.Run(testCase.name, func(t *testing.T) {
			listAccount := store.Account{982, "Lucas", "08131391043", Hash("lixo"), 5000, "06/01/2020"}
			listAccounts := store.Account{981, "Rafael", "38453162093", Hash("call"), 6000, "06/01/2020"}

			accountStorage := store.NewStoredAccount()
			accountToken := store.NewStoredToked()
			usecase := AccountUsecase{
				Store: accountStorage,
				Token: accountToken,
			}

			usecase.Store.TransferredAccount(listAccount)
			usecase.Store.TransferredAccount(listAccounts)

			gotErr, gotToken := usecase.AuthenticatedLogin(testCase.in.CPF, testCase.in.Secret)

			if !testCase.wantErr && gotErr != nil {
				t.Errorf("unexpected error, wantErr=%v; gotErr=%s", testCase.wantErr, gotErr)
			}

			if testCase.wantErr && gotErr == nil {
				t.Error("wanted err but got nil")
			}

			if gotToken == "" && !testCase.wantErr && gotErr != nil {
				t.Errorf("expected an Token but got %s", gotToken)
			}
		})
	}
}
