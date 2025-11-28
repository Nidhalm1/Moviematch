package main

import (
    "log"

    "github.com/gin-gonic/gin"

    "moviematch/internal/config"
    "moviematch/internal/handler"
    "moviematch/internal/repository"
    "moviematch/internal/service"
)

func main() {
    cfg := config.Load()

    db, err := config.ConnectDB(cfg)
    if err != nil {
        log.Fatalf("failed to connect to database: %v", err)
    }

    r := gin.Default()

    movieRepo := repository.NewMovieRepository(db)
    userRepo := repository.NewUserRepository(db)

    authService := service.NewAuthService(cfg, userRepo)
    movieService := service.NewMovieService(movieRepo)

    authHandler := handler.NewAuthHandler(authService)
    movieHandler := handler.NewMovieHandler(movieService)

    api := r.Group("/api")

    api.GET("/health", func(c *gin.Context) {
        c.JSON(200, gin.H{"status": "ok"})
    })

    api.POST("/register", authHandler.Register)
    api.POST("/login", authHandler.Login)

    protected := api.Group("/")
    protected.Use(authService.AuthMiddleware())

    protected.GET("/movies", movieHandler.GetAll)
    protected.GET("/movies/recommend", movieHandler.Recommend)

    log.Println("Server running on port", cfg.Port)
    if err := r.Run(":" + cfg.Port); err != nil {
        log.Fatalf("server failed: %v", err)
    }
}
