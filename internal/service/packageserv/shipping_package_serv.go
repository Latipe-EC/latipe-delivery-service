package packageserv

import (
	"context"
	"delivery-service/internal/domain/entities"
	"delivery-service/internal/domain/message"
	"delivery-service/internal/domain/repos"
	"delivery-service/internal/publisher"
	"delivery-service/pkgs/mapper"
	"github.com/gofiber/fiber/v2/log"
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
	// Initialize new shipping package
	newPackage := entities.ShippingPackage{}

	// Initialize reply message
	msgReply := message.CreateOrderReplyMessage{
		OrderID: msg.OrderID,
		Status:  message.COMMIT_SUCCESS, // Assume success until an error occurs
	}

	// Defer the reply message publishing
	defer func(replyPub *publisher.ReplyPurchasePublisher, replyMsg *message.CreateOrderReplyMessage) {
		err := replyPub.ReplyPurchaseMessage(replyMsg)
		if err != nil {
			log.Error(err)
		}
	}(sr.replyPub, &msgReply)

	// Bind the message to the new package
	err := mapper.BindingStruct(msg, &newPackage)
	if err != nil {
		// If binding fails, log the error and return immediately
		log.Error("Failed to bind message to struct:", err)
		return err
	}

	// Save the data into the repository
	_, err = sr.packagesRepos.CreateShippingPackages(ctx, &newPackage)
	if err != nil {
		// If saving fails, update the reply message status
		msgReply.Status = message.COMMIT_FAIL
		log.Error("Failed to create shipping package:", err)
	}

	return err
}
