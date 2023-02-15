package main

import (
	"log"
	"net/http"

	"uacademy/blogpost/api_gateway/clients"
	"uacademy/blogpost/api_gateway/config"
	docs "uacademy/blogpost/api_gateway/docs" // docs is generated by Swag CLI, you have to import it.
	"uacademy/blogpost/api_gateway/handlers"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @license.name Apache 2.0
// @license.url  http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	cfg := config.Load()
	if cfg.Environment != "development" {
		gin.SetMode(gin.ReleaseMode)
	}

	docs.SwaggerInfo.Title = cfg.App
	docs.SwaggerInfo.Version = cfg.AppVersion

	r := gin.New()
	if cfg.Environment != "production" {
		r.Use(gin.Logger(), gin.Recovery()) // Later they will be replaced by custom Logger and Recovery
	}

	grpcClients, err := clients.NewGrpcClients(cfg)
	if err != nil {
		panic(err)
	}
	defer grpcClients.Close()

	h := handlers.NewHandler(cfg, grpcClients)

	r.GET("/ping", MyCORSMiddleware(), h.AuthMiddleware("*"), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	v1 := r.Group("/v1")
	{
		v1.Use(MyCORSMiddleware())
		v1.POST("/login", h.Login)
		v1.POST("/article", h.AuthMiddleware("*"), h.CreateArticle)
		v1.GET("/article/:id", h.AuthMiddleware("*"), h.GetArticleByID)
		v1.GET("/article", h.AuthMiddleware("*"), h.GetArticleList)
		v1.PUT("/article", h.AuthMiddleware("*"), h.UpdateArticle)
		v1.DELETE("/article/:id", h.AuthMiddleware("ADMIN"), h.DeleteArticle)

		v1.GET("/my-articles", h.AuthMiddleware("*"), h.SearchArticleByMyUsername)

		// v1.POST("/author", h.CreateAuthor)
		// v1.GET("/author/:id", h.GetAuthorByID)
		// v1.GET("/author", h.GetAuthorList)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(cfg.HTTPPort) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

// MyCORSMiddleware ...
func MyCORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("MyCORSMiddleware...")

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Max-Age", "3600")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}