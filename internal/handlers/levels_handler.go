package handlers

import (
	"net/http"
	"strconv"

	"oxo_game/internal/models"
	"oxo_game/internal/services"

	"github.com/gin-gonic/gin"
)

type LevelsHandler struct {
	service services.LevelService
}

func NewLevelsHandler(service services.LevelService) *LevelsHandler {
	return &LevelsHandler{service: service}
}

func (h *LevelsHandler) GetAllLevels(c *gin.Context) {
	levels, err := h.service.GetAllLevels()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, levels)
}

func (h *LevelsHandler) GetLevelByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid level ID"})
		return
	}
	level, err := h.service.GetLevelByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, level)
}

func (h *LevelsHandler) CreateLevel(c *gin.Context) {
	var level models.Level
	if err := c.BindJSON(&level); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request payload"})
		return
	}
	id, err := h.service.CreateLevel(level.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": id})
}
