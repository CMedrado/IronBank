package account

import (
	storeaccount "github.com/CMedrado/DesafioStone/storage/file/account"
	"io"
	"io/ioutil"
	"os"
	"testing"
)

type CreateAccountTestInput struct {
	ID        int
	Name      string
	CPF       string
	Secret    string
	Balance   int
	CreatedAt string
}

func createTemporaryFile(t *testing.T, Accounts string) (io.ReadWriteSeeker, func()) {
	filetmp, err := ioutil.TempFile("", "db")

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
				CPF:     "38453162093",
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
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			dataBase, clenDataBase := createTemporaryFile(t, `[{"id":981,"name":"Rafael","cpf":"38453162093","secret":"53b9e9679a8ea25880376080b76f98ad","balance":6000,"created_at":"06/01/2020"},{"id":982,"name":"Lucas","cpf":"08131391043","secret":"c74af74c69d81831a5703aefe9cb4199","balance":5000,"created_at":"06/01/2020"}]`)
			defer clenDataBase()
			accountStorage := storeaccount.NewStoredAccount(dataBase)
			usecase := UseCase{
				StoredAccount: accountStorage,
			}

			//test
			gotID, gotErr := usecase.CreateAccount(testCase.in.Name, testCase.in.CPF, testCase.in.Secret, testCase.in.Balance)

			//assert
			if !testCase.wantErr && gotErr != nil { // O teste falhará pois não queremos erro e obtivemos um
				t.Errorf("unexpected error, wantErr=%v; gotErr=%s", testCase.wantErr, gotErr)
			}

			if testCase.wantErr && gotErr == nil { // O teste falhará pois queremos erro e não obtivemos um
				t.Error("wanted err but got nil")
			}

			if gotID == 0 && !testCase.wantErr && gotErr != nil {
				t.Errorf("expected an ID but got %d", gotID)
			}
		})
	}
}

func TestGetBalance(t *testing.T) {
	tt := []struct {
		name    string
		in      int
		wantErr bool
		want    int
	}{
		{
			name:    "should successfully get balance with formatted CPF",
			in:      982,
			wantErr: false,
			want:    5000,
		},
		{
			name:    "should successfully get balance with unformulated CPF",
			in:      981,
			wantErr: false,
			want:    6000,
		},
		{
			name:    "should unsuccessfully get balance when CPF is invalid",
			in:      398 - 6,
			wantErr: true,
		},
		{
			name:    "should unsuccessfully get balance when dont exist account",
			in:      06237,
			wantErr: true,
		},
	}

	for _, testCase := range tt {
		t.Run(testCase.name, func(t *testing.T) {
			dataBase, clenDataBase := createTemporaryFile(t, `[{"id":981,"name":"Rafael","cpf":"38453162093","secret":"53b9e9679a8ea25880376080b76f98ad","balance":6000,"created_at":"06/01/2020"},{"id":982,"name":"Lucas","cpf":"08131391043","secret":"c74af74c69d81831a5703aefe9cb4199","balance":5000,"created_at":"06/01/2020"}]`)
			defer clenDataBase()
			accountStorage := storeaccount.NewStoredAccount(dataBase)
			usecase := UseCase{
				StoredAccount: accountStorage,
			}
			//test
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
