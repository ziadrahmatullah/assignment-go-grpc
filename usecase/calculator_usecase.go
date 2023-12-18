package usecase

import (
	"context"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/client"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/dto"
	proto "git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/pb"
)

type EmergencyFundsUsecase interface {
	CalculateEmergencyFunds(context.Context, dto.EmergencyFundsReq) (*proto.EmergencyFundsRes, error)
}

type emergencyFundsUsecase struct {
}

func NewEmergencyFundsUsecase() EmergencyFundsUsecase {
	return &emergencyFundsUsecase{}
}

func (eu *emergencyFundsUsecase) CalculateEmergencyFunds(ctx context.Context, req dto.EmergencyFundsReq) (*proto.EmergencyFundsRes, error) {
	res, err := client.CalculateThirdService(req)
	if err != nil {
		return nil, err
	}
	return &proto.EmergencyFundsRes{
		RecommendedFunds: res.RecommendedFunds,
		MaritalStatus:    res.MaritalStatus.String(),
	}, nil
}
