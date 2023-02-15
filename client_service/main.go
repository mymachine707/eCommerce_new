package main

import (
	"eCommerce/config"
	"eCommerce/services/article"
	"eCommerce/services/author"
	"eCommerce/storage"
	"eCommerce/storage/postgres"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	cfg := config.Load()

	psqlConnString := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresDatabase,
	)
	var err error
	var stg storage.StorageI
	stg, err = postgres.InitDB(psqlConnString)
	if err != nil {
		panic(err)
	}

	println("gRPC server tutorial in Go")

	listener, err := net.Listen("tcp", cfg.GRPCPort)
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()

	authorService := &author.AuthorService{}
	blogpost.RegisterAuthorServiceServer(srv, authorService)

	articleService := article.NewArticleService(stg)
	blogpost.RegisterArticleServiceServer(srv, articleService)

	reflection.Register(srv)

	if err := srv.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
