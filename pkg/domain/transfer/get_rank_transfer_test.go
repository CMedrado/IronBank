package transfer

//
//import (
//	"context"
//	"fmt"
//	"github.com/redis/go-redis/v9"
//	"testing"
//)
//
//func TestMakeGetRankTransfers(t *testing.T) {
//	rdb := redis.NewClient(&redis.Options{
//		Addr: fmt.Sprintf("%s:%s", "localhost", "6379"),
//	})
//	tt := []struct {
//		name                    string
//		in                      CreateTransferInput
//		wantErr                 bool
//		expectedUpdateCallCount int
//	}{
//		{
//			name: "should successfully transfer amount",
//			in: CreateTransferInput{
//				Token:                "MjAvMDcvMjAyMSAxNToxNzoyNTo2YjE5NDFkYi1jZTE3LTRmZmUtYTdlZC0yMjQ5M2E5MjZiYmM6MzlhNzBhOTQtYTgyZC00ZGI4LTg3YWUtYmQ5MDBjNmE3YzA0",
//				AccountDestinationID: "a61227cf-a857-4bc6-8fcd-ad97cdad382a",
//				Amount:               300,
//			},
//			wantErr: false,
//		},
//		{
//			name: "should unsuccessfully transfer amount when there is wrong token",
//			in: CreateTransferInput{
//				Token:                "MDIvMDgvMjAyMSAwOToyNzo0NDo2YjE5NDFkYi1jZTE3LTRmZmUtYTdlZC0yMjQ5M3E5MjZiYmM6YmQxODIxZTQtM2I5YS00M2RjLWJkZGUtNjBiM2QyMTRiYzdm",
//				AccountDestinationID: "75432539-c5ba-46d3-9690-44985b516da7",
//				Amount:               300,
//			},
//			wantErr:                 true,
//			expectedUpdateCallCount: 0,
//		},
//		{
//			name: "should unsuccessfully transfer amount when there is no account destination id",
//			in: CreateTransferInput{
//				Token:                "MDIvMDgvMjAyMSAwOToyNzo0NDo2YjE5NDFkYi1jZTE3LTRmZmUtYTdlZC0yMjQ5M2E5MjZiYmM6YmQxODIxZTQtM2I5YS00M2RjLWJkZGUtNjBiM2QyMTRhYzdm",
//				AccountDestinationID: "c5424440-4737-4e03-86d2-3adac90ddd20",
//				Amount:               300,
//			},
//			wantErr:                 true,
//			expectedUpdateCallCount: 0,
//		},
//		{
//			name: "should unsuccessfully transfer amount when the amount is too slow",
//			in: CreateTransferInput{
//				Token:                "MDIvMDgvMjAyMSAwOToyNzo0NDo2YjE5NDFkYi1jZTE3LTRmZmUtYTdlZC0yMjQ5M2E5MjZiYmM6YmQxODIxZTQtM2I5YS00M2RjLWJkZGUtNjBiM2QyMTRhYzdm",
//				AccountDestinationID: "75432539-c5ba-46d3-9690-44985b516da7",
//				Amount:               0,
//			},
//			wantErr:                 true,
//			expectedUpdateCallCount: 0,
//		},
//		{
//			name: "should unsuccessfully transfer amount when the amount is greater than the balance",
//			in: CreateTransferInput{
//				Token:                "MDIvMDgvMjAyMSAwOToyNzo0NDo2YjE5NDFkYi1jZTE3LTRmZmUtYTdlZC0yMjQ5M2E5MjZiYmM6YmQxODIxZTQtM2I5YS00M2RjLWJkZGUtNjBiM2QyMTRhYzdm",
//				AccountDestinationID: "75432539-c5ba-46d3-9690-44985b516da7",
//				Amount:               52000,
//			},
//			wantErr:                 true,
//			expectedUpdateCallCount: 0,
//		},
//	}
//
//	for _, testCase := range tt {
//		t.Run(testCase.name, func(t *testing.T) {
//			usecase := UseCase{
//				StoredTransfer: TransferRepoMock{},
//				redis:          rdb,
//			}
//			_, _ = usecase.CreateTransfers(context.Background(), testCase.in.Token, testCase.in.Amount, testCase.in.AccountDestinationID)
//			gotRank, gotErr := usecase.CreateTransfers(context.Background(), testCase.in.Token, testCase.in.Amount, testCase.in.AccountDestinationID)
//			if !testCase.wantErr && gotErr != nil {
//				t.Errorf("unexpected error, wantErr=%v; gotErr=%s", testCase.wantErr, gotErr)
//			}
//
//			if testCase.wantErr && gotErr == nil {
//				t.Error("wanted err but got nil")
//			}
//
//			if (gotRank == 1) && !testCase.wantErr && gotErr != nil {
//				t.Errorf("expected an Token but got %d", gotTransfer)
//			}
//		})
//	}
//}
