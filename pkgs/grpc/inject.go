package grpcclient

import "github.com/google/wire"

var Set = wire.NewSet(NewGrpcServerConnection)
