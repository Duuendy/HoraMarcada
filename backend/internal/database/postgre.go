package database

import (
	"database/sql"

	"github.com/Duuendy/HoraMarcada/backend/internal/domain"
)

type ServiceRepo struct {
	DB *sql.DB
}

func (r *ServiceRepo) Create(service domain.ServiceModel) (int, error) {
	const q = `INSERT INTO tblservices (name, price_cent, time_minutes, is_maintenance) VALUES ($1, $2, $3, $4) RETURNING id`
	var id int
	err := r.DB.QueryRow(q, service.Name, service.PriceCent, service.TimeMinutes, service.IsMaintenance).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *ServiceRepo) List() ([]domain.ServiceModel, error) {
	query := `SELECT id, name, price_cent, time_minutes, is_maintenance FROM tblservices`
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var services []domain.ServiceModel
	for rows.Next() {
		var s domain.ServiceModel
		if err := rows.Scan(&s.ID, &s.Name, &s.PriceCent, &s.TimeMinutes, &s.IsMaintenance); err != nil {
			return nil, err
		}
		services = append(services, s)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return services, nil
}
