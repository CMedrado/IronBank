package transfer

import (
	store_account "github.com/CMedrado/DesafioStone/store/account"
	store_token "github.com/CMedrado/DesafioStone/store/token"
	store_transfer "github.com/CMedrado/DesafioStone/store/transfer"
	"io"
	"io/ioutil"
	"os"
	"testing"
)

type CreateTransferInput struct {
	AccountOriginID      int
	Token                string
	AccountDestinationID int
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
				Token:                "MjEvMDYvMjAyMSAyMzo1OTowMDoxOTcyNzg4Nw==",
				AccountDestinationID: 98498081,
				Amount:               300,
			},
			wantErr:                 false,
			expectedUpdateCallCount: 1,
		},
		{
			name: "should unsuccessfully transfer amount when there is wrong token",
			in: CreateTransferInput{
				Token:                "MTgvMDYvMjAyMSAxNjozNDozMjo5ODQ5ODA4MQ==",
				AccountDestinationID: 98498081,
				Amount:               300,
			},
			wantErr:                 true,
			expectedUpdateCallCount: 0,
		},
		{
			name: "should unsuccessfully transfer amount when there is no account destination id",
			in: CreateTransferInput{
				Token:                "MjEvMDYvMjAyMSAyMzo1OTowMDoxOTcyNzg4Nw==",
				AccountDestinationID: 19727887,
				Amount:               300,
			},
			wantErr:                 true,
			expectedUpdateCallCount: 0,
		},
		{
			name: "should unsuccessfully transfer amount when the amount is too slow",
			in: CreateTransferInput{
				Token:                "MjEvMDYvMjAyMSAyMzo1OTowMDoxOTcyNzg4Nw==",
				AccountDestinationID: 98498081,
				Amount:               0,
			},
			wantErr:                 true,
			expectedUpdateCallCount: 0,
		},
		{
			name: "should unsuccessfully transfer amount when the amount is greater than the balance",
			in: CreateTransferInput{
				Token:                "MjEvMDYvMjAyMSAyMzo1OTowMDoxOTcyNzg4Nw==",
				AccountDestinationID: 98498081,
				Amount:               52000,
			},
			wantErr:                 true,
			expectedUpdateCallCount: 0,
		},
	}

	for _, testCase := range tt {
		t.Run(testCase.name, func(t *testing.T) {
			dataBaseAccount, clenDataBaseAccount := createTemporaryFileAccount(t, `[{"id":19727887,"name":"Rafael","cpf":"38453162093","secret":"53b9e9679a8ea25880376080b76f98ad","balance":6000,"created_at":"06/01/2020"},{"id":98498081,"name":"Lucas","cpf":"08131391043","secret":"c74af74c69d81831a5703aefe9cb4199","balance":5000,"created_at":"06/01/2020"}]`)
			defer clenDataBaseAccount()
			dataBaseToken, clenDataBaseToken := createTemporaryFileToken(t, `[{"id":19727887,"token":"MjEvMDYvMjAyMSAyMzo1OTowMDoxOTcyNzg4Nw=="}]`)
			defer clenDataBaseToken()
			dataBaseTransfer, clenDataBaseTransfer := createTemporaryFileTransfer(t, ``)
			defer clenDataBaseTransfer()
			accountAccount := store_account.NewStoredAccount(dataBaseAccount)
			accountUsecase := &AccountUsecaseMock{AccountList: accountAccount}
			accountToken := store_token.NewStoredToked(dataBaseToken)
			tokenUseCase := &TokenUseCaseMock{accountToken}
			storagedTransfer := store_transfer.NewStoredTransferAccountID(dataBaseTransfer)
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

			if gotTransfer == 0 && !testCase.wantErr && gotErr != nil {
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
				Token: "MjEvMDYvMjAyMSAyMzo1OTowMDoxOTcyNzg4Nw==",
			},
			wantErr: false,
			want:    []store_transfer.Transfer{{ID: 47278511, AccountOriginID: 98498081, AccountDestinationID: 19727887, Amount: 500, CreatedAt: "13/05/2021 09:09:16"}, {ID: 6410694, AccountOriginID: 98498081, AccountDestinationID: 19727887, Amount: 200, CreatedAt: "13/05/2021 09:09:16"}},
		},
		{
			name: "should unsuccessfully get transfer when there is wrong token",
			in: CreateTransferInput{
				Token: "MTgvMDYvMjAyMSAxNjozNDozMjo5ODQ5ODA4MQ==",
			},
			wantErr: true,
		},
	}
	for _, testCase := range tt {
		t.Run(testCase.name, func(t *testing.T) {
			dataBaseAccount, clenDataBaseAccount := createTemporaryFileAccount(t, `[{"id":19727887,"name":"Rafael","cpf":"38453162093","secret":"53b9e9679a8ea25880376080b76f98ad","balance":6000,"created_at":"06/01/2020"},{"id":98498081,"name":"Lucas","cpf":"08131391043","secret":"c74af74c69d81831a5703aefe9cb4199","balance":5000,"created_at":"06/01/2020"}]`)
			defer clenDataBaseAccount()
			dataBaseToken, clenDataBaseToken := createTemporaryFileToken(t, `[{"id":19727887,"token":"MjEvMDYvMjAyMSAyMzo1OTowMDoxOTcyNzg4Nw=="}]`)
			defer clenDataBaseToken()
			dataBaseTransfer, clenDataBaseTransfer := createTemporaryFileTransfer(t, `[[{"id":39984059,"account_origin_id":19727887,"account_destination_id":98498081,"amount":60,"created_at":"18/06/2021 17:26:26"},{"id":27131847,"account_origin_id":19727887,"account_destination_id":98498081,"amount":50,"created_at":"18/06/2021 17:26:26"},{"id":74941318,"account_origin_id":11902081,"account_destination_id":98498081,"amount":900,"created_at":"18/06/2021 17:26:26"},{"id":11902081,"account_origin_id":98498081,"account_destination_id":19727887,"amount":90,"created_at":"18/06/2021 17:26:26"}]}`)
			defer clenDataBaseTransfer()
			accountAccount := store_account.NewStoredAccount(dataBaseAccount)
			accountUsecase := &AccountUsecaseMock{AccountList: accountAccount}
			accountToken := store_token.NewStoredToked(dataBaseToken)
			tokenUseCase := &TokenUseCaseMock{accountToken}
			storagedTransfer := store_transfer.NewStoredTransferAccountID(dataBaseTransfer)
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

func (uc AccountUsecaseMock) CreateAccount(_ string, _ string, _ string, _ int) (int, error) {
	return 0, nil
}

func (uc AccountUsecaseMock) GetBalance(_ int) (int, error) {
	return 0, nil
}

func (uc AccountUsecaseMock) GetAccounts() []store_account.Account {
	return nil
}

func (uc AccountUsecaseMock) SearchAccount(id int) store_account.Account {
	account := store_account.Account{}

	for _, a := range uc.AccountList.GetAccounts() {
		if a.ID == id {
			account = a
		}
	}

	return account
}

func (uc *AccountUsecaseMock) UpdateBalance(_ store_account.Account, _ store_account.Account) {
	uc.UpdateCallCount++
}

func (uc AccountUsecaseMock) GetAccountCPF(cpf string) store_account.Account {
	account := store_account.Account{}
	for _, a := range uc.AccountList.GetAccounts() {
		if a.CPF == cpf {
			account = a
		}
	}

	return account
}

func (uc AccountUsecaseMock) GetAccount() []store_account.Account {
	return nil
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

func (uc TokenUseCaseMock) GetTokenID(id int) store_token.Token {
	token := store_token.Token{}

	for _, a := range uc.TokenList.GetTokens() {
		if a.ID == id {
			token = a
		}
	}

	return token
}
