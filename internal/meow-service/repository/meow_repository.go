package repository

import (
	"github.com/kshvyryaev/cyber-meower/internal/meow-service/domain"
)

type MeowRepository interface {
	Create(meow *domain.Meow) (int, error)
}
