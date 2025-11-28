package service

import (
    "errors"
    "sort"

    "moviematch/internal/domain"
    "moviematch/internal/repository"
)

type MovieService struct {
    repo repository.MovieRepository
}

func NewMovieService(repo repository.MovieRepository) *MovieService {
    return &MovieService{repo: repo}
}

func (s *MovieService) GetAll() ([]domain.Movie, error) {
    return s.repo.FindAll()
}

// Recommend returns the top N movies by rating (default 5)
func (s *MovieService) Recommend() ([]domain.Movie, error) {
    movies, err := s.repo.FindAll()
    if err != nil {
        return nil, err
    }

    if len(movies) == 0 {
        return nil, errors.New("no movies available")
    }

    sort.Slice(movies, func(i, j int) bool {
        return movies[i].Rating > movies[j].Rating
    })

    n := 5
    if len(movies) < n {
        n = len(movies)
    }

    return movies[:n], nil
}
