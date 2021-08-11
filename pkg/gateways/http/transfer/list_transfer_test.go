package transfer

import (
	"bytes"
	"github.com/sirupsen/logrus"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestHandler_ListTransfers(t *testing.T) {
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
			name:         "should successfully get transfers",
			method:       "GET",
			path:         "/transfers",
			response:     http.StatusOK,
			token:        "MDIvMDgvMjAyMSAwOToyNzo0NDo2YjE5NDFkYi1jZTE3LTRmZmUtYTdlZC0yMjQ5M2E5MjZiYmM6YmQxODIxZTQtM2I5YS00M2RjLWJkZGUtNjBiM2QyMTRhYzdm",
			responsebody: `{"transfers":[{"id":"47399f23-2093-4dde-b32f-990cac27630e","origin_account_id":"6b1941db-ce17-4ffe-a7ed-22493a926bbc","destination_account_id":"a61227cf-a857-4bc6-8fcd-ad97cdad382a","amount":150,"created_at":"2021-07-20T15:17:25.933365Z"}]}` + "\n",
		},
		{
			name:         "should unsuccessfully get transfer when there is wrong token",
			method:       "GET",
			path:         "/transfers",
			response:     http.StatusUnauthorized,
			token:        "MDIvMDgvMjAyMSAwOToyNzo0NDo2YjE5NDFkYi1jZTE3LTRmZmUtYTdlZC0yMeQ5M2E5MjZiYmM6YmQxODIxZTQtM2I5YS00M2RjLWJkZGUtNjBiM2QyMTRhYzdm",
			responsebody: `{"errors":"given token is invalid"}` + "\n",
		},
	}
	for _, tc := range gettransfer {
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
			respondRecorder := httptest.NewRecorder()

			request.Header.Add("Authorization", tc.token)

			s.ListTransfers(respondRecorder, request)

			if tc.response != respondRecorder.Code { // O teste falhará pois não queremos erro e obtivemos um
				t.Errorf("unexpected error, wantErr= %d; gotErr= %d", tc.response, respondRecorder.Code)
			}

			if respondRecorder.Body.String() != tc.responsebody && tc.responsebody != "" {
				t.Errorf("expected an %s but got %s", tc.responsebody, respondRecorder.Body.String())
			}
		})
	}
}
