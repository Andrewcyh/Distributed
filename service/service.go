package service

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

func Start(ctx context.Context, serviceName, host, port string,
	registerHandlersFunc func()) (context.Context, error) {
	ctx = startService(ctx, serviceName, host, port)
	registerHandlersFunc()

	return ctx, nil
}

func startService(ctx context.Context, serviceName, host, port string) context.Context {
	ctx, cancel := context.WithCancel(ctx)

	var srv http.Server
	srv.Addr = host + ":" + port

	go func() {
		log.Println(srv.ListenAndServe())
		cancel()
	}()

	go func() {
		fmt.Printf("%v started. Press any key to stop.\n", serviceName)
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

	return ctx
}
