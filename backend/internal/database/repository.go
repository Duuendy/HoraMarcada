package database

import (
	"github.com/Duuendy/HoraMarcada/backend/internal/domain"
)

type ServiceRepository interface {
	List() ([]domain.ServiceModel, error)
	Create(domain.ServiceModel) (int, error)
}
