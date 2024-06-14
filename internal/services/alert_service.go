package Services

import (
	"Naska_tg_alertbot/internal/entities"
	"Naska_tg_alertbot/internal/notifier"
	"log"
)

// AlertService предоставляет методы для обработки алертов
type AlertService struct {
	Notifier *Notifier.TelegramNotifier
}

// NewAlertService создает новый экземпляр AlertService с передачей botToken и chatIDs
func NewAlertService(botToken string, chatIDs map[string]int64) *AlertService {
	tgNotifier, err := Notifier.NewTelegramNotifier(botToken, chatIDs)
	if err != nil {
		log.Fatalf("Error creating Telegram notifier: %v", err)
	}

	return &AlertService{
		Notifier: tgNotifier,
	}
}

// ProcessAlert обрабатывает алерт
func (as *AlertService) ProcessAlert(alert Entities.Alert) {
	err := as.Notifier.Notify(alert)
	if err != nil {
		log.Printf("Error notifying about alert: %v", err)
	}
}
