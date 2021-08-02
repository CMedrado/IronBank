package transfer

import (
	"github.com/CMedrado/DesafioStone/domain"
	store_transfer "github.com/CMedrado/DesafioStone/storage/file/transfer"
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

var i = 0

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
				Token:                "MDIvMDgvMjAyMSAwOToyNzo0NDo2YjE5NDFkYi1jZTE3LTRmZmUtYTdlZC0yMjQ5M2E5MjZiYmM6YmQxODIxZTQtM2I5YS00M2RjLWJkZGUtNjBiM2QyMTRhYzdm",
				AccountDestinationID: "a61227cf-a857-4bc6-8fcd-ad97cdad382a",
				Amount:               300,
			},
			wantErr: false,
		},
		{
			name: "should unsuccessfully transfer amount when there is wrong token",
			in: CreateTransferInput{
				Token:                "MDIvMDgvMjAyMSAwOToyNzo0NDo2YjE5NDFkYi1jZTE3LTRmZmUtYTdlZC0yMjQ5M3E5MjZiYmM6YmQxODIxZTQtM2I5YS00M2RjLWJkZGUtNjBiM2QyMTRhYzdm",
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
			accountUsecase := &AccountUsecaseMock{}
			tokenUseCase := &TokenUseCaseMock{}
			usecase := UseCase{
				AccountUseCase: accountUsecase,
				TokenUseCase:   tokenUseCase,
				StoredTransfer: TransferRepoMock{},
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
				Token: "MDIvMDgvMjAyMSAwOToyNzo0NDo2YjE5NDFkYi1jZTE3LTRmZmUtYTdlZC0yMjQ5M2E5MjZiYmM6YmQxODIxZTQtM2I5YS00M2RjLWJkZGUtNjBiM2QyMTRhYzdm",
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
			accountUsecase := &AccountUsecaseMock{}
			tokenUseCase := &TokenUseCaseMock{}
			usecase := UseCase{
				AccountUseCase: accountUsecase,
				TokenUseCase:   tokenUseCase,
				StoredTransfer: TransferRepoMock{},
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
}

func (uc AccountUsecaseMock) CreateAccount(_ string, _ string, _ string, _ int) (uuid.UUID, error) {
	return uuid.UUID{}, nil
}

func (uc AccountUsecaseMock) GetBalance(_ string) (int, error) {
	return 0, nil
}

func (uc AccountUsecaseMock) GetAccounts() ([]domain.Account, error) {
	return []domain.Account{}, nil
}

func (uc AccountUsecaseMock) SearchAccount(id uuid.UUID) (domain.Account, error) {
	time1, _ := time.Parse("2006-01-02T15:04:05.999999999Z07:00", "2021-07-20T15:17:25.933365Z")
	time2, _ := time.Parse("2006-01-02T15:04:05.999999999Z07:00", "2021-07-20T15:15:58.201088Z")
	if id == uuid.MustParse("6b1941db-ce17-4ffe-a7ed-22493a926bbc") {
		return domain.Account{
			ID:        uuid.MustParse("6b1941db-ce17-4ffe-a7ed-22493a926bbc"),
			Name:      "Lucas",
			CPF:       "38453162093",
			Secret:    "7e65a9b554bbc9817aa049ce38c84a72",
			Balance:   1000,
			CreatedAt: time1,
		}, nil
	}
	if id == uuid.MustParse("a61227cf-a857-4bc6-8fcd-ad97cdad382a") {
		return domain.Account{
			ID:        uuid.MustParse("a61227cf-a857-4bc6-8fcd-ad97cdad382a"),
			Name:      "Rafael",
			CPF:       "08131391043",
			Secret:    "3467e121a1a109628e0a5b0cebba361b",
			Balance:   6000,
			CreatedAt: time2,
		}, nil
	}
	return domain.Account{}, nil
}

func (uc *AccountUsecaseMock) UpdateBalance(_ domain.Account, _ domain.Account) error {
	return nil
}

func (uc AccountUsecaseMock) GetAccountCPF(cpf string) (domain.Account, error) {
	time1, _ := time.Parse("2006-01-02T15:04:05.999999999Z07:00", "2021-07-20T15:17:25.933365Z")
	time2, _ := time.Parse("2006-01-02T15:04:05.999999999Z07:00", "2021-07-20T15:15:58.201088Z")
	if cpf == "38453162093" {
		return domain.Account{
			ID:        uuid.MustParse("6b1941db-ce17-4ffe-a7ed-22493a926bbc"),
			Name:      "Lucas",
			CPF:       "38453162093",
			Secret:    "7e65a9b554bbc9817aa049ce38c84a72",
			Balance:   1000,
			CreatedAt: time1,
		}, nil
	}
	if cpf == "08131391043" {
		return domain.Account{
			ID:        uuid.MustParse("a61227cf-a857-4bc6-8fcd-ad97cdad382a"),
			Name:      "Rafael",
			CPF:       "08131391043",
			Secret:    "3467e121a1a109628e0a5b0cebba361b",
			Balance:   6000,
			CreatedAt: time2,
		}, nil
	}
	return domain.Account{}, nil
}

type TokenUseCaseMock struct {
}

func (uc TokenUseCaseMock) AuthenticatedLogin(_, _ string) (error, string) {
	return nil, ""
}

func (uc TokenUseCaseMock) GetTokenID(id uuid.UUID) (domain.Token, error) {
	time1, _ := time.Parse("2006-01-02T15:04:05.999999999Z07:00", "2021-08-02T09:27:44.933365Z")
	time2, _ := time.Parse("2006-01-02T15:04:05.999999999Z07:00", "2021-08-02T15:15:58.201088Z")
	if i == 0 {
		i++
		return domain.Token{
			ID:        uuid.MustParse("39a70a94-a82d-4db8-87ae-bd900c6a7c04"),
			IdAccount: uuid.MustParse("6b1941db-ce17-4ffe-a7ed-22493a926bbc"),
			CreatedAt: time1,
		}, nil
	}
	if id == uuid.MustParse("a61227cf-a857-4bc6-8fcd-ad97cdad382a") {
		return domain.Token{
			ID:        uuid.MustParse("40ccb980-538f-4a1d-b1c8-566da5888f45"),
			IdAccount: uuid.MustParse("a61227cf-a857-4bc6-8fcd-ad97cdad382a"),
			CreatedAt: time2,
		}, nil
	}
	return domain.Token{}, nil
}

type TransferRepoMock struct {
}

func (uc TransferRepoMock) ReturnTransfer() ([]domain.Transfer, error) {
	time1, _ := time.Parse("2006-01-02T15:04:05.999999999Z07:00", "2021-07-20T15:17:25.933365Z")
	return []domain.Transfer{
		{
			ID:                   uuid.MustParse("47399f23-2093-4dde-b32f-990cac27630e"),
			OriginAccountID:      uuid.MustParse("6b1941db-ce17-4ffe-a7ed-22493a926bbc"),
			DestinationAccountID: uuid.MustParse("a61227cf-a857-4bc6-8fcd-ad97cdad382a"),
			Amount:               150,
			CreatedAt:            time1,
		},
	}, nil
}

func (uc TransferRepoMock) SaveTransfer(_ domain.Transfer) error {
	return nil
}
