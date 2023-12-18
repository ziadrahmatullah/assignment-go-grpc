package handler

import (
	"context"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/apperror"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/appvalidator"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/dto"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/pb"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/usecase"
)

type AuthGRPCHandler struct {
	pb.UnimplementedAuthServiceServer
	usecase   usecase.UserUsecase
	validator appvalidator.AppValidator
}

func NewAuthGRPCHandler(uu usecase.UserUsecase, val appvalidator.AppValidator) *AuthGRPCHandler {
	return &AuthGRPCHandler{
		usecase:   uu,
		validator: val,
	}
}

func (h *AuthGRPCHandler) Register(ctx context.Context, req *pb.RegisterReq) (*pb.RegisterRes, error) {
	userReq := dto.RegisterReq{
		Name:      req.Name,
		Birthdate: req.Birthdate,
		Email:     req.Email,
		Password:  req.Password,
	}
	err := h.validator.Validate(userReq)
	if err != nil {
		return nil, apperror.ErrInvalidBody
	}
	res, err := h.usecase.CreateUser(ctx, userReq)
	if err != nil {
		return nil, err
	}
	userRes := &pb.RegisterRes{
		Id:        uint32(res.ID),
		Name:      res.Name,
		Birthdate: res.Birthdate,
		Email:     res.Email,
	}
	return userRes, nil
}

func (h *AuthGRPCHandler) Login(ctx context.Context, data *pb.LoginReq) (*pb.LoginRes, error) {
	userReq := dto.LoginReq{
		Email:    data.Email,
		Password: data.Password,
	}
	err := h.validator.Validate(userReq)
	if err != nil {
		return nil, apperror.ErrInvalidBody
	}
	res, err := h.usecase.UserLogin(ctx, userReq)
	if err != nil {
		return nil, err
	}
	return &pb.LoginRes{AccessToken: res.AccessToken}, nil
}
