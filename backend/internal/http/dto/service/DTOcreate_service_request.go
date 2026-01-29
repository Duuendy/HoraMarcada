package service

type DTOCreateServiceRequest struct {
	Name		 	string	`json:"name"`
	PriceCent		int		`json:"price_cent"`
	TimeMinutes    	int		`json:"time_minutes"` // in minutes
	IsMaintenance	bool	`json:"is_maintenance"`
}