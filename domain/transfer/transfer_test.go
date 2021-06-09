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
	msg := base64.StdEncoding.EncodeToString([]byte("10/02/2009 02:02:00 : 19727887"))
	msgs := base64.StdEncoding.EncodeToString([]byte("10/03/2009 02:02:00 : 19727887"))
	tt := []struct {
		name    string
		in      CreateTransferInput
		wantErr bool
	}{
		{
			name: "should successfully transfer amount",
			in: CreateTransferInput{
				Token:                msg,
				AccountDestinationID: 19727887,
				Amount:               300,
			},
			wantErr: false,
		},
		{
			name: "should unsuccessfully transfer amount when there is wrong token",
			in: CreateTransferInput{
				Token:                msgs,
				AccountDestinationID: 19727887,
				Amount:               300,
			},
			wantErr: true,
		},
		{
			name: "should unsuccessfully transfer amount when there is no account destination id",
			in: CreateTransferInput{
				Token:                msg,
				AccountDestinationID: 19727807,
				Amount:               300,
			},
			wantErr: true,
		},
		{
			name: "should unsuccessfully transfer amount when the amount is too slow",
			in: CreateTransferInput{
				Token:                msg,
				AccountDestinationID: 19727807,
				Amount:               0,
			},
			wantErr: true,
		},
		{
			name: "should unsuccessfully transfer amount when the amount is greater than the balance",
			in: CreateTransferInput{
				Token:                msg,
				AccountDestinationID: 19727807,
				Amount:               5200,
			},
			wantErr: true,
		},
	}

	for _, testCase := range tt {
		t.Run(testCase.name, func(t *testing.T) {
			listAccount := store.Account{ID: 19727887, Name: "Lucas", CPF: "08131391043", Secret: domain.CreateHash("lixo"), Balance: 5000, CreatedAt: "06/01/2020"}
			listAccounts := store.Account{ID: 98498081, Name: "Rafael", CPF: "38453162093", Secret: domain.CreateHash("call"), Balance: 6000, CreatedAt: "06/01/2020"}

			accountStorage := store.NewStoredAccount()
			accountToken := store.NewStoredToked()
			usecase := UseCase{
				StoredAccount: accountStorage,
				StoredToken:   accountToken,
			}

			usecase.StoredAccount.PostAccount(listAccount)
			usecase.StoredAccount.PostAccount(listAccounts)
			usecase.StoredToken.PostToken(19727887, msg)

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
		})
	}
}

func TestMakeGetTransfers(t *testing.T) {
	msg := base64.StdEncoding.EncodeToString([]byte("10/02/2009 02:02:00:98498081"))
	msgs := base64.StdEncoding.EncodeToString([]byte("10/03/2009 02:02:00:98498081"))

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
			listAccount := store.Account{ID: 19727887, Name: "Lucas", CPF: "08131391043", Secret: domain.CreateHash("lixo"), Balance: 5000, CreatedAt: "06/01/2020"}
			listAccounts := store.Account{ID: 98498081, Name: "Rafael", CPF: "38453162093", Secret: domain.CreateHash("call"), Balance: 6000, CreatedAt: "06/01/2020"}
			listTransfer := store.Transfer{ID: 47278511, AccountOriginID: 98498081, AccountDestinationID: 19727887, Amount: 500, CreatedAt: "13/05/2021 09:09:16"}
			listTransfers := store.Transfer{ID: 6410694, AccountOriginID: 98498081, AccountDestinationID: 19727887, Amount: 200, CreatedAt: "13/05/2021 09:09:16"}

			accountStorage := store.NewStoredAccount()
			accountToken := store.NewStoredToked()
			accountTransfer := store.NewStoredTransferAccountID()
			usecase := UseCase{
				StoredAccount:  accountStorage,
				StoredToken:    accountToken,
				StoredTransfer: accountTransfer,
			}

			usecase.StoredAccount.PostAccount(listAccount)
			usecase.StoredAccount.PostAccount(listAccounts)
			usecase.StoredToken.PostToken(98498081, msg)
			usecase.StoredTransfer.PostTransferID(listTransfer, 98498081)
			usecase.StoredTransfer.PostTransferID(listTransfers, 98498081)

			testss, gotErr := usecase.GetTransfers(testCase.in.Token)

			if !testCase.wantErr && gotErr != nil {
				t.Errorf("unexpected error, wantErr=%v; gotErr=%s", testCase.wantErr, gotErr)
			}

			if testCase.wantErr && gotErr == nil {
				t.Error("wanted err but got nil")
			}
			if &testCase.want != &testss {
				t.Errorf("unexpected error, wantErr=%v; gotErr=%v", testCase.want, testss)
			}
		})
	}
}
