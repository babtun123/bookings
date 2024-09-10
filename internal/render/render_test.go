package render

import (
	"net/http"
	"testing"

	"github.com/babtun123/bookings/internal/models"
)

func TestAddDefaultData(t *testing.T) {
	var td models.TemplateData

	r, err := getSession()
	if err != nil {
		t.Error(err)
	}

	session.Put(r.Context(), "flash", "123")

	result := AddDefaultData(&td, r)

	if result.Flash != "123" {
		t.Error("Flash value of 123 not found in session")
	}

}

func TestRenderTemplate(t *testing.T) {
	result := "Yes"
	expected := "No"
	if result != expected {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func getSession() (*http.Request, error) {
	r, err := http.NewRequest("GET", "/some-url", nil)
	if err != nil {
		return nil, err
	}

	ctx := r.Context()
	ctx, _ = session.Load(ctx, r.Header.Get("X-session"))
	r = r.WithContext(ctx)

	return r, nil
}
