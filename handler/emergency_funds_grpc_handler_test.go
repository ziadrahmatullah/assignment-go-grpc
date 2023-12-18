package handler_test

import (
	"context"
	"testing"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/apperror"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/constant"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/dto"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/handler"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/mocks"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/pb"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

var protoEfReq = pb.EmergencyFundsReq{
	MonthlyIncome:             "100000",
	MonthlyExpense:            "20000",
	FinancialResponsibilities: "1000",
	MaritalStatus:             "Married",
	NumberOfChildren:          1,
}

var monthlyIncome, _ = decimal.NewFromString(protoEfReq.MonthlyIncome)
var monthlyExpense, _ = decimal.NewFromString(protoEfReq.MonthlyExpense)
var financialRes, _ = decimal.NewFromString(protoEfReq.FinancialResponsibilities)

var efReq = dto.EmergencyFundsReq{
	MonthlyIncome:             monthlyIncome,
	MonthlyExpense:            monthlyExpense,
	FinancialResponsibilities: financialRes,
	MaritalStatus:             constant.MaritalStatuses(protoEfReq.MaritalStatus),
	NumberOfChildren:          &protoEfReq.NumberOfChildren,
}

var efRes = pb.EmergencyFundsRes{
	RecommendedFunds: "10000",
	MaritalStatus:    "Married",
}

func TestCalculateEmergencyFunds(t *testing.T) {
	t.Run("should return something if emergency funds success", func(t *testing.T) {
		v := mocks.NewAppValidator(t)
		uu := mocks.NewEmergencyFundsUsecase(t)
		uh := handler.NewEmergencyFundsGRPCHandler(uu, v)
		ctx := context.Background()
		v.On("Validate", efReq).Return(nil)
		uu.On("CalculateEmergencyFunds", ctx, efReq).Return(&efRes, nil)

		res, _ := uh.CalculateEmergencyFunds(ctx, &protoEfReq)

		assert.NotNil(t, res)
	})
	t.Run("should return err when invalid body", func(t *testing.T) {
		expectedErr := apperror.ErrInvalidBody
		v := mocks.NewAppValidator(t)
		uu := mocks.NewEmergencyFundsUsecase(t)
		uh := handler.NewEmergencyFundsGRPCHandler(uu, v)
		ctx := context.Background()
		v.On("Validate", efReq).Return(expectedErr)

		_, err := uh.CalculateEmergencyFunds(ctx, &protoEfReq)

		assert.ErrorIs(t, err, expectedErr)
	})

	t.Run("should return err when error in query", func(t *testing.T) {
		expectedErr := apperror.ErrCantConnectToThirdParty
		v := mocks.NewAppValidator(t)
		uu := mocks.NewEmergencyFundsUsecase(t)
		uh := handler.NewEmergencyFundsGRPCHandler(uu, v)
		ctx := context.Background()
		v.On("Validate", efReq).Return(nil)
		uu.On("CalculateEmergencyFunds", ctx, efReq).Return(nil, expectedErr)


		_, err := uh.CalculateEmergencyFunds(ctx, &protoEfReq)

		assert.ErrorIs(t, err, expectedErr)
	})
}
