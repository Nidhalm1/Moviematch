package config

import (
    "log"
    "os"

    "gorm.io/driver/postgres"
    "gorm.io/gorm"

    "moviematch/internal/domain"
)

type Config struct {
    Port      string
    DBUrl     string
    JWTSecret string
}

func Load() *Config {
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    dbURL := os.Getenv("DB_URL")
    if dbURL == "" {
        dbURL = "postgres://user:password@db:5432/moviematch?sslmode=disable"
    }

    jwtSecret := os.Getenv("JWT_SECRET")
    if jwtSecret == "" {
        jwtSecret = "dev-secret-change-me"
    }

    return &Config{
        Port:      port,
        DBUrl:     dbURL,
        JWTSecret: jwtSecret,
    }
}

func ConnectDB(cfg *Config) (*gorm.DB, error) {
    db, err := gorm.Open(postgres.Open(cfg.DBUrl), &gorm.Config{})
    if err != nil {
        return nil, err
    }

    if err := db.AutoMigrate(&domain.Movie{}, &domain.User{}); err != nil {
        return nil, err
    }

    log.Println("Database connected and migrated")
    return db, nil
}
