package handler

import (
	"context"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/apperror"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/appvalidator"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/dto"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/model"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/pb"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/usecase"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/util"
	"github.com/shopspring/decimal"
)

type TransactionGRPCHandler struct {
	pb.UnimplementedTransactionServiceServer
	usecase   usecase.TransactionUsecase
	validator appvalidator.AppValidator
}

func NewTransactionGRPCHandler(uu usecase.TransactionUsecase, val appvalidator.AppValidator) *TransactionGRPCHandler {
	return &TransactionGRPCHandler{
		usecase:   uu,
		validator: val,
	}
}

func (h *TransactionGRPCHandler) GetAllTransactions(ctx context.Context, req *pb.TransactionsReq) (*pb.TransactionPaginationRes, error) {
	var dto dto.ListTransactionsReq
	dto.Search = &req.Search
	dto.FilterStart = &req.FilterStart
	dto.FilterEnd = &req.FilterEnd
	dto.SortBy = &req.SortBy
	dto.SortType = &req.SortType
	dto.PaginationLimit = &req.PaginationLimit
	dto.PaginationPage = &req.PaginationPage
	userId := ctx.Value("id").(uint)
	res, err := h.usecase.GetTransactions(ctx, dto, userId)
	if err != nil {
		return nil, err
	}
	var txRes pb.TransactionPaginationRes
	for _, tx := range res.Data {
		txG := pb.TransactionRes{
			Id:              uint32(tx.ID),
			CreatedAt:       tx.CreatedAt.Format("2006-01-02"),
			TransactionType: util.ToTransactionTypeEnum(tx.TransactionType),
			Sender:          tx.Sender,
			Receiver:        tx.Receiver,
			Amount:          tx.Amount.String(),
			Description:     tx.Description,
		}
		if tx.SourceOfFund != nil {
			txG.SourceOfFund = util.ToSourceOfFundEnum(*tx.SourceOfFund)
		}
		txRes.Data = append(txRes.Data, &txG)
	}

	return &pb.TransactionPaginationRes{
		Data:      txRes.Data,
		TotalDate: int32(res.TotalData),
		TotalPage: int32(res.TotalPage),
		Page:      int32(res.Page),
	}, nil
}
func (h *TransactionGRPCHandler) Transfer(ctx context.Context, req *pb.TransferReq) (*pb.TransactionRes, error) {
	amount, err := decimal.NewFromString(req.Amount)
	if err != nil {
		return nil, apperror.ErrInvalidBody
	}
	txReq := dto.TransferReq{
		WalletNumber: req.WalletNumber,
		Amount:       amount,
		Description:  req.Description,
	}
	err = h.validator.Validate(txReq)
	if err != nil {
		return nil, apperror.ErrInvalidBody
	}
	if !util.IsTransferAmountValid(txReq.Amount) {
		return nil, apperror.ErrInvalidAmount
	}
	userId := ctx.Value("id").(uint)
	res, err := h.usecase.Transfer(ctx, txReq, userId)
	if err != nil {
		return nil, err
	}
	return &pb.TransactionRes{
		Id:              uint32(res.ID),
		CreatedAt:       res.CreatedAt.Format("2006-01-02"),
		TransactionType: util.ToTransactionTypeEnum(res.TransactionType),
		Sender:          res.Sender,
		Receiver:        res.Receiver,
		Amount:          res.Amount.String(),
		Description:     res.Description,
	}, nil

}
func (h *TransactionGRPCHandler) TopUp(ctx context.Context, req *pb.TopUpReq) (*pb.TransactionRes, error) {
	amount, err := decimal.NewFromString(req.Amount)
	if err != nil {
		return nil, apperror.ErrInvalidBody
	}
	txReq := dto.TopUpReq{
		Amount:       amount,
		SourceOfFund: req.SourceOfFund,
	}
	err = h.validator.Validate(txReq)
	if err != nil {
		return nil, apperror.ErrInvalidBody
	}
	if !util.IsTopUpAmountValid(txReq.Amount) {
		return nil, apperror.ErrInvalidAmount
	}
	if !model.IsSourceOfFundValid(txReq.SourceOfFund) {
		return nil, apperror.ErrInvalidSourceOfFund
	}
	userId := ctx.Value("id").(uint)
	res, err := h.usecase.TopUp(ctx, txReq, userId)
	if err != nil {
		return nil, err
	}
	return &pb.TransactionRes{
		Id:              uint32(res.ID),
		CreatedAt:       res.CreatedAt.Format("2006-01-02"),
		TransactionType: util.ToTransactionTypeEnum(res.TransactionType),
		SourceOfFund:    util.ToSourceOfFundEnum(*res.SourceOfFund),
		Receiver:        res.Receiver,
		Amount:          res.Amount.String(),
		Description:     res.Description,
	}, nil
}
