package outline

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

type VpnUrlGetter interface {
	VpnUrl() string
}

type OutlineHttpClient struct {
	client *http.Client
	url    string
}

func NewOutlineHttpClient(vpnUrlGetter VpnUrlGetter) (*OutlineHttpClient, error) {
	client := &http.Client{
		Timeout: 5 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	url := vpnUrlGetter.VpnUrl()

	return &OutlineHttpClient{
		client: client,
		url:    url,
	}, nil
}

func (c *OutlineHttpClient) CreateAccessKey() (string, int64, error) {
	url := fmt.Sprintf("%s/access-keys", c.url)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, nil)
	if err != nil {
		return "", 0, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return "", 0, fmt.Errorf("failed to create access key: %v", err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", 0, fmt.Errorf("failed to read response body: %v", err)
	}

	var r createAccessKeyResp

	if err := json.Unmarshal(body, &r); err != nil {
		return "", 0, fmt.Errorf("failed to unmarshal body: %v", err)
	}

	keyId, err := strconv.ParseInt(r.Id, 10, 64)
	if err != nil {
		return "", 0, fmt.Errorf("failed to parse key id: %v", err)
	}

	return r.AccessUrl, keyId, nil
}

func (c *OutlineHttpClient) DeleteKey(keyId int64) error {
	url := fmt.Sprintf("%s/access-keys/%d", c.url, keyId)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, url, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("failed to delete access key: %v", err)
	}

	return nil
}
