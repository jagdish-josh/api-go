package internal

import (
	"net/http"
)

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	result := HealthService()
	w.Write([]byte(result))
}

