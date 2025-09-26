package discord

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Color constants for common Discord colors
const (
	ColorRed      = 0xff0000
	ColorGreen    = 0x00ff00
	ColorBlue     = 0x0000ff
	ColorYellow   = 0xffff00
	ColorOrange   = 0xffa500
	ColorPurple   = 0x800080
	ColorPink     = 0xffc0cb
	ColorCyan     = 0x00ffff
	ColorGray     = 0x808080
	ColorDarkGray = 0x404040
)

// QuickWebhook creates a webhook with URL and sends it immediately
func QuickWebhook(url, content string) error {
	return WebhookBuilder().
		SetURL(url).
		SetContent(content).
		Send()
}

// QuickEmbed creates and sends an embed webhook immediately
func QuickEmbed(url, title, description string, color int) error {
	embed := EmbedBuilder().
		SetTitle(title).
		SetDescription(description).
		SetColor(color).
		SetCurrentTimestamp()

	return WebhookBuilder().
		SetURL(url).
		AddEmbed(embed).
		Send()
}

// CreateMarketReport creates a market report embed
func CreateMarketReport(kospi, kosdaq, exchangeRate string) *Embed {
	return EmbedBuilder().
		SetTitle("📊 Market Status").
		SetColor(ColorBlue).
		AddField("KOSPI", kospi, true).
		AddField("KOSDAQ", kosdaq, true).
		AddField("Exchange Rate", exchangeRate, true).
		SetCurrentTimestamp()
}

// CreateNewsEmbed creates a news embed
func CreateNewsEmbed(title, content, source string) *Embed {
	return EmbedBuilder().
		SetTitle("📰 "+title).
		SetDescription(content).
		SetColor(ColorRed).
		SetFooter(source, "").
		SetCurrentTimestamp()
}

// CreateErrorEmbed creates an error notification embed
func CreateErrorEmbed(errorCode, message, details string) *Embed {
	return EmbedBuilder().
		SetTitle("⚠️ System Error").
		SetDescription(message).
		SetColor(ColorRed).
		AddField("Error Code", errorCode, true).
		AddField("Occurrence Time", time.Now().Format("2006-01-02 15:04:05"), true).
		AddField("Details", details, false).
		SetFooter("QuiredLab System", "").
		SetCurrentTimestamp()
}

// CreateSuccessEmbed creates a success notification embed
func CreateSuccessEmbed(title, message string) *Embed {
	return EmbedBuilder().
		SetTitle("✅ " + title).
		SetDescription(message).
		SetColor(ColorGreen).
		SetCurrentTimestamp()
}

// CreateWarningEmbed creates a warning notification embed
func CreateWarningEmbed(title, message string) *Embed {
	return EmbedBuilder().
		SetTitle("⚠️ " + title).
		SetDescription(message).
		SetColor(ColorYellow).
		SetCurrentTimestamp()
}

// CreateInfoEmbed creates an info notification embed
func CreateInfoEmbed(title, message string) *Embed {
	return EmbedBuilder().
		SetTitle("ℹ️ " + title).
		SetDescription(message).
		SetColor(ColorCyan).
		SetCurrentTimestamp()
}

// FormatNumber formats a number with Korean style (e.g., 1,234,567)
func FormatNumber(num int64) string {
	str := strconv.FormatInt(num, 10)
	if len(str) <= 3 {
		return str
	}

	var result strings.Builder
	for i, char := range str {
		if i > 0 && (len(str)-i)%3 == 0 {
			result.WriteString(",")
		}
		result.WriteRune(char)
	}
	return result.String()
}

// FormatPercentage formats a percentage with sign
func FormatPercentage(value float64) string {
	sign := "+"
	if value < 0 {
		sign = ""
	}
	return fmt.Sprintf("%s%.2f%%", sign, value)
}

// CreateStockEmbed creates a stock information embed
func CreateStockEmbed(symbol, name string, price float64, change float64, volume int64) *Embed {
	color := ColorGreen
	if change < 0 {
		color = ColorRed
	}

	return EmbedBuilder().
		SetTitle(fmt.Sprintf("📈 %s (%s)", name, symbol)).
		SetColor(color).
		AddField("현재가", fmt.Sprintf("%.0f원", price), true).
		AddField("등락률", FormatPercentage(change), true).
		AddField("거래량", FormatNumber(volume), true).
		SetCurrentTimestamp()
}

// CreateUserActivityEmbed creates a user activity embed
func CreateUserActivityEmbed(username, action, details string) *Embed {
	return EmbedBuilder().
		SetTitle("👤 사용자 활동").
		SetColor(ColorPurple).
		AddField("사용자", username, true).
		AddField("활동", action, true).
		AddField("상세", details, false).
		SetCurrentTimestamp()
}

// CreateSystemStatusEmbed creates a system status embed
func CreateSystemStatusEmbed(status, uptime, cpu, memory string) *Embed {
	color := ColorGreen
	if status != "정상" {
		color = ColorRed
	}

	return EmbedBuilder().
		SetTitle("🖥️ 시스템 상태").
		SetColor(color).
		AddField("상태", status, true).
		AddField("가동시간", uptime, true).
		AddField("CPU 사용률", cpu, true).
		AddField("메모리 사용률", memory, true).
		SetCurrentTimestamp()
}

// SendMarketReport sends a market report webhook
func SendMarketReport(webhookURL, kospi, kosdaq, exchangeRate string) error {
	embed := CreateMarketReport(kospi, kosdaq, exchangeRate)
	return WebhookBuilder().
		SetURL(webhookURL).
		SetUsername("Market Bot").
		AddEmbed(embed).
		Send()
}

// SendNewsAlert sends a news alert webhook
func SendNewsAlert(webhookURL, title, content, source string) error {
	embed := CreateNewsEmbed(title, content, source)
	return WebhookBuilder().
		SetURL(webhookURL).
		SetUsername("News Bot").
		AddEmbed(embed).
		Send()
}

// SendErrorNotification sends an error notification webhook
func SendErrorNotification(webhookURL, errorCode, message, details string) error {
	embed := CreateErrorEmbed(errorCode, message, details)
	return WebhookBuilder().
		SetURL(webhookURL).
		SetUsername("System Monitor").
		AddEmbed(embed).
		Send()
}

// SendSuccessNotification sends a success notification webhook
func SendSuccessNotification(webhookURL, title, message string) error {
	embed := CreateSuccessEmbed(title, message)
	return WebhookBuilder().
		SetURL(webhookURL).
		SetUsername("System Monitor").
		AddEmbed(embed).
		Send()
}
