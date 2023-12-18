package usecase

import (
	"context"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/client"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/dto"
	pb "git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/pb/emergency-funds"
)

type EmergencyFundsUsecase interface {
	CalculateEmergencyFuns(context.Context, dto.EmergencyFundsReq) (*pb.CalculateEmergencyFundsResponse, error)
}

type emergencyFundsUsecase struct {
}

func NewEmergencyFundsUsecase() EmergencyFundsUsecase {
	return &emergencyFundsUsecase{}
}

func (eu *emergencyFundsUsecase) CalculateEmergencyFuns(ctx context.Context, req dto.EmergencyFundsReq) (*pb.CalculateEmergencyFundsResponse, error) {
	return client.CalculateThirdService(req)
}
