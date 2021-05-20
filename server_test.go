package main

import (
	"bytes"
	"encoding/base64"
	"github.com/CMedrado/DesafioStone/domain"
	"github.com/CMedrado/DesafioStone/https"
	"github.com/CMedrado/DesafioStone/store"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestNewServerAccount(t *testing.T) { // Fazer

	accountTransfer := store.NewStoredTransferAccountID()
	accountToken := store.NewStoredToked()
	accountLogin := store.NewStoredLogin()
	accountStorage := store.NewStoredAccount()
	armazenamento := domain.AccountUseCase{Store: accountStorage, Login: accountLogin, Token: accountToken, Transfer: accountTransfer}
	servidor := https.NewServerAccount(&armazenamento)

	createt := []struct {
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
			response:     http.StatusAccepted,
			responsebody: `{"id":98498081}` + "\n",
		},
		{
			name:         "should successfully create an account with unformatted CPF",
			method:       "POST",
			path:         "/accounts",
			body:         `{"name": "Lucas", "cpf": "38453162093", "secret": "jax", "balance": 3000}`,
			response:     http.StatusAccepted,
			responsebody: `{"id":19727887}` + "\n",
		},
		{
			name:     "should unsuccessfully create an account when CPF is invalid",
			method:   "POST",
			path:     "/accounts",
			body:     `{"name": "Rafael", "cpf": "08131.391043", "secret": "tatatal", "balance": 5000}`,
			response: http.StatusUnauthorized,
		},
	}
	for _, tc := range createt {
		t.Run(tc.name, func(t *testing.T) {
			bodyBytes := []byte(tc.body)
			request, _ := http.NewRequest(tc.method, tc.path, bytes.NewReader(bodyBytes))
			responseRecorder := httptest.NewRecorder()

			servidor.ServeHTTP(responseRecorder, request)

			if tc.response != responseRecorder.Code {
				t.Errorf("unexpected error, wantErr= %d; gotErr= %d", tc.response, responseRecorder.Code)
			}

			if responseRecorder.Body.String() != tc.responsebody {
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
			response:     http.StatusAccepted,
			responsebody: `[{"id":98498081,"name":"Rafael","cpf":"08131391043","secret":"3467e121a1a109628e0a5b0cebba361b","balance":5000,"created_at":"19/05/2021 11:11:40"},{"id":19727887,"name":"Lucas","cpf":"38453162093","secret":"7e65a9b554bbc9817aa049ce38c84a72","balance":3000,"created_at":"19/05/2021 10:10:12"}]` + "\n",
		},
	}
	for _, tc := range accountst {
		t.Run(tc.name, func(t *testing.T) {
			bodyBytes := []byte(tc.body)
			request, _ := http.NewRequest(tc.method, tc.path, bytes.NewReader(bodyBytes))
			responseRecorder := httptest.NewRecorder()

			servidor.ServeHTTP(responseRecorder, request)

			if tc.response != responseRecorder.Code { // O teste falhará pois não queremos erro e obtivemos um
				t.Errorf("unexpected error, wantErr= %d; gotErr= %d", tc.response, responseRecorder.Code)
			}

			//if responseRecorder.Body.String() != tc.responsebody {
			//	t.Errorf("expected an %s but got %s", tc.responsebody, responseRecorder.Body.String())
			//}
		})
	}
	firstID := armazenamento.Store.ReturnCPF("38453162093")
	secondID := armazenamento.Store.ReturnCPF("08131391043")
	firstIDString := strconv.Itoa(firstID)
	secondIDString := strconv.Itoa(secondID)
	balancet := []struct {
		name         string
		method       string
		path         string
		body         string
		response     int
		responsebody string
	}{
		{
			name:         "should successfully get balance with formatted CPF",
			method:       "GET",
			path:         "/accounts/" + firstIDString + "/balance",
			response:     http.StatusAccepted,
			responsebody: `{"balance":3000}` + "\n",
		},
		{
			name:         "should successfully get balance with unformatted CPF",
			method:       "GET",
			path:         "/accounts/" + secondIDString + "/balance",
			response:     http.StatusAccepted,
			responsebody: `{"balance":5000}` + "\n",
		},
		{
			name:     "should unsuccessfully get balance when CPF is invalid",
			method:   "GET",
			path:     "/accounts/3848/balance",
			response: http.StatusBadRequest,
		},
		{
			name:     "should unsuccessfully get balance when dont exist account",
			method:   "GET",
			path:     "/accounts/398-6/balance",
			response: http.StatusBadRequest,
		},
	}
	for _, tc := range balancet {
		t.Run(tc.name, func(t *testing.T) {
			bodyBytes := []byte(tc.body)
			request, _ := http.NewRequest(tc.method, tc.path, bytes.NewReader(bodyBytes))
			responseRecorder := httptest.NewRecorder()

			servidor.ServeHTTP(responseRecorder, request)

			if tc.response != responseRecorder.Code { // O teste falhará pois não queremos erro e obtivemos um
				t.Errorf("unexpected error, wantErr= %d; gotErr= %d", tc.response, responseRecorder.Code)
			}

			if responseRecorder.Body.String() != tc.responsebody {
				t.Errorf("expected an %s but got %s", tc.responsebody, responseRecorder.Body.String())
			}
		})
	}
	logint := []struct {
		name         string
		method       string
		path         string
		body         string
		response     int
		responsebody string
	}{
		{
			name:     "should successfully authenticated login with formatted CPF",
			method:   "POST",
			path:     "/login",
			body:     `{"cpf": "08131391043", "Secret": "tatatal"}`,
			response: http.StatusAccepted,
		},
		{
			name:     "should successfully authenticated login with unformatted CPF",
			method:   "POST",
			path:     "/login",
			body:     `{"cpf": "38453162093", "Secret": "jax"}`,
			response: http.StatusAccepted,
		},
		{
			name:     "should unsuccessfully authenticated login when CPF is invalid",
			method:   "POST",
			path:     "/login",
			body:     `{"cpf": "384531620.93", "Secret": "jax"}`,
			response: http.StatusNotAcceptable,
		},
		{
			name:     "should unsuccessfully authenticated login when cpf is not registered",
			method:   "POST",
			path:     "/login",
			body:     `{"cpf": "38453162793", "Secret": "jax"}`,
			response: http.StatusNotAcceptable,
		},
		{
			name:     "should unsuccessfully authenticated login when secret is not correct",
			method:   "POST",
			path:     "/login",
			body:     `{"cpf": "081.313.910-43", "Secret": "call"}`,
			response: http.StatusUnauthorized,
		},
	}
	for _, tc := range logint {
		t.Run(tc.name, func(t *testing.T) {
			bodyBytes := []byte(tc.body)
			request, _ := http.NewRequest(tc.method, tc.path, bytes.NewReader(bodyBytes))
			responseRecorder := httptest.NewRecorder()

			servidor.ServeHTTP(responseRecorder, request)

			if tc.response != responseRecorder.Code {
				t.Errorf("unexpected error, wantErr= %d; gotErr= %d", tc.response, responseRecorder.Code)
			}

			//if responseRecorder.Body.String() != tc.body {
			//	t.Errorf("expected an %s but got %s", responseRecorder.Body.String(), tc.body)
			//}
		})
	}
	firstToken := armazenamento.Token.ReturnToken(firstID)
	secondToken := "03/02/2000 03:05:55:" + strconv.Itoa(secondID)
	secondTokenD := base64.StdEncoding.EncodeToString([]byte(secondToken))
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
			name:     "should successfully transfer amount",
			method:   "POST",
			path:     "/transfers",
			body:     `{"account_destination_id":` + secondIDString + `,"amount": 500}`,
			response: http.StatusAccepted,
			token:    firstToken,
		},
		{
			name:     "should successfully transfer amount",
			method:   "POST",
			path:     "/transfers",
			body:     `{"account_destination_id":` + secondIDString + `,"amount": 300}`,
			response: http.StatusAccepted,
			token:    firstToken,
		},
		{
			name:     "should unsuccessfully transfer amount when there is wrong token",
			method:   "POST",
			path:     "/transfers",
			body:     `{"account_destination_id":` + secondIDString + `,"amount": 300}`,
			response: http.StatusUnauthorized,
			token:    secondTokenD,
		},
	}
	for _, tc := range createtransfer {
		t.Run(tc.name, func(t *testing.T) {
			bodyBytes := []byte(tc.body)
			request, _ := http.NewRequest(tc.method, tc.path, bytes.NewReader(bodyBytes))
			respondeRecorder := httptest.NewRecorder()

			request.Header.Add("Authorization", tc.token)

			servidor.ServeHTTP(respondeRecorder, request)

			if tc.response != respondeRecorder.Code { // O teste falhará pois não queremos erro e obtivemos um
				t.Errorf("unexpected error, wantErr= %d; gotErr= %d", tc.response, respondeRecorder.Code)
			}
		})
	}
	gettransfer := []struct {
		name         string
		method       string
		path         string
		body         string
		response     int
		responsebody string
		token        string
	}{
		{
			name:     "should successfully get transfers",
			method:   "GET",
			path:     "/transfers",
			response: http.StatusAccepted,
			token:    firstToken,
		},
		{
			name:     "should unsuccessfully get transfer when there is wrong token",
			method:   "GET",
			path:     "/transfers",
			response: http.StatusUnauthorized,
			token:    secondTokenD,
		},
	}
	for _, tc := range gettransfer {
		t.Run(tc.name, func(t *testing.T) {
			bodyBytes := []byte(tc.body)
			request, _ := http.NewRequest(tc.method, tc.path, bytes.NewReader(bodyBytes))
			respondRecorder := httptest.NewRecorder()

			request.Header.Add("Authorization", tc.token)

			servidor.ServeHTTP(respondRecorder, request)

			if tc.response != respondRecorder.Code { // O teste falhará pois não queremos erro e obtivemos um
				t.Errorf("unexpected error, wantErr= %d; gotErr= %d", tc.response, respondRecorder.Code)
			}
		})
	}
}
