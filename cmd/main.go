package main

import (
	server "delivery-service/internal"
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	"net"
	"sync"
)

func main() {
	fmt.Println("Init application")

	serv, err := server.New()
	if err != nil {
		log.Fatalf("%s", err)
	}

	var wg sync.WaitGroup
	//subscriber
	wg.Add(1)
	go runWithRecovery(func() { serv.PurchaseCreatedSub().ListenPurchaseEvent(&wg) })

	//api handler
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := serv.App().Listen(serv.Config().Server.RestPort); err != nil {
			fmt.Printf("%s", err)
		}
	}()

	//grpc handler
	wg.Add(1)
	go runWithRecovery(func() {
		defer wg.Done()
		log.Infof("GRPC server run on localhost%v", serv.Config().GRPC.Port)
		lis, err := net.Listen("tcp", serv.Config().GRPC.Port)
		if err != nil {
			log.Fatalf("failed to listen: %v\n", err)
		}

		if err := serv.DeliServ().Serve(lis); err != nil {
			log.Infof("%s", err)
		}
	})

	wg.Wait()
}

func runWithRecovery(fn func()) {
	defer func() {
		if r := recover(); r != nil {
			log.Debugf("Recovered from panic: %v", r)
		}
	}()
	fn()
}
