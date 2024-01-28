package packageserv

import (
	"context"
	"delivery-service/internal/domain/entities"
	"delivery-service/internal/domain/message"
	"delivery-service/internal/domain/repos"
	"delivery-service/internal/publisher"
	"delivery-service/pkgs/mapper"
)

type ShippingPackageService struct {
	deliRepos     *repos.DeliveryRepos
	packagesRepos *repos.ShippingPackageRepos
	replyPub      *publisher.ReplyPurchasePublisher
}

func NewShippingPackageService(deliRepos *repos.DeliveryRepos,
	packagesRepos *repos.ShippingPackageRepos, replyPub *publisher.ReplyPurchasePublisher) *ShippingPackageService {
	return &ShippingPackageService{
		deliRepos:     deliRepos,
		packagesRepos: packagesRepos,
		replyPub:      replyPub,
	}
}
func (sr *ShippingPackageService) CreateShippingPackage(ctx context.Context, msg *message.CreateOrderMessage) error {
	newPackage := entities.ShippingPackage{}

	if err := mapper.BindingStruct(msg, &newPackage); err != nil {
		return err
	}

	//init reply message
	msgReply := message.CreateOrderReplyMessage{
		OrderID: msg.OrderID,
		Status:  message.COMMIT_SUCCESS,
	}

	// If saving the data into the repository fails, update the message status.
	_, err := sr.packagesRepos.CreateShippingPackages(ctx, &newPackage)
	if err != nil {
		msgReply.Status = message.COMMIT_FAIL
	}

	err = sr.replyPub.ReplyPurchaseMessage(&msgReply)

	return err
}
