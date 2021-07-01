package account

import (
	"bytes"
	store_account "github.com/CMedrado/DesafioStone/storage/file/account"
	"github.com/sirupsen/logrus"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestHandler_GetBalance(t *testing.T) {
	dataBaseAccount, clenDataBaseAccount := createTemporaryFileAccount(t, `[{"id":"f7ee7351-4c96-40ca-8cd8-37434810ddfa","name":"Rafael","cpf":"38453162093","secret":"53b9e9679a8ea25880376080b76f98ad","balance":6000,"created_at":"06/01/2020"},{"id":"a505b1f9-ac4c-45aa-be43-8614a227a9d4","name":"Lucas","cpf":"08131391043","secret":"c74af74c69d81831a5703aefe9cb4199","balance":5000,"created_at":"06/01/2020"}]`)
	defer clenDataBaseAccount()
	accountStorage := store_account.NewStoredAccount(dataBaseAccount)
	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{TimestampFormat: time.RFC3339})
	Lentry := logrus.NewEntry(logger)
	AccountUsecase := &AccountUsecaseMock{AccountList: accountStorage}
	S := new(Handler)
	S.account = AccountUsecase
	S.logger = Lentry
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

			S.GetBalance(responseRecorder, request)

			if tc.response != responseRecorder.Code { // O teste falhará pois não queremos erro e obtivemos um
				t.Errorf("unexpected error, wantErr= %d; gotErr= %d", tc.response, responseRecorder.Code)
			}

			if responseRecorder.Body.String() != tc.responsebody && tc.responsebody != "" {
				t.Errorf("expected an %s but got %s", tc.responsebody, responseRecorder.Body.String())
			}
		})
	}
}
