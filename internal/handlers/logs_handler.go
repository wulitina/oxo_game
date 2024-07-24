package handlers

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"oxo_game/internal/models"
	"oxo_game/internal/services"
)

var (
	ErrLogNotFound = errors.New("log not found")
)

type LogHandler struct {
	service services.LogService
}

func NewLogsHandler(service services.LogService) *LogHandler {
	return &LogHandler{
		service: service,
	}
}

func (h *LogHandler) GetAllLogs(c *gin.Context) {
	logs, err := h.service.GetAllLogs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, logs)
}

func (h *LogHandler) GetLogByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid log ID"})
		return
	}

	log, err := h.service.GetLogByID(id)
	if err != nil {
		if err == ErrLogNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "log not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, log)
}

func (h *LogHandler) CreateLog(c *gin.Context) {
	var log models.Log
	if err := c.BindJSON(&log); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request payload"})
		return
	}

	// Set timestamps
	now := time.Now().Unix()
	log.Timestamp = now
	log.CreatedAt = now
	log.UpdatedAt = now

	id, err := h.service.CreateLog(log)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id})
}

func (h *LogHandler) GetLogsByPlayerID(c *gin.Context) {
	playerID, err := strconv.Atoi(c.Query("player_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid player ID"})
		return
	}

	logs, err := h.service.GetLogsByPlayerID(playerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, logs)
}

func (h *LogHandler) GetLogsByAction(c *gin.Context) {
	action := c.Query("action")
	if action == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "action parameter is required"})
		return
	}

	logs, err := h.service.GetLogsByAction(action)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, logs)
}

func (h *LogHandler) GetLogsByTimeRange(c *gin.Context) {
	startTime, err := strconv.ParseInt(c.Query("start_time"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid start_time parameter"})
		return
	}

	endTime, err := strconv.ParseInt(c.Query("end_time"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid end_time parameter"})
		return
	}

	logs, err := h.service.GetLogsByTimeRange(startTime, endTime)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, logs)
}

func (h *LogHandler) DeleteLog(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid log ID"})
		return
	}

	if err := h.service.DeleteLog(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
