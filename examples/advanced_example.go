package main

import (
	"fmt"
	"log"

	discord "github.com/bankusy/discord-go/model"
)

func main() {
	webhookURL := "https://discord.com/api/webhooks/1420932562959859762/AiQESd6MT5AS_dfDbKPH7-KZweqErz0uRyBkqySCpVNGKFvkZPzcEh_anwUozxOzvKHT"

	// Advanced Example 1: Market Report
	fmt.Println("Sending market report...")
	err := discord.SendMarketReport(
		webhookURL,
		"2,450.12 (+1.2%)",
		"850.45 (+0.8%)",
		"1,320 KRW/USD",
	)
	if err != nil {
		log.Printf("Error sending market report: %v", err)
	}

	// Advanced Example 2: News Alert
	fmt.Println("Sending news alert...")
	err = discord.SendNewsAlert(
		webhookURL,
		"Samsung Electronics Stock Rise",
		"Stock price continues to rise due to increased semiconductor sales.",
		"Korea Economic Daily",
	)
	if err != nil {
		log.Printf("Error sending news alert: %v", err)
	}

	// Advanced Example 3: Stock Information Embed
	fmt.Println("Sending stock information...")
	stockEmbed := discord.CreateStockEmbed(
		"005930", // Samsung Electronics
		"Samsung Electronics",
		75000.0,  // Current price
		2.5,      // Change rate
		15000000, // Volume
	)

	err = discord.WebhookBuilder().
		SetURL(webhookURL).
		SetUsername("Stock Bot").
		AddEmbed(stockEmbed).
		Send()
	if err != nil {
		log.Printf("Error sending stock info: %v", err)
	}

	// Advanced Example 4: User Activity Log
	fmt.Println("Sending user activity...")
	activityEmbed := discord.CreateUserActivityEmbed(
		"John Doe",
		"Login",
		"Logged in from profile page.",
	)

	err = discord.WebhookBuilder().
		SetURL(webhookURL).
		SetUsername("Activity Logger").
		AddEmbed(activityEmbed).
		Send()
	if err != nil {
		log.Printf("Error sending user activity: %v", err)
	}

	// Advanced Example 5: System Status Monitoring
	fmt.Println("Sending system status...")
	statusEmbed := discord.CreateSystemStatusEmbed(
		"Normal",
		"7 days 12 hours 30 minutes",
		"45%",
		"67%",
	)

	err = discord.WebhookBuilder().
		SetURL(webhookURL).
		SetUsername("System Monitor").
		AddEmbed(statusEmbed).
		Send()
	if err != nil {
		log.Printf("Error sending system status: %v", err)
	}

	// Advanced Example 6: Comprehensive Report (Multiple Embeds)
	fmt.Println("Sending comprehensive report...")

	// Market Summary
	marketSummary := discord.CreateMarketReport(
		"2,450.12 (+1.2%)",
		"850.45 (+0.8%)",
		"1,320 KRW/USD",
	)

	// Major News
	news1 := discord.CreateNewsEmbed(
		"Samsung Electronics Stock Rise",
		"Stock price continues to rise due to semiconductor sales increase",
		"Maeil Business Newspaper",
	)

	news2 := discord.CreateNewsEmbed(
		"SK Hynix Earnings Announcement",
		"Improved performance due to memory semiconductor price increase",
		"Korea Economic Daily",
	)

	// System Status
	systemStatus := discord.CreateSystemStatusEmbed(
		"Normal",
		"7 days 12 hours 30 minutes",
		"45%",
		"67%",
	)

	err = discord.WebhookBuilder().
		SetURL(webhookURL).
		SetUsername("QuiredLab Daily Report").
		SetContent("🌅 **Daily Comprehensive Report**").
		AddEmbed(marketSummary).
		AddEmbed(news1).
		AddEmbed(news2).
		AddEmbed(systemStatus).
		Send()
	if err != nil {
		log.Printf("Error sending comprehensive report: %v", err)
	}

	// Advanced Example 7: Error Notification
	fmt.Println("Sending error notification...")
	err = discord.SendErrorNotification(
		webhookURL,
		"API_001",
		"External API Connection Failed",
		"Connection to data collection service failed 3 times in a row.",
	)
	if err != nil {
		log.Printf("Error sending error notification: %v", err)
	}

	// Advanced Example 8: Success Notification
	fmt.Println("Sending success notification...")
	err = discord.SendSuccessNotification(
		webhookURL,
		"Data Collection Complete",
		"Today's market data collection has been completed successfully.",
	)
	if err != nil {
		log.Printf("Error sending success notification: %v", err)
	}

	fmt.Println("All advanced webhook examples completed!")
}
