package protobuf

import (
	"delivery-service/internal/grpc-service/protobuf/deliveryGrpc"
	"github.com/google/wire"
)

var Set = wire.NewSet(deliveryGrpc.NewDeliveryServerGRPC)
