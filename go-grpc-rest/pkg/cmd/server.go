package cmd

import (
	"context"
	"flag"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"go-grpc-rest/pkg/protocol/grpc"
	"go-grpc-rest/pkg/protocol/rest"
	v1 "go-grpc-rest/pkg/service"
)

// configuration for Server
type Config struct {
	// gRPC server start parameters section
	GRPCPort string
	HTTPPort string

	// DB Datastore parameters section
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

// runs gRPC server and HTTP gateway
func RunServer() error {
	ctx := context.Background()

	// get configuration
	var cfg Config
	flag.StringVar(&cfg.GRPCPort, "grpc-port", "9000", "gRPC port to bind")
	flag.StringVar(&cfg.HTTPPort, "http-port", "8000", "HTTP port to bind")
	flag.StringVar(&cfg.DBHost, "db-host", "localhost", "Database host")
	flag.StringVar(&cfg.DBPort, "db-port", "27017", "Database port")
	flag.StringVar(&cfg.DBUser, "db-user", "yoo", "Database user")
	flag.StringVar(&cfg.DBPassword, "db-password", "integra1", "Database password")
	flag.StringVar(&cfg.DBName, "db-name", "test", "Database name")
	flag.Parse()

	if len(cfg.GRPCPort) == 0 {
		return fmt.Errorf("invalid TCP port for gRPC server: '%s'", cfg.GRPCPort)
	}

	if len(cfg.HTTPPort) == 0 {
		return fmt.Errorf("invalid TCP port for HTTP gateway: '%s'", cfg.HTTPPort)
	}

	dsn := fmt.Sprintf("mongodb://%s:%s", cfg.DBHost, cfg.DBPort)
	// credential := options.Credential{
	// 	Username: cfg.DBUser,
	// 	Password: cfg.DBPassword,
	// }

	// clientOptions := options.Client().ApplyURI(dsn).SetAuth(credential)

	clientOptions := options.Client().ApplyURI(dsn)

	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Could not connect to MongoDB: %v\n", err)
	} else {
		fmt.Println("Connected to Mongodb")
	}

	db := client.Database(cfg.DBName)

	v1API := v1.NewToDoServiceServer(db)

	// run HTTP gateway
	go func() {
		_ = rest.RunServer(ctx, cfg.GRPCPort, cfg.HTTPPort)
	}()

	return grpc.RunServer(ctx, v1API, cfg.GRPCPort)
}
