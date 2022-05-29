package rest

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path"
	"strings"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

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

	// Serve the swagger-ui and swagger file
	mux := http.NewServeMux()
	mux.Handle("/", rmux)
	mux.HandleFunc("/swagger/", serveSwaggerFile)
	serveSwaggerUI(mux)

	// mux.HandleFunc("/swagger.json", serveSwaggerFile)
	// fs := http.FileServer(http.Dir("www/swagger-ui"))
	// mux.Handle("/swagger-ui", http.StripPrefix("/swagger-ui", fs))

	srv := &http.Server{
		Addr:    ":" + httpPort,
		Handler: mux,
	}

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

func serveSwaggerFile(w http.ResponseWriter, r *http.Request) {
	log.Println("start serveSwaggerFile")

	if !strings.HasSuffix(r.URL.Path, "swagger.json") {
		log.Printf("Not Found: %s", r.URL.Path)
		http.NotFound(w, r)
		return
	}

	p := strings.TrimPrefix(r.URL.Path, "/swagger/")
	p = path.Join("api/swagger", p)

	log.Printf("Serving swagger-file: %s", p)

	http.ServeFile(w, r, p)
}

func serveSwaggerUI(mux *http.ServeMux) {
	fileServer := http.FileServer(http.Dir("www/swagger-ui"))
	prefix := "/swagger-ui/"

	mux.Handle(prefix, http.StripPrefix(prefix, fileServer))
}
