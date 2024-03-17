package main

import (
	"encoding/json"
	"math/rand"
	"os"
	"time"
)

// Status struct untuk menyimpan nilai air dan angin
type Status struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

func main() {
	// Set interval untuk mengupdate file JSON setiap 15 detik
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		// Generate nilai acak untuk air dan angin
		water := rand.Intn(100) + 1
		wind := rand.Intn(100) + 1

		// Tentukan status berdasarkan nilai air dan angin
		waterStatus := determineWaterStatus(water)
		windStatus := determineWindStatus(wind)

		// Simpan data ke file JSON
		status := Status{Water: water, Wind: wind}
		saveStatusToFile(status, waterStatus, windStatus)
	}
}

func determineWaterStatus(water int) string {
	if water < 5 {
		return "Aman"
	} else if water >= 6 && water <= 8 {
		return "Siaga"
	} else {
		return "Bahaya"
	}
}

func determineWindStatus(wind int) string {
	if wind < 6 {
		return "Aman"
	} else if wind >= 7 && wind <= 15 {
		return "Siaga"
	} else {
		return "Bahaya"
	}
}

func saveStatusToFile(status Status, waterStatus, windStatus string) {
	data := map[string]interface{}{
		"status": map[string]interface{}{
			"water": map[string]interface{}{
				"value":  status.Water,
				"status": waterStatus,
			},
			"wind": map[string]interface{}{
				"value":  status.Wind,
				"status": windStatus,
			},
		},
	}

	file, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		panic(err)
	}

	err = os.WriteFile("status.json", file, 0644)
	if err != nil {
		panic(err)
	}
}
