package http

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Duuendy/HoraMarcada/backend/internal/domain"
)

type fakeRepo struct{}

func (fakeRepo) List() ([]domain.ServiceModel, error) { return nil, nil }
func (fakeRepo) Create(domain.ServiceModel) (int, error) { return 42, nil }

func TestRouter_ServicesCreate_GET_returns405Not404(t *testing.T) {
	mux := Router(fakeRepo{})
	req := httptest.NewRequest(http.MethodGet, "/services/create", nil)
	rr := httptest.NewRecorder()
	mux.ServeHTTP(rr, req)
	if rr.Code == http.StatusNotFound {
		t.Fatalf("unexpected 404 for GET /services/create: body=%q", rr.Body.String())
	}
	if rr.Code != http.StatusMethodNotAllowed {
		t.Fatalf("want StatusMethodNotAllowed got %d body=%q", rr.Code, rr.Body.String())
	}
}
