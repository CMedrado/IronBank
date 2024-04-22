package account

import (
	"fmt"
	"testing"

	"github.com/redis/go-redis/v9"
)

func TestGetBalance(t *testing.T) {
	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", "localhost", "6379"),
	})
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
			in:      "f7ee7351-4c96-40ca-8cd8-37434810ddfs",
			wantErr: true,
		},
	}

	for _, testCase := range tt {
		t.Run(testCase.name, func(t *testing.T) {
			usecase := UseCase{StoredAccount: AccountRepoMock{}, redis: rdb}
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
