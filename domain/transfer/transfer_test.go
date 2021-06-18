package transfer

import (
	"encoding/base64"
	"github.com/CMedrado/DesafioStone/domain"
	"github.com/CMedrado/DesafioStone/store"
	"testing"
)

type CreateTransferInput struct {
	AccountOriginID      int
	Token                string
	AccountDestinationID int
	Amount               int
}

func TestMakeTransfers(t *testing.T) {
	msg := base64.StdEncoding.EncodeToString([]byte("10/02/2009 02:02:00 : 0"))
	msgs := base64.StdEncoding.EncodeToString([]byte("10/03/2009 02:02:00 : 1"))
	tt := []struct {
		name                    string
		in                      CreateTransferInput
		wantErr                 bool
		expectedUpdateCallCount int
	}{
		{
			name: "should successfully transfer amount",
			in: CreateTransferInput{
				Token:                msg,
				AccountDestinationID: 1,
				Amount:               300,
			},
			wantErr:                 false,
			expectedUpdateCallCount: 1,
		},
		{
			name: "should unsuccessfully transfer amount when there is wrong token",
			in: CreateTransferInput{
				Token:                msgs,
				AccountDestinationID: 1,
				Amount:               300,
			},
			wantErr:                 true,
			expectedUpdateCallCount: 0,
		},
		{
			name: "should unsuccessfully transfer amount when there is no account destination id",
			in: CreateTransferInput{
				Token:                msg,
				AccountDestinationID: 2,
				Amount:               300,
			},
			wantErr:                 true,
			expectedUpdateCallCount: 0,
		},
		{
			name: "should unsuccessfully transfer amount when the amount is too slow",
			in: CreateTransferInput{
				Token:                msg,
				AccountDestinationID: 1,
				Amount:               0,
			},
			wantErr:                 true,
			expectedUpdateCallCount: 0,
		},
		{
			name: "should unsuccessfully transfer amount when the amount is greater than the balance",
			in: CreateTransferInput{
				Token:                msg,
				AccountDestinationID: 1,
				Amount:               5200,
			},
			wantErr:                 true,
			expectedUpdateCallCount: 0,
		},
	}

	for _, testCase := range tt {
		t.Run(testCase.name, func(t *testing.T) {
			accountStorage := make(map[string]store.Account)
			accountToken := make(map[int]store.Token)
			originAccount := store.Account{ID: 0, Name: "Lucas", CPF: "08131391043", Secret: domain.CreateHash("lixo"), Balance: 5000, CreatedAt: "06/01/2020"}
			destinationAccount := store.Account{ID: 1, Name: "Rafael", CPF: "38453162093", Secret: domain.CreateHash("call"), Balance: 6000, CreatedAt: "06/01/2020"}
			originToken := store.Token{Token: msg}
			accountStorage[originAccount.CPF] = originAccount
			accountStorage[destinationAccount.CPF] = destinationAccount
			accountToken[originAccount.ID] = originToken

			tokenUseCase := &TokenUseCaseMock{TokenList: accountToken}
			accountUsecase := &AccountUsecaseMock{AccountList: accountStorage}
			storagedTransfer := store.NewStoredTransferAccountID()
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
	msg := base64.StdEncoding.EncodeToString([]byte("10/02/2009 02:02:00:0"))
	msgs := base64.StdEncoding.EncodeToString([]byte("10/03/2009 02:02:00:1"))

	var tt = []struct {
		name    string
		in      CreateTransferInput
		wantErr bool
		want    []store.Transfer
	}{
		{
			name: "should successfully get transfers",
			in: CreateTransferInput{
				Token: msg,
			},
			wantErr: false,
			want:    []store.Transfer{{ID: 47278511, AccountOriginID: 98498081, AccountDestinationID: 19727887, Amount: 500, CreatedAt: "13/05/2021 09:09:16"}, {ID: 6410694, AccountOriginID: 98498081, AccountDestinationID: 19727887, Amount: 200, CreatedAt: "13/05/2021 09:09:16"}},
		},
		{
			name: "should unsuccessfully get transfer when there is wrong token",
			in: CreateTransferInput{
				Token: msgs,
			},
			wantErr: true,
		},
	}
	for _, testCase := range tt {
		t.Run(testCase.name, func(t *testing.T) {
			accountStorage := make(map[string]store.Account)
			accountToken := make(map[int]store.Token)
			originAccount := store.Account{ID: 0, Name: "Lucas", CPF: "08131391043", Secret: domain.CreateHash("lixo"), Balance: 5000, CreatedAt: "06/01/2020"}
			destinationAccount := store.Account{ID: 1, Name: "Rafael", CPF: "38453162093", Secret: domain.CreateHash("call"), Balance: 6000, CreatedAt: "06/01/2020"}
			listTransfer := store.Transfer{ID: 47278511, AccountOriginID: 98498081, AccountDestinationID: 19727887, Amount: 500, CreatedAt: "13/05/2021 09:09:16"}
			listTransfers := store.Transfer{ID: 6410694, AccountOriginID: 98498081, AccountDestinationID: 19727887, Amount: 200, CreatedAt: "13/05/2021 09:09:16"}
			originToken := store.Token{Token: msg}
			accountStorage[originAccount.CPF] = originAccount
			accountStorage[destinationAccount.CPF] = destinationAccount
			accountToken[originAccount.ID] = originToken

			tokenUseCase := &TokenUseCaseMock{TokenList: accountToken}
			storagedTransfer := store.NewStoredTransferAccountID()
			accountUsecase := &AccountUsecaseMock{AccountList: accountStorage}
			usecase := UseCase{
				AccountUseCase: accountUsecase,
				TokenUseCase:   tokenUseCase,
				StoredTransfer: storagedTransfer,
			}

			usecase.StoredTransfer.PostTransferID(listTransfer, destinationAccount.ID)
			usecase.StoredTransfer.PostTransferID(listTransfers, destinationAccount.ID)

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
	AccountList map[string]store.Account

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

func (uc AccountUsecaseMock) GetAccounts() []store.Account {
	return nil
}

func (uc AccountUsecaseMock) SearchAccount(id int) store.Account {
	account := store.Account{}

	for _, a := range uc.AccountList {
		if a.ID == id {
			account = a
		}
	}

	return account
}

func (uc *AccountUsecaseMock) UpdateBalance(_ store.Account, _ store.Account) {
	uc.UpdateCallCount++
}

func (uc AccountUsecaseMock) GetAccountCPF(cpf string) store.Account {
	return uc.AccountList[cpf]
}

func (uc AccountUsecaseMock) GetAccount() map[string]store.Account {
	return nil
}

type TokenUseCaseMock struct {
	TokenList map[int]store.Token
}

func (uc TokenUseCaseMock) AuthenticatedLogin(_, _ string) (error, string) {
	return nil, ""
}

func (uc TokenUseCaseMock) ReturnToken(_ int) string {
	return ""
}

func (uc TokenUseCaseMock) GetTokenID(id int) store.Token {
	return uc.TokenList[id]
}
