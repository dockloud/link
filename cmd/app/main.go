package main

//import (
//	"context"
//	"fmt"
//	"os"
//	"time"
//
//	cache "github.com/dockloud/link/internal/infrastructure/cache/redis"
//	"github.com/dockloud/link/internal/infrastructure/database/postgres"
//	"github.com/dockloud/link/internal/infrastructure/document"
//	"github.com/dockloud/link/internal/infrastructure/link"
//	"github.com/dockloud/link/internal/transport/grpc/client"
//	usecase "github.com/dockloud/link/internal/usecase/link"
//)
//
//func main() {
//	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
//	defer cancel()
//
//	docClient, err := client.NewDocumentClient(os.Getenv("DOC_CONN_ADDR"))
//	if err != nil {
//		panic(fmt.Errorf("failed to create document client: %w", err))
//	}
//
//	database, err := database.NewPostgresDB(os.Getenv("DATABASE_URL"))
//	if err != nil {
//		panic(fmt.Errorf("failed to connect to database: %w", err))
//	}
//
//	cache, err := cache.NewRedisCache(os.Getenv("CACHE_URL"))
//	if err != nil {
//		panic(fmt.Errorf("failed to connect to cache: %w", err))
//	}
//
//	linkRepo := link.NewRepository(database, cache)
//	docRepo := document.NewRepository(*docClient)
//	linkServ := usecase.NewService(linkRepo, docRepo)
//}
