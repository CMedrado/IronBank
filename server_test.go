package main

import (
	"bytes"
	"github.com/CMedrado/DesafioStone/domain"
	"github.com/CMedrado/DesafioStone/https"
	"github.com/CMedrado/DesafioStone/store"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

func TestNewServerAccount(t *testing.T) {

	accountTransfer := store.NewStoredTransferTwo()
	accountToken := store.NewStoredToked()
	accountLogin := store.NewStoredLogin()
	accountStorage := store.NewStoredAccount()
	armazenamento := domain.AccountUsecase{accountStorage, accountLogin, accountToken, accountTransfer}
	servidor := https.NewServerAccount(&armazenamento)

	tt := []struct {
		name         string
		method       string
		path         string
		body         string
		response     int
		responsebody string
	}{
		{
			name:     "Should successfully create an account with formatted CPF",
			method:   "POST",
			path:     "/accounts/",
			body:     `{"name": "Rafael", "cpf": "081.313.910-43", "secret": "tatatal", "balance": 5000}`,
			response: http.StatusAccepted,
		},
		{
			name:     "should successfully create an account with unformatted CPF",
			method:   "POST",
			path:     "/accounts/",
			body:     `{"name": "Lucas", "cpf": "38453162093", "secret": "jax", "balance": 3000}`,
			response: http.StatusAccepted,
		},
		{
			name:     "should unsuccessfully create an account when CPF is invalid",
			method:   "POST",
			path:     "/accounts/",
			body:     `{"name": "Rafael", "cpf": "08131.391043", "secret": "tatatal", "balance": 5000}`,
			response: http.StatusUnauthorized,
		},
		{
			name:     "should successfully get accounts",
			method:   "GET",
			path:     "/accounts",
			response: http.StatusAccepted,
		},
		{
			name:         "should successfully get balance with formatted CPF",
			method:       "GET",
			path:         "/accounts/081.313.910-43/balance",
			response:     http.StatusAccepted,
			responsebody: `5000`,
		},
		{
			name:         "should successfully get balance with unformatted CPF",
			method:       "GET",
			path:         "/accounts/38453162093/balance",
			response:     http.StatusAccepted,
			responsebody: `6000`,
		},
		{
			name:     "should unsuccessfully get balance when CPF is invalid",
			method:   "GET",
			path:     "/accounts/384.53162-093/balance",
			response: http.StatusBadRequest,
		},
		{
			name:     "should unsuccessfully get balance when dont exist account",
			method:   "GET",
			path:     "/accounts/062.136.280-37/balance",
			response: http.StatusBadRequest,
		},
		{
			name:     "should successfully authenticated login with formatted CPF",
			method:   "POST",
			path:     "/login",
			body:     `{"cpf:": "08131391043", "Secret": "tatatal"}`,
			response: http.StatusAccepted,
		},
		{
			name:     "should successfully authenticated login with unformatted CPF",
			method:   "POST",
			path:     "/login",
			body:     `{"cpf:": "38453162093", "Secret": "jax"}`,
			response: http.StatusAccepted,
		},
		{
			name:     "should unsuccessfully authenticated login when CPF is invalid",
			method:   "POST",
			path:     "/login",
			body:     `{"cpf:": "384531620.93", "Secret": "jax"}`,
			response: http.StatusNotAcceptable,
		},
		{
			name:     "should unsuccessfully authenticated login when cpf is not registered",
			method:   "POST",
			path:     "/login",
			body:     `{"cpf:": "38453162793", "Secret": "jax"}`,
			response: http.StatusNotAcceptable,
		},
		{
			name:     "should unsuccessfully authenticated login when secret is not correct",
			method:   "POST",
			path:     "/login",
			body:     `{"cpf:": "081.313.910-43", "Secret": "call"}`,
			response: http.StatusNotAcceptable,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			bodyBytes := []byte(tc.body)
			request, _ := http.NewRequest(tc.method, tc.path, bytes.NewReader(bodyBytes))
			responseRecorder := httptest.NewRecorder()

			servidor.ServeHTTP(responseRecorder, request)

			if tc.response != responseRecorder.Code { // O teste falhará pois não queremos erro e obtivemos um
				t.Errorf("unexpected error, wantErr= %d; gotErr= %d", tc.response, responseRecorder.Code)
			}

			if strings.TrimSpace(responseRecorder.Body.String()) == "0" {
				t.Errorf("expected an ID but got %s", responseRecorder.Body.String())
			}
		})
	}

	firstID := armazenamento.Store.VoltaCPF("38453162093")
	secondID := armazenamento.Store.VoltaCPF("08131391043")
	firstToken := armazenamento.Token.VoltaToken(firstID)
	firstIDSTring := strconv.Itoa(firstID)
	secondIDString := strconv.Itoa(secondID)
	firstTokenString := strconv.Itoa(firstToken)

	secondtt := []struct {
		name         string
		method       string
		path         string
		body         string
		response     int
		responsebody string
	}{
		{
			name:     "should successfully transfer amount",
			method:   "POST",
			path:     "/transfers",
			body:     `"accountoriginid":` + firstIDSTring + `"token":` + firstTokenString + `"accountdestinationid":` + secondIDString + `"amount": 500`,
			response: http.StatusAccepted,
		},
		{
			name:     "should successfully transfer amount",
			method:   "POST",
			path:     "/transfers",
			body:     `"accountoriginid":` + firstIDSTring + `"token":` + firstTokenString + `"accountdestinationid":` + secondIDString + `"amount": 300`,
			response: http.StatusAccepted,
		},
		{
			name:     "should unsuccessfully transfer amount when there is no account origin id",
			method:   "POST",
			path:     "/transfers",
			body:     `"accountoriginid": "98498761", "token":` + firstTokenString + `"accountdestinationid":` + secondIDString + `"amount": 300`,
			response: http.StatusNotAcceptable,
		},
		{
			name:     "should unsuccessfully transfer amount when there is wrong token",
			method:   "POST",
			path:     "/transfers",
			body:     `"accountoriginid":` + firstIDSTring + `"token": 27131807 ,"accountdestinationid":` + secondIDString + `"amount": 300`,
			response: http.StatusUnauthorized,
		},
		{
			name:     "should unsuccessfully transfer amount when there is no account destination id",
			method:   "POST",
			path:     "/transfers",
			body:     `"accountoriginid":` + firstIDSTring + `"token":` + firstTokenString + `"accountdestinationid":19727807, "amount": 300`,
			response: http.StatusNotAcceptable,
		},
		{
			name:     "should unsuccessfully transfer amount when the amount is too slow",
			method:   "POST",
			path:     "/transfers",
			body:     `"accountoriginid":` + firstIDSTring + `"token":` + firstTokenString + `"accountdestinationid":` + secondIDString + `"amount": 0`,
			response: http.StatusPaymentRequired,
		},
		{
			name:     "should unsuccessfully transfer amount when the amount is greater than the balance",
			method:   "POST",
			path:     "/transfers",
			body:     `"accountoriginid":` + firstIDSTring + `"token":` + firstTokenString + `"accountdestinationid":` + secondIDString + `"amount": 9000`,
			response: http.StatusPaymentRequired,
		},
		{
			name:   "should successfully get transfers",
			method: "GET",
			path:   "/transfers",
			body:   `"accountoriginid":` + firstIDSTring + `"token":` + firstTokenString,
		},
		{
			name:     "should unsuccessfully get transfers when there is no account origin id",
			method:   "GET",
			path:     "/transfers",
			body:     `"accountoriginid":98498981, "token":` + firstTokenString,
			response: http.StatusNotAcceptable,
		},
		{
			name:     "should unsuccessfully get transfer when there is wrong token",
			method:   "GET",
			path:     "/transfers",
			body:     `"accountoriginid":` + firstIDSTring + `"token":27131047`,
			response: http.StatusUnauthorized,
		},
	}
	for _, tc := range secondtt {
		t.Run(tc.name, func(t *testing.T) {
			bodyBytes := []byte(tc.body)
			request, _ := http.NewRequest(tc.method, tc.path, bytes.NewReader(bodyBytes))
			respondeRecorder := httptest.NewRecorder()

			servidor.ServeHTTP(respondeRecorder, request)

			if tc.response != respondeRecorder.Code { // O teste falhará pois não queremos erro e obtivemos um
				t.Errorf("unexpected error, wantErr= %d; gotErr= %d", tc.response, respondeRecorder.Code)
			}

			if strings.TrimSpace(respondeRecorder.Body.String()) == "0" {
				t.Errorf("expected an ID but got %s", respondeRecorder.Body.String())
			}
		})
	}

}
