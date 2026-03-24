package domain

type ServiceModel struct {
	ID            int
	Name          string
	PriceCent     int
	TimeMinutes   int
	IsMaintenance bool
}