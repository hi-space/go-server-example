package cmd

import (
	"context"
	"flag"
	"log"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/hi-space/go-grpc-mongodb/pkg/protocol/grpc"
	"github.com/hi-space/go-grpc-mongodb/pkg/service/v1"
)

// configuration for Server
type Config struct {
	// gRPC server start parameters section
	GRPCPort string

	// DB Datastore parameters section
	DBHost string
	DBPort string
	DBUser string
	DBPassword string
	DBName string
}

// runs gRPC server and HTTP gateway
func RunServer() error {
	ctx := context.Background()

	// get configuration
	var cfg Config
	flag.StringVar(&cfg.GRPCPort, "grpc-port", "9000", "gRPC port to bind")
	flag.StringVar(&cfg.DBHost, "db-host", "localhost", "Database host")
	flag.StringVar(&cfg.DBPort, "db-port", "27017", "Database port")
	flag.StringVar(&cfg.DBUser, "db-user", "yoo", "Database user")
	flag.StringVar(&cfg.DBPassword, "db-password", "integra1", "Database password")
	flag.StringVar(&cfg.DBName, "db-name", "test", "Database name")
	flag.Parse()

	if len(cfg.GRPCPort) == 0 {
		return fmt.Errorf("invalid TCP port for gRPC server: '%s'", cfg.GRPCPort)
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

	return grpc.RunServer(ctx, v1API, cfg.GRPCPort)
}