package userserv

import (
	"context"
	dto2 "delivery-service/internal/adapter/userserv/dto"
	"github.com/stretchr/testify/mock"
)

type UserServiceMock struct {
	mock.Mock
}

func (us *UserServiceMock) GetAddressById(ctx context.Context, request *dto2.GetAddressRequest) (*dto2.GetAddressResponse, error) {
	args := us.Called(request)
	return args.Get(0).(*dto2.GetAddressResponse), args.Error(1)
}

func (us *UserServiceMock) Authorization(ctx context.Context, req *dto2.AuthorizationRequest) (*dto2.AuthorizationResponse, error) {
	args := us.Called(req)
	return args.Get(0).(*dto2.AuthorizationResponse), args.Error(1)
}

func (us *UserServiceMock) CreateNewAccount(ctx context.Context, req *dto2.CreateAccountRequest) (*dto2.CreateAccountResponse, error) {
	args := us.Called(req)
	return args.Get(0).(*dto2.CreateAccountResponse), args.Error(1)
}
