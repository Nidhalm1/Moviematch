package domain

import "time"

type Movie struct {
    ID        uint      `gorm:"primaryKey" json:"id"`
    Title     string    `json:"title"`
    Genre     string    `json:"genre"`
    Rating    float64   `json:"rating"`
    CreatedAt time.Time `json:"created_at"`
}
