package transfer

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"

	domain2 "github.com/CMedrado/DesafioStone/pkg/domain"
	"github.com/CMedrado/DesafioStone/pkg/domain/entities"
)

func TestHandler_CreateTransfer(t *testing.T) {
	createtransfer := []struct {
		name         string
		method       string
		path         string
		body         string
		response     int
		responsebody string
		token        string
	}{
		{
			name:         "should successfully transfer amount",
			method:       "POST",
			path:         "/transfers",
			body:         `{"account_destination_id":"a61227cf-a857-4bc6-8fcd-ad97cdad382a","amount": 500}`,
			response:     http.StatusCreated,
			responsebody: `{"id":"c5424440-4737-4e03-86d2-3adac90ddd20"}` + "\n",
			token:        "Basic MDIvMDgvMjAyMSAwOToyNzo0NDo2YjE5NDFkYi1jZTE3LTRmZmUtYTdlZC0yMjQ5M2E5MjZiYmM6YmQxODIxZTQtM2I5YS00M2RjLWJkZGUtNjBiM2QyMTRhYzdm",
		},
		{
			name:         "should unsuccessfully transfer amount when there is wrong token",
			method:       "POST",
			path:         "/transfers",
			body:         `{"account_destination_id":"75432539-c5ba-46d3-9690-44985b516da7","amount": 300}`,
			response:     http.StatusUnauthorized,
			token:        "Basic MDIvMDgvMjAyMSAwOToyNzo0NDo2YjE5NDFkYi1jZTE3LTRmZmUtYTdlZC0yMeQ5M2E5MjZiYmM6YmQxODIxZTQtM2I5YS00M2RjLWJkZGUtNjBiM2QyMTRhYzdm",
			responsebody: `{"errors":"given token is invalid"}` + "\n",
		},
		{
			name:         "should unsuccessfully transfer amount when there is wrong destination ID",
			method:       "POST",
			path:         "/transfers",
			body:         `{"account_destination_id":"75432539-c5ba-46d3-9690-44985b516da5","amount": 300}`,
			response:     http.StatusNotFound,
			token:        "Basic MDIvMDgvMjAyMSAwOToyNzo0NDo2YjE5NDFkYi1jZTE3LTRmZmUtYTdlZC0yMjQ5M2E5MjZiYmM6YmQxODIxZTQtM2I5YS00M2RjLWJkZGUtNjBiM2QyMTRhYzdm",
			responsebody: `{"errors":"given account destination id is invalid"}` + "\n",
		},
		{
			name:         "should unsuccessfully transfer amount when there is invalid amount",
			method:       "POST",
			path:         "/transfers",
			body:         `{"account_destination_id":"75432539-c5ba-46d3-9690-44985b516da7","amount": -5}`,
			response:     http.StatusBadRequest,
			token:        "Basic MDIvMDgvMjAyMSAwOToyNzo0NDo2YjE5NDFkYi1jZTE3LTRmZmUtYTdlZC0yMjQ5M2E5MjZiYmM6YmQxODIxZTQtM2I5YS00M2RjLWJkZGUtNjBiM2QyMTRhYzdm",
			responsebody: `{"errors":"given amount is invalid"}` + "\n",
		},
		{
			name:         "should unsuccessfully transfer amount when there without balance ",
			method:       "POST",
			path:         "/transfers",
			body:         `{"account_destination_id":"75432539-c5ba-46d3-9690-44985b516da7","amount": 60000}`,
			response:     http.StatusBadRequest,
			token:        "Basic MDIvMDgvMjAyMSAwOToyNzo0NDo2YjE5NDFkYi1jZTE3LTRmZmUtYTdlZC0yMjQ5M2E5MjZiYmM6YmQxODIxZTQtM2I5YS00M2RjLWJkZGUtNjBiM2QyMTRhYzdm",
			responsebody: `{"errors":"given account without balance"}` + "\n",
		},
		{
			name:         "should unsuccessfully transfer amount when there same account ",
			method:       "POST",
			path:         "/transfers",
			body:         `{"account_destination_id":"6b1941db-ce17-4ffe-a7ed-22493a926bbc","amount": 300}`,
			response:     http.StatusBadRequest,
			token:        "Basic MDIvMDgvMjAyMSAwOToyNzo0NDo2YjE5NDFkYi1jZTE3LTRmZmUtYTdlZC0yMjQ5M2E5MjZiYmM6YmQxODIxZTQtM2I5YS00M2RjLWJkZGUtNjBiM2QyMTRhYzdm",
			responsebody: `{"errors":"given account is the same as the account destination"}` + "\n",
		},
		{
			name:     "should unsuccessfully transfer amount when json is invalid",
			method:   "POST",
			path:     "/transfers",
			body:     `{"account"0}`,
			response: http.StatusBadRequest,
			token:    "Basic MDIvMDgvMjAyMSAwOToyNzo0NDo2YjE5NDFkYi1jZTE3LTRmZmUtYTdlZC0yMjQ5M2E5MjZiYmM6YmQxODIxZTQtM2I5YS00M2RjLWJkZGUtNjBiM2QyMTRhYzdm",
		},
		{
			name:         "should successfully transfer amount",
			method:       "POST",
			path:         "/transfers",
			body:         `{"account_destination_id":"a61227cf-a857-4bc6-8fcd-ad97cdad382a","amount": 500}`,
			response:     http.StatusBadRequest,
			responsebody: `{"errors":"given the authorization header type is not basic"}` + "\n",
			token:        "MDIvMDgvMjAyMSAwOToyNzo0NDo2YjE5NDFkYi1jZTE3LTRmZmUtYTdlZC0yMjQ5M2E5MjZiYmM6YmQxODIxZTQtM2I5YS00M2RjLWJkZGUtNjBiM2QyMTRhYzdm",
		},
	}
	for _, tc := range createtransfer {
		t.Run(tc.name, func(t *testing.T) {
			s := new(Handler)
			s.transfer = &TransferUsecaseMock{}
			s.account = &AccountUsecaseMock{}
			s.login = &TokenUseCaseMock{}
			logger := logrus.New()
			logger.SetFormatter(&logrus.TextFormatter{TimestampFormat: time.RFC3339})
			Lentry := logrus.NewEntry(logger)
			s.logger = Lentry
			bodyBytes := []byte(tc.body)
			request, _ := http.NewRequest(tc.method, tc.path, bytes.NewReader(bodyBytes))
			respondeRecorder := httptest.NewRecorder()

			request.Header.Add("Authorization", tc.token)

			s.CreateTransfer(respondeRecorder, request)

			if tc.response != respondeRecorder.Code { // O teste falhará pois não queremos erro e obtivemos um
				t.Errorf("unexpected error, wantErr= %d; gotErr= %d", tc.response, respondeRecorder.Code)
			}
			if respondeRecorder.Body.String() != tc.responsebody && tc.responsebody != "" {
				t.Errorf("expected an %s but got %s", tc.responsebody, respondeRecorder.Body.String())
			}
		})
	}
}

