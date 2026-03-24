package http

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Duuendy/HoraMarcada/backend/internal/domain"
	dto "github.com/Duuendy/HoraMarcada/backend/internal/http/dto/dto_service"
	resp "github.com/Duuendy/HoraMarcada/backend/internal/http/response"
)

type listFakeRepo struct{}

func (listFakeRepo) Create(domain.ServiceModel) (int, error) { return 0, nil }

func (listFakeRepo) List() ([]domain.ServiceModel, error) {
	return []domain.ServiceModel{
		{ID: 1, Name: "A", PriceCent: 100, TimeMinutes: 10, IsMaintenance: false},
	}, nil
}

func TestList_wrapsItemsInDTO(t *testing.T) {
	h := &ServiceHandler{Repository: listFakeRepo{}}
	req := httptest.NewRequest(http.MethodGet, "/services/list", nil)
	rr := httptest.NewRecorder()
	h.List(rr, req)
	if rr.Code != http.StatusOK {
		t.Fatalf("status %d body %s", rr.Code, rr.Body.String())
	}
	var wrap resp.APIResponse
	if err := json.Unmarshal(rr.Body.Bytes(), &wrap); err != nil {
		t.Fatal(err)
	}
	data, err := json.Marshal(wrap.Data)
	if err != nil {
		t.Fatal(err)
	}
	var out dto.DTOListServiceResponse
	if err := json.Unmarshal(data, &out); err != nil {
		t.Fatal(err)
	}
	if len(out.Items) != 1 || out.Items[0].Name != "A" {
		t.Fatalf("unexpected payload: %+v", out)
	}
}
