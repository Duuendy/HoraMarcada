package service

type ServiceItem struct {
	Id            int    `json:"id"`
	Name          string `json:"name"`
	PriceCent     int    `json:"price_cent"`
	TimeMinutes   int    `json:"time_minutes"`
	IsMaintenance bool   `json:"is_maintenance"`
}

type DTOListServiceResponse struct {
	Items []ServiceItem `json:"items"`
}