package main

import (
	"log"
	"net"
	"os"

	pb "github.com/dockloud/contract/gen/link/v1"
	cache "github.com/dockloud/link/internal/infrastructure/cache/redis"
	database "github.com/dockloud/link/internal/infrastructure/database/postgres"
	"github.com/dockloud/link/internal/infrastructure/document"
	"github.com/dockloud/link/internal/infrastructure/link"
	"github.com/dockloud/link/internal/transport/grpc/auth"
	"github.com/dockloud/link/internal/transport/grpc/client"
	"github.com/dockloud/link/internal/transport/grpc/server"
	usecase "github.com/dockloud/link/internal/usecase/link"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

func main() {
	docClient, err := client.NewDocumentClient(
		os.Getenv("DOC_CONN_ADDR"),
		[]byte(os.Getenv("LINK_SERVICE_KEY")),
	)
	if err != nil {
		log.Fatalf("failed to create document client: %v", err)
	}

	db, err := database.NewPostgresDB(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}

	cch, err := cache.NewRedisCache(os.Getenv("CACHE_URL"))
	if err != nil {
		log.Fatalf("failed to connect to cache: %v", err)
	}

	linkRepo := link.NewRepository(db, cch)
	docRepo := document.NewRepository(*docClient)
	linkService := usecase.NewService(linkRepo, docRepo)

	lis, err := net.Listen("tcp", ":50150")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	keyring := auth.Keyring{
		"dck-gateway-v1": {
			ID:      "dck-gateway-v1",
			Service: "dockloud-gateway",
			Secret:  []byte(os.Getenv("GATEWAY_SERVICE_KEY")),
		},
	}
	authInterceptor, err := auth.NewKeyringInterceptor(keyring)
	if err != nil {
		log.Fatalf("failed to create interceptor: %v", err)
	}

	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(authInterceptor))
	healthServer := health.NewServer()

	healthpb.RegisterHealthServer(grpcServer, healthServer)
	healthServer.SetServingStatus("", healthpb.HealthCheckResponse_SERVING)

	pb.RegisterLinkServiceServer(grpcServer, server.NewLinkServer(*linkService))
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
