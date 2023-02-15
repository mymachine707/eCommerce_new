package clients

import (
	"context"
	"eCommerce/eCommerce"
	"eCommerce/storage"
	"log"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type clientsService struct {
	stg storage.StorageI
	eCommerce.UnimplementedClientServiceServer
}

// NewArticleService ...
func NewArticleService(stg storage.StorageI) *clientsService {
	return &clientsService{
		stg: stg,
	}
}

// Ping ...
func (s *clientsService) Ping(ctx context.Context, req *eCommerce.Empty) (*eCommerce.Pong, error) {
	log.Println("Ping")

	return &eCommerce.Pong{
		Message: "OK",
	}, nil
}

// CreateArticle ...
func (s *clientsService) CreateArticle(ctx context.Context, req *eCommerce.CreateArticleRequest) (*eCommerce.Article, error) {
	id := uuid.New()

	err := s.stg.AddArticle(id.String(), req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.AddArticle: %s", err.Error())
	}

	article, err := s.stg.GetArticleByID(id.String())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.Stg.GetArticleByID: %s", err.Error())
	}

	return &eCommerce.Article{
		Id:        article.Id,
		Content:   article.Content,
		AuthorId:  article.Author.Id,
		CreatedAt: article.CreatedAt,
		UpdatedAt: article.UpdatedAt,
	}, nil
}

// UpdateArticle ....
func (s *clientsService) UpdateArticle(ctx context.Context, req *eCommerce.UpdateArticleRequest) (*eCommerce.Article, error) {
	err := s.stg.UpdateArticle(req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.UpdateArticle: %s", err.Error())
	}

	article, err := s.stg.GetArticleByID(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.GetArticleByID: %s", err.Error())
	}

	return &eCommerce.Article{
		Id:        article.Id,
		Content:   article.Content,
		AuthorId:  article.Author.Id,
		CreatedAt: article.CreatedAt,
		UpdatedAt: article.UpdatedAt,
	}, nil
}

// DeleteArticle ....
func (s *clientsService) DeleteArticle(ctx context.Context, req *eCommerce.DeleteArticleRequest) (*eCommerce.Article, error) {
	article, err := s.stg.GetArticleByID(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.GetArticleByID: %s", err.Error())
	}

	err = s.stg.DeleteArticle(article.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.DeleteArticle: %s", err.Error())
	}

	return &eCommerce.Article{
		Id:        article.Id,
		Content:   article.Content,
		AuthorId:  article.Author.Id,
		CreatedAt: article.CreatedAt,
		UpdatedAt: article.UpdatedAt,
	}, nil
}

// GetArticleList ....
func (s *clientsService) GetArticleList(ctx context.Context, req *eCommerce.GetArticleListRequest) (*eCommerce.GetArticleListResponse, error) {
	res, err := s.stg.GetArticleList(int(req.Offset), int(req.Limit), req.Search)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.DeleteArticle: %s", err.Error())
	}

	return res, nil
}

// GetArticleByID ....
func (s *clientsService) GetArticleByID(ctx context.Context, req *eCommerce.GetArticleByIDRequest) (*eCommerce.GetArticleByIDResponse, error) {
	article, err := s.stg.GetArticleByID(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.GetArticleByID: %s", err.Error())
	}

	return article, nil
}
