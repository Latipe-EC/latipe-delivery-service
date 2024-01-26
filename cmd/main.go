package main

import (
	server "delivery-service/internal"
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	"sync"
)

func main() {
	fmt.Println("Init application")

	serv, err := server.New()
	if err != nil {
		log.Fatalf("%s", err)
	}

	//subscriber
	var wg sync.WaitGroup

	//api handler
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := serv.App().Listen(serv.Config().Server.Port); err != nil {
			fmt.Printf("%s", err)
		}
	}()

	wg.Wait()
}
