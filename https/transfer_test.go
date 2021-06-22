package https

import (
	"bytes"
	"errors"
	"github.com/CMedrado/DesafioStone/domain"
	"github.com/CMedrado/DesafioStone/domain/transfer"
	store_account "github.com/CMedrado/DesafioStone/store/account"
	store_login "github.com/CMedrado/DesafioStone/store/login"
	store_transfer "github.com/CMedrado/DesafioStone/store/transfer"
	"github.com/sirupsen/logrus"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"
)

func TestTransferHandler(t *testing.T) {
	Account1 := store_account.Account{ID: 98498081, Name: "Rafael", CPF: "38453162093", Secret: domain.CreateHash("call"), Balance: 5000, CreatedAt: "06/01/2020"}
	Account2 := store_account.Account{ID: 19727887, Name: "Lucas", CPF: "08131391043", Secret: domain.CreateHash("lixo"), Balance: 5000, CreatedAt: "06/01/2020"}
	Token1 := store_login.Token{Token: "MTgvMDYvMjAyMSAxNjozNDozMjo5ODQ5ODA4MQ=="}
	Token2 := store_login.Token{Token: "MTgvMDYvMjAyMSAxNjoxMzoyNDo5ODQ5ODA4MQ=="}
	Transfer1 := store_transfer.Transfer{ID: 74941318, AccountOriginID: 98498081, AccountDestinationID: 19727887, Amount: 900, CreatedAt: "18/06/2021 17:26:26"}
	Transfer2 := store_transfer.Transfer{ID: 27131847, AccountOriginID: 98498081, AccountDestinationID: 19727887, Amount: 50, CreatedAt: "18/06/2021 17:26:26"}
	Transfer3 := store_transfer.Transfer{ID: 39984059, AccountOriginID: 98498081, AccountDestinationID: 19727887, Amount: 60, CreatedAt: "18/06/2021 17:26:26"}
	Transfer4 := store_transfer.Transfer{ID: 11902081, AccountOriginID: 98498081, AccountDestinationID: 19727887, Amount: 90, CreatedAt: "18/06/2021 17:26:26"}
	AccountStorage := make(map[string]store_account.Account)
	AccountTransferID := make(map[int]store_transfer.Transfer)
	TokenStorage := make(map[int]store_login.Token)
	AccountStorage[Account1.CPF] = Account1
	AccountStorage[Account2.CPF] = Account2
	TokenStorage[Account1.ID] = Token1
	TokenStorage[Account2.ID] = Token2
	AccountTransferID[Transfer3.ID] = Transfer3
	AccountTransferID[Transfer2.ID] = Transfer2
	AccountTransferID[Transfer1.ID] = Transfer1
	AccountTransferID[Transfer4.ID] = Transfer4
	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{TimestampFormat: time.RFC3339})
	Lentry := logrus.NewEntry(logger)
	TokenUseCase := &TransferUsecaseMock{AccountList: AccountStorage, TokenList: TokenStorage, TransferList: AccountTransferID}
	S := new(ServerAccount)
	S.transfer = TokenUseCase
	S.logger = Lentry
	secondIDString := strconv.Itoa(Account2.ID)
	firstIDString := strconv.Itoa(Account1.ID)
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
			token:        "MTgvMDYvMjAyMSAxNjozNDozMjo5ODQ5ODA4MQ==",
		},
		{
			name:         "should successfully transfer amount",
			method:       "POST",
			path:         "/transfers",
			body:         `{"account_destination_id":` + secondIDString + `,"amount": 300}`,
			response:     http.StatusCreated,
			responsebody: `{"id":19878}` + "\n",
			token:        "MTgvMDYvMjAyMSAxNjozNDozMjo5ODQ5ODA4MQ==",
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
			token:        "MTgvMDYvMjAyMSAxNjozNDozMjo5ODQ5ODA4MQ==",
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
			body:         `{"account_destination_id":` + secondIDString + `,"amount": 6000}`,
			response:     http.StatusBadRequest,
			token:        "MTgvMDYvMjAyMSAxNjozNDozMjo5ODQ5ODA4MQ==",
			responsebody: `{"errors":"given account without balance"}` + "\n",
		},
		{
			name:         "should unsuccessfully transfer amount when there same account ",
			method:       "POST",
			path:         "/transfers",
			body:         `{"account_destination_id":` + firstIDString + `,"amount": 300}`,
			response:     http.StatusBadRequest,
			token:        "MTgvMDYvMjAyMSAxNjozNDozMjo5ODQ5ODA4MQ==",
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
			token:        "MTgvMDYvMjAyMSAxNjozNDozMjo5ODQ5ODA4MQ==",
			responsebody: `{"transfers":[{"id":39984059,"account_origin_id":98498081,"account_destination_id":19727887,"amount":60,"created_at":"18/06/2021 17:26:26"},{"id":27131847,"account_origin_id":98498081,"account_destination_id":19727887,"amount":50,"created_at":"18/06/2021 17:26:26"},{"id":74941318,"account_origin_id":98498081,"account_destination_id":19727887,"amount":900,"created_at":"18/06/2021 17:26:26"},{"id":11902081,"account_origin_id":98498081,"account_destination_id":19727887,"amount":90,"created_at":"18/06/2021 17:26:26"}]}` + "\n",
		},
		{
			name:         "should unsuccessfully get transfer when there is wrong token",
			method:       "GET",
			path:         "/transfers",
			response:     http.StatusUnauthorized,
			token:        "MTgvMDYvMjAyMSAxNjoxMzoyNDo5ODQ5ODA4MQ==",
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
	AccountList  map[string]store_account.Account
	TokenList    map[int]store_login.Token
	TransferList map[int]store_transfer.Transfer
}

func (uc TransferUsecaseMock) GetTransfers(token string) ([]store_transfer.Transfer, error) {
	accountOriginID := transfer.DecoderToken(token)
	transfers := uc.TransferList
	if token != uc.TokenList[accountOriginID].Token {
		return []store_transfer.Transfer{}, errors.New("given token is invalid")
	}
	var transfer []store_transfer.Transfer
	for _, a := range transfers {
		transfer = append(transfer, a)
	}
	return transfer, nil
}

func (uc TransferUsecaseMock) CreateTransfers(token string, accountDestinationID int, amount int) (error, int) {
	if amount <= 0 {
		return errors.New("given amount is invalid"), 0
	}
	accountOriginID := transfer.DecoderToken(token)
	if token != uc.TokenList[accountOriginID].Token {
		return errors.New("given token is invalid"), 0
	}
	if accountOriginID == accountDestinationID {
		return errors.New("given account is the same as the account destination"), 0
	}
	accountOrigin := store_account.Account{}
	for _, a := range uc.AccountList {
		if a.ID == accountOriginID {
			accountOrigin = a
		}
	}
	accountDestination := store_account.Account{}
	for _, a := range uc.AccountList {
		if a.ID == accountDestinationID {
			accountDestination = a
		}
	}
	if accountOrigin.Balance < amount {
		return errors.New("given account without balance"), 0
	}
	if (accountDestination == store_account.Account{}) {
		return errors.New("given account destination id is invalid"), 0
	}
	return nil, 19878
}
