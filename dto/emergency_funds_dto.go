package dto

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/constant"
	"github.com/shopspring/decimal"
)

type EmergencyFundsReq struct {
	MonthlyIncome             decimal.Decimal          `binding:"required" validate:"required"`
	MonthlyExpense            decimal.Decimal          `binding:"required" validate:"required"`
	FinancialResponsibilities decimal.Decimal          `binding:"required" validate:"required"`
	MaritalStatus             constant.MaritalStatuses `binding:"required" validate:"required"`
	NumberOfChildren          *uint32                  `binding:"required,min=0" validate:"required,min=0"`
}
