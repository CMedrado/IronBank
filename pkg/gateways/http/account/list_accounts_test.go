package account

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler_ListAccounts(t *testing.T) {
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
			response:     http.StatusOK,
			responsebody: `{"accounts":[{"id":"f7ee7351-4c96-40ca-8cd8-37434810ddfa","name":"Rafael","cpf":"38453162093","secret":"53b9e9679a8ea25880376080b76f98ad","balance":6000,"created_at":"2021-08-02T09:41:46.813816-03:00"},{"id":"a505b1f9-ac4c-45aa-be43-8614a227a9d4","name":"Lucas","cpf":"08131391043","secret":"c74af74c69d81831a5703aefe9cb4199","balance":5000,"created_at":"2021-08-02T09:41:46.813816-03:00"}]}` + "\n",
		},
	}
	for _, tc := range accountst {
		t.Run(tc.name, func(t *testing.T) {
			s := new(Handler)
			s.account = &AccountUsecaseMock{}
			bodyBytes := []byte(tc.body)
			request, _ := http.NewRequest(tc.method, tc.path, bytes.NewReader(bodyBytes))
			responseRecorder := httptest.NewRecorder()

			s.ListAccounts(responseRecorder, request)

			if tc.response != responseRecorder.Code { // O teste falhará pois não queremos erro e obtivemos um
				t.Errorf("unexpected error, wantErr= %d; gotErr= %d", tc.response, responseRecorder.Code)
			}

			if responseRecorder.Body.String() != tc.responsebody {
				t.Errorf("expected an %s but got %s", tc.responsebody, responseRecorder.Body.String())
			}
		})
	}
}
