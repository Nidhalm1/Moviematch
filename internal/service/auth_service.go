package service

import (
    "errors"
    "net/http"
    "strings"
    "time"

    "github.com/gin-gonic/gin"
    "golang.org/x/crypto/bcrypt"

    "moviematch/internal/config"
    "moviematch/internal/repository"
    jwtpkg "moviematch/pkg/jwt"
)

type AuthService struct {
    cfg       *config.Config
    userRepo  repository.UserRepository
}

func NewAuthService(cfg *config.Config, repo repository.UserRepository) *AuthService {
    return &AuthService{
        cfg:      cfg,
        userRepo: repo,
    }
}

func (s *AuthService) RegisterUser(email, password string) error {
    if email == "" || password == "" {
        return errors.New("email and password are required")
    }
    hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }
    return s.userRepo.Create(email, string(hash))
}

func (s *AuthService) LoginUser(email, password string) (string, error) {
    user, err := s.userRepo.FindByEmail(email)
    if err != nil {
        return "", errors.New("invalid credentials")
    }

    if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
        return "", errors.New("invalid credentials")
    }

    token, err := jwtpkg.GenerateToken(user.Email, s.cfg.JWTSecret, time.Hour*24)
    if err != nil {
        return "", err
    }
    return token, nil
}

func (s *AuthService) AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing Authorization header"})
            return
        }

        tokenString := strings.TrimPrefix(authHeader, "Bearer ")
        tokenString = strings.TrimSpace(tokenString)
        if tokenString == "" {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid Authorization header"})
            return
        }

        claims, err := jwtpkg.ValidateToken(tokenString, s.cfg.JWTSecret)
        if err != nil {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
            return
        }

        c.Set("user_email", claims.Subject)
        c.Next()
    }
}
