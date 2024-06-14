package Handlers

import (
	Entities "Naska_tg_alertbot/internal/entities"
	Services "Naska_tg_alertbot/internal/services"
	"encoding/json"
	"net/http"
)

var alertService *Services.AlertService

// InitAlertHandler инициализирует обработчик алертов HTTP
func InitAlertHandler(service *Services.AlertService) {
	alertService = service
	http.HandleFunc("/alert", AlertHandler)
}

// AlertHandler обрабатывает HTTP запросы с алертами
func AlertHandler(w http.ResponseWriter, r *http.Request) {
	var alert Entities.Alert
	if err := json.NewDecoder(r.Body).Decode(&alert); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Обработка алерта через сервис
	alertService.ProcessAlert(alert)

	w.WriteHeader(http.StatusOK)
}
