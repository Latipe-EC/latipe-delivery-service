package userserv

import (
	"context"
	"delivery-service/adapter/userserv/dto"
	"delivery-service/config"
	"delivery-service/pkgs/mapper"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v2/log"
)

type UserService struct {
	restyClient *resty.Client
	config      *config.Config
}

func NewUserService(client *resty.Client, cfg *config.Config) UserService {
	return UserService{restyClient: client, config: cfg}
}

func (us UserService) GetAddressById(ctx context.Context, request *dto.GetAddressRequest) (*dto.GetAddressResponse, error) {
	resp, err := us.restyClient.
		SetBaseURL(us.config.AdapterService.UserService.UserURL).
		R().
		SetContext(ctx).
		SetDebug(false).
		Get(request.URL() + fmt.Sprintf("/%v", request.AddressId))

	if err != nil {
		log.Errorf("[%s] [Get address]: %s", "ERROR", err)
		return nil, err
	}

	if resp.StatusCode() >= 500 {
		log.Errorf("[%s] [Get address]: %s", "ERROR", resp.Body())
		return nil, errors.New("get address internal")
	}

	var regResp *dto.GetAddressResponse
	err = mapper.BindingStruct(resp.Body(), &regResp)
	if err != nil {
		log.Errorf("[%s] [Get product]: %s", "ERROR", err)
		return nil, err
	}

	return regResp, nil
}

func (us UserService) Authorization(ctx context.Context, req *dto.AuthorizationRequest) (*dto.AuthorizationResponse, error) {
	resp, err := us.restyClient.
		SetBaseURL(us.config.AdapterService.UserService.AuthURL).
		R().
		SetBody(req).
		SetContext(ctx).
		SetDebug(false).
		Post(req.URL())

	if err != nil {
		log.Errorf("[%s] [Authorize token]: %s", "ERROR", err)
		return nil, err
	}

	if resp.StatusCode() >= 500 {
		log.Errorf("[%s] [Authorize token]: %s", "ERROR", resp.Body())
		return nil, err
	}

	var regResp *dto.AuthorizationResponse

	if err := json.Unmarshal(resp.Body(), &regResp); err != nil {
		log.Errorf("[%s] [Authorize token]: %s", "ERROR", err)
		return nil, err
	}

	return regResp, nil
}

func (us UserService) CreateNewAccount(ctx context.Context, req *dto.CreateAccountRequest) (*dto.CreateAccountResponse, error) {
	resp, err := us.restyClient.
		SetBaseURL(us.config.AdapterService.UserService.UserURL).
		R().
		SetBody(req.Body).
		SetContext(ctx).
		SetDebug(true).
		SetHeader("Authorization", fmt.Sprintf("Bearer %v", req.BearerToken)).
		Post(req.URL())

	if err != nil {
		log.Errorf("[Create account]: %s", err)
		return nil, err
	}

	if resp.StatusCode() >= 500 {
		log.Errorf("[Create account]: %s", resp.Body())
		return nil, errors.New("internal server error")
	}

	if resp.StatusCode() >= 400 {
		log.Errorf(" [Create account]: %s", resp.Body())
		return nil, errors.New("bad request error")
	}

	var regResp *dto.CreateAccountResponse

	if err := json.Unmarshal(resp.Body(), &regResp); err != nil {
		log.Errorf("[Create account]: %s", err)
		return nil, err
	}

	return regResp, nil
}
