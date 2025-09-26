package main

import (
	"fmt"
	"log"

	discord "github.com/bankusy/discord-go/model"
)

func main() {
	// 웹훅 URL 설정
	webhookURL := "https://discord.com/api/webhooks/1420932562959859762/AiQESd6MT5AS_dfDbKPH7-KZweqErz0uRyBkqySCpVNGKFvkZPzcEh_anwUozxOzvKHT"

	// Example 1: Simple text message
	fmt.Println("Sending simple text message...")
	err := discord.WebhookBuilder().
		SetURL(webhookURL).
		SetContent("Hello! This is a test message.").
		Send()
	if err != nil {
		log.Printf("Error sending simple message: %v", err)
	}

	// Example 2: Message with embed
	fmt.Println("Sending embed message...")
	embed := discord.EmbedBuilder().
		SetTitle("📊 Market Analysis Report").
		SetDescription("Analyzing today's major market trends.").
		SetColorHex("#00ff00"). // Green
		AddField("KOSPI", "2,450.12 (+1.2%)", true).
		AddField("NASDAQ", "14,250.45 (+0.8%)", true).
		AddField("Dow Jones", "34,500.78 (+0.5%)", true).
		AddField("Key Issue", "Semiconductor sector continues to rise", false).
		SetFooter("QuiredLab", "").
		SetCurrentTimestamp()

	err = discord.WebhookBuilder().
		SetURL(webhookURL).
		SetUsername("QuiredLab Bot").
		AddEmbed(embed).
		Send()
	if err != nil {
		log.Printf("Error sending embed message: %v", err)
	}

	// Example 3: Message with multiple embeds
	fmt.Println("Sending multiple embeds message...")

	// First embed - Market summary
	marketEmbed := discord.EmbedBuilder().
		SetTitle("📈 Market Summary").
		SetColorHex("#3498db"). // Blue
		AddField("KOSPI", "2,450.12 (+1.2%)", true).
		AddField("KOSDAQ", "850.45 (+0.8%)", true).
		AddField("Exchange Rate", "1,320 KRW/USD", true).
		SetCurrentTimestamp()

	// Second embed - Major news
	newsEmbed := discord.EmbedBuilder().
		SetTitle("📰 Major News").
		SetColorHex("#e74c3c"). // Red
		AddField("Samsung Electronics", "Stock price rises due to semiconductor sales increase", false).
		AddField("SK Hynix", "Memory semiconductor price uptrend", false).
		AddField("LG Chem", "Battery business expansion announcement", false).
		SetFooter("Source: Various News", "")

	err = discord.WebhookBuilder().
		SetURL(webhookURL).
		SetUsername("QuiredLab News").
		SetContent("🌅 **Today's Market Report**").
		AddEmbed(marketEmbed).
		AddEmbed(newsEmbed).
		Send()
	if err != nil {
		log.Printf("Error sending multiple embeds message: %v", err)
	}

	// Example 4: Error notification message
	fmt.Println("Sending error notification...")
	errorEmbed := discord.EmbedBuilder().
		SetTitle("⚠️ System Alert").
		SetDescription("An error occurred during data collection.").
		SetColorHex("#ff0000"). // Red
		AddField("Error Code", "ERR_001", true).
		AddField("Occurrence Time", "2024-01-15 14:30:00", true).
		AddField("Details", "API connection failed", false).
		SetFooter("QuiredLab System", "").
		SetCurrentTimestamp()

	err = discord.WebhookBuilder().
		SetURL(webhookURL).
		SetUsername("System Monitor").
		AddEmbed(errorEmbed).
		Send()
	if err != nil {
		log.Printf("Error sending error notification: %v", err)
	}

	fmt.Println("All webhook examples completed!")
}
