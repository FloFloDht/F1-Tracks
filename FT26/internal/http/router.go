package http

import (
	"github.com/gin-gonic/gin"

	"FT26/internal/config"
	"FT26/internal/tracks"

	"go.mongodb.org/mongo-driver/mongo"
)

func NewRouter(client *mongo.Client, cfg *config.Config) *gin.Engine {
	r := gin.Default()

	repo := tracks.NewRepository(client, cfg.MongoDBName)
	service := tracks.NewService(repo)
	handler := tracks.NewHandler(service)

	circuitsRoutes := r.Group("/circuits")
	{
		circuitsRoutes.GET("", handler.GetAll)
		circuitsRoutes.GET("/:id", handler.GetByID)
	}

	return r
}
