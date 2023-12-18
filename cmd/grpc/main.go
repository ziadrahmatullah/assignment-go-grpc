package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/appvalidator"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/database"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/handler"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/middleware"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/pb"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/repository"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/usecase"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("no env got")
	}
	db, err := database.ConnectDB()
	if err != nil{
		log.Println(err.Error())
	}
	v := appvalidator.NewAppValidatorImpl()
	appvalidator.SetValidator(v)

	wr := repository.NewWalletRepository(db)
	ar := repository.NewAttemptRepository(db)

	ur := repository.NewUserRepository(db)
	uu := usecase.NewUserUsecase(ur, wr, ar)
	uh := handler.NewAuthGRPCHandler(uu, v)

	tr := repository.NewTransactionRepository(db)
	tu := usecase.NewTransactionUsecase(tr, wr)
	th := handler.NewTransactionGRPCHandler(tu, v)

	eu := usecase.NewEmergencyFundsUsecase()
	eh := handler.NewEmergencyFundsGRPCHandler(eu, v)

	list, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Println("error starting tcp server")
	}

	server := grpc.NewServer(grpc.ChainUnaryInterceptor(
		middleware.LoggerInterceptor,
		middleware.ErrorInterceptor,
		middleware.AuthInterceptor,
	))

	pb.RegisterAuthServiceServer(server, uh)
	pb.RegisterTransactionServiceServer(server, th)
	pb.RegisterEmergencyFundsServiceServer(server, eh)

	log.Println("starting grpc server")

	signCh := make(chan os.Signal, 1)
	signal.Notify(signCh, os.Interrupt, syscall.SIGTERM)

	go func() {
		if err := server.Serve(list); err != nil {
			signCh <- syscall.SIGINT
		}
	}()
	log.Println("server started")
	<-signCh
	log.Println("server stopped")
}
