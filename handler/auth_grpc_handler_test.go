package handler_test

import (
	"context"
	"testing"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/apperror"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/handler"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/mocks"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/pb"
	"github.com/stretchr/testify/assert"
)

var protoLoginReq = &pb.LoginReq{
	Email:    "alice@gmail.com",
	Password: "alice123",
}

var protoRegisterReq = &pb.RegisterReq{
	Name:      "Alice",
	Birthdate: "2001-03-03",
	Email:     "alice@gmail.com",
	Password:  "alice123",
}

func TestHandleGRPCUserRegister(t *testing.T) {
	t.Run("should return something if register success", func(t *testing.T) {
		v := mocks.NewAppValidator(t)
		uu := mocks.NewUserUsecase(t)
		uh := handler.NewAuthGRPCHandler(uu, v)
		ctx := context.Background()
		v.On("Validate", registerReq[0]).Return(nil)
		uu.On("CreateUser", ctx, registerReq[0]).Return(&registerRes[0], nil)

		res, _ := uh.Register(ctx, protoRegisterReq)

		assert.NotNil(t, res)
	})
	t.Run("should return err when invalid body", func(t *testing.T) {
		expectedErr := apperror.ErrInvalidBody
		v := mocks.NewAppValidator(t)
		uu := mocks.NewUserUsecase(t)
		uh := handler.NewAuthGRPCHandler(uu, v)
		ctx := context.Background()
		v.On("Validate", registerReq[0]).Return(expectedErr)

		_, err := uh.Register(ctx, protoRegisterReq)

		assert.ErrorIs(t, err, expectedErr)
	})

	t.Run("should return err when error in query", func(t *testing.T) {
		expectedErr := apperror.ErrNewUserQuery
		v := mocks.NewAppValidator(t)
		uu := mocks.NewUserUsecase(t)
		uh := handler.NewAuthGRPCHandler(uu, v)
		ctx := context.Background()
		v.On("Validate", registerReq[0]).Return(nil)
		uu.On("CreateUser", ctx, registerReq[0]).Return(nil, expectedErr)

		_, err := uh.Register(ctx, protoRegisterReq)

		assert.ErrorIs(t, err, expectedErr)
	})
}

func TestHandleGRPCUserLogin(t *testing.T) {
	t.Run("should return something if register success", func(t *testing.T) {
		v := mocks.NewAppValidator(t)
		uu := mocks.NewUserUsecase(t)
		uh := handler.NewAuthGRPCHandler(uu, v)
		ctx := context.Background()
		v.On("Validate", loginReq[0]).Return(nil)
		uu.On("UserLogin", ctx, loginReq[0]).Return(&loginRes[0], nil)

		res, _ := uh.Login(ctx, protoLoginReq)

		assert.NotNil(t, res)
	})
	t.Run("should return err when invalid body", func(t *testing.T) {
		expectedErr := apperror.ErrInvalidBody
		v := mocks.NewAppValidator(t)
		uu := mocks.NewUserUsecase(t)
		uh := handler.NewAuthGRPCHandler(uu, v)
		ctx := context.Background()
		v.On("Validate", loginReq[0]).Return(expectedErr)

		_, err := uh.Login(ctx, protoLoginReq)

		assert.ErrorIs(t, err, expectedErr)
	})

	t.Run("should return err when error in query", func(t *testing.T) {
		expectedErr := apperror.ErrNewUserQuery
		v := mocks.NewAppValidator(t)
		uu := mocks.NewUserUsecase(t)
		uh := handler.NewAuthGRPCHandler(uu, v)
		ctx := context.Background()
		v.On("Validate", loginReq[0]).Return(nil)
		uu.On("UserLogin", ctx, loginReq[0]).Return(nil, expectedErr)

		_, err := uh.Login(ctx, protoLoginReq)

		assert.ErrorIs(t, err, expectedErr)
	})
}
