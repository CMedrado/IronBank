package authentication

import (
	"github.com/CMedrado/DesafioStone/domain"
	account2 "github.com/CMedrado/DesafioStone/domain/account"
	store_account "github.com/CMedrado/DesafioStone/storage/file/account"
	store_token "github.com/CMedrado/DesafioStone/storage/file/token"
	"github.com/google/uuid"
	"io"
	"io/ioutil"
	"os"
	"testing"
)

type CreateLoginInput struct {
	CPF    string
	Secret string
}

func createTemporaryFileToken(t *testing.T, Tokens string) (io.ReadWriteSeeker, func()) {
	filetmp, err := ioutil.TempFile("", "dbtoken")

	if err != nil {
		t.Fatalf("it is not possible to write the temporary file %v", err)
	}

	filetmp.Write([]byte(Tokens))

	removeArquivo := func() {
		filetmp.Close()
		os.Remove(filetmp.Name())
	}

	return filetmp, removeArquivo
}

func createTemporaryFileAccount(t *testing.T, Accounts string) (io.ReadWriteSeeker, func()) {
	filetmp, err := ioutil.TempFile("", "dbaccount")

	if err != nil {
		t.Fatalf("it is not possible to write the temporary file %v", err)
	}

	filetmp.Write([]byte(Accounts))

	removeArquivo := func() {
		filetmp.Close()
		os.Remove(filetmp.Name())
	}

	return filetmp, removeArquivo
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
			dataBaseAccount, clenDataBaseAccount := createTemporaryFileAccount(t, `[{"id":981,"name":"Rafael","cpf":"38453162093","secret":"53b9e9679a8ea25880376080b76f98ad","balance":6000,"created_at":"06/01/2020"},{"id":982,"name":"Lucas","cpf":"08131391043","secret":"c74af74c69d81831a5703aefe9cb4199","balance":5000,"created_at":"06/01/2020"}]`)
			defer clenDataBaseAccount()
			dataBaseToken, clenDataBaseToken := createTemporaryFileToken(t, `[{0,0}]`)
			defer clenDataBaseToken()
			accountAccount := store_account.NewStoredAccount(dataBaseAccount)
			accountUsecase := &AccountUseCaseMock{accountAccount}
			accountToken := store_token.NewStoredToked(dataBaseToken)
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
	AccountList *store_account.StoredAccount
}

func (uc AccountUseCaseMock) CreateAccount(_ string, _ string, _ string, _ int) (uuid.UUID, error) {
	return uuid.UUID{}, nil
}

func (uc AccountUseCaseMock) GetBalance(_ string) (int, error) {
	return 0, nil
}

func (uc AccountUseCaseMock) GetAccounts() []domain.Account {
	return nil
}

func (uc AccountUseCaseMock) SearchAccount(_ uuid.UUID) domain.Account {
	return domain.Account{}
}

func (uc *AccountUseCaseMock) UpdateBalance(_ domain.Account, _ domain.Account) {
}

func (uc AccountUseCaseMock) GetAccountCPF(cpf string) domain.Account {
	account := domain.Account{}
	for _, a := range uc.AccountList.ReturnAccounts() {
		if a.CPF == cpf {
			account = account2.ChangeAccountStorage(a)
		}
	}

	return account
}
