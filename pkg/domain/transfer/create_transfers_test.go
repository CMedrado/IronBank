package transfer

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"

	"github.com/CMedrado/DesafioStone/pkg/domain/entities"
)

type CreateTransferInput struct {
	AccountOriginID      int
	Token                string
	AccountDestinationID string
	Amount               int
}

func TestMakeTransfers(t *testing.T) {
	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", "localhost", "6379"),
	})
	tt := []struct {
		name                    string
		in                      CreateTransferInput
		wantErr                 bool
		expectedUpdateCallCount int
	}{
		{
			name: "should successfully transfer amount",
			in: CreateTransferInput{
				Token:                "MjAvMDcvMjAyMSAxNToxNzoyNTo2YjE5NDFkYi1jZTE3LTRmZmUtYTdlZC0yMjQ5M2E5MjZiYmM6MzlhNzBhOTQtYTgyZC00ZGI4LTg3YWUtYmQ5MDBjNmE3YzA0",
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
				redis:          rdb,
			}
			gotErr, gotTransfer := usecase.CreateTransfers(context.Background(), testCase.in.Token, testCase.in.Amount, testCase.in.AccountDestinationID)
			if !testCase.wantErr && gotErr != nil {
				t.Errorf("unexpected error, wantErr=%v; gotErr=%s", testCase.wantErr, gotErr)
			}

			if testCase.wantErr && gotErr == nil {
				t.Error("wanted err but got nil")
			}

			if (gotTransfer == uuid.UUID{}) && !testCase.wantErr && gotErr != nil {
				t.Errorf("expected an Token but got %d", gotTransfer)
			}
		})
	}
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

func (rm TransferRepoMock) ReturnTokenID(id uuid.UUID) (entities.Token, error) {
	time1, _ := time.Parse("2006-01-02T15:04:05.999999999Z07:00", "2021-07-20T15:17:25.933365Z")
	if id.String() == "39a70a94-a82d-4db8-87ae-bd900c6a7c04" {
		return entities.Token{
			ID:        uuid.MustParse("39a70a94-a82d-4db8-87ae-bd900c6a7c04"),
			IdAccount: uuid.MustParse("6b1941db-ce17-4ffe-a7ed-22493a926bbc"),
			CreatedAt: time1,
		}, nil

	}
	if id == uuid.MustParse("f7ee7351-4c96-40ca-8cd8-37434810ddfa") {
		return entities.Token{
			ID:        uuid.MustParse("40ccb980-538f-4a1d-b1c8-566da5888f45"),
			IdAccount: uuid.MustParse("f7ee7351-4c96-40ca-8cd8-37434810ddfa"),
			CreatedAt: time.Now(),
		}, nil
	}
	return entities.Token{}, nil
}

func (uc TransferRepoMock) ChangeBalance(person1, person2 entities.Account) error {
	return nil
}

func (uc TransferRepoMock) ReturnAccountID(id uuid.UUID) (entities.Account, error) {
	if id.String() == "6b1941db-ce17-4ffe-a7ed-22493a926bbc" {
		return entities.Account{
			ID:        uuid.MustParse("6b1941db-ce17-4ffe-a7ed-22493a926bbc"),
			Name:      "Lucas",
			CPF:       "08131391043",
			Secret:    "c74af74c69d81831a5703aefe9cb4199",
			Balance:   5000,
			CreatedAt: time.Now(),
		}, nil
	}
	if id == uuid.MustParse("a505b1f9-ac4c-45aa-be43-8614a227a9d4") {
		return entities.Account{
			ID:        uuid.MustParse("a505b1f9-ac4c-45aa-be43-8614a227a9d4"),
			Name:      "Rafael",
			CPF:       "38453162093",
			Secret:    "53b9e9679a8ea25880376080b76f98ad",
			Balance:   6000,
			CreatedAt: time.Now(),
		}, nil
	}
	if id.String() == "a61227cf-a857-4bc6-8fcd-ad97cdad382a" {
		return entities.Account{
			ID:        uuid.MustParse("a61227cf-a857-4bc6-8fcd-ad97cdad382a"),
			Name:      "Lucas",
			CPF:       "08131391043",
			Secret:    "c74af74c69d81831a5703aefe9cb4199",
			Balance:   5000,
			CreatedAt: time.Now(),
		}, nil
	}
	return entities.Account{}, nil
}
