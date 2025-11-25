package api

import (
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
