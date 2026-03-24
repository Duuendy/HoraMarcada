package mapper

import (
	domain "github.com/Duuendy/HoraMarcada/backend/internal/domain"
	dto "github.com/Duuendy/HoraMarcada/backend/internal/http/dto/dto_service"
)

func ToServiceItem(s domain.ServiceModel) dto.ServiceItem {
	return dto.ServiceItem{
		Id:            s.ID,
		Name:          s.Name,
		PriceCent:     s.PriceCent,
		TimeMinutes:   s.TimeMinutes,
		IsMaintenance: s.IsMaintenance,
	}
}
