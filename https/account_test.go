package https

import (
	"bytes"
	"errors"
	"github.com/CMedrado/DesafioStone/domain"
	account2 "github.com/CMedrado/DesafioStone/domain/account"
	store_account "github.com/CMedrado/DesafioStone/storage/file/account"
	"github.com/sirupsen/logrus"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

var Aux = 0

func TestAccountHandler(t *testing.T) {
	dataBaseAccount, clenDataBaseAccount := createTemporaryFileAccount(t, `[{"id":19727887,"name":"Rafael","cpf":"38453162093","secret":"53b9e9679a8ea25880376080b76f98ad","balance":6000,"created_at":"06/01/2020"},{"id":98498081,"name":"Lucas","cpf":"08131391043","secret":"c74af74c69d81831a5703aefe9cb4199","balance":5000,"created_at":"06/01/2020"}]`)
	defer clenDataBaseAccount()
	accountStorage := store_account.NewStoredAccount(dataBaseAccount)
	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{TimestampFormat: time.RFC3339})
	Lentry := logrus.NewEntry(logger)
	AccountUsecase := &AccountUsecaseMock{AccountList: accountStorage}
	S := new(ServerAccount)
	S.account = AccountUsecase
	S.logger = Lentry

	tt := []struct {
		name         string
		method       string
		path         string
		body         string
		response     int
		responsebody string
	}{
		{
			name:         "Should successfully create an account with formatted CPF",
			method:       "POST",
			path:         "/accounts",
			body:         `{"name": "Rafael", "cpf": "081.313.910-43", "secret": "tatatal", "balance": 5000}`,
			responsebody: `{"id":981}` + "\n",
			response:     http.StatusCreated,
		},
		{
			name:         "should successfully create an account with unformatted CPF",
			method:       "POST",
			path:         "/accounts",
			body:         `{"name": "Lucas", "cpf": "38453162093", "secret": "jax", "balance": 3000}`,
			responsebody: `{"id":982}` + "\n",
			response:     http.StatusCreated,
		},
		{
			name:         "should unsuccessfully create an account when CPF is invalid",
			method:       "POST",
			path:         "/accounts",
			body:         `{"name": "Rafael", "cpf": "08131.391043", "secret": "tatatal", "balance": 5000}`,
			response:     http.StatusNotAcceptable,
			responsebody: `{"errors":"given cpf is invalid"}` + "\n",
		},
		{
			name:         "should unsuccessfully create an account when CPF is invalid",
			method:       "POST",
			path:         "/accounts",
			body:         `{"name": "Rafael", "cpf": "08131.391.0-43", "secret": "tatatal", "balance": 5000}`,
			response:     http.StatusNotAcceptable,
			responsebody: `{"errors":"given cpf is invalid"}` + "\n",
		},
		{
			name:         "should unsuccessfully create an account when balance is invalid",
			method:       "POST",
			path:         "/accounts",
			body:         `{"name": "Rafael", "cpf": "08131391043", "secret": "tatatal", "balance": -5}`,
			response:     http.StatusBadRequest,
			responsebody: `{"errors":"given the balance amount is invalid"}` + "\n",
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
			bodyBytes := []byte(tc.body)
			request, _ := http.NewRequest(tc.method, tc.path, bytes.NewReader(bodyBytes))
			responseRecorder := httptest.NewRecorder()

			S.processAccount(responseRecorder, request)

			if tc.response != responseRecorder.Code {
				t.Errorf("unexpected error, wantErr= %d; gotErr= %d", tc.response, responseRecorder.Code)
			}

			if responseRecorder.Body.String() != tc.responsebody && tc.responsebody != "" {
				t.Errorf("expected an %s but got %s", tc.responsebody, responseRecorder.Body.String())
			}
		})
	}
	accountst := []struct {
		name         string
		method       string
		path         string
		body         string
		response     int
		responsebody string
	}{
		{
			name:         "should successfully get accounts",
			method:       "GET",
			path:         "/accounts",
			response:     http.StatusOK,
			responsebody: `{"accounts":[{"id":19727887,"name":"Rafael","cpf":"38453162093","secret":"53b9e9679a8ea25880376080b76f98ad","balance":6000,"created_at":"06/01/2020"},{"id":98498081,"name":"Lucas","cpf":"08131391043","secret":"c74af74c69d81831a5703aefe9cb4199","balance":5000,"created_at":"06/01/2020"}]}` + "\n",
		},
	}
	for _, tc := range accountst {
		t.Run(tc.name, func(t *testing.T) {
			bodyBytes := []byte(tc.body)
			request, _ := http.NewRequest(tc.method, tc.path, bytes.NewReader(bodyBytes))
			responseRecorder := httptest.NewRecorder()

			S.handleAccounts(responseRecorder, request)

			if tc.response != responseRecorder.Code { // O teste falhará pois não queremos erro e obtivemos um
				t.Errorf("unexpected error, wantErr= %d; gotErr= %d", tc.response, responseRecorder.Code)
			}

			if responseRecorder.Body.String() != tc.responsebody {
				t.Errorf("expected an %s but got %s", tc.responsebody, responseRecorder.Body.String())
			}
		})
	}
	balancet := []struct {
		name         string
		method       string
		path         string
		body         string
		response     int
		responsebody string
	}{
		{
			name:         "should successfully get balance with ID",
			method:       "GET",
			response:     http.StatusOK,
			responsebody: `{"balance":6000}` + "\n",
		},
		{
			name:         "should unsuccessfully get balance when ID is invalid",
			method:       "GET",
			path:         "/accounts/3848/balance",
			response:     http.StatusNotAcceptable,
			responsebody: `{"errors":"given id is invalid"}` + "\n",
		},
	}
	for _, tc := range balancet {
		t.Run(tc.name, func(t *testing.T) {
			bodyBytes := []byte(tc.body)
			request, _ := http.NewRequest(tc.method, tc.path, bytes.NewReader(bodyBytes))
			responseRecorder := httptest.NewRecorder()

			S.handleBalance(responseRecorder, request)

			if tc.response != responseRecorder.Code { // O teste falhará pois não queremos erro e obtivemos um
				t.Errorf("unexpected error, wantErr= %d; gotErr= %d", tc.response, responseRecorder.Code)
			}

			if responseRecorder.Body.String() != tc.responsebody && tc.responsebody != "" {
				t.Errorf("expected an %s but got %s", tc.responsebody, responseRecorder.Body.String())
			}
		})
	}
}

