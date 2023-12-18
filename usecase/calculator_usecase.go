package usecase

import (
	"context"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/dto"
	proto "git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/pb"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/repository"
)

type EmergencyFundsUsecase interface {
	CalculateEmergencyFunds(context.Context, dto.EmergencyFundsReq) (*proto.EmergencyFundsRes, error)
}

type emergencyFundsUsecase struct {
	r repository.CalculatorRepository
}

func NewEmergencyFundsUsecase(r repository.CalculatorRepository) EmergencyFundsUsecase {
	return &emergencyFundsUsecase{
		r: r,
	}
}

func (eu *emergencyFundsUsecase) CalculateEmergencyFunds(ctx context.Context, req dto.EmergencyFundsReq) (*proto.EmergencyFundsRes, error) {
	res, err := eu.r.CalculateThirdService(req)
	if err != nil {
		return nil, err
	}
	return &proto.EmergencyFundsRes{
		RecommendedFunds: res.RecommendedFunds,
		MaritalStatus:    res.MaritalStatus.String(),
	}, nil
}
