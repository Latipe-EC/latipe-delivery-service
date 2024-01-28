package subscribers

import "github.com/google/wire"

var Set = wire.NewSet(NewPurchaseCreatedSub)
