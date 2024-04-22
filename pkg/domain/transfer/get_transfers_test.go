package transfer

import (
	"fmt"
	"testing"

	"github.com/redis/go-redis/v9"
)

func TestMakeGetTransfers(t *testing.T) {
	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", "localhost", "6379"),
	})
	var tt = []struct {
		name    string
		in      CreateTransferInput
		wantErr bool
	}{
		{
			name: "should successfully get transfers",
			in: CreateTransferInput{
				Token: "MjAvMDcvMjAyMSAxNToxNzoyNTo2YjE5NDFkYi1jZTE3LTRmZmUtYTdlZC0yMjQ5M2E5MjZiYmM6MzlhNzBhOTQtYTgyZC00ZGI4LTg3YWUtYmQ5MDBjNmE3YzA0",
			},
			wantErr: false,
		},
		{
			name: "should unsuccessfully get transfer when there is wrong token",
			in: CreateTransferInput{
				Token: "MDIvMDgvMjAyMSAwOToyNzo0NDo2YjE5NDFkYi1jZTE3LTRmZmUtYTdlZC0yMsQ5M2E5MjZiYmM6YmQxODIxZTQtM2I5YS00M2RjLWJkZGUtNjBiM2QyMTRhYzdm",
			},
			wantErr: true,
		},
	}
	for _, testCase := range tt {
		t.Run(testCase.name, func(t *testing.T) {
			usecase := UseCase{
				StoredTransfer: TransferRepoMock{},
				redis:          rdb,
			}
			gotTransfer, gotErr := usecase.GetTransfers(testCase.in.Token)

			if !testCase.wantErr && gotErr != nil {
				t.Errorf("unexpected error, wantErr=%v; gotErr=%s", testCase.wantErr, gotErr)
			}

			if testCase.wantErr && gotErr == nil {
				t.Error("wanted err but got nil")
			}

			if gotTransfer == nil && !testCase.wantErr && gotErr != nil {
				t.Errorf("expected an Token but got %v", gotTransfer)
			}
		})
	}
}
