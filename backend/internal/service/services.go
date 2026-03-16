package service

import (
	"database/sql"
	"fmt"
	// "github.com/Duuendy/HoraMarcada/backend/internal/http/dto/service"
)

func CreateService(name string, priceCent int, timeMinutes int, isMaintenance bool) int {

	// Lógica para criar o serviço (exemplo simplificado)

	id := 1 // Simula a criação e retorna um ID fixo
	return id
}

type ServiceModel struct {
	ID            int
	Name          string
	PriceCent     int
	TimeMinutes   int
	IsMaintenance bool
}

func ListServices(db *sql.DB) ([]ServiceModel, error) {
	//Logica para buscar direto no banco
	query := `SELECT id, name, price_cent, time_minutes, is_maintenance FROM tblservices` 

	rows, err := db.Query(query)
 	if err != nil {
		fmt.Println("Erro na query", err)
		return nil, err
	}
	defer rows.Close()

	services := make([]ServiceModel, 0)

	for rows.Next() {
		var s ServiceModel
		err:= rows.Scan(&s.ID, &s.Name, &s.PriceCent, &s.TimeMinutes, &s.IsMaintenance)
		if err != nil {
			fmt.Println("ListService DB Error:", err)
			return nil, err
		}
		services = append(services, s)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return services, nil

	// Lógica para listar serviços (exemplo simplificado)
	// services := []ServiceModel{
	// 	{ID: 1, Name: "Service A", PriceCent: 1000, TimeMinutes: 60, IsMaintenance: false},
	// 	{ID: 2, Name: "Service B", PriceCent: 2000, TimeMinutes: 120, IsMaintenance: true},
	// }
	// return services

}

// func GetServiceByID(id int) (ServiceModel, bool) {
// 	services := ListServices()
// 	for _, s := range services {
// 		if s.ID == id {
// 			return s, true
// 		}
// 	}

// 	return ServiceModel{}, false
// }
