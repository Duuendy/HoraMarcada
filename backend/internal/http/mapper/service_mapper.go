package mapper

import (
	dto "github.com/Duuendy/HoraMarcada/backend/internal/http/dto/service"
	domain "github.com/Duuendy/HoraMarcada/backend/internal/service"
)

func ToServiceItem(s domain.ServiceModel) dto.ServiceItem {
	return dto.ServiceItem{
		Id:             s.ID,
		Name:           s.Name,
		PriceCent:      s.PriceCent,
		TimeMinutes:    s.TimeMinutes,
		IsMaintenance:  s.IsMaintenance,
	}
}
