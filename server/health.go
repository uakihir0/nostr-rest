package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// HealthResponse
type HealthResponse struct {
	Status string `json:"status"`
}

// GetHealth
func GetHealth(c echo.Context) error {
	return c.JSON(
		http.StatusOK,
		HealthResponse{
			Status: "up",
		},
	)
}
