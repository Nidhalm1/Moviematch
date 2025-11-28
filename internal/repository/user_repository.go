package repository

import (
    "moviematch/internal/domain"

    "gorm.io/gorm"
)

type UserRepository interface {
    Create(email, passwordHash string) error
    FindByEmail(email string) (*domain.User, error)
}

type GormUserRepository struct {
    db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
    return &GormUserRepository{db: db}
}

func (r *GormUserRepository) Create(email, passwordHash string) error {
    user := &domain.User{
        Email:        email,
        PasswordHash: passwordHash,
    }
    return r.db.Create(user).Error
}

func (r *GormUserRepository) FindByEmail(email string) (*domain.User, error) {
    var user domain.User
    if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
        return nil, err
    }
    return &user, nil
}
