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

func TestLoginHandler(t *testing.T) {

	dataBaseAccount, clenDataBaseAccount := createTemporaryFileAccount(t, `[{"id":19727887,"name":"Rafael","cpf":"38453162093","secret":"53b9e9679a8ea25880376080b76f98ad","balance":6000,"created_at":"06/01/2020"},{"id":98498081,"name":"Lucas","cpf":"08131391043","secret":"c74af74c69d81831a5703aefe9cb4199","balance":5000,"created_at":"06/01/2020"}]`)
	defer clenDataBaseAccount()
	accountStorage := store_account.NewStoredAccount(dataBaseAccount)
	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{TimestampFormat: time.RFC3339})
	Lentry := logrus.NewEntry(logger)
	LoginUseCase := &TokenUseCaseMock{AccountList: accountStorage}
	S := new(ServerAccount)
	S.login = LoginUseCase
	S.logger = Lentry
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
			body:         `{"cpf": "38453162793", "Secret": "jax"}`,
			response:     http.StatusUnauthorized,
			responsebody: `{"errors":"given secret or CPF are incorrect"}` + "\n",
		},
		{
			name:         "should unsuccessfully create an account when CPF is invalid",
			method:       "POST",
			path:         "/login",
			body:         `{"cpf": "384531.62793", "Secret": "jax"}`,
			response:     http.StatusUnauthorized,
			responsebody: `{"errors":"given secret or CPF are incorrect"}` + "\n",
		},
		{
			name:         "should unsuccessfully authenticated login when secret is not correct",
			method:       "POST",
			path:         "/login",
			body:         `{"cpf": "081.313.910-43", "Secret": "call"}`,
			response:     http.StatusUnauthorized,
			responsebody: `{"errors":"given secret or CPF are incorrect"}` + "\n",
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
			bodyBytes := []byte(tc.body)
			request, _ := http.NewRequest(tc.method, tc.path, bytes.NewReader(bodyBytes))
			responseRecorder := httptest.NewRecorder()

			S.processLogin(responseRecorder, request)

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
	AccountList *store_account.StoredAccount
}

func (uc TokenUseCaseMock) AuthenticatedLogin(cpf, secret string) (error, string) {
	secretHash := domain.CreateHash(secret)
	account := domain.Account{}
	for _, a := range uc.AccountList.ReturnAccounts() {
		if a.CPF == cpf {
			account = account2.ChangeAccountStorage(a)
		}
	}
	if len(cpf) != 11 && len(cpf) != 14 {
		return errors.New("given secret or CPF are incorrect"), ""
	}
	if len(cpf) == 14 {
		if string([]rune(cpf)[3]) == "." && string([]rune(cpf)[7]) == "." && string([]rune(cpf)[11]) == "-" {
			if account.CPF != cpf {
				return errors.New("given secret or CPF are incorrect"), ""
			}
			if account.Secret != secret {
				return errors.New("given secret or CPF are incorrect"), ""
			}
			if account == (domain.Account{}) {
				return errors.New("given secret or CPF are incorrect"), ""
			}
			return nil, "passou"
		} else {
			return errors.New("given secret or CPF are incorrect"), ""
		}
	}
	if account == (domain.Account{}) {
		return errors.New("given secret or CPF are incorrect"), ""
	}
	if account.CPF != cpf {
		return errors.New("given secret or CPF are incorrect"), ""
	}
	if account.Secret != secretHash {
		return errors.New("given secret or CPF are incorrect"), ""
	}
	return nil, "passou"
}

func (uc TokenUseCaseMock) ReturnToken(_ int) string {
	return ""
}

func (uc TokenUseCaseMock) GetTokenID(_ int) domain.Token {
	return domain.Token{}
}
