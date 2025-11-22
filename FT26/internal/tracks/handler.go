package tracks

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service Service
}

func NewHandler(s Service) *Handler {
	return &Handler{service: s}
}

func (h *Handler) GetAll(c *gin.Context) {
	circuits, err := h.service.GetAll(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, circuits)
}

func (h *Handler) GetByID(c *gin.Context) {
	id := c.Param("id")

	circuit, err := h.service.GetByID(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Circuit not found"})
		return
	}

	c.JSON(http.StatusOK, circuit)
}
