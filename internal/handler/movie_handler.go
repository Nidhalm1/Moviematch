package handler

import (
    "net/http"

    "github.com/gin-gonic/gin"

    "moviematch/internal/domain"
    "moviematch/internal/service"
)

type MovieHandler struct {
    service *service.MovieService
}

func NewMovieHandler(service *service.MovieService) *MovieHandler {
    return &MovieHandler{service: service}
}

func (h *MovieHandler) GetAll(c *gin.Context) {
    movies, err := h.service.GetAll()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, movies)
}

func (h *MovieHandler) Recommend(c *gin.Context) {
    movies, err := h.service.Recommend()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, movies)
}

// (Optionnel) création de film pour peupler la base
type createMovieRequest struct {
    Title  string  `json:"title" binding:"required"`
    Genre  string  `json:"genre" binding:"required"`
    Rating float64 `json:"rating" binding:"required"`
}

func (h *MovieHandler) Create(c *gin.Context) {
    var req createMovieRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    movie := &domain.Movie{
        Title:  req.Title,
        Genre:  req.Genre,
        Rating: req.Rating,
    }

    // on accède au repo via le service actuel: simplification
    // dans un projet plus gros, on créerait une méthode dédiée
    _ = movie // placeholder pour éviter un import circulaire
    c.JSON(http.StatusNotImplemented, gin.H{"message": "not implemented in demo"})
}
