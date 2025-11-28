package tests

import (
    "testing"

    "moviematch/internal/domain"
    "moviematch/internal/repository"
    "moviematch/internal/service"
)

type mockMovieRepo struct {
    movies []domain.Movie
}

func (m *mockMovieRepo) FindAll() ([]domain.Movie, error) {
    return m.movies, nil
}

func (m *mockMovieRepo) Create(movie *domain.Movie) error {
    m.movies = append(m.movies, *movie)
    return nil
}

func TestRecommendOrdersByRating(t *testing.T) {
    repo := &mockMovieRepo{
        movies: []domain.Movie{
            {Title: "A", Rating: 3.0},
            {Title: "B", Rating: 5.0},
            {Title: "C", Rating: 4.0},
        },
    }

    svc := service.NewMovieService(repo)

    recos, err := svc.Recommend()
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }

    if len(recos) == 0 || recos[0].Title != "B" {
        t.Fatalf("expected B as first recommendation, got %+v", recos[0])
    }
}
