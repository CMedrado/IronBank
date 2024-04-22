package transfer

import (
	"bytes"
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/google/uuid"

	"github.com/CMedrado/DesafioStone/pkg/domain"
	"github.com/CMedrado/DesafioStone/pkg/domain/entities"
	"github.com/CMedrado/DesafioStone/pkg/domain/transfer"
)

var (
	secret = "MDIvMDgvMjAyMSAwOToyNzo0NDo2YjE5NDFkYi1jZTE3LTRmZmUtYTdlZC0yMjQ5M2E5MjZiYmM6YmQxODIxZTQtM2I5YS00M2RjLWJkZGUtNjBiM2QyMTRhYzdm" // #nosec
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

func (uc *TransferUsecaseMock) GetTransfers(token string) ([]entities.Transfer, error) {
	time1, _ := time.Parse("2006-01-02T15:04:05.999999999Z07:00", "2021-07-20T15:17:25.933365Z")
	if secret == token {
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

	if token == "MDIvMDgvMjAyMSAwOToyNzo0NDo2YjE5NDFkYi1jZTE3LTRmZmUtYTdlZC0yMeQ5M2E5MjZiYmM6YmQxODIxZTQtM2I5YS00M2RjLWJkZGUtNjBiM2QyMTRhYzdm" { // #nosec
		return []entities.Transfer{}, domain.ErrInvalidToken
	}

	return []entities.Transfer{}, errors.New("given token is invalid")
}

func (uc TransferUsecaseMock) CreateTransfers(_ context.Context, token string, amount int, accountDestinationIdUUID string) (error, uuid.UUID) {
	if amount <= 0 {
		return errors.New("given amount is invalid"), uuid.UUID{}
	}
	if secret != token {
		return domain.ErrInvalidToken, uuid.UUID{}
	}
	if "6b1941db-ce17-4ffe-a7ed-22493a926bbc" == accountDestinationIdUUID {
		return errors.New("given account is the same as the account destination"), uuid.UUID{}
	}
	if 6000 < amount {
		return errors.New("given account without balance"), uuid.UUID{}
	}
	if "75432539-c5ba-46d3-9690-44985b516da5" == accountDestinationIdUUID {
		return transfer.ErrInvalidDestinationID, uuid.UUID{}
	}
	return nil, uuid.MustParse("c5424440-4737-4e03-86d2-3adac90ddd20")
}

func (uc *TransferUsecaseMock) GetCountTransfer(_ context.Context) (int64, error) {
	return 10, nil
}

func (uc *TransferUsecaseMock) GetRankTransfer(_ context.Context) ([]string, error) {
	return []string{}, nil
}
