package https

import (
	"bytes"
	"errors"
	"github.com/CMedrado/DesafioStone/domain"
	"github.com/CMedrado/DesafioStone/store"
	"github.com/sirupsen/logrus"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestLoginHandler(t *testing.T) {

	Account1 := store.Account{ID: 981, Name: "Rafael", CPF: "38453162093", Secret: domain.CreateHash("call"), Balance: 6000, CreatedAt: "06/01/2020"}
	Account2 := store.Account{ID: 982, Name: "Lucas", CPF: "08131391043", Secret: domain.CreateHash("lixo"), Balance: 5000, CreatedAt: "06/01/2020"}
	AccountStorage := make(map[string]store.Account)
	AccountStorage[Account1.CPF] = Account1
	AccountStorage[Account2.CPF] = Account2
	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{TimestampFormat: time.RFC3339})
	Lentry := logrus.NewEntry(logger)
	LoginUseCase := &TokenUseCaseMock{AccountList: AccountStorage}
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
	AccountList map[string]store.Account
}

func (uc TokenUseCaseMock) AuthenticatedLogin(cpf, secret string) (error, string) {
	secretHash := domain.CreateHash(secret)
	if len(cpf) != 11 && len(cpf) != 14 {
		return errors.New("given secret or CPF are incorrect"), ""
	}
	if len(cpf) == 14 {
		if string([]rune(cpf)[3]) == "." && string([]rune(cpf)[7]) == "." && string([]rune(cpf)[11]) == "-" {
			if uc.AccountList[cpf].CPF != cpf {
				return errors.New("given secret or CPF are incorrect"), ""
			}
			if uc.AccountList[cpf].Secret != secret {
				return errors.New("given secret or CPF are incorrect"), ""
			}
			if uc.AccountList[cpf] == (store.Account{}) {
				return errors.New("given secret or CPF are incorrect"), ""
			}
			return nil, "passou"
		} else {
			return errors.New("given secret or CPF are incorrect"), ""
		}
	}
	if uc.AccountList[cpf] == (store.Account{}) {
		return errors.New("given secret or CPF are incorrect"), ""
	}
	if uc.AccountList[cpf].CPF != cpf {
		return errors.New("given secret or CPF are incorrect"), ""
	}
	if uc.AccountList[cpf].Secret != secretHash {
		return errors.New("given secret or CPF are incorrect"), ""
	}
	return nil, "passou"
}

func (uc TokenUseCaseMock) ReturnToken(_ int) string {
	return ""
}

func (uc TokenUseCaseMock) GetTokenID(_ int) store.Token {
	return store.Token{}
}
