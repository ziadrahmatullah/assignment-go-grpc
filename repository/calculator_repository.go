package repository

import (
	"context"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/apperror"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/dto"
	pb "git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/pb/emergency-funds"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/util"
	"google.golang.org/grpc"
)

type CalculatorRepository interface {
	CalculateThirdService(dto.EmergencyFundsReq) (*pb.CalculateEmergencyFundsResponse, error)
}

type calulatorRepository struct {
}

func NewCalculatorRepository() CalculatorRepository {
	return &calulatorRepository{}
}

func (cr *calulatorRepository) CalculateThirdService(req dto.EmergencyFundsReq) (*pb.CalculateEmergencyFundsResponse, error) {
	conn, err := grpc.Dial("localhost:8081", grpc.WithInsecure())
	if err != nil {
		return nil, apperror.ErrCantConnectToThirdParty
	}
	defer conn.Close()

	client := pb.NewCalculatorServiceClient(conn)

	userReq := &pb.CalculateEmergencyFundsRequest{
		MonthlyIncome:             req.MonthlyIncome.String(),
		MonthlyExpense:            req.MonthlyExpense.String(),
		FinancialResponsibilities: req.FinancialResponsibilities.String(),
		MaritalStatus:             util.ToMaritalStatusEnum(req.MaritalStatus),
	}

	ctx := context.Background()

	res, err := client.CalculateEmergencyFunds(ctx, userReq)
	if err != nil {
		return nil, apperror.ErrFailedToRequstThirdParty
	}

	return res, nil
}
