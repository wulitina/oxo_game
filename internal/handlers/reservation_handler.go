package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"oxo_game/internal/models"
	"oxo_game/internal/services"
)

type ReservationHandler struct {
	reservationService services.ReservationService
}

func NewReservationHandler(service services.ReservationService) *ReservationHandler {
	return &ReservationHandler{
		reservationService: service,
	}
}

func (h *ReservationHandler) ListReservations(c *gin.Context) {
	var (
		roomIDParam = c.Query("room_id")
		dateParam   = c.Query("date")
		limitParam  = c.Query("limit")
	)

	var roomID int
	if roomIDParam != "" {
		id, err := strconv.Atoi(roomIDParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid room_id parameter"})
			return
		}
		roomID = id
	}

	var date time.Time
	if dateParam != "" {
		parsedDate, err := time.Parse("2006-01-02", dateParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid date parameter (format: yyyy-mm-dd)"})
			return
		}
		date = parsedDate
	}

	var limit int
	if limitParam != "" {
		limit, err := strconv.Atoi(limitParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid limit parameter"})
			return
		}
		limit = limit
	}

	reservations := h.reservationService.ListReservationsByRoomAndDate(roomID, date)

	if limit > 0 && limit < len(reservations) {
		reservations = reservations[:limit]
	}

	c.JSON(http.StatusOK, reservations)
}

func (h *ReservationHandler) CreateReservation(c *gin.Context) {
	var reservation models.Reservation
	if err := c.ShouldBindJSON(&reservation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request payload"})
		return
	}

	id, err := h.reservationService.CreateReservation(reservation.RoomID, reservation.Date, reservation.Time, reservation.PlayerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id})
}
