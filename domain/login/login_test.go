package login

import (
	"github.com/CMedrado/DesafioStone/domain"
	store_account "github.com/CMedrado/DesafioStone/store/account"
	store_login "github.com/CMedrado/DesafioStone/store/login"
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
			accountStorage := make(map[string]store_account.Account)
			listAccount := store_account.Account{ID: 982, Name: "Lucas", CPF: "08131391043", Secret: domain.CreateHash("lixo"), Balance: 5000, CreatedAt: "06/01/2020"}
			listAccounts := store_account.Account{ID: 981, Name: "Rafael", CPF: "38453162093", Secret: domain.CreateHash("call"), Balance: 6000, CreatedAt: "06/01/2020"}

			accountStorage[listAccount.CPF] = listAccount
			accountStorage[listAccounts.CPF] = listAccounts
			accountUsecase := &AccountUseCaseMock{AccountList: accountStorage}
			accountToken := store_login.NewStoredToked()
			usecase := UseCase{
				AccountUseCase: accountUsecase,
				StoredToken:    accountToken,
			}

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

type AccountUseCaseMock struct {
	AccountList map[string]store_account.Account
}

func (uc AccountUseCaseMock) ReturnCPF(_ string) int {
	return 0
}

func (uc AccountUseCaseMock) CreateAccount(_ string, _ string, _ string, _ int) (int, error) {
	return 0, nil
}

func (uc AccountUseCaseMock) GetBalance(_ int) (int, error) {
	return 0, nil
}

func (uc AccountUseCaseMock) GetAccounts() []store_account.Account {
	return nil
}

func (uc AccountUseCaseMock) SearchAccount(id int) store_account.Account {
	return store_account.Account{}
}

func (uc *AccountUseCaseMock) UpdateBalance(_ store_account.Account, _ store_account.Account) {
}

func (uc AccountUseCaseMock) GetAccountCPF(cpf string) store_account.Account {
	return uc.AccountList[cpf]
}

func (uc AccountUseCaseMock) GetAccount() map[string]store_account.Account {
	return nil
}
