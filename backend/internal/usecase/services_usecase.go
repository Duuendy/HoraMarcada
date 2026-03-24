package service

import (
	"github.com/Duuendy/HoraMarcada/backend/internal/database"
	"github.com/Duuendy/HoraMarcada/backend/internal/domain"
)

func ListServices(repo database.ServiceRepository) ([]domain.ServiceModel, error) {
	return repo.List()
}
