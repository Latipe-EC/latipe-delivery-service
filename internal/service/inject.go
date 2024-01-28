package service

import (
	"delivery-service/internal/service/deliveryserv"
	"delivery-service/internal/service/packageserv"
	"delivery-service/internal/service/shippingserv"
	"github.com/google/wire"
)

var Set = wire.NewSet(
	shippingserv.NewShippingCostService,
	deliveryserv.NewDeliveryService,
	packageserv.NewShippingPackageService,
)
