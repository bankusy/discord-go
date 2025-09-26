# Discord Webhook Builder for Go

A Discord webhook builder written in Go. Create intuitive and flexible webhook messages using the chaining pattern.

## Key Features

-   🔗 **Chaining Pattern**: Intuitive API through method chaining
-   🎨 **Rich Embeds**: Various embed options including title, description, fields, colors, and images
-   🚀 **Quick Sending**: Send webhooks instantly with simple functions
-   🛠️ **Helper Functions**: Pre-defined templates for market reports, news alerts, error notifications, etc.
-   🎯 **Type Safety**: Safe API leveraging Go's type system

## Installation

```bash
go mod tidy
```

## Basic Usage

### 1. Simple Text Message

```go
import discord "github.com/bankusy/discord-go/model"

err := discord.WebhookBuilder().
    SetURL("YOUR_WEBHOOK_URL").
    SetContent("Hello! This is a test message.").
    Send()
```

### 2. Embed Message

```go
embed := discord.EmbedBuilder().
    SetTitle("📊 Market Analysis Report").
    SetDescription("Analyzing today's major market trends.").
    SetColorHex("#00ff00").
    AddField("KOSPI", "2,450.12 (+1.2%)", true).
    AddField("NASDAQ", "14,250.45 (+0.8%)", true).
    SetFooter("QuiredLab", "").
    SetCurrentTimestamp()

err := discord.WebhookBuilder().
    SetURL("YOUR_WEBHOOK_URL").
    SetUsername("QuiredLab Bot").
    AddEmbed(embed).
    Send()
```

### 3. Multiple Embeds Message

```go
marketEmbed := discord.EmbedBuilder().
    SetTitle("📈 Market Summary").
    SetColorHex("#3498db").
    AddField("KOSPI", "2,450.12 (+1.2%)", true).
    SetCurrentTimestamp()

newsEmbed := discord.EmbedBuilder().
    SetTitle("📰 Major News").
    SetColorHex("#e74c3c").
    AddField("Samsung Electronics", "Stock price rises due to semiconductor sales increase", false).
    SetFooter("Source: Various News", "")

err := discord.WebhookBuilder().
    SetURL("YOUR_WEBHOOK_URL").
    SetContent("🌅 **Today's Market Report**").
    AddEmbed(marketEmbed).
    AddEmbed(newsEmbed).
    Send()
```

## Helper Functions Usage

### Market Report

```go
err := discord.SendMarketReport(
    "YOUR_WEBHOOK_URL",
    "2,450.12 (+1.2%)",  // KOSPI
    "850.45 (+0.8%)",    // KOSDAQ
    "1,320 KRW/USD",     // Exchange Rate
)
```

### News Alert

```go
err := discord.SendNewsAlert(
    "YOUR_WEBHOOK_URL",
    "Samsung Electronics Stock Rise",
    "Stock price continues to rise due to increased semiconductor sales.",
    "Korea Economic Daily",
)
```

### Error Notification

```go
err := discord.SendErrorNotification(
    "YOUR_WEBHOOK_URL",
    "API_001",
    "External API Connection Failed",
    "Connection to data collection service failed 3 times in a row.",
)
```

### Success Notification

```go
err := discord.SendSuccessNotification(
    "YOUR_WEBHOOK_URL",
    "Data Collection Complete",
    "Today's market data collection has been completed successfully.",
)
```

## Embed Builder Methods

### Basic Settings

-   `SetTitle(title string)` - Embed title
-   `SetDescription(description string)` - Embed description
-   `SetURL(url string)` - Embed URL
-   `SetColor(color int)` - Color (decimal)
-   `SetColorHex(hexColor string)` - Color (hexadecimal)

### Fields and Content

-   `AddField(name, value string, inline bool)` - Add field
-   `SetFooter(text, iconURL string)` - Set footer
-   `SetAuthor(name, url, iconURL string)` - Set author
-   `SetThumbnail(url string)` - Set thumbnail
-   `SetImage(url string)` - Set image

### Time Settings

-   `SetTimestamp(timestamp time.Time)` - Set specific time
-   `SetCurrentTimestamp()` - Set current time

## Webhook Builder Methods

-   `SetURL(url string)` - Set webhook URL
-   `SetContent(content string)` - Set message content
-   `SetUsername(username string)` - Set username
-   `SetAvatarURL(avatarURL string)` - Set avatar URL
-   `SetTTS(tts bool)` - Set TTS
-   `AddEmbed(embed *Embed)` - Add embed
-   `Send()` - Send webhook

## Color Constants

```go
discord.ColorRed      // Red
discord.ColorGreen    // Green
discord.ColorBlue     // Blue
discord.ColorYellow   // Yellow
discord.ColorOrange   // Orange
discord.ColorPurple   // Purple
discord.ColorPink     // Pink
discord.ColorCyan     // Cyan
discord.ColorGray     // Gray
discord.ColorDarkGray // Dark Gray
```

## Running Examples

```bash
# Run basic examples
go run examples/webhook_example.go

# Run advanced examples
go run examples/advanced_example.go
```

## Error Handling

All `Send()` methods return an `error`. Check the return value for proper error handling:

```go
err := discord.WebhookBuilder().
    SetURL("YOUR_WEBHOOK_URL").
    SetContent("Test message").
    Send()

if err != nil {
    log.Printf("Webhook send failed: %v", err)
}
```

## License

MIT License

## Contributing

Please submit issues for bug reports or feature suggestions.
