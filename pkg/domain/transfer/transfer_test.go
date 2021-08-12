package transfer

import (
	"github.com/CMedrado/DesafioStone/pkg/domain"
	"github.com/CMedrado/DesafioStone/pkg/domain/entities"
	"github.com/CMedrado/DesafioStone/pkg/gateways/db/file/transfer"
	transfer2 "github.com/CMedrado/DesafioStone/pkg/gateways/http/transfer"
	"github.com/google/uuid"
	"testing"
	"time"
)

type CreateTransferInput struct {
	AccountOriginID      int
	Token                string
	AccountDestinationID string
	Amount               int
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
				Token:                "MDIvMDgvMjAyMSAwOToyNzo0NDo2YjE5NDFkYi1jZTE3LTRmZmUtYTdlZC0yMjQ5M2E5MjZiYmM6MzlhNzBhOTQtYTgyZC00ZGI4LTg3YWUtYmQ5MDBjNmE3YzA0",
				AccountDestinationID: "a61227cf-a857-4bc6-8fcd-ad97cdad382a",
				Amount:               300,
			},
			wantErr: false,
		},
		{
			name: "should unsuccessfully transfer amount when there is wrong token",
			in: CreateTransferInput{
				Token:                "MDIvMDgvMjAyMSAwOToyNzo0NDo2YjE5NDFkYi1jZTE3LTRmZmUtYTdlZC0yMjQ5M3E5MjZiYmM6YmQxODIxZTQtM2I5YS00M2RjLWJkZGUtNjBiM2QyMTRiYzdm",
				AccountDestinationID: "75432539-c5ba-46d3-9690-44985b516da7",
				Amount:               300,
			},
			wantErr:                 true,
			expectedUpdateCallCount: 0,
		},
		{
			name: "should unsuccessfully transfer amount when there is no account destination id",
			in: CreateTransferInput{
				Token:                "MDIvMDgvMjAyMSAwOToyNzo0NDo2YjE5NDFkYi1jZTE3LTRmZmUtYTdlZC0yMjQ5M2E5MjZiYmM6YmQxODIxZTQtM2I5YS00M2RjLWJkZGUtNjBiM2QyMTRhYzdm",
				AccountDestinationID: "c5424440-4737-4e03-86d2-3adac90ddd20",
				Amount:               300,
			},
			wantErr:                 true,
			expectedUpdateCallCount: 0,
		},
		{
			name: "should unsuccessfully transfer amount when the amount is too slow",
			in: CreateTransferInput{
				Token:                "MDIvMDgvMjAyMSAwOToyNzo0NDo2YjE5NDFkYi1jZTE3LTRmZmUtYTdlZC0yMjQ5M2E5MjZiYmM6YmQxODIxZTQtM2I5YS00M2RjLWJkZGUtNjBiM2QyMTRhYzdm",
				AccountDestinationID: "75432539-c5ba-46d3-9690-44985b516da7",
				Amount:               0,
			},
			wantErr:                 true,
			expectedUpdateCallCount: 0,
		},
		{
			name: "should unsuccessfully transfer amount when the amount is greater than the balance",
			in: CreateTransferInput{
				Token:                "MDIvMDgvMjAyMSAwOToyNzo0NDo2YjE5NDFkYi1jZTE3LTRmZmUtYTdlZC0yMjQ5M2E5MjZiYmM6YmQxODIxZTQtM2I5YS00M2RjLWJkZGUtNjBiM2QyMTRhYzdm",
				AccountDestinationID: "75432539-c5ba-46d3-9690-44985b516da7",
				Amount:               52000,
			},
			wantErr:                 true,
			expectedUpdateCallCount: 0,
		},
	}

	for _, testCase := range tt {
		t.Run(testCase.name, func(t *testing.T) {
			usecase := UseCase{
				StoredTransfer: TransferRepoMock{},
			}
			accountOriginID, tokenOriginID, gotErr := transfer2.DecoderToken(testCase.in.Token)
			if gotErr == nil {
				accountOrigin, gotErr := SearchAccount(accountOriginID)
				if gotErr == nil {
					accountToken, gotErr := GetTokenID(tokenOriginID)
					if gotErr == nil {
						accountDestinationIdUUID, gotErr := uuid.Parse(testCase.in.AccountDestinationID)
						if gotErr == nil {
							accountDestination, gotErr := SearchAccount(accountDestinationIdUUID)
							if gotErr == nil {
								gotErr, gotTransfer, _, _ := usecase.CreateTransfers(accountOriginID, accountToken, testCase.in.Token, accountOrigin, accountDestination, testCase.in.Amount, accountDestinationIdUUID)
								if !testCase.wantErr && gotErr != nil {
									t.Errorf("unexpected error, wantErr=%v; gotErr=%s", testCase.wantErr, gotErr)
								}

								if testCase.wantErr && gotErr == nil {
									t.Error("wanted err but got nil")
								}

								if (gotTransfer == uuid.UUID{}) && !testCase.wantErr && gotErr != nil {
									t.Errorf("expected an Token but got %d", gotTransfer)
								}
							}
						}
					}
				}
			}
		})
	}
}

