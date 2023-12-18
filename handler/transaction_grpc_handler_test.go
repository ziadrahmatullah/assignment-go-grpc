package handler_test

import (
	"context"
	"testing"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/apperror"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/handler"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/mocks"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/pb"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var pbtransReq = pb.TransactionsReq{
	Search:          "",
	FilterStart:     "",
	FilterEnd:       "",
	SortBy:          "",
	SortType:        "",
	PaginationLimit: "",
	PaginationPage:  "",
}

var pbtfReq = pb.TransferReq{
	WalletNumber: "7000000000002",
	Amount:       "10000",
}

var pbtuReq = pb.TopUpReq{
	Amount:       "100000",
	SourceOfFund: "Cash",
}

func TestGRPCGetAllTransactions(t *testing.T) {
	t.Run("should return something if get transactions success", func(t *testing.T) {
		v := mocks.NewAppValidator(t)
		uu := mocks.NewTransactionUsecase(t)
		uh := handler.NewTransactionGRPCHandler(uu, v)
		ctx := context.Background()
		ctx2 := context.WithValue(ctx, "id", uint(1))
		uu.On("GetTransactions", ctx2, mock.Anything, uint(1)).Return(&transactionsRes, nil)

		res, _ := uh.GetAllTransactions(ctx2, &pbtransReq)

		assert.NotNil(t, res)
	})

	t.Run("should return err when error in query", func(t *testing.T) {
		expectedErr := apperror.ErrCantConnectToThirdParty
		v := mocks.NewAppValidator(t)
		uu := mocks.NewTransactionUsecase(t)
		uh := handler.NewTransactionGRPCHandler(uu, v)
		ctx := context.Background()
		ctx2 := context.WithValue(ctx, "id", uint(1))
		uu.On("GetTransactions", ctx2, mock.Anything, uint(1)).Return(nil, expectedErr)

		_, err := uh.GetAllTransactions(ctx2, &pbtransReq)

		assert.ErrorIs(t, err, expectedErr)
	})
}

func TestGRPCTransfer(t *testing.T) {
	t.Run("should return something if transfer success", func(t *testing.T) {
		v := mocks.NewAppValidator(t)
		uu := mocks.NewTransactionUsecase(t)
		uh := handler.NewTransactionGRPCHandler(uu, v)
		ctx := context.Background()
		ctx2 := context.WithValue(ctx, "id", uint(1))
		v.On("Validate", mock.Anything).Return(nil)
		uu.On("Transfer", ctx2, mock.Anything, uint(1)).Return(&transactions[0], nil)

		res, _ := uh.Transfer(ctx2, &pbtfReq)

		assert.NotNil(t, res)
	})

	t.Run("should return err when error in query", func(t *testing.T) {
		expectedErr := apperror.ErrCantConnectToThirdParty
		v := mocks.NewAppValidator(t)
		uu := mocks.NewTransactionUsecase(t)
		uh := handler.NewTransactionGRPCHandler(uu, v)
		ctx := context.Background()
		ctx2 := context.WithValue(ctx, "id", uint(1))
		v.On("Validate", mock.Anything).Return(nil)
		uu.On("Transfer", ctx2, mock.Anything, uint(1)).Return(nil, expectedErr)

		_, err := uh.Transfer(ctx2, &pbtfReq)

		assert.ErrorIs(t, err, expectedErr)
	})

	t.Run("should return err when invalid body", func(t *testing.T) {
		expectedErr := apperror.ErrInvalidBody
		v := mocks.NewAppValidator(t)
		uu := mocks.NewTransactionUsecase(t)
		uh := handler.NewTransactionGRPCHandler(uu, v)
		ctx := context.Background()
		ctx2 := context.WithValue(ctx, "id", uint(1))
		v.On("Validate", mock.Anything).Return(expectedErr)

		_, err := uh.Transfer(ctx2, &pbtfReq)

		assert.ErrorIs(t, err, expectedErr)
	})
}

func TestGRPCTopUp(t *testing.T) {

	t.Run("should return err when error in query", func(t *testing.T) {
		expectedErr := apperror.ErrCantConnectToThirdParty
		v := mocks.NewAppValidator(t)
		uu := mocks.NewTransactionUsecase(t)
		uh := handler.NewTransactionGRPCHandler(uu, v)
		ctx := context.Background()
		ctx2 := context.WithValue(ctx, "id", uint(1))
		v.On("Validate", mock.Anything).Return(nil)
		uu.On("TopUp", ctx2, mock.Anything, uint(1)).Return(nil, expectedErr)

		_, err := uh.TopUp(ctx2, &pbtuReq)

		assert.ErrorIs(t, err, expectedErr)
	})

	t.Run("should return err when invalid body", func(t *testing.T) {
		expectedErr := apperror.ErrInvalidBody
		v := mocks.NewAppValidator(t)
		uu := mocks.NewTransactionUsecase(t)
		uh := handler.NewTransactionGRPCHandler(uu, v)
		ctx := context.Background()
		ctx2 := context.WithValue(ctx, "id", uint(1))
		v.On("Validate", mock.Anything).Return(expectedErr)

		_, err := uh.TopUp(ctx2, &pbtuReq)

		assert.ErrorIs(t, err, expectedErr)
	})
}
