package adapter

import (
	"delivery-service/internal/adapter/userserv"
	"github.com/google/wire"
)

var Set = wire.NewSet(userserv.NewUserService)
