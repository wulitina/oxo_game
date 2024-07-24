package handlers

import (
	"net/http"
	"strconv"

	"oxo_game/internal/models"
	"oxo_game/internal/services"

	"github.com/gin-gonic/gin"
)

type RoomsHandler struct {
	service services.RoomService
}

func NewRoomsHandler(service services.RoomService) *RoomsHandler {
	return &RoomsHandler{
		service: service,
	}
}

func (h *RoomsHandler) GetAllRooms(c *gin.Context) {
	rooms, err := h.service.GetAllRooms()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, rooms)
}

func (h *RoomsHandler) GetRoomByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid room ID"})
		return
	}
	room, err := h.service.GetRoomByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, room)
}

func (h *RoomsHandler) CreateRoom(c *gin.Context) {
	var room models.Room
	if err := c.BindJSON(&room); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request payload"})
		return
	}

	id, err := h.service.CreateRoom(room.Name, room.Description)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id})
}

func (h *RoomsHandler) UpdateRoom(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid room ID"})
		return
	}

	var room models.Room
	if err := c.BindJSON(&room); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request payload"})
		return
	}

	if err := h.service.UpdateRoom(id, room.Name, room.Description); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *RoomsHandler) DeleteRoom(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid room ID"})
		return
	}

	if err := h.service.DeleteRoom(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
