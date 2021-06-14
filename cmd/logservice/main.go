package main

import (
	"Andrew/Distributed/log"
	"Andrew/Distributed/service"
	"context"
	"fmt"
	stlog "log"
)

func main() {
	log.Run("./distributed.log")
	host, port := "localhost", "8080"
	ctx, err := service.Start(
		context.Background(),
		"Log Service",
		host,
		port,
		log.RegisterHandlers,
	)
	if err != nil {
		stlog.Fatalln(err)
	}
	<-ctx.Done()

	fmt.Println("Shutting down log service.")
}
