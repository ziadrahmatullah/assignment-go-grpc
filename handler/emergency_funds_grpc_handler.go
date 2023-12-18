package handler

import (
	"context"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/apperror"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/appvalidator"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/constant"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/dto"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/pb"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/usecase"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/util"
	"github.com/shopspring/decimal"
)

type EmerGencyFundsGRPCHandler struct {
	pb.UnimplementedEmergencyFundsServiceServer
	usecase   usecase.EmergencyFundsUsecase
	validator appvalidator.AppValidator
}

func NewEmergencyFundsGRPCHandler(uu usecase.EmergencyFundsUsecase, valid appvalidator.AppValidator) *EmerGencyFundsGRPCHandler {
	return &EmerGencyFundsGRPCHandler{
		usecase:   uu,
		validator: valid,
	}
}

func (eh *EmerGencyFundsGRPCHandler) CalculateEmergencyFunds(ctx context.Context, req *pb.EmergencyFundsReq) (*pb.EmergencyFundsRes, error) {
	monthlyIncome, err := decimal.NewFromString(req.MonthlyIncome)
	if err != nil {
		return nil, apperror.ErrInvalidBody
	}
	monthlyExpense, err := decimal.NewFromString(req.MonthlyExpense)
	if err != nil {
		return nil, apperror.ErrInvalidBody
	}
	financialRes, err := decimal.NewFromString(req.FinancialResponsibilities)
	if err != nil {
		return nil, apperror.ErrInvalidBody
	}
	userReq := dto.EmergencyFundsReq{
		MonthlyIncome:             monthlyIncome,
		MonthlyExpense:            monthlyExpense,
		FinancialResponsibilities: financialRes,
		MaritalStatus:             constant.MaritalStatuses(req.MaritalStatus),
		NumberOfChildren:          uint(req.NumberOfChildren),
	}
	err = eh.validator.Validate(userReq)
	if err != nil {
		return nil, apperror.ErrInvalidBody
	}
	if !util.IsValidNumChildren(userReq.NumberOfChildren, userReq.MaritalStatus) {
		return nil, apperror.ErrInvalidChildren
	}
	if !util.IsValidDecimal(userReq.MonthlyIncome) {
		return nil, apperror.ErrInvalidMonthlyIncome
	}
	if !util.IsValidDecimal(userReq.MonthlyExpense) {
		return nil, apperror.ErrInvalidMonthlyExpense
	}
	if !util.IsValidDecimal(userReq.FinancialResponsibilities) {
		return nil, apperror.ErrInvalidFinancialAbilities
	}
	res, err := eh.usecase.CalculateEmergencyFuns(ctx, userReq)
	if err != nil {
		return nil, err
	}
	return &pb.EmergencyFundsRes{
		RecommendedFunds: res.RecommendedFunds,
	}, nil
}
