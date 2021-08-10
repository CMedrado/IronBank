package transfer

import (
	"bytes"
	"github.com/CMedrado/DesafioStone/pkg/gateways/db/file/account"
	"github.com/CMedrado/DesafioStone/pkg/gateways/db/file/token"
	"github.com/CMedrado/DesafioStone/pkg/gateways/db/file/transfer"
	"github.com/sirupsen/logrus"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestHandler_ListTransfers(t *testing.T) {
	dataBaseAccount, clenDataBaseAccount := createTemporaryFileAccount(t, `[{"id":"c5424440-4737-4e03-86d2-3adac90ddd20","name":"Lucas","cpf":"38453162093","secret":"7e65a9b554bbc9817aa049ce38c84a72","balance":3000,"created_at":"29/06/2021 12:45:35"},{"id":"75432539-c5ba-46d3-9690-44985b516da7","name":"Rafael","cpf":"08131391043","secret":"3467e121a1a109628e0a5b0cebba361b","balance":5000,"created_at":"29/06/2021 12:46:28"}]`)
	defer clenDataBaseAccount()
	dataBaseToken, clenDataBaseToken := createTemporaryFileToken(t, `[{"id":"c5424440-4737-4e03-86d2-3adac90ddd20","token":"MjkvMDYvMjAyMSAxMjo0NzowNzpjNTQyNDQ0MC00NzM3LTRlMDMtODZkMi0zYWRhYzkwZGRkMjA="}]`)
	defer clenDataBaseToken()
	dataBaseTransfer, clenDataBaseTransfer := createTemporaryFileTransfer(t, `[{"id":"47399f23-2093-4dde-b32f-990cac27630e","account_origin_id":"c5424440-4737-4e03-86d2-3adac90ddd20","account_destination_id":"75432539-c5ba-46d3-9690-44985b516da7","amount":150,"created_at":"29/06/2021 12:48:06"}]`)
	defer clenDataBaseTransfer()
	accountStorage := account.NewStoredAccount(dataBaseAccount)
	tokenStorage := token.NewStoredToked(dataBaseToken)
	transferStorage := transfer.NewStoredTransfer(dataBaseTransfer)
	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{TimestampFormat: time.RFC3339})
	Lentry := logrus.NewEntry(logger)
	TransferUseCase := &TransferUsecaseMock{AccountList: accountStorage, TokenList: tokenStorage, TransferList: transferStorage}
	S := new(Handler)
	S.transfer = TransferUseCase
	S.logger = Lentry
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
			token:        "MjkvMDYvMjAyMSAxMjo0NzowNzpjNTQyNDQ0MC00NzM3LTRlMDMtODZkMi0zYWRhYzkwZGRkMjA=",
			responsebody: `{"transfers":[{"id":"47399f23-2093-4dde-b32f-990cac27630e","account_origin_id":"c5424440-4737-4e03-86d2-3adac90ddd20","account_destination_id":"75432539-c5ba-46d3-9690-44985b516da7","amount":150,"created_at":"29/06/2021 12:48:06"}]}` + "\n",
		},
		{
			name:         "should unsuccessfully get transfer when there is wrong token",
			method:       "GET",
			path:         "/transfers",
			response:     http.StatusUnauthorized,
			token:        "MjkvMDYvMjAyMSAxMzoxNjo1NTo3NTQzMjUzOS1jNWJhLTQ2ZDMtOTY5MC00NDk4NWI1MTZkYTc=",
			responsebody: `{"errors":"given token is invalid"}` + "\n",
		},
	}
	for _, tc := range gettransfer {
		t.Run(tc.name, func(t *testing.T) {
			bodyBytes := []byte(tc.body)
			request, _ := http.NewRequest(tc.method, tc.path, bytes.NewReader(bodyBytes))
			respondRecorder := httptest.NewRecorder()

			request.Header.Add("Authorization", tc.token)

			S.ListTransfers(respondRecorder, request)

			if tc.response != respondRecorder.Code { // O teste falhará pois não queremos erro e obtivemos um
				t.Errorf("unexpected error, wantErr= %d; gotErr= %d", tc.response, respondRecorder.Code)
			}

			if respondRecorder.Body.String() != tc.responsebody && tc.responsebody != "" {
				t.Errorf("expected an %s but got %s", tc.responsebody, respondRecorder.Body.String())
			}
		})
	}
}
