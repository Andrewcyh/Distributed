package service

import (
	"Andrew/Distributed/registry"
	"context"
	"fmt"
	"log"
	"net/http"
)

func Start(ctx context.Context, host, port string, reg registry.Registration,
	registerHandlersFunc func()) (context.Context, error) {
	registerHandlersFunc()
	ctx = startService(ctx, reg.ServiceName, host, port)
	err := registry.RegisterService(reg)
	if err != nil {
		return ctx, err
	}

	return ctx, nil
}

func startService(ctx context.Context, serviceName registry.ServiceName, host, port string) context.Context {
	ctx, cancel := context.WithCancel(ctx)
	serviceAddress := fmt.Sprintf("http://%s:%s", host, port)
	var srv http.Server
	srv.Addr = host + ":" + port

	go func() {
		log.Println(srv.ListenAndServe())
		err := registry.ShutdownService(serviceAddress)
		if err != nil {
			log.Println(err)
		}
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

		err = registry.ShutdownService(serviceAddress)
		if err != nil {
			log.Println(err)
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
