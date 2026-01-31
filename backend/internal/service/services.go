package service

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

func ListServices() []ServiceModel {

	// Lógica para listar serviços (exemplo simplificado)
	services := []ServiceModel{
		{ID: 1, Name: "Service A", PriceCent: 1000, TimeMinutes: 60, IsMaintenance: false},
		{ID: 2, Name: "Service B", PriceCent: 2000, TimeMinutes: 120, IsMaintenance: true},
	}
	return services
}