type TransferUsecaseMock struct {
}

func (uc *TransferUsecaseMock) GetTransfers(_ entities.Account, _ entities.Token, token string) ([]entities.Transfer, error) {
	time1, _ := time.Parse("2006-01-02T15:04:05.999999999Z07:00", "2021-07-20T15:17:25.933365Z")
	if "MDIvMDgvMjAyMSAwOToyNzo0NDo2YjE5NDFkYi1jZTE3LTRmZmUtYTdlZC0yMjQ5M2E5MjZiYmM6YmQxODIxZTQtM2I5YS00M2RjLWJkZGUtNjBiM2QyMTRhYzdm" == token {
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

	return []entities.Transfer{}, errors.New("given token is invalid")
}

func (uc TransferUsecaseMock) CreateTransfers(accountOriginID uuid.UUID, _ entities.Token, token string, accountOrigin entities.Account, accountDestination entities.Account, amount int, accountDestinationIdUUID uuid.UUID) (error, uuid.UUID, entities.Account, entities.Account) {
	if amount <= 0 {
		return errors.New("given amount is invalid"), uuid.UUID{}, entities.Account{}, entities.Account{}
	}
	if "MDIvMDgvMjAyMSAwOToyNzo0NDo2YjE5NDFkYi1jZTE3LTRmZmUtYTdlZC0yMjQ5M2E5MjZiYmM6YmQxODIxZTQtM2I5YS00M2RjLWJkZGUtNjBiM2QyMTRhYzdm" != token {
		return errors.New("given token is invalid"), uuid.UUID{}, entities.Account{}, entities.Account{}
	}
	if accountOriginID == accountDestinationIdUUID {
		return errors.New("given account is the same as the account destination"), uuid.UUID{}, entities.Account{}, entities.Account{}
	}
	if accountOrigin.Balance < amount {
		return errors.New("given account without balance"), uuid.UUID{}, entities.Account{}, entities.Account{}
	}
	if (accountDestination == entities.Account{}) {
		return errors.New("given account destination id is invalid"), uuid.UUID{}, entities.Account{}, entities.Account{}
	}
	return nil, uuid.MustParse("c5424440-4737-4e03-86d2-3adac90ddd20"), accountOrigin, accountDestination
}

type TokenUseCaseMock struct {
	AccountList AccountUsecaseMock
}

func (uc TokenUseCaseMock) AuthenticatedLogin(secret string, account entities.Account) (error, string) {
	secretHash := domain2.CreateHash(secret)
	if account == (entities.Account{}) {
		return errors.New("given secret or CPF are incorrect"), ""
	}
	if account.CPF != account.CPF {
		return errors.New("given secret or CPF are incorrect"), ""
	}
	if account.Secret != secretHash {
		return errors.New("given secret or CPF are incorrect"), ""
	}
	return nil, "passou"
}

func (uc TokenUseCaseMock) GetTokenID(id uuid.UUID) (entities.Token, error) {
	time1, _ := time.Parse("2006-01-02T15:04:05.999999999Z07:00", "2021-08-02T09:27:44.933365Z")
	if id == uuid.MustParse("39a70a94-a82d-4db8-87ae-bd900c6a7c04") {
		return entities.Token{
			ID:        uuid.MustParse("39a70a94-a82d-4db8-87ae-bd900c6a7c04"),
			IdAccount: uuid.MustParse("6b1941db-ce17-4ffe-a7ed-22493a926bbc"),
			CreatedAt: time1,
		}, nil
	}
	return entities.Token{}, nil
}

type AccountUsecaseMock struct {
}

func (uc AccountUsecaseMock) CreateAccount(name string, cpf string, _ string, balance int) (uuid.UUID, error) {
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
			ID:        uuid.MustParse("a61227cf-a857-4bc6-8fcd-ad97cdad382a"),
			Name:      "Rafael",
			CPF:       "38453162093",
			Secret:    "53b9e9679a8ea25880376080b76f98ad",
			Balance:   6000,
			CreatedAt: time1,
		},
		{
			ID:        uuid.MustParse("6b1941db-ce17-4ffe-a7ed-22493a926bbc"),
			Name:      "Lucas",
			CPF:       "08131391043",
			Secret:    "c74af74c69d81831a5703aefe9cb4199",
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

func (uc AccountUsecaseMock) GetAccountCPF(cpf string) (entities.Account, error) {
	account := entities.Account{}
	accounts, _ := uc.GetAccounts()
	for _, a := range accounts {
		if a.CPF == cpf {
			account = a
		}
	}

	return account, nil
}
