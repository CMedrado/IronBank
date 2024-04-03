package authentication

import (
	"bytes"
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	domain2 "github.com/CMedrado/DesafioStone/pkg/domain"
	"github.com/CMedrado/DesafioStone/pkg/domain/authentication"
	"github.com/CMedrado/DesafioStone/pkg/domain/entities"
	"github.com/google/uuid"
)

var (
	cpfIncorrect = `{"errors":"given secret or CPF are incorrect"}` + "\n"
)

func TestHandler_Login(t *testing.T) {
	logint := []struct {
		name         string
		method       string
		path         string
		body         string
		response     int
		responsebody string
	}{
		{
			name:         "should successfully authenticated login with formatted CPF",
			method:       "POST",
			path:         "/login",
			body:         `{"cpf": "08131391043", "Secret": "lixo"}`,
			response:     http.StatusOK,
			responsebody: `{"token":"passou"}` + "\n",
		},
		{
			name:         "should successfully authenticated login with unformatted CPF",
			method:       "POST",
			path:         "/login",
			body:         `{"cpf": "38453162093", "Secret": "call"}`,
			response:     http.StatusOK,
			responsebody: `{"token":"passou"}` + "\n",
		},
		{
			name:         "should unsuccessfully authenticated login when cpf is not registered",
			method:       "POST",
			path:         "/login",
			body:         `{"cpf": "38453162723", "Secret": "jax"}`,
			response:     http.StatusUnauthorized,
			responsebody: cpfIncorrect,
		},
		{
			name:         "should unsuccessfully create an account when CPF is invalid",
			method:       "POST",
			path:         "/login",
			body:         `{"cpf": "384531.62793", "Secret": "jax"}`,
			response:     http.StatusUnauthorized,
			responsebody: cpfIncorrect,
		},
		{
			name:         "should unsuccessfully authenticated login when secret is not correct",
			method:       "POST",
			path:         "/login",
			body:         `{"cpf": "081.313.910-43", "Secret": "call"}`,
			response:     http.StatusUnauthorized,
			responsebody: cpfIncorrect,
		},
		{
			name:     "should unsuccessfully authenticated login when json is invalid",
			method:   "POST",
			path:     "/login",
			body:     `{"Secret" "jax"}`,
			response: http.StatusBadRequest,
		},
	}
	for _, tc := range logint {
		t.Run(tc.name, func(t *testing.T) {
			s := new(Handler)
			s.login = &TokenUseCaseMock{}
			s.account = &AccountUsecaseMock{}
			bodyBytes := []byte(tc.body)
			request, _ := http.NewRequest(tc.method, tc.path, bytes.NewReader(bodyBytes))
			responseRecorder := httptest.NewRecorder()

			s.Login(responseRecorder, request)

			if tc.response != responseRecorder.Code {
				t.Errorf("unexpected error, wantErr= %d; gotErr= %d", tc.response, responseRecorder.Code)
			}

			if responseRecorder.Body.String() != tc.responsebody && tc.responsebody != "" {
				t.Errorf("expected an %s but got %s", tc.responsebody, responseRecorder.Body.String())
			}
		})
	}
}

type TokenUseCaseMock struct {
}

func (uc TokenUseCaseMock) AuthenticatedLogin(secret string, account entities.Account) (error, string) {
	secretHash := domain2.CreateHash(secret)
	if account == (entities.Account{}) {
		return authentication.ErrLogin, ""
	}
	//if account.CPF != account.CPF {
	//	return errors.New("given secret or CPF are incorrect"), ""
	//}
	if account.Secret != secretHash {
		return authentication.ErrLogin, ""
	}
	return nil, "passou"
}

func (uc TokenUseCaseMock) GetTokenID(_ uuid.UUID) (entities.Token, error) {
	return entities.Token{}, nil
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
	return 0, nil
}

func (uc AccountUsecaseMock) GetAccounts() ([]entities.Account, error) {
	time1, _ := time.Parse("2006-01-02T15:04:05.999999999Z07:00", "2021-08-02T09:41:46.813816-03:00")
	return []entities.Account{
		{
			ID:        uuid.MustParse("f7ee7351-4c96-40ca-8cd8-37434810ddfa"),
			Name:      "Rafael",
			CPF:       "38453162093",
			Secret:    "7edb360f06acaef2cc80dba16cf563f199d347db4443da04da0c8173e3f9e4ed",
			Balance:   6000,
			CreatedAt: time1,
		},
		{
			ID:        uuid.MustParse("a505b1f9-ac4c-45aa-be43-8614a227a9d4"),
			Name:      "Lucas",
			CPF:       "08131391043",
			Secret:    "3cf4897608d8ae2d9ccd9e087be19fc7ba962b9ae38e784f1b770eea3394645f",
			Balance:   5000,
			CreatedAt: time1,
		},
	}, nil
}

func (uc AccountUsecaseMock) SearchAccount(id uuid.UUID) (entities.Account, error) {
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
