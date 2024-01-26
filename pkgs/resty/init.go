package restyclient

import (
	"delivery-service/config"
	"github.com/go-resty/resty/v2"
	"github.com/google/wire"
	"time"
)

var Set = wire.NewSet(InitRestyGoClient)

func InitRestyGoClient(config *config.Config) *resty.Client {
	cli := resty.New().
		SetDebug(false).
		SetTimeout(5 * time.Second)
	return cli
}
