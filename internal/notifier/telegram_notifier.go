package Notifier

import (
	"Naska_tg_alertbot/internal/entities"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

	"log"
)

// TelegramNotifier представляет уведомление в Telegram
type TelegramNotifier struct {
	Bot     *tgbotapi.BotAPI
	ChatIDs map[string]int64
}

// NewTelegramNotifier создает новый экземпляр TelegramNotifier
func NewTelegramNotifier(botToken string, chatIDs map[string]int64) (*TelegramNotifier, error) {
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		return nil, err
	}

	return &TelegramNotifier{
		Bot:     bot,
		ChatIDs: chatIDs,
	}, nil
}

// Notify отправляет уведомление в Telegram
func (tn *TelegramNotifier) Notify(alert Entities.Alert) error {
	chatID, ok := tn.ChatIDs[alert.Labels["severity"]]
	if !ok {
		log.Printf("Unknown severity level: %v", alert.Labels["severity"])
		return fmt.Errorf("unknown severity level: %v", alert.Labels["severity"])
	}

	var emoji string
	if alert.Status == "firing" {
		emoji = "🔥"
		if alert.Labels["severity"] == "warning" {
			emoji = "⚠️"
		} else {
			emoji = "❗️"
		}
	} else if alert.Status == "resolved" {
		emoji = "✅"
	}

	messageText := fmt.Sprintf(
		"%s %s\n\n🔔 Summary: %s\n📝 Description: %s\n⚠️ Severity: %s\n🕒 Started at: %s",
		emoji,
		alert.Status,
		alert.Annotations.Summary,
		alert.Annotations.Description,
		alert.Labels["severity"],
		alert.StartsAt.Format("Jan 02, 15:04:05 UTC"),
	)

	if alert.Status == "resolved" {
		messageText += fmt.Sprintf("\n✅ Resolved at: %s", alert.EndsAt.Format("Jan 02, 15:04:05 UTC"))
	}

	msg := tgbotapi.NewMessage(chatID, messageText)
	_, err := tn.Bot.Send(msg)
	if err != nil {
		log.Printf("Error sending message to Telegram: %v", err)
		return err
	}

	return nil
}
