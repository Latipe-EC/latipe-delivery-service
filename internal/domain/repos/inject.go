package repos

import "github.com/google/wire"

var Set = wire.NewSet(
	InitWardRepository,
	InitProvinceRepository,
	InitDistrictRepository,
	NewDeliveryRepos,
	NewShippingPackageRepos,
)
