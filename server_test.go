package main

import (
	"github.com/CMedrado/DesafioStone/domain"
	"github.com/CMedrado/DesafioStone/https"
	"github.com/CMedrado/DesafioStone/store"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type TestCreatedRequest struct {
	Name    string `json:"name"`
	CPF     string `json:"cpf"`
	Secret  string `json:"secret"`
	Balance uint   `json:"balance"`
}

func TestNewServerAccount(t *testing.T) {

	tt := []struct {
		name       string
		method     string
		in         store.Account
		body       string
		statusCode int
		wantErr    bool
		want       uint
	}{
		{
			name:   "Should successfully create an account with formatted CPF",
			method: http.MethodPost,
			in: store.Account{
				Name:    "Rafael",
				CPF:     "081.313.910-43",
				Secret:  "tatatal",
				Balance: 5000,
			},
			body:       `[{"name": "Rafael", "cpf": "081.313.910-43", "secret": "tatatal", "balance": 5000}`,
			statusCode: http.StatusAccepted,
			wantErr:    false,
		},
		{
			name:   "should successfully create an account with unformatted CPF",
			method: http.MethodPost,
			in: store.Account{
				Name:    "Rafael",
				CPF:     "08131391043",
				Secret:  "tatatal",
				Balance: 5000,
			},
			body:       `[{"name": "Rafael", "cpf": "08131391043", "secret": "tatatal", "balance": 5000}`,
			statusCode: http.StatusAccepted,
			wantErr:    false,
		},
		{
			name:   "should unsuccessfully create an account when CPF is invalid",
			method: http.MethodPost,
			in: store.Account{
				Name:    "Rafael",
				CPF:     "08131.391043",
				Secret:  "tatatal",
				Balance: 5000,
			},
			body:       `[{"name": "Rafael", "cpf": "08131.391043, "secret": "tatatal", "balance": 5000}`,
			statusCode: http.StatusUnauthorized,
			wantErr:    true,
		},
	}
	for _, tc := range tt {
		request, _ := http.NewRequest("POST", "/accounts/", strings.NewReader(tc.body))
		respondeRecorder := httptest.NewRecorder()

		accountStorage := store.NewStoredAccount()
		armazenamento := domain.AccountUsecase{Store: accountStorage}
		servidor := https.NewServerAccount(&armazenamento)
		servidor.ServeHTTP(respondeRecorder, request)

		if tc.statusCode != respondeRecorder.Code && tc.wantErr { // O teste falhará pois não queremos erro e obtivemos um
			t.Errorf("unexpected error, wantErr= %d; gotErr= %d", tc.statusCode, respondeRecorder.Code)
		}

		if strings.TrimSpace(respondeRecorder.Body.String()) == "0" {
			t.Errorf("expected an ID but got %s", respondeRecorder.Body.String())
		}
	}

	//req, _ = http.NewRequest(http.MethodGet, fmt.Sprintf("/accounts"), nil)
	//servidor.ServeHTTP(httptest.NewRecorder(), req)
	//req, _ = http.NewRequest(http.MethodPost, fmt.Sprintf("/accounts/"), nil)
	//servidor.ServeHTTP(httptest.NewRecorder(), req)
	//req, _ = http.NewRequest(http.MethodPost, fmt.Sprintf("/login"), nil)
	//servidor.ServeHTTP(httptest.NewRecorder(), req)
	//req, _ = http.NewRequest(http.MethodGet, fmt.Sprintf("/transfers"), nil)
	//servidor.ServeHTTP(httptest.NewRecorder(), req)
	//req, _ = http.NewRequest(http.MethodPost, fmt.Sprintf("/transfers"), nil)
	//servidor.ServeHTTP(httptest.NewRecorder(), req)

}
