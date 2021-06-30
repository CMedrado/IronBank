package https

import (
	"bytes"
	"errors"
	"github.com/CMedrado/DesafioStone/domain"
	"github.com/CMedrado/DesafioStone/domain/account"
	"github.com/CMedrado/DesafioStone/domain/login"
	"github.com/CMedrado/DesafioStone/domain/transfer"
	store_account "github.com/CMedrado/DesafioStone/storage/file/account"
	store_token "github.com/CMedrado/DesafioStone/storage/file/token"
	store_transfer "github.com/CMedrado/DesafioStone/storage/file/transfer"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

func createTemporaryFileToken(t *testing.T, Tokens string) (io.ReadWriteSeeker, func()) {
	filetmp, err := ioutil.TempFile("", "dbtoken")

	if err != nil {
		t.Fatalf("it is not possible to write the temporary file %v", err)
	}

	filetmp.Write([]byte(Tokens))

	removeArquivo := func() {
		filetmp.Close()
		os.Remove(filetmp.Name())
	}

	return filetmp, removeArquivo
}

func createTemporaryFileAccount(t *testing.T, Accounts string) (io.ReadWriteSeeker, func()) {
	filetmp, err := ioutil.TempFile("", "dbaccount")

	if err != nil {
		t.Fatalf("it is not possible to write the temporary file %v", err)
	}

	filetmp.Write([]byte(Accounts))

	removeArquivo := func() {
		filetmp.Close()
		os.Remove(filetmp.Name())
	}

	return filetmp, removeArquivo
}

func createTemporaryFileTransfer(t *testing.T, Transfers string) (io.ReadWriteSeeker, func()) {
	filetmp, err := ioutil.TempFile("", "dbtransfer")

	if err != nil {
		t.Fatalf("it is not possible to write the temporary file %v", err)
	}

	filetmp.Write([]byte(Transfers))

	removeArquivo := func() {
		filetmp.Close()
		os.Remove(filetmp.Name())
	}

	return filetmp, removeArquivo
}
func TestTransferHandler(t *testing.T) {
	dataBaseAccount, clenDataBaseAccount := createTemporaryFileAccount(t, `[{"id":"c5424440-4737-4e03-86d2-3adac90ddd20","name":"Lucas","cpf":"38453162093","secret":"7e65a9b554bbc9817aa049ce38c84a72","balance":3000,"created_at":"29/06/2021 12:45:35"},{"id":"75432539-c5ba-46d3-9690-44985b516da7","name":"Rafael","cpf":"08131391043","secret":"3467e121a1a109628e0a5b0cebba361b","balance":5000,"created_at":"29/06/2021 12:46:28"}]`)
	defer clenDataBaseAccount()
	dataBaseToken, clenDataBaseToken := createTemporaryFileToken(t, `[{"id":"c5424440-4737-4e03-86d2-3adac90ddd20","token":"MjkvMDYvMjAyMSAxMjo0NzowNzpjNTQyNDQ0MC00NzM3LTRlMDMtODZkMi0zYWRhYzkwZGRkMjA="}]`)
	defer clenDataBaseToken()
	dataBaseTransfer, clenDataBaseTransfer := createTemporaryFileTransfer(t, `[{"id":"47399f23-2093-4dde-b32f-990cac27630e","account_origin_id":"c5424440-4737-4e03-86d2-3adac90ddd20","account_destination_id":"75432539-c5ba-46d3-9690-44985b516da7","amount":150,"created_at":"29/06/2021 12:48:06"}]`)
	defer clenDataBaseTransfer()
	accountStorage := store_account.NewStoredAccount(dataBaseAccount)
	tokenStorage := store_token.NewStoredToked(dataBaseToken)
	transferStorage := store_transfer.NewStoredTransfer(dataBaseTransfer)
	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{TimestampFormat: time.RFC3339})
	Lentry := logrus.NewEntry(logger)
	TransferUseCase := &TransferUsecaseMock{AccountList: accountStorage, TokenList: tokenStorage, TransferList: transferStorage}
	S := new(ServerAccount)
	S.transfer = TransferUseCase
	S.logger = Lentry
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
			body:         `{"account_destination_id":"75432539-c5ba-46d3-9690-44985b516da7","amount": 500}`,
			response:     http.StatusCreated,
			responsebody: `{"id":"c5424440-4737-4e03-86d2-3adac90ddd20"}` + "\n",
			token:        "MjkvMDYvMjAyMSAxMjo0NzowNzpjNTQyNDQ0MC00NzM3LTRlMDMtODZkMi0zYWRhYzkwZGRkMjA=",
		},
		{
			name:         "should unsuccessfully transfer amount when there is wrong token",
			method:       "POST",
			path:         "/transfers",
			body:         `{"account_destination_id":"75432539-c5ba-46d3-9690-44985b516da7","amount": 300}`,
			response:     http.StatusUnauthorized,
			token:        "MjkvMDYvMjAyMSAxMzoxNjo1NTo3NTQzMjUzOS1jNWJhLTQ2ZDMtOTY5MC00NDk4NWI1MTZkYTc=",
			responsebody: `{"errors":"given token is invalid"}` + "\n",
		},
		{
			name:         "should unsuccessfully transfer amount when there is wrong destination ID",
			method:       "POST",
			path:         "/transfers",
			body:         `{"account_destination_id":"75432539-c5ba-46d3-9690-44985b516da5","amount": 300}`,
			response:     http.StatusNotAcceptable,
			token:        "MjkvMDYvMjAyMSAxMjo0NzowNzpjNTQyNDQ0MC00NzM3LTRlMDMtODZkMi0zYWRhYzkwZGRkMjA=",
			responsebody: `{"errors":"given account destination id is invalid"}` + "\n",
		},
		{
			name:         "should unsuccessfully transfer amount when there is invalid amount",
			method:       "POST",
			path:         "/transfers",
			body:         `{"account_destination_id":"75432539-c5ba-46d3-9690-44985b516da7","amount": -5}`,
			response:     http.StatusBadRequest,
			token:        "MjkvMDYvMjAyMSAxMjo0NzowNzpjNTQyNDQ0MC00NzM3LTRlMDMtODZkMi0zYWRhYzkwZGRkMjA=",
			responsebody: `{"errors":"given amount is invalid"}` + "\n",
		},
		{
			name:         "should unsuccessfully transfer amount when there without balance ",
			method:       "POST",
			path:         "/transfers",
			body:         `{"account_destination_id":"75432539-c5ba-46d3-9690-44985b516da7","amount": 60000}`,
			response:     http.StatusBadRequest,
			token:        "MjkvMDYvMjAyMSAxMjo0NzowNzpjNTQyNDQ0MC00NzM3LTRlMDMtODZkMi0zYWRhYzkwZGRkMjA=",
			responsebody: `{"errors":"given account without balance"}` + "\n",
		},
		{
			name:         "should unsuccessfully transfer amount when there same account ",
			method:       "POST",
			path:         "/transfers",
			body:         `{"account_destination_id":"c5424440-4737-4e03-86d2-3adac90ddd20","amount": 300}`,
			response:     http.StatusBadRequest,
			token:        "MjkvMDYvMjAyMSAxMjo0NzowNzpjNTQyNDQ0MC00NzM3LTRlMDMtODZkMi0zYWRhYzkwZGRkMjA=",
			responsebody: `{"errors":"given account is the same as the account destination"}` + "\n",
		},
		{
			name:     "should unsuccessfully transfer amount when json is invalid",
			method:   "POST",
			path:     "/transfers",
			body:     `{"account"0}`,
			response: http.StatusBadRequest,
			token:    "MTgvMDYvMjAyMSAxNjozNDozMjo5ODQ5ODA4MQ==",
		},
	}
	for _, tc := range createtransfer {
		t.Run(tc.name, func(t *testing.T) {
			bodyBytes := []byte(tc.body)
			request, _ := http.NewRequest(tc.method, tc.path, bytes.NewReader(bodyBytes))
			respondeRecorder := httptest.NewRecorder()

			request.Header.Add("Authorization", tc.token)

			S.processTransfer(respondeRecorder, request)

			if tc.response != respondeRecorder.Code { // O teste falhará pois não queremos erro e obtivemos um
				t.Errorf("unexpected error, wantErr= %d; gotErr= %d", tc.response, respondeRecorder.Code)
			}
			if respondeRecorder.Body.String() != tc.responsebody && tc.responsebody != "" {
				t.Errorf("expected an %s but got %s", tc.responsebody, respondeRecorder.Body.String())
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

			S.handleTransfers(respondRecorder, request)

			if tc.response != respondRecorder.Code { // O teste falhará pois não queremos erro e obtivemos um
				t.Errorf("unexpected error, wantErr= %d; gotErr= %d", tc.response, respondRecorder.Code)
			}

			if respondRecorder.Body.String() != tc.responsebody && tc.responsebody != "" {
				t.Errorf("expected an %s but got %s", tc.responsebody, respondRecorder.Body.String())
			}
		})
	}
}

