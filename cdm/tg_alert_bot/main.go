package main

import (
	Handlers "Naska_tg_alertbot/internal/adapter/handlers"
	Services "Naska_tg_alertbot/internal/services"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	var botToken string
	var criticalChatID string
	var warningChatID string

	flag.StringVar(&botToken, "bot_token", "", "Telegram Bot Token")
	flag.StringVar(&criticalChatID, "critical_chat_id", "", "Telegram Chat ID for Critical alerts")
	flag.StringVar(&warningChatID, "warning_chat_id", "", "Telegram Chat ID for Warning alerts")
	flag.Parse()

	if botToken == "" || criticalChatID == "" || warningChatID == "" {
		log.Fatalf("All flags are required: bot_token, critical_chat_id, warning_chat_id")
	}

	// Преобразуем chat ID из string в int64
	criticalChatIDInt, err := strconv.ParseInt(criticalChatID, 10, 64)
	if err != nil {
		log.Fatalf("Error parsing critical chat ID: %v", err)
	}

	warningChatIDInt, err := strconv.ParseInt(warningChatID, 10, 64)
	if err != nil {
		log.Fatalf("Error parsing warning chat ID: %v", err)
	}

	chatIDs := map[string]int64{
		"critical": criticalChatIDInt,
		"warning":  warningChatIDInt,
	}

	// Пример использования прочитанной конфигурации
	fmt.Printf("Bot Token: %s\n", botToken)
	fmt.Printf("Chat IDs: %v\n", chatIDs)

	// Инициализируем обработчик алертов HTTP и запускаем сервер
	alertService := Services.NewAlertService(botToken, chatIDs)
	Handlers.InitAlertHandler(alertService)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
