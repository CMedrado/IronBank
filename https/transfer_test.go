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
	"github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
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
	dataBaseAccount, clenDataBaseAccount := createTemporaryFileAccount(t, `[{"id":19727887,"name":"Rafael","cpf":"38453162093","secret":"53b9e9679a8ea25880376080b76f98ad","balance":6000,"created_at":"06/01/2020"},{"id":98498081,"name":"Lucas","cpf":"08131391043","secret":"c74af74c69d81831a5703aefe9cb4199","balance":5000,"created_at":"06/01/2020"}]`)
	defer clenDataBaseAccount()
	dataBaseToken, clenDataBaseToken := createTemporaryFileToken(t, `[{"id":19727887,"token":"MjEvMDYvMjAyMSAyMzo1OTowMDoxOTcyNzg4Nw=="}]`)
	defer clenDataBaseToken()
	dataBaseTransfer, clenDataBaseTransfer := createTemporaryFileTransfer(t, `[{"id":39984059,"account_origin_id":98498081,"account_destination_id":19727887,"amount":60,"created_at":"18/06/2021 17:26:26"},{"id":27131847,"account_origin_id":98498081,"account_destination_id":19727887,"amount":50,"created_at":"18/06/2021 17:26:26"},{"id":74941318,"account_origin_id":98498081,"account_destination_id":19727887,"amount":900,"created_at":"18/06/2021 17:26:26"},{"id":11902081,"account_origin_id":98498081,"account_destination_id":19727887,"amount":90,"created_at":"18/06/2021 17:26:26"}]`)
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
	secondIDString := strconv.Itoa(98498081)
	firstIDString := strconv.Itoa(19727887)
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
			body:         `{"account_destination_id":` + secondIDString + `,"amount": 500}`,
			response:     http.StatusCreated,
			responsebody: `{"id":19878}` + "\n",
			token:        "MjEvMDYvMjAyMSAyMzo1OTowMDoxOTcyNzg4Nw==",
		},
		{
			name:         "should unsuccessfully transfer amount when there is wrong token",
			method:       "POST",
			path:         "/transfers",
			body:         `{"account_destination_id":` + secondIDString + `,"amount": 300}`,
			response:     http.StatusUnauthorized,
			token:        "MTgvMDYvMjAyMSAxNjozNDozMjo5OD5ODA4MQ==",
			responsebody: `{"errors":"given token is invalid"}` + "\n",
		},
		{
			name:         "should unsuccessfully transfer amount when there is wrong destination ID",
			method:       "POST",
			path:         "/transfers",
			body:         `{"account_destination_id":7568497,"amount": 300}`,
			response:     http.StatusNotAcceptable,
			token:        "MjEvMDYvMjAyMSAyMzo1OTowMDoxOTcyNzg4Nw==",
			responsebody: `{"errors":"given account destination id is invalid"}` + "\n",
		},
		{
			name:         "should unsuccessfully transfer amount when there is invalid amount",
			method:       "POST",
			path:         "/transfers",
			body:         `{"account_destination_id":` + secondIDString + `,"amount": -5}`,
			response:     http.StatusBadRequest,
			token:        "MTgvMDYvMjAyMSAxNjozNDozMjo5ODQ5ODA4MQ==",
			responsebody: `{"errors":"given amount is invalid"}` + "\n",
		},
		{
			name:         "should unsuccessfully transfer amount when there without balance ",
			method:       "POST",
			path:         "/transfers",
			body:         `{"account_destination_id":` + secondIDString + `,"amount": 60000}`,
			response:     http.StatusBadRequest,
			token:        "MjEvMDYvMjAyMSAyMzo1OTowMDoxOTcyNzg4Nw==",
			responsebody: `{"errors":"given account without balance"}` + "\n",
		},
		{
			name:         "should unsuccessfully transfer amount when there same account ",
			method:       "POST",
			path:         "/transfers",
			body:         `{"account_destination_id":` + firstIDString + `,"amount": 300}`,
			response:     http.StatusBadRequest,
			token:        "MjEvMDYvMjAyMSAyMzo1OTowMDoxOTcyNzg4Nw==",
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
			token:        "MjEvMDYvMjAyMSAyMzo1OTowMDoxOTcyNzg4Nw==",
			responsebody: `{"transfers":[{"id":39984059,"account_origin_id":98498081,"account_destination_id":19727887,"amount":60,"created_at":"18/06/2021 17:26:26"},{"id":27131847,"account_origin_id":98498081,"account_destination_id":19727887,"amount":50,"created_at":"18/06/2021 17:26:26"},{"id":74941318,"account_origin_id":98498081,"account_destination_id":19727887,"amount":900,"created_at":"18/06/2021 17:26:26"},{"id":11902081,"account_origin_id":98498081,"account_destination_id":19727887,"amount":90,"created_at":"18/06/2021 17:26:26"}]}` + "\n",
		},
		{
			name:         "should unsuccessfully get transfer when there is wrong token",
			method:       "GET",
			path:         "/transfers",
			response:     http.StatusUnauthorized,
			token:        "MjEvMDYvMjAyMSAyMzo1OTowMDoxOTcyNzg4w==",
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
	accountOriginID := transfer.DecoderToken(token)
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

func (uc TransferUsecaseMock) CreateTransfers(token string, accountDestinationID int, amount int) (error, int) {
	if amount <= 0 {
		return errors.New("given amount is invalid"), 0
	}
	accountOriginID := transfer.DecoderToken(token)
	tokens := domain.Token{}
	for _, a := range uc.TokenList.ReturnTokens() {
		if a.ID == accountOriginID {
			tokens = login.ChangeTokenStorage(a)
		}
	}
	if tokens.Token != token {
		return errors.New("given token is invalid"), 0
	}
	if accountOriginID == accountDestinationID {
		return errors.New("given account is the same as the account destination"), 0
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
		return errors.New("given account without balance"), 0
	}
	if (accountDestination == domain.Account{}) {
		return errors.New("given account destination id is invalid"), 0
	}
	return nil, 19878
}
