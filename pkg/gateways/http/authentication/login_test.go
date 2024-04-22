package authentication

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bmizerany/assert"
	"github.com/google/uuid"

	domain2 "github.com/CMedrado/DesafioStone/pkg/domain"
	"github.com/CMedrado/DesafioStone/pkg/domain/authentication"
	"github.com/CMedrado/DesafioStone/pkg/domain/entities"
)

var (
	cpfIncorrect = `{"errors":"given secret or CPF are incorrect"}` + "\n"
)

func TestHandler_Login(t *testing.T) {
	logint := []struct {
		name         string
		method       string
		path         string
		body         string
		response     int
		responsebody string
	}{
		{
			name:         "should successfully authenticated login with formatted CPF",
			method:       "POST",
			path:         "/login",
			body:         `{"cpf": "08131391043", "Secret": "lixo"}`,
			response:     http.StatusOK,
			responsebody: `{"token":"passou"}` + "\n",
		},
		{
			name:         "should successfully authenticated login with unformatted CPF",
			method:       "POST",
			path:         "/login",
			body:         `{"cpf": "38453162093", "Secret": "call"}`,
			response:     http.StatusOK,
			responsebody: `{"token":"passou"}` + "\n",
		},
		{
			name:         "should unsuccessfully authenticated login when cpf is not registered",
			method:       "POST",
			path:         "/login",
			body:         `{"cpf": "38453162723", "Secret": "jax"}`,
			response:     http.StatusUnauthorized,
			responsebody: cpfIncorrect,
		},
		{
			name:         "should unsuccessfully create an account when CPF is invalid",
			method:       "POST",
			path:         "/login",
			body:         `{"cpf": "384531.62793", "Secret": "jax"}`,
			response:     http.StatusUnauthorized,
			responsebody: cpfIncorrect,
		},
		{
			name:         "should unsuccessfully authenticated login when secret is not correct",
			method:       "POST",
			path:         "/login",
			body:         `{"cpf": "081.313.910-43", "Secret": "calls"}`,
			response:     http.StatusUnauthorized,
			responsebody: cpfIncorrect,
		},
		{
			name:     "should unsuccessfully authenticated login when json is invalid",
			method:   "POST",
			path:     "/login",
			body:     `{"Secret" "jax"}`,
			response: http.StatusBadRequest,
		},
	}
	for _, tc := range logint {
		t.Run(tc.name, func(t *testing.T) {
			s := new(Handler)
			s.login = &TokenUseCaseMock{}
			bodyBytes := []byte(tc.body)
			request, _ := http.NewRequest(tc.method, tc.path, bytes.NewReader(bodyBytes))
			responseRecorder := httptest.NewRecorder()

			s.Login(responseRecorder, request)

			assert.Equal(t, tc.response, responseRecorder.Code)
			assert.Equal(t, responseRecorder.Body.String(), tc.responsebody)
		})
	}
}

type TokenUseCaseMock struct {
}

func (uc TokenUseCaseMock) AuthenticatedLogin(secret string, cpf string) (error, string) {
	secretHash := domain2.CreateHash(secret)
	secret1 := domain2.CreateHash("call")
	secret2 := domain2.CreateHash("lixo")
	if cpf == "" {
		return authentication.ErrLogin, ""
	}
	//if account.CPF != account.CPF {
	//	return errors.New("given secret or CPF are incorrect"), ""
	//}
	switch {
	case secretHash == secret1:
		return nil, "passou"
	case secretHash == secret2:
		return nil, "passou"
	default:
		return authentication.ErrLogin, ""
	}
}

func (uc TokenUseCaseMock) GetTokenID(_ uuid.UUID) (entities.Token, error) {
	return entities.Token{}, nil
}