func TestMakeGetTransfers(t *testing.T) {

	var tt = []struct {
		name    string
		in      CreateTransferInput
		wantErr bool
		want    []transfer.Transfer
	}{
		{
			name: "should successfully get transfers",
			in: CreateTransferInput{
				Token: "MDIvMDgvMjAyMSAwOToyNzo0NDo2YjE5NDFkYi1jZTE3LTRmZmUtYTdlZC0yMjQ5M2E5MjZiYmM6MzlhNzBhOTQtYTgyZC00ZGI4LTg3YWUtYmQ5MDBjNmE3YzA0",
			},
			wantErr: false,
		},
		{
			name: "should unsuccessfully get transfer when there is wrong token",
			in: CreateTransferInput{
				Token: "MDIvMDgvMjAyMSAwOToyNzo0NDo2YjE5NDFkYi1jZTE3LTRmZmUtYTdlZC0yMsQ5M2E5MjZiYmM6YmQxODIxZTQtM2I5YS00M2RjLWJkZGUtNjBiM2QyMTRhYzdm",
			},
			wantErr: true,
		},
	}
	for _, testCase := range tt {
		t.Run(testCase.name, func(t *testing.T) {
			usecase := UseCase{
				StoredTransfer: TransferRepoMock{},
			}
			accountOriginID, tokenID, gotErr := transfer2.DecoderToken(testCase.in.Token)
			accountToken, gotErr := GetTokenID(tokenID)
			gotTransfer, gotErr := usecase.GetTransfers(accountOriginID, accountToken, testCase.in.Token)

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

func SearchAccount(id uuid.UUID) (entities.Account, error) {
	time1, _ := time.Parse("2006-01-02T15:04:05.999999999Z07:00", "2021-07-20T15:17:25.933365Z")
	time2, _ := time.Parse("2006-01-02T15:04:05.999999999Z07:00", "2021-07-20T15:15:58.201088Z")
	if id == uuid.MustParse("6b1941db-ce17-4ffe-a7ed-22493a926bbc") {
		return entities.Account{
			ID:        uuid.MustParse("6b1941db-ce17-4ffe-a7ed-22493a926bbc"),
			Name:      "Lucas",
			CPF:       "38453162093",
			Secret:    "7e65a9b554bbc9817aa049ce38c84a72",
			Balance:   1000,
			CreatedAt: time1,
		}, nil
	}
	if id == uuid.MustParse("a61227cf-a857-4bc6-8fcd-ad97cdad382a") {
		return entities.Account{
			ID:        uuid.MustParse("a61227cf-a857-4bc6-8fcd-ad97cdad382a"),
			Name:      "Rafael",
			CPF:       "08131391043",
			Secret:    "3467e121a1a109628e0a5b0cebba361b",
			Balance:   6000,
			CreatedAt: time2,
		}, nil
	}
	return entities.Account{}, domain.ErrAccountExists
}

func GetAccountCPF(cpf string) (entities.Account, error) {
	time1, _ := time.Parse("2006-01-02T15:04:05.999999999Z07:00", "2021-07-20T15:17:25.933365Z")
	time2, _ := time.Parse("2006-01-02T15:04:05.999999999Z07:00", "2021-07-20T15:15:58.201088Z")
	if cpf == "38453162093" {
		return entities.Account{
			ID:        uuid.MustParse("6b1941db-ce17-4ffe-a7ed-22493a926bbc"),
			Name:      "Lucas",
			CPF:       "38453162093",
			Secret:    "7e65a9b554bbc9817aa049ce38c84a72",
			Balance:   1000,
			CreatedAt: time1,
		}, nil
	}
	if cpf == "08131391043" {
		return entities.Account{
			ID:        uuid.MustParse("a61227cf-a857-4bc6-8fcd-ad97cdad382a"),
			Name:      "Rafael",
			CPF:       "08131391043",
			Secret:    "3467e121a1a109628e0a5b0cebba361b",
			Balance:   6000,
			CreatedAt: time2,
		}, nil
	}
	return entities.Account{}, nil
}

func GetTokenID(id uuid.UUID) (entities.Token, error) {
	time1, _ := time.Parse("2006-01-02T15:04:05.999999999Z07:00", "2021-08-02T09:27:44.933365Z")
	time2, _ := time.Parse("2006-01-02T15:04:05.999999999Z07:00", "2021-08-02T15:15:58.201088Z")
	if id == uuid.MustParse("39a70a94-a82d-4db8-87ae-bd900c6a7c04") {
		return entities.Token{
			ID:        uuid.MustParse("39a70a94-a82d-4db8-87ae-bd900c6a7c04"),
			IdAccount: uuid.MustParse("6b1941db-ce17-4ffe-a7ed-22493a926bbc"),
			CreatedAt: time1,
		}, nil
	}
	if id == uuid.MustParse("a61227cf-a857-4bc6-8fcd-ad97cdad382a") {
		return entities.Token{
			ID:        uuid.MustParse("40ccb980-538f-4a1d-b1c8-566da5888f45"),
			IdAccount: uuid.MustParse("a61227cf-a857-4bc6-8fcd-ad97cdad382a"),
			CreatedAt: time2,
		}, nil
	}
	return entities.Token{}, nil
}

type TransferRepoMock struct {
}

func (uc TransferRepoMock) ReturnTransfer(id uuid.UUID) ([]entities.Transfer, error) {
	time1, _ := time.Parse("2006-01-02T15:04:05.999999999Z07:00", "2021-07-20T15:17:25.933365Z")
	return []entities.Transfer{
		{
			ID:                   uuid.MustParse("47399f23-2093-4dde-b32f-990cac27630e"),
			OriginAccountID:      uuid.MustParse("6b1941db-ce17-4ffe-a7ed-22493a926bbc"),
			DestinationAccountID: uuid.MustParse("a61227cf-a857-4bc6-8fcd-ad97cdad382a"),
			Amount:               150,
			CreatedAt:            time1,
		},
	}, nil
}

func (uc TransferRepoMock) SaveTransfer(_ entities.Transfer) error {
	return nil
}
