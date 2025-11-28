package repository

import (
    "moviematch/internal/domain"

    "gorm.io/gorm"
)

type MovieRepository interface {
    FindAll() ([]domain.Movie, error)
    Create(movie *domain.Movie) error
}

type GormMovieRepository struct {
    db *gorm.DB
}

func NewMovieRepository(db *gorm.DB) MovieRepository {
    return &GormMovieRepository{db: db}
}

func (r *GormMovieRepository) FindAll() ([]domain.Movie, error) {
    var movies []domain.Movie
    err := r.db.Order("rating DESC").Find(&movies).Error
    return movies, err
}

func (r *GormMovieRepository) Create(movie *domain.Movie) error {
    return r.db.Create(movie).Error
}
