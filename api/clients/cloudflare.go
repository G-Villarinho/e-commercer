package clients

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"

	"github.com/g-villarinho/xp-life-api/config"
	"github.com/g-villarinho/xp-life-api/pkgs"
)

type CloudflareResponse struct {
	Success bool `json:"success"`
	Result  struct {
		Variants []string `json:"variants"`
	} `json:"result"`
	Errors []struct {
		Message string `json:"message"`
	} `json:"errors"`
}

type CloudflareClient interface {
	UploadImage(ctx context.Context, image []byte, filename string) (string, error)
}

type cloudflareClient struct {
	di         *pkgs.Di
	httpClient *http.Client
}

func NewCloudflareClient(di *pkgs.Di) (CloudflareClient, error) {
	return &cloudflareClient{
		di:         di,
		httpClient: &http.Client{Timeout: config.Env.CloudFlareImage.Timeout},
	}, nil
}

func (c *cloudflareClient) UploadImage(ctx context.Context, image []byte, filename string) (string, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	fileWriter, err := writer.CreateFormFile("file", filename)
	if err != nil {
		return "", fmt.Errorf("error creating form file: %w", err)
	}

	if _, err := io.Copy(fileWriter, bytes.NewReader(image)); err != nil {
		return "", fmt.Errorf("error copying file: %w", err)
	}

	if err := writer.Close(); err != nil {
		return "", fmt.Errorf("error closing writer: %w", err)
	}

	req, err := http.NewRequest("POST", config.Env.CloudFlareImage.URL, body)
	if err != nil {
		return "", fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", config.Env.CloudFlareImage.Token))
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("upload error with status code: %d", resp.StatusCode)
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response: %w", err)
	}

	var cfResp CloudflareResponse
	if err := json.Unmarshal(respBody, &cfResp); err != nil {
		return "", fmt.Errorf("error decoding JSON response: %w", err)
	}

	if !cfResp.Success {
		return "", fmt.Errorf("cloudflare response error: %+v", cfResp.Errors)
	}

	if len(cfResp.Result.Variants) == 0 {
		return "", fmt.Errorf("no image URL returned")
	}

	return cfResp.Result.Variants[0], nil
}
