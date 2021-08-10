package account

import (
	domain2 "github.com/CMedrado/DesafioStone/pkg/domain"
	"github.com/google/uuid"
	"testing"
	"time"
)

type CreateAccountTestInput struct {
	ID        uuid.UUID
	Name      string
	CPF       string
	Secret    string
	Balance   int
	CreatedAt string
}

var I = 0

func TestCreateAccount(t *testing.T) {
	//prepare
	testTable := []struct { // tt := ....
		name    string                 //Nome do teste
		in      CreateAccountTestInput //Entrada da Função
		wantErr bool                   //Pra dizer se espera ou não um err
		want    int
	}{
		{
			name: "should successfully create an account with formatted CPF",
			in: CreateAccountTestInput{
				Name:    "Rafaels",
				CPF:     "081.313.920-43",
				Secret:  "lucas",
				Balance: 50000,
			},
			wantErr: false,
		},
		{
			name: "should successfully create an account with unformulated CPF",
			in: CreateAccountTestInput{
				Name:    "Lucas",
				CPF:     "98634575498",
				Secret:  "teo90",
				Balance: 60000,
			},
			wantErr: false,
		},
		{
			name: "should unsuccessfully create an account when CPF is invalid",
			in: CreateAccountTestInput{
				Name:    "Marcos",
				CPF:     "398.176200-26",
				Secret:  "marcos35",
				Balance: 7000,
			},
			wantErr: true,
		},
		{
			name: "should unsuccessfully create an account when CPF is already used",
			in: CreateAccountTestInput{
				Name:    "Rafaels",
				CPF:     "081.313.920-43",
				Secret:  "lucas",
				Balance: 50000,
			},
			wantErr: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			//test
			useCase := UseCase{AccountRepoMock{}}
			gotID, gotErr := useCase.CreateAccount(testCase.in.Name, testCase.in.CPF, testCase.in.Secret, testCase.in.Balance)

			//assert
			if !testCase.wantErr && gotErr != nil { // O teste falhará pois não queremos erro e obtivemos um
				t.Errorf("unexpected error, wantErr=%v; gotErr=%s", testCase.wantErr, gotErr)
			}

			if testCase.wantErr && gotErr == nil { // O teste falhará pois queremos erro e não obtivemos um
				t.Error("wanted err but got nil")
			}

			if (gotID == uuid.UUID{}) && !testCase.wantErr && gotErr != nil {
				t.Errorf("expected an ID but got %d", gotID)
			}
		})
	}
}

func TestGetBalance(t *testing.T) {
	tt := []struct {
		name    string
		in      string
		wantErr bool
		want    int
	}{
		{
			name:    "should successfully get balance with ID",
			in:      "a505b1f9-ac4c-45aa-be43-8614a227a9d4",
			wantErr: false,
			want:    6000,
		},
		{
			name:    "should unsuccessfully get balance when ID is invalid",
			in:      "f7ee7351-4c96-40ca-8cd8-37434810ddfa",
			wantErr: true,
		},
	}

	for _, testCase := range tt {
		t.Run(testCase.name, func(t *testing.T) {
			usecase := UseCase{AccountRepoMock{}}
			gotBalance, gotErr := usecase.GetBalance(testCase.in)

			//assert
			if !testCase.wantErr && gotErr != nil { // O teste falhará pois não queremos erro e obtivemos um
				t.Errorf("unexpected error, wantErr=%v; gotErr=%s", testCase.wantErr, gotErr)
			}

			if testCase.wantErr && gotErr == nil { // O teste falhará pois queremos erro e não obtivemos um
				t.Error("wanted err but got nil")
			}

			if gotBalance != testCase.want {
				t.Errorf("expected an ID but got %d", gotBalance)
			}
		})
	}
}

type AccountRepoMock struct {
}

func (uc AccountRepoMock) SaveAccount(_ domain2.Account) error {
	return nil
}

func (uc AccountRepoMock) ReturnAccounts() ([]domain2.Account, error) {
	return []domain2.Account{
		{
			ID:        uuid.MustParse("a505b1f9-ac4c-45aa-be43-8614a227a9d4"),
			Name:      "Rafael",
			CPF:       "38453162093",
			Secret:    "53b9e9679a8ea25880376080b76f98ad",
			Balance:   6000,
			CreatedAt: time.Now(),
		},
		{
			ID:        uuid.MustParse("f7ee7351-4c96-40ca-8cd8-37434810ddfa"),
			Name:      "Lucas",
			CPF:       "08131391043",
			Secret:    "c74af74c69d81831a5703aefe9cb4199",
			Balance:   5000,
			CreatedAt: time.Now(),
		},
	}, nil
}

func (uc AccountRepoMock) ChangeBalance(person1, person2 domain2.Account) error {
	return nil
}

func (uc AccountRepoMock) ReturnAccountID(id uuid.UUID) (domain2.Account, error) {
	if id == uuid.MustParse("f7ee7351-4c96-40ca-8cd8-37434810ddfa") {
		return domain2.Account{
			ID:        uuid.MustParse("f7ee7351-4c96-40ca-8cd8-37434810ddfa"),
			Name:      "Lucas",
			CPF:       "08131391043",
			Secret:    "c74af74c69d81831a5703aefe9cb4199",
			Balance:   5000,
			CreatedAt: time.Now(),
		}, nil
	}
	if id == uuid.MustParse("f7ee7351-4c96-40ca-8cd8-37434810ddfa") {
		return domain2.Account{
			ID:        uuid.MustParse("a505b1f9-ac4c-45aa-be43-8614a227a9d4"),
			Name:      "Rafael",
			CPF:       "38453162093",
			Secret:    "53b9e9679a8ea25880376080b76f98ad",
			Balance:   6000,
			CreatedAt: time.Now(),
		}, nil
	}
	return domain2.Account{}, nil
}

func (uc AccountRepoMock) ReturnAccountCPF(cpf string) (domain2.Account, error) {
	cpf2 := "08131392043"
	cpf3 := "38453162093"
	if cpf == cpf2 {
		if I == 0 {
			I = 1
			return domain2.Account{}, nil
		}
		if I != 0 {
			return domain2.Account{
				ID:        uuid.MustParse("f7ee7351-4c96-40ca-8cd8-37434810ddfa"),
				Name:      "Lucas",
				CPF:       "08131391043",
				Secret:    "c74af74c69d81831a5703aefe9cb4199",
				Balance:   5000,
				CreatedAt: time.Now(),
			}, nil
		}
	}
	if cpf == cpf3 {
		return domain2.Account{
			ID:        uuid.MustParse("a505b1f9-ac4c-45aa-be43-8614a227a9d4"),
			Name:      "Rafael",
			CPF:       "38453162093",
			Secret:    "53b9e9679a8ea25880376080b76f98ad",
			Balance:   6000,
			CreatedAt: time.Now(),
		}, nil

	}
	return domain2.Account{}, nil
}
