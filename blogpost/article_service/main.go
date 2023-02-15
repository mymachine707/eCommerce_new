package main

import (
	"fmt"
	"log"
	"net"

	"uacademy/blogpost/article_service/config" // docs is generated by Swag CLI, you have to import it.
	"uacademy/blogpost/article_service/protogen/blogpost"
	"uacademy/blogpost/article_service/services/article"
	"uacademy/blogpost/article_service/services/author"
	"uacademy/blogpost/article_service/storage"
	"uacademy/blogpost/article_service/storage/postgres"

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