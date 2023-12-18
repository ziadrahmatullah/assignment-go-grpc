package usecase_test

import (
	"context"
	"testing"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/apperror"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/constant"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/dto"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/mocks"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/pb"
	emergency_funds "git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/pb/emergency-funds"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/usecase"
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

var newEf = emergency_funds.CalculateEmergencyFundsResponse{
	RecommendedFunds: "10000",
	MaritalStatus:    emergency_funds.MaritalStatus_MARITAL_STATUS_MARRIED,
}

func TestCalculateEmergencyFunds(t *testing.T) {
	t.Run("should return not nil when success", func(t *testing.T) {
		r := mocks.NewCalculatorRepository(t)
		u := usecase.NewEmergencyFundsUsecase(r)
		ctx := context.Background()
		r.On("CalculateThirdService", efReq).Return(&newEf, nil)

		resUsers, _ := u.CalculateEmergencyFunds(ctx, efReq)

		assert.NotNil(t, resUsers)
	})

	t.Run("should return err when fail", func(t *testing.T) {
		expectedErr := apperror.ErrCantConnectToThirdParty
		r := mocks.NewCalculatorRepository(t)
		u := usecase.NewEmergencyFundsUsecase(r)
		ctx := context.Background()
		r.On("CalculateThirdService", efReq).Return(nil, expectedErr)

		_, err := u.CalculateEmergencyFunds(ctx, efReq)

		assert.ErrorIs(t, err, expectedErr)
	})
}
