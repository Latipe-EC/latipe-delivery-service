package userserv

import (
	"context"
	"delivery-service/adapter/userserv/dto"
	"github.com/stretchr/testify/mock"
)

type UserServiceMock struct {
	mock.Mock
}

func (us *UserServiceMock) GetAddressById(ctx context.Context, request *dto.GetAddressRequest) (*dto.GetAddressResponse, error) {
	args := us.Called(request)
	return args.Get(0).(*dto.GetAddressResponse), args.Error(1)
}

func (us *UserServiceMock) Authorization(ctx context.Context, req *dto.AuthorizationRequest) (*dto.AuthorizationResponse, error) {
	args := us.Called(req)
	return args.Get(0).(*dto.AuthorizationResponse), args.Error(1)
}

func (us *UserServiceMock) CreateNewAccount(ctx context.Context, req *dto.CreateAccountRequest) (*dto.CreateAccountResponse, error) {
	args := us.Called(req)
	return args.Get(0).(*dto.CreateAccountResponse), args.Error(1)
}
