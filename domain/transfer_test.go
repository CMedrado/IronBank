package domain

import (
	"encoding/base64"
	"github.com/CMedrado/DesafioStone/store"
	"testing"
)

type CreateTransferInput struct {
	AccountOriginID      int
	Token                string
	AccountDestinationID int
	Amount               uint
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
			listAccount := store.Account{19727887, "Lucas", "08131391043", Hash("lixo"), 5000, "06/01/2020"}
			listAccounts := store.Account{98498081, "Rafael", "38453162093", Hash("call"), 6000, "06/01/2020"}

			accountStorage := store.NewStoredAccount()
			accountToken := store.NewStoredToked()
			usecase := AccountUsecase{
				Store: accountStorage,
				Token: accountToken,
			}

			usecase.Store.TransferredAccount(listAccount)
			usecase.Store.TransferredAccount(listAccounts)
			usecase.Token.CreatedToken(19727887, msg)

			gotErr, gotTransfer := usecase.MakeTransfers(testCase.in.Token, testCase.in.AccountDestinationID, testCase.in.Amount)

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
	msg := base64.StdEncoding.EncodeToString([]byte("10/02/2009 02:02:00 : 98498081"))
	msgs := base64.StdEncoding.EncodeToString([]byte("10/03/2009 02:02:00 : 98498081"))

	tt := []struct {
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
			want:    []store.Transfer{{6410694, 98498081, 19727887, 200, "13/05/2021 09:09:16"}, {47278511, 98498081, 19727887, 500, "13/05/2021 09:09:16"}},
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
			listAccount := store.Account{19727887, "Lucas", "08131391043", Hash("lixo"), 5000, "06/01/2020"}
			listAccounts := store.Account{98498081, "Rafael", "38453162093", Hash("call"), 6000, "06/01/2020"}
			listTransfer := store.Transfer{47278511, 98498081, 19727887, 500, "13/05/2021 09:09:16"}
			listTransfers := store.Transfer{6410694, 98498081, 19727887, 200, "13/05/2021 09:09:16"}

			accountStorage := store.NewStoredAccount()
			accountToken := store.NewStoredToked()
			accountTransfer := store.NewStoredTransferTwo()
			usecase := AccountUsecase{
				Store:    accountStorage,
				Token:    accountToken,
				Transfer: accountTransfer,
			}

			usecase.Store.TransferredAccount(listAccount)
			usecase.Store.TransferredAccount(listAccounts)
			usecase.Token.CreatedToken(98498081, msg)
			usecase.Transfer.CreatedTransferTwo(listTransfer, 98498081)
			usecase.Transfer.CreatedTransferTwo(listTransfers, 98498081)

			_, gotErr := usecase.GetTransfers(testCase.in.Token)

			if !testCase.wantErr && gotErr != nil {
				t.Errorf("unexpected error, wantErr=%v; gotErr=%s", testCase.wantErr, gotErr)
			}

			if testCase.wantErr && gotErr == nil {
				t.Error("wanted err but got nil")
			}
		})
	}
}
