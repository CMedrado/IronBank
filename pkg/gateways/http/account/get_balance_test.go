package account

import (
	"bytes"
	"github.com/sirupsen/logrus"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestHandler_GetBalance(t *testing.T) {

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
			response:     http.StatusNotFound,
			responsebody: `{"errors":"given id is invalid"}` + "\n",
		},
	}
	for _, tc := range balancet {
		t.Run(tc.name, func(t *testing.T) {
			s := new(Handler)
			s.account = &AccountUsecaseMock{}
			logger := logrus.New()
			logger.SetFormatter(&logrus.TextFormatter{TimestampFormat: time.RFC3339})
			Lentry := logrus.NewEntry(logger)
			s.logger = Lentry
			bodyBytes := []byte(tc.body)
			request, _ := http.NewRequest(tc.method, tc.path, bytes.NewReader(bodyBytes))
			responseRecorder := httptest.NewRecorder()

			s.GetBalance(responseRecorder, request)

			if tc.response != responseRecorder.Code { // O teste falhará pois não queremos erro e obtivemos um
				t.Errorf("unexpected error, wantErr= %d; gotErr= %d", tc.response, responseRecorder.Code)
			}

			if responseRecorder.Body.String() != tc.responsebody && tc.responsebody != "" {
				t.Errorf("expected an %s but got %s", tc.responsebody, responseRecorder.Body.String())
			}
		})
	}
}
