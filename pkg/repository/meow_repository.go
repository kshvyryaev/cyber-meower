package repository

import "github.com/kshvyryaev/cyber-meower-meower-service/pkg/domain"

type MeowRepository interface {
	Create(meow *domain.Meow) (int, error)
}
