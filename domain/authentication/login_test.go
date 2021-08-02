package authentication

import (
	"github.com/CMedrado/DesafioStone/domain"
	"github.com/google/uuid"
	"testing"
	"time"
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
			accountUsecase := &AccountUseCaseMock{}
			usecase := UseCase{
				AccountUseCase: accountUsecase,
				StoredToken:    LoginRepoMock{},
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
}

func (uc AccountUseCaseMock) CreateAccount(_ string, _ string, _ string, _ int) (uuid.UUID, error) {
	return uuid.UUID{}, nil
}

func (uc AccountUseCaseMock) GetBalance(_ string) (int, error) {
	return 0, nil
}

func (uc AccountUseCaseMock) GetAccounts() ([]domain.Account, error) {
	return []domain.Account{}, nil
}

func (uc AccountUseCaseMock) SearchAccount(_ uuid.UUID) (domain.Account, error) {
	return domain.Account{}, nil
}

func (uc *AccountUseCaseMock) UpdateBalance(_ domain.Account, _ domain.Account) error {
	return nil
}

func (uc AccountUseCaseMock) GetAccountCPF(cpf string) (domain.Account, error) {
	if cpf == "08131391043" {
		return domain.Account{
			ID:        uuid.MustParse("f7ee7351-4c96-40ca-8cd8-37434810ddfa"),
			Name:      "Lucas",
			CPF:       "08131391043",
			Secret:    "c74af74c69d81831a5703aefe9cb4199",
			Balance:   5000,
			CreatedAt: time.Now(),
		}, nil
	}
	if cpf == "38453162093" {
		return domain.Account{
			ID:        uuid.MustParse("a505b1f9-ac4c-45aa-be43-8614a227a9d4"),
			Name:      "Rafael",
			CPF:       "38453162093",
			Secret:    "53b9e9679a8ea25880376080b76f98ad",
			Balance:   6000,
			CreatedAt: time.Now(),
		}, nil
	}
	return domain.Account{}, nil
}

type LoginRepoMock struct {
}

func (rm LoginRepoMock) SaveToken(_ domain.Token) error {
	return nil
}

func (rm LoginRepoMock) ReturnTokenID(id uuid.UUID) (domain.Token, error) {
	if id == uuid.MustParse("a505b1f9-ac4c-45aa-be43-8614a227a9d4") {
		return domain.Token{
			ID:        uuid.MustParse("39a70a94-a82d-4db8-87ae-bd900c6a7c04"),
			IdAccount: uuid.MustParse("a505b1f9-ac4c-45aa-be43-8614a227a9d4"),
			CreatedAt: time.Now(),
		}, nil

	}
	if id == uuid.MustParse("f7ee7351-4c96-40ca-8cd8-37434810ddfa") {
		return domain.Token{
			ID:        uuid.MustParse("40ccb980-538f-4a1d-b1c8-566da5888f45"),
			IdAccount: uuid.MustParse("f7ee7351-4c96-40ca-8cd8-37434810ddfa"),
			CreatedAt: time.Now(),
		}, nil
	}
	return domain.Token{}, nil
}
