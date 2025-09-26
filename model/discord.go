package discord

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	logger "github.com/bankusy/logger/model"
)

type Webhook struct {
	URL  string
	Body *WebhookBody
}

type EmbedField struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline,omitempty"`
}

type Footer struct {
	Text    string `json:"text"`
	IconURL string `json:"icon_url,omitempty"`
}

type Author struct {
	Name    string `json:"name"`
	URL     string `json:"url,omitempty"`
	IconURL string `json:"icon_url,omitempty"`
}

type Thumbnail struct {
	URL string `json:"url"`
}

type Image struct {
	URL string `json:"url"`
}

type Embed struct {
	Title       string       `json:"title,omitempty"`
	URL         string       `json:"url,omitempty"`
	Description string       `json:"description,omitempty"`
	Color       int          `json:"color,omitempty"`
	Fields      []EmbedField `json:"fields,omitempty"`
	Footer      *Footer      `json:"footer,omitempty"`
	Author      *Author      `json:"author,omitempty"`
	Thumbnail   *Thumbnail   `json:"thumbnail,omitempty"`
	Image       *Image       `json:"image,omitempty"`
	Timestamp   *time.Time   `json:"timestamp,omitempty"`
}

type WebhookBody struct {
	Content   string  `json:"content,omitempty"`
	Username  string  `json:"username,omitempty"`
	AvatarURL string  `json:"avatar_url,omitempty"`
	TTS       bool    `json:"tts,omitempty"`
	Embeds    []Embed `json:"embeds,omitempty"`
}

func WebhookBuilder() *Webhook {
	return &Webhook{
		Body: &WebhookBody{
			Embeds: make([]Embed, 0),
		},
	}
}

func (wh *Webhook) SetURL(url string) *Webhook {
	wh.URL = url
	return wh
}

func (wh *Webhook) SetContent(content string) *Webhook {
	wh.Body.Content = content
	return wh
}

func (wh *Webhook) SetUsername(username string) *Webhook {
	wh.Body.Username = username
	return wh
}

func (wh *Webhook) SetAvatarURL(avatarURL string) *Webhook {
	wh.Body.AvatarURL = avatarURL
	return wh
}

func (wh *Webhook) SetTTS(tts bool) *Webhook {
	wh.Body.TTS = tts
	return wh
}

func (wh *Webhook) AddEmbed(embed *Embed) *Webhook {
	wh.Body.Embeds = append(wh.Body.Embeds, *embed)
	return wh
}

func (wh *Webhook) Send() error {
	if wh.URL == "" {
		return fmt.Errorf("webhook URL is required")
	}

	webhookBody, err := json.Marshal(wh.Body)
	if err != nil {
		logger.Error("Failed to marshal webhook body:", err)
		return err
	}

	resp, err := http.Post(wh.URL, "application/json", bytes.NewBuffer(webhookBody))
	if err != nil {
		logger.Error("Failed to send webhook:", err)
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return nil
	}

	return fmt.Errorf("[%d] Failed to send webhook", resp.StatusCode)
}

func EmbedBuilder() *Embed {
	return &Embed{
		Fields: make([]EmbedField, 0),
	}
}

func (e *Embed) SetTitle(title string) *Embed {
	e.Title = title
	return e
}

func (e *Embed) SetURL(url string) *Embed {
	e.URL = url
	return e
}

func (e *Embed) SetDescription(description string) *Embed {
	e.Description = description
	return e
}

func (e *Embed) SetColor(color int) *Embed {
	e.Color = color
	return e
}

func (e *Embed) SetColorHex(hexColor string) *Embed {
	// Convert hex to decimal
	var color int
	fmt.Sscanf(hexColor, "#%x", &color)
	e.Color = color
	return e
}

func (e *Embed) AddField(name, value string, inline bool) *Embed {
	e.Fields = append(e.Fields, EmbedField{
		Name:   name,
		Value:  value,
		Inline: inline,
	})
	return e
}

func (e *Embed) SetFooter(text, iconURL string) *Embed {
	e.Footer = &Footer{
		Text:    text,
		IconURL: iconURL,
	}
	return e
}

func (e *Embed) SetAuthor(name, url, iconURL string) *Embed {
	e.Author = &Author{
		Name:    name,
		URL:     url,
		IconURL: iconURL,
	}
	return e
}

func (e *Embed) SetThumbnail(url string) *Embed {
	e.Thumbnail = &Thumbnail{URL: url}
	return e
}

func (e *Embed) SetImage(url string) *Embed {
	e.Image = &Image{URL: url}
	return e
}

func (e *Embed) SetTimestamp(timestamp time.Time) *Embed {
	e.Timestamp = &timestamp
	return e
}

func (e *Embed) SetCurrentTimestamp() *Embed {
	now := time.Now()
	e.Timestamp = &now
	return e
}
