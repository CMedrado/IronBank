package account

import (
	"bytes"
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/CMedrado/DesafioStone/pkg/domain/entities"
	"github.com/google/uuid"
)

var Aux = 0

func TestHandler_CreateAccount(t *testing.T) {

	tt := []struct {
		name         string
		method       string
		path         string
		body         string
		response     int
		responseBody string
	}{
		{
			name:         "Should successfully create an account with formatted CPF",
			method:       "POST",
			path:         "/accounts",
			body:         `{"name": "Rafael", "cpf": "081.313.910-43", "secret": "tatatal", "balance": 5000}`,
			responseBody: `{"id":"f7ee7351-4c96-40ca-8cd8-37434810ddfa"}` + "\n",
			response:     http.StatusCreated,
		},
		{
			name:         "should successfully create an account with unformatted CPF",
			method:       "POST",
			path:         "/accounts",
			body:         `{"name": "Lucas", "cpf": "38453162093", "secret": "jax", "balance": 3000}`,
			responseBody: `{"id":"a505b1f9-ac4c-45aa-be43-8614a227a9d4"}` + "\n",
			response:     http.StatusCreated,
		},
		{
			name:         "should unsuccessfully create an account when CPF is invalid",
			method:       "POST",
			path:         "/accounts",
			body:         `{"name": "Rafael", "cpf": "08131.391043", "secret": "tatatal", "balance": 5000}`,
			response:     http.StatusBadRequest,
			responseBody: `{"errors":"given cpf is invalid"}` + "\n",
		},
		{
			name:         "should unsuccessfully create an account when balance is invalid",
			method:       "POST",
			path:         "/accounts",
			body:         `{"name": "Rafael", "cpf": "08131391043", "secret": "tatatal", "balance": -5}`,
			response:     http.StatusBadRequest,
			responseBody: `{"errors":"given the balance amount is invalid"}` + "\n",
		},
		{
			name:     "should unsuccessfully create an account when json is invalid",
			method:   "POST",
			path:     "/accounts",
			body:     `{"ne" "Lucas", "cpf" "38453162093", "secret""jax", "balance" 3000}`,
			response: http.StatusBadRequest,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			s := new(Handler)
			s.account = &AccountUsecaseMock{}
			bodyBytes := []byte(tc.body)
			request, _ := http.NewRequest(tc.method, tc.path, bytes.NewReader(bodyBytes))
			responseRecorder := httptest.NewRecorder()

			s.CreateAccount(responseRecorder, request)

			if tc.response != responseRecorder.Code {
				t.Errorf("unexpected error, wantErr= %d; gotErr= %d", tc.response, responseRecorder.Code)
			}

			if responseRecorder.Body.String() != tc.responseBody && tc.responseBody != "" {
				t.Errorf("expected an %s but got %s", tc.responseBody, responseRecorder.Body.String())
			}
		})
	}
}

type AccountUsecaseMock struct {
}

func (uc AccountUsecaseMock) CreateAccount(_ context.Context, name string, cpf string, _ string, balance int) (uuid.UUID, error) {
	if len(cpf) != 11 && len(cpf) != 14 {
		return uuid.UUID{}, errors.New("given cpf is invalid")
	}
	if len(cpf) == 14 {
		if string([]rune(cpf)[3]) == "." && string([]rune(cpf)[7]) == "." && string([]rune(cpf)[11]) == "-" {
			if balance <= 0 {
				return uuid.UUID{}, errors.New("given the balance amount is invalid")
			}
			if name == "Rafael" {
				return uuid.MustParse("f7ee7351-4c96-40ca-8cd8-37434810ddfa"), nil
			}
			if name == "Lucas" {
				return uuid.MustParse("a505b1f9-ac4c-45aa-be43-8614a227a9d4"), nil
			}
		} else {
			return uuid.UUID{}, errors.New("given cpf is invalid")
		}
	}
	if balance <= 0 {
		return uuid.UUID{}, errors.New("given the balance amount is invalid")
	}
	if name == "Rafael" {
		return uuid.MustParse("f7ee7351-4c96-40ca-8cd8-37434810ddfa"), nil
	}
	if name == "Lucas" {
		return uuid.MustParse("a505b1f9-ac4c-45aa-be43-8614a227a9d4"), nil
	}
	return uuid.UUID{}, nil
}

func (uc AccountUsecaseMock) GetBalance(_ string) (int, error) {
	if Aux == 0 {
		Aux++
		return 6000, nil
	} else {
		return 0, errors.New("given id is invalid")
	}
}

func (uc AccountUsecaseMock) GetAccounts() ([]entities.Account, error) {
	time1, _ := time.Parse("2006-01-02T15:04:05.999999999Z07:00", "2021-08-02T09:41:46.813816-03:00")
	return []entities.Account{
		{
			ID:        uuid.MustParse("f7ee7351-4c96-40ca-8cd8-37434810ddfa"),
			Name:      "Rafael",
			CPF:       "38453162093",
			Secret:    "53b9e9679a8ea25880376080b76f98ad",
			Balance:   6000,
			CreatedAt: time1,
		},
		{
			ID:        uuid.MustParse("a505b1f9-ac4c-45aa-be43-8614a227a9d4"),
			Name:      "Lucas",
			CPF:       "08131391043",
			Secret:    "c74af74c69d81831a5703aefe9cb4199",
			Balance:   5000,
			CreatedAt: time1,
		},
	}, nil
}

func (uc AccountUsecaseMock) GetAccountID(id uuid.UUID) (entities.Account, error) {
	account := entities.Account{}
	accounts, _ := uc.GetAccounts()
	for _, a := range accounts {
		if a.ID == id {
			account = a
		}
	}

	return account, nil
}

func (uc *AccountUsecaseMock) UpdateBalance(_ entities.Account, _ entities.Account) error {
	return nil
}

func (uc AccountUsecaseMock) GetAccountCPF(_ context.Context, cpf string) (entities.Account, error) {
	account := entities.Account{}
	accounts, _ := uc.GetAccounts()
	for _, a := range accounts {
		if a.CPF == cpf {
			account = a
		}
	}

	return account, nil
}
