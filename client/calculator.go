package client

import (
	"context"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/apperror"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/dto"
	pb "git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/pb/emergency-funds"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/util"
	"google.golang.org/grpc"
)

func CalculateThirdService(req dto.EmergencyFundsReq) (*pb.CalculateEmergencyFundsResponse, error) {
	// Set up the gRPC client to connect to the server
	conn, err := grpc.Dial("localhost:8081", grpc.WithInsecure())
	if err != nil {
		return nil, apperror.ErrCantConnectToThirdParty
	}
	defer conn.Close()

	// Create a gRPC client
	client := pb.NewCalculatorServiceClient(conn)

	// Prepare the request message
	userReq := &pb.CalculateEmergencyFundsRequest{
		MonthlyIncome:             req.MonthlyIncome.String(),
		MonthlyExpense:            req.MonthlyExpense.String(),
		FinancialResponsibilities: req.FinancialResponsibilities.String(),
		MaritalStatus:             util.ToMaritalStatusEnum(req.MaritalStatus),
	}

	ctx := context.Background()

	// Send the request to the server
	res, err := client.CalculateEmergencyFunds(ctx, userReq)
	if err != nil {
		return nil, apperror.ErrFailedToRequstThirdParty
	}

	return res, nil
}