package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// client.GetDevices() retrieves all devices linked to organization
func (c *client) GetDevices() ([]DeviceInfo, error) {
	// Refresh tokens if needed
	claims, err := c.RefreshTokens()
	if err != nil {
		log.Fatalf("Failed to refresh tokens, please re-login.\n%w\n", err)
	}

	// Call API to get devices
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/worker/%s", c.baseURL, claims.OrgID), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	c.addHeaders(req)
	res, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("API request failed: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Fatalf("Get devices command failed.\n%w\n", res.StatusCode)
	}

	// Unmarshal response body into DeviceInfo slice
	BodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("Error reading from reader: %v", err)
	}
	var devices []DeviceInfo
	err = json.Unmarshal(BodyBytes, &devices)
	if err != nil {
		log.Fatalf("Error unmarshaling response body: %v", err)
	}

	return devices, nil
}

// client.GetDevice() retrieves information for a specific device by AgentID
func (c *client) GetDevice(devID string) (DeviceInfo, error) {
	// Refresh tokens if needed
	claims, err := c.RefreshTokens()
	if err != nil {
		log.Fatalf("Failed to refresh tokens, please re-login.\n%w\n", err)
	}

	// Call API to get devices
	var device DeviceInfo
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/worker/%s/%s",
		c.baseURL, claims.OrgID, devID), nil)
	if err != nil {
		return device, fmt.Errorf("failed to create request: %w", err)
	}
	c.addHeaders(req)
	res, err := c.http.Do(req)
	if err != nil {
		return device, fmt.Errorf("API request failed: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Fatalf("Get device command failed.\n%w\n", res.StatusCode)
	}

	// Unmarshal response body into DeviceInfo slice
	BodyBytes, err := io.ReadAll(res.Body)
	//fmt.Printf(string(BodyBytes))
	if err != nil {
		log.Fatalf("Error reading from reader: %v", err)
	}
	err = json.Unmarshal(BodyBytes, &device)
	if err != nil {
		log.Fatalf("Error unmarshaling response body: %v", err)
	}

	return device, nil
}

// client.AuthBootBox() authorizes a BootBox using the provided OTP
func (c *client) AuthBootBox(ctx context.Context, otp string) error {
	// Configure request
	url := fmt.Sprintf("%s/device/auth/%s", c.baseURL, otp)
	req, err := http.NewRequestWithContext(ctx, "POST", url, nil)
	if err != nil {
		return fmt.Errorf("failed to create login request: %v", err)
	}
	c.addHeaders(req)

	// Make OTP request
	res, err := c.http.Do(req)
	if err != nil {
		return fmt.Errorf("API request failed: %w", err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		log.Fatalf("Command failed.\n%w\n", res.StatusCode)
	}

	return nil
}
