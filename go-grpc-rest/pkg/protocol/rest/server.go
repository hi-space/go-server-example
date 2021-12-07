package rest

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"

	v1 "go-grpc-rest/pkg/api"
)

func serveSwagger(w http.ResponseWriter, r *http.Request) {
	log.Println("serverSwagger")
	http.ServeFile(w, r, "go-grpc-rest/api/swagger/todo.swagger.json")
}

// RunServer runs HTTP/REST gateway
func RunServer(ctx context.Context, grpcPort, httpPort string) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	rmux := runtime.NewServeMux()

	conn, err := grpc.DialContext(
		context.Background(),
		"localhost:"+grpcPort,
		grpc.WithBlock(),
		grpc.WithInsecure(),
	)

	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	if err := v1.RegisterToDoServiceHandler(ctx, rmux, conn); err != nil {
		log.Fatalf("failed to start HTTP gateway: %v", err)
	}

	srv := &http.Server{
		Addr:    ":" + httpPort,
		Handler: rmux,
	}

	// Serve the swagger-ui and swagger file
	// mux := http.NewServeMux()
	// mux.Handle("/", rmux)
	// mux.HandleFunc("/swagger.json", serveSwagger)
	// fs := http.FileServer(http.Dir("www/swagger-ui"))
	// mux.Handle("/swagger-ui", http.StripPrefix("/swagger-ui", fs))

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// sig is a ^C, handle it
		}

		_, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		_ = srv.Shutdown(ctx)
	}()

	log.Println("starting HTTP/REST gateway...")
	return srv.ListenAndServe()
}
