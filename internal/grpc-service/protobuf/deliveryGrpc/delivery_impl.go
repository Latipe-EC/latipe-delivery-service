package deliveryGrpc

import (
	"context"
	"delivery-service/internal/domain/dto"
	"delivery-service/internal/service/shippingserv"
	"delivery-service/pkgs/mapper"
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type deliveryServer struct {
	deliService *shippingserv.ShippingCostService
	UnimplementedDeliveryServiceServer
}

func NewDeliveryServerGRPC(deliService *shippingserv.ShippingCostService) DeliveryServiceServer {
	return &deliveryServer{
		deliService: deliService,
	}
}

func (d deliveryServer) CalculateShippingCost(ctx context.Context, request *GetShippingCostRequest) (*GetShippingCostResponse, error) {
	req := dto.OrderShippingCostRequest{}
	response := GetShippingCostResponse{}

	if err := mapper.BindingStruct(request, &req); err != nil {
		log.Errorf("%v", err)
		return nil, status.Error(codes.Internal, fmt.Sprintf("%v", err))
	}

	if err := mapper.BindingStruct(request, &req); err != nil {
		log.Errorf("%v", err)
		return nil, status.Error(codes.Internal, fmt.Sprintf("%v", err))
	}

	resp, err := d.deliService.CalculateOrderShippingCost(ctx, &req)
	if err != nil {
		log.Errorf("%v", err)
		return nil, status.Error(codes.Internal, fmt.Sprintf("%v", err))
	}

	if err := mapper.BindingStructGrpc(resp, &response); err != nil {
		log.Errorf("%v", err)
		return nil, status.Error(codes.Internal, fmt.Sprintf("%v", err))
	}

	return &response, nil

}