type AccountUsecaseMock struct {
	AccountList *store_account.StoredAccount

	UpdateCallCount int
}

func (uc AccountUsecaseMock) CreateAccount(name string, cpf string, _ string, balance int) (int, error) {
	if len(cpf) != 11 && len(cpf) != 14 {
		return 0, errors.New("given cpf is invalid")
	}
	if len(cpf) == 14 {
		if string([]rune(cpf)[3]) == "." && string([]rune(cpf)[7]) == "." && string([]rune(cpf)[11]) == "-" {
			if balance <= 0 {
				return 0, errors.New("given the balance amount is invalid")
			}
			if name == "Rafael" {
				return 981, nil
			}
			if name == "Lucas" {
				return 982, nil
			}
		} else {
			return 0, errors.New("given cpf is invalid")
		}
	}
	if balance <= 0 {
		return 0, errors.New("given the balance amount is invalid")
	}
	if name == "Rafael" {
		return 981, nil
	}
	if name == "Lucas" {
		return 982, nil
	}
	return 1, nil
}

func (uc AccountUsecaseMock) GetBalance(id int) (int, error) {
	if Aux == 0 {
		id = 19727887
		Aux++
	} else {
		id = 0
	}
	accountOrigin := uc.SearchAccount(id)
	if (accountOrigin == domain.Account{}) {
		return 0, errors.New("given id is invalid")
	}
	return accountOrigin.Balance, nil
}

func (uc AccountUsecaseMock) GetAccounts() []domain.Account {
	var account []domain.Account

	for _, a := range uc.AccountList.ReturnAccounts() {
		account = append(account, account2.ChangeAccountStorage(a))
	}

	return account
}

func (uc AccountUsecaseMock) SearchAccount(id int) domain.Account {
	account := domain.Account{}

	for _, a := range uc.AccountList.ReturnAccounts() {
		if a.ID == id {
			account = account2.ChangeAccountStorage(a)
		}
	}

	return account
}

func (uc *AccountUsecaseMock) UpdateBalance(_ domain.Account, _ domain.Account) {
	uc.UpdateCallCount++
}

func (uc AccountUsecaseMock) GetAccountCPF(cpf string) domain.Account {
	account := domain.Account{}
	for _, a := range uc.AccountList.ReturnAccounts() {
		if a.CPF == cpf {
			account = account2.ChangeAccountStorage(a)
		}
	}

	return account
}

func (uc AccountUsecaseMock) GetAccount() []domain.Account {
	return nil
}
