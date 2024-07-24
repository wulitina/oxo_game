package handlers

import (
	"net/http"
	"strconv"

	"oxo_game/internal/models"
	"oxo_game/internal/services"

	"github.com/gin-gonic/gin"
)

type PlayersHandler struct {
	service *services.PlayerService
}

func NewPlayersHandler(service *services.PlayerService) *PlayersHandler {
	return &PlayersHandler{service: service}
}

func (h *PlayersHandler) GetAllPlayers(c *gin.Context) {
	players, err := h.service.GetAllPlayers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, players)
}

func (h *PlayersHandler) GetPlayerByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid player ID"})
		return
	}
	player, err := h.service.GetPlayerByID(id)
	if err != nil {
		if err == services.ErrPlayerNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "player not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, player)
}

func (h *PlayersHandler) CreatePlayer(c *gin.Context) {
	var player models.Player
	if err := c.BindJSON(&player); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := h.service.CreatePlayer(player)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": id})
}

func (h *PlayersHandler) UpdatePlayer(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var updatedPlayer models.Player
	if err := c.BindJSON(&updatedPlayer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request payload"})
		return
	}
	if err := h.service.UpdatePlayer(id, updatedPlayer); err != nil {
		if err == services.ErrPlayerNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "player updated successfully"})
}

func (h *PlayersHandler) DeletePlayer(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.DeletePlayer(id); err != nil {
		if err == services.ErrPlayerNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "player not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "player deleted successfully"})
}
