package main

import (
	"Andrew/Distributed/registry"
	"context"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.Handle("/services", &registry.RegistryService{})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var srv http.Server
	srv.Addr = registry.ServerPort

	go func() {
		log.Println(srv.ListenAndServe())
		cancel()
	}()

	go func() {
		fmt.Println("Registry service started. Press any key to stop.")
		var s string
		_, err := fmt.Scanln(&s)
		if err != nil {
			log.Println(err)
			cancel()
		}
		err = srv.Shutdown(ctx)
		if err != nil {
			log.Println(err)
			cancel()
		}
		cancel()
	}()

	<-ctx.Done()
	fmt.Println("Shutting down registry service")
}
