package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"oxo_game/internal/handlers"
	"oxo_game/internal/repositories"
	"oxo_game/internal/services"
	"syscall"

	"github.com/gin-gonic/gin"
)

func main() {

	//db := db.GetDB()
	//fmt.Println(db)
	// Initialize repositories
	playerRepo := repositories.NewInMemoryPlayerRepository()
	levelRepo := repositories.NewInMemoryLevelRepository()
	roomRepo := repositories.NewInMemoryRoomRepository()
	// 初始化预约系统的服务、处理器和仓库
	reservationRepo := repositories.NewInMemoryReservationRepository()
	logRepo := repositories.NewInMemoryLogRepository()
	challengeRepo := repositories.NewInMemoryChallengeRepository()

	// Initialize services
	playerService := services.NewPlayerService(playerRepo)
	levelService := services.NewLevelService(levelRepo)
	roomService := services.NewRoomService(roomRepo)
	reservationService := services.NewReservationService(reservationRepo)
	logService := services.NewLogService(logRepo)
	challengeService := services.NewChallengeService(challengeRepo, playerRepo)

	// Initialize handlers
	playersHandler := handlers.NewPlayersHandler(playerService)
	levelsHandler := handlers.NewLevelsHandler(levelService)
	roomsHandler := handlers.NewRoomsHandler(roomService)
	logsHandler := handlers.NewLogsHandler(logService)
	challengeHandler := handlers.NewChallengeHandler(challengeService)

	reservationHandler := handlers.NewReservationHandler(reservationService)

	// Setup Gin router
	router := gin.Default()

	// Routes for players
	router.GET("/players", playersHandler.GetAllPlayers)
	router.GET("/players/:id", playersHandler.GetPlayerByID)
	router.POST("/players", playersHandler.CreatePlayer)
	router.PUT("/players/:id", playersHandler.UpdatePlayer)
	router.DELETE("/players/:id", playersHandler.DeletePlayer)

	// Routes for levels
	router.GET("/levels", levelsHandler.GetAllLevels)
	router.POST("/levels", levelsHandler.CreateLevel)

	router.GET("/rooms", roomsHandler.GetAllRooms)
	router.GET("/rooms/:id", roomsHandler.GetRoomByID)
	router.POST("/rooms", roomsHandler.CreateRoom)
	router.PUT("/rooms/:id", roomsHandler.UpdateRoom)
	router.DELETE("/rooms/:id", roomsHandler.DeleteRoom)

	router.GET("/reservations", reservationHandler.ListReservations)
	router.POST("/reservations", reservationHandler.CreateReservation)

	router.POST("/challenges", challengeHandler.ParticipateChallenge)
	router.GET("/challenges/results", challengeHandler.ListLatestChallenges)

	// Logs endpoints
	router.GET("/logs", logsHandler.GetAllLogs)
	router.POST("/logs", logsHandler.CreateLog)

	// Start server
	router.Run(":8080")

	// Start HTTP server
	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error starting server: %v", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	if err := server.Shutdown(nil); err != nil {
		log.Fatalf("Error shutting down server: %v", err)
	}
	log.Println("Server stopped.")
}