type TransferUsecaseMock struct {
	AccountList  *store_account.StoredAccount
	TokenList    *store_token.StoredToken
	TransferList *store_transfer.StoredTransferAccount
}

func (uc *TransferUsecaseMock) GetTransfers(token string) ([]domain.Transfer, error) {
	accountOriginID, err := transfer.DecoderToken(token)
	if err != nil {
		return []domain.Transfer{}, domain.ErrParse
	}
	tokens := domain.Token{}
	for _, a := range uc.TokenList.ReturnTokens() {
		if a.ID == accountOriginID {
			tokens = login.ChangeTokenStorage(a)
		}
	}
	if token != tokens.Token {
		return []domain.Transfer{}, errors.New("given token is invalid")
	}
	var transfers []domain.Transfer
	for _, a := range uc.TransferList.ReturnTransfers() {
		transfers = append(transfers, transfer.ChangeTransferStorage(a))
	}
	return transfers, nil
}

func (uc TransferUsecaseMock) CreateTransfers(token string, accountDestinationIDString string, amount int) (error, uuid.UUID) {
	if amount <= 0 {
		return errors.New("given amount is invalid"), uuid.UUID{}
	}
	accountDestinationID := uuid.MustParse(accountDestinationIDString)
	accountOriginID, err := transfer.DecoderToken(token)
	if err != nil {
		return domain.ErrParse, uuid.UUID{}
	}

	tokens := domain.Token{}
	for _, a := range uc.TokenList.ReturnTokens() {
		if a.ID == accountOriginID {
			tokens = login.ChangeTokenStorage(a)
		}
	}
	if tokens.Token != token {
		return errors.New("given token is invalid"), uuid.UUID{}
	}
	if accountOriginID == accountDestinationID {
		return errors.New("given account is the same as the account destination"), uuid.UUID{}
	}
	accountOrigin := domain.Account{}
	for _, a := range uc.AccountList.ReturnAccounts() {
		if a.ID == accountOriginID {
			accountOrigin = account.ChangeAccountStorage(a)
		}
	}
	accountDestination := domain.Account{}
	for _, a := range uc.AccountList.ReturnAccounts() {
		if a.ID == accountDestinationID {
			accountDestination = account.ChangeAccountStorage(a)
		}
	}
	if accountOrigin.Balance < amount {
		return errors.New("given account without balance"), uuid.UUID{}
	}
	if (accountDestination == domain.Account{}) {
		return errors.New("given account destination id is invalid"), uuid.UUID{}
	}
	returnID := uuid.MustParse("c5424440-4737-4e03-86d2-3adac90ddd20")
	return nil, returnID
}
