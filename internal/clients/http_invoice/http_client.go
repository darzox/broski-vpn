package http_invoice

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

const (
	baseURL = "https://api.telegram.org"
)

type TokenGetter interface {
	Token() string
}

type TelegramHTTPClient struct {
	client *http.Client
	token  string
	url    string
}

func NewTelegramHTTPClient(tokenGetter TokenGetter) (*TelegramHTTPClient, error) {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	url := fmt.Sprintf("%s/bot%s/sendInvoice", baseURL, tokenGetter.Token())

	return &TelegramHTTPClient{
		client: client,
		token:  tokenGetter.Token(),
		url:    url,
	}, nil
}

func (c *TelegramHTTPClient) SendInvoice(userId int64, amount int, dayAmount int) error {
	body := invoice{
		ChatID:         userId,
		Title:          fmt.Sprintf("Ключ на %d дней", dayAmount),
		Description:    fmt.Sprintf("Ключ на %d дней для доступа к VPN", dayAmount),
		Payload:        fmt.Sprintf(`{chat_id: %d, amount: %d, period: %d}`, userId, amount, dayAmount),
		ProviderToken:  "",
		StartParameter: "",
		ProviderData:   "{}",
		Prices: []price{
			{
				Label:  fmt.Sprintf("%d дней", dayAmount),
				Amount: amount,
			},
		},
		Currency: "XTR",
	}

	jsonData, err := json.Marshal(body)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(context.Background(), http.MethodPost, c.url, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("payment error: %v")
	}

	return nil
}
