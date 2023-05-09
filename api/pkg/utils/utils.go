package utils

import (
	"cadet-project/pkg/models"
	"html"
	"strings"
)

func CleanUserData(str string) string {
	return strings.TrimSpace(html.EscapeString(str))
}

func CheckCity(clientRequest string, start string) bool {
	var err error

	start = strings.Title(strings.ToLower(start))
	if err != nil || clientRequest == "" || clientRequest != start {
		return false
	}

	return true
}

func MaxMileage(vehicles []*models.Vehicle) float32 {
	max := vehicles[0].Mileage

	for _, m := range vehicles {
		if max < m.Mileage {
			max = m.Mileage
		}
	}

	return max
}
