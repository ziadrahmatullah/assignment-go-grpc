package dto

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/constant"
	"github.com/shopspring/decimal"
)

type EmergencyFundsReq struct {
	MonthlyIncome             decimal.Decimal          `validate:"required"`
	MonthlyExpense            decimal.Decimal          `validate:"required"`
	FinancialResponsibilities decimal.Decimal          `validate:"required"`
	MaritalStatus             constant.MaritalStatuses `validate:"required"`
	NumberOfChildren          uint                     `validate:"required"`
}
