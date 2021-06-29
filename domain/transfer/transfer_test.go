package transfer

import (
	"github.com/CMedrado/DesafioStone/domain"
	account2 "github.com/CMedrado/DesafioStone/domain/account"
	"github.com/CMedrado/DesafioStone/domain/login"
	store_account "github.com/CMedrado/DesafioStone/storage/file/account"
	store_token "github.com/CMedrado/DesafioStone/storage/file/token"
	store_transfer "github.com/CMedrado/DesafioStone/storage/file/transfer"
	"github.com/google/uuid"
	"io"
	"io/ioutil"
	"os"
	"testing"
)

type CreateTransferInput struct {
	AccountOriginID      int
	Token                string
	AccountDestinationID string
	Amount               int
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

func createTemporaryFileTransfer(t *testing.T, Transfers string) (io.ReadWriteSeeker, func()) {
	filetmp, err := ioutil.TempFile("", "dbtransfer")

	if err != nil {
		t.Fatalf("it is not possible to write the temporary file %v", err)
	}

	filetmp.Write([]byte(Transfers))

	removeArquivo := func() {
		filetmp.Close()
		os.Remove(filetmp.Name())
	}

	return filetmp, removeArquivo
}

func TestMakeTransfers(t *testing.T) {
	tt := []struct {
		name                    string
		in                      CreateTransferInput
		wantErr                 bool
		expectedUpdateCallCount int
	}{
		{
			name: "should successfully transfer amount",
			in: CreateTransferInput{
				Token:                "MjkvMDYvMjAyMSAxMjo0NzowNzpjNTQyNDQ0MC00NzM3LTRlMDMtODZkMi0zYWRhYzkwZGRkMjA=",
				AccountDestinationID: "75432539-c5ba-46d3-9690-44985b516da7",
				Amount:               300,
			},
			wantErr:                 false,
			expectedUpdateCallCount: 1,
		},
		{
			name: "should unsuccessfully transfer amount when there is wrong token",
			in: CreateTransferInput{
				Token:                "MjkvMDYvMjAyMSAxMzoxNjo1NTo3NTQzMjUzOS1jNWJhLTQ2ZDMtOTY5MC00NDk4NWI1MTZkYTc=",
				AccountDestinationID: "75432539-c5ba-46d3-9690-44985b516da7",
				Amount:               300,
			},
			wantErr:                 true,
			expectedUpdateCallCount: 0,
		},
		{
			name: "should unsuccessfully transfer amount when there is no account destination id",
			in: CreateTransferInput{
				Token:                "MjkvMDYvMjAyMSAxMjo0NzowNzpjNTQyNDQ0MC00NzM3LTRlMDMtODZkMi0zYWRhYzkwZGRkMjA=",
				AccountDestinationID: "c5424440-4737-4e03-86d2-3adac90ddd20",
				Amount:               300,
			},
			wantErr:                 true,
			expectedUpdateCallCount: 0,
		},
		{
			name: "should unsuccessfully transfer amount when the amount is too slow",
			in: CreateTransferInput{
				Token:                "MjkvMDYvMjAyMSAxMjo0NzowNzpjNTQyNDQ0MC00NzM3LTRlMDMtODZkMi0zYWRhYzkwZGRkMjA=",
				AccountDestinationID: "75432539-c5ba-46d3-9690-44985b516da7",
				Amount:               0,
			},
			wantErr:                 true,
			expectedUpdateCallCount: 0,
		},
		{
			name: "should unsuccessfully transfer amount when the amount is greater than the balance",
			in: CreateTransferInput{
				Token:                "MjkvMDYvMjAyMSAxMjo0NzowNzpjNTQyNDQ0MC00NzM3LTRlMDMtODZkMi0zYWRhYzkwZGRkMjA=",
				AccountDestinationID: "75432539-c5ba-46d3-9690-44985b516da7",
				Amount:               52000,
			},
			wantErr:                 true,
			expectedUpdateCallCount: 0,
		},
	}

	for _, testCase := range tt {
		t.Run(testCase.name, func(t *testing.T) {
			dataBaseAccount, clenDataBaseAccount := createTemporaryFileAccount(t, `[{"id":"c5424440-4737-4e03-86d2-3adac90ddd20","name":"Lucas","cpf":"38453162093","secret":"7e65a9b554bbc9817aa049ce38c84a72","balance":3000,"created_at":"29/06/2021 12:45:35"},{"id":"75432539-c5ba-46d3-9690-44985b516da7","name":"Rafael","cpf":"08131391043","secret":"3467e121a1a109628e0a5b0cebba361b","balance":5000,"created_at":"29/06/2021 12:46:28"}]`)
			defer clenDataBaseAccount()
			dataBaseToken, clenDataBaseToken := createTemporaryFileToken(t, `[{"id":"c5424440-4737-4e03-86d2-3adac90ddd20","token":"MjkvMDYvMjAyMSAxMjo0NzowNzpjNTQyNDQ0MC00NzM3LTRlMDMtODZkMi0zYWRhYzkwZGRkMjA="}]`)
			defer clenDataBaseToken()
			dataBaseTransfer, clenDataBaseTransfer := createTemporaryFileTransfer(t, ``)
			defer clenDataBaseTransfer()
			accountAccount := store_account.NewStoredAccount(dataBaseAccount)
			accountUsecase := &AccountUsecaseMock{AccountList: accountAccount}
			accountToken := store_token.NewStoredToked(dataBaseToken)
			tokenUseCase := &TokenUseCaseMock{accountToken}
			storagedTransfer := store_transfer.NewStoredTransfer(dataBaseTransfer)
			usecase := UseCase{
				AccountUseCase: accountUsecase,
				TokenUseCase:   tokenUseCase,
				StoredTransfer: storagedTransfer,
			}

			gotErr, gotTransfer := usecase.CreateTransfers(testCase.in.Token, testCase.in.AccountDestinationID, testCase.in.Amount)

			if !testCase.wantErr && gotErr != nil {
				t.Errorf("unexpected error, wantErr=%v; gotErr=%s", testCase.wantErr, gotErr)
			}

			if testCase.wantErr && gotErr == nil {
				t.Error("wanted err but got nil")
			}

			if (gotTransfer == uuid.UUID{}) && !testCase.wantErr && gotErr != nil {
				t.Errorf("expected an Token but got %d", gotTransfer)
			}

			if accountUsecase.UpdateCallCount != testCase.expectedUpdateCallCount {

			}
		})
	}
}

func TestMakeGetTransfers(t *testing.T) {

	var tt = []struct {
		name    string
		in      CreateTransferInput
		wantErr bool
		want    []store_transfer.Transfer
	}{
		{
			name: "should successfully get transfers",
			in: CreateTransferInput{
				Token: "MjkvMDYvMjAyMSAxMjo0NzowNzpjNTQyNDQ0MC00NzM3LTRlMDMtODZkMi0zYWRhYzkwZGRkMjA=",
			},
			wantErr: false,
		},
		{
			name: "should unsuccessfully get transfer when there is wrong token",
			in: CreateTransferInput{
				Token: "MjkvMDYvMjAyMSAxMzoxNjo1NTo3NTQzMjUzOS1jNWJhLTQ2ZDMtOTY5MC00NDk4NWI1MTZkYTc=",
			},
			wantErr: true,
		},
	}
	for _, testCase := range tt {
		t.Run(testCase.name, func(t *testing.T) {
			dataBaseAccount, clenDataBaseAccount := createTemporaryFileAccount(t, `[{"id":"c5424440-4737-4e03-86d2-3adac90ddd20","name":"Lucas","cpf":"38453162093","secret":"7e65a9b554bbc9817aa049ce38c84a72","balance":3000,"created_at":"29/06/2021 12:45:35"},{"id":"75432539-c5ba-46d3-9690-44985b516da7","name":"Rafael","cpf":"08131391043","secret":"3467e121a1a109628e0a5b0cebba361b","balance":5000,"created_at":"29/06/2021 12:46:28"}]`)
			defer clenDataBaseAccount()
			dataBaseToken, clenDataBaseToken := createTemporaryFileToken(t, `[{"id":"c5424440-4737-4e03-86d2-3adac90ddd20","token":"MjkvMDYvMjAyMSAxMjo0NzowNzpjNTQyNDQ0MC00NzM3LTRlMDMtODZkMi0zYWRhYzkwZGRkMjA="}]`)
			defer clenDataBaseToken()
			dataBaseTransfer, clenDataBaseTransfer := createTemporaryFileTransfer(t, `[{"id":"47399f23-2093-4dde-b32f-990cac27630e","account_origin_id":"c5424440-4737-4e03-86d2-3adac90ddd20","account_destination_id":"75432539-c5ba-46d3-9690-44985b516da7","amount":150,"created_at":"29/06/2021 12:48:06"}]`)
			defer clenDataBaseTransfer()
			accountAccount := store_account.NewStoredAccount(dataBaseAccount)
			accountUsecase := &AccountUsecaseMock{AccountList: accountAccount}
			accountToken := store_token.NewStoredToked(dataBaseToken)
			tokenUseCase := &TokenUseCaseMock{accountToken}
			storagedTransfer := store_transfer.NewStoredTransfer(dataBaseTransfer)
			usecase := UseCase{
				AccountUseCase: accountUsecase,
				TokenUseCase:   tokenUseCase,
				StoredTransfer: storagedTransfer,
			}

			gotTransfer, gotErr := usecase.GetTransfers(testCase.in.Token)

			if !testCase.wantErr && gotErr != nil {
				t.Errorf("unexpected error, wantErr=%v; gotErr=%s", testCase.wantErr, gotErr)
			}

			if testCase.wantErr && gotErr == nil {
				t.Error("wanted err but got nil")
			}

			if gotTransfer == nil && !testCase.wantErr && gotErr != nil {
				t.Errorf("expected an Token but got %v", gotTransfer)
			}
		})
	}
}

type AccountUsecaseMock struct {
	AccountList *store_account.StoredAccount

	UpdateCallCount int
}

func (uc AccountUsecaseMock) ReturnCPF(_ string) int {
	return 0
}

func (uc AccountUsecaseMock) CreateAccount(_ string, _ string, _ string, _ int) (uuid.UUID, error) {
	return uuid.UUID{}, nil
}

func (uc AccountUsecaseMock) GetBalance(_ string) (int, error) {
	return 0, nil
}

func (uc AccountUsecaseMock) GetAccounts() []domain.Account {
	return nil
}

func (uc AccountUsecaseMock) SearchAccount(id uuid.UUID) domain.Account {
	account := domain.Account{}

	for _, a := range uc.AccountList.ReturnAccounts() {
		if a.ID == id {
			account = account2.ChangeAccountStorage(a)
		}
	}

	return account
}

func (uc *AccountUsecaseMock) UpdateBalance(_ domain.Account, _ domain.Account) {
	uc.UpdateCallCount++
}

func (uc AccountUsecaseMock) GetAccountCPF(cpf string) domain.Account {
	account := domain.Account{}
	for _, a := range uc.AccountList.ReturnAccounts() {
		if a.CPF == cpf {
			account = account2.ChangeAccountStorage(a)
		}
	}

	return account
}

type TokenUseCaseMock struct {
	TokenList *store_token.StoredToken
}

func (uc TokenUseCaseMock) AuthenticatedLogin(_, _ string) (error, string) {
	return nil, ""
}

func (uc TokenUseCaseMock) ReturnToken(_ int) string {
	return ""
}

func (uc TokenUseCaseMock) GetTokenID(id uuid.UUID) domain.Token {
	token := domain.Token{}

	for _, a := range uc.TokenList.ReturnTokens() {
		if a.ID == id {
			token = login.ChangeTokenStorage(a)
		}
	}

	return token
}
