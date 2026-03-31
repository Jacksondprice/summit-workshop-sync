package handlers

import (
	"net/http/httptest"
	"testing"
)

func TestListSchedules(t *testing.T) {
	rec := httptest.NewRecorder()
	ListSchedules(rec, httptest.NewRequest("GET", "/api/schedules", nil))
	if rec.Code != 200 {
		t.Fatalf("status %d", rec.Code)
	}
}
