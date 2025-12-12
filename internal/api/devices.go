/*
 * Anasazi Precision Engineering LLC CONFIDENTIAL
 *
 * Unpublished Copyright (c) 2025 Anasazi Precision Engineering LLC. All Rights Reserved.
 *
 * Proprietary to Anasazi Precision Engineering LLC and may be covered by patents, patents
 * in process, and trade secret or copyright law. Dissemination of this information or
 * reproduction of this material is strictly forbidden unless prior written
 * permission is obtained from Anasazi Precision Engineering LLC.
 */
package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// client.GetWorkers() retrieves all devices linked to organization
func (c *client) GetWorkers() ([]DeviceInfo, error) {
	// Refresh tokens if needed
	claims, err := c.RefreshTokens()
	if err != nil {
		return nil, fmt.Errorf("Failed to refresh tokens, please re-login.%w", err)
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
		return nil, fmt.Errorf("Get workers command failed, status code: %w", res.StatusCode)
	}

	// Unmarshal response body into DeviceInfo slice
	BodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("Error reading from reader: %v", err)
	}
	var devices []DeviceInfo
	err = json.Unmarshal(BodyBytes, &devices)
	if err != nil {
		return nil, fmt.Errorf("Error unmarshaling response body: %w", err)
	}

	return devices, nil
}

// client.GetWorker() retrieves information for a specific device by AgentID
func (c *client) GetWorker(devID string) (DeviceInfo, error) {
	// Refresh tokens if needed
	claims, err := c.RefreshTokens()
	if err != nil {
		return DeviceInfo{}, fmt.Errorf("Failed to refresh tokens, please re-login.%w", err)
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
		return device, fmt.Errorf("Get worker command failed, status code: %v", res.StatusCode)
	}

	// Unmarshal response body into DeviceInfo slice
	BodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return DeviceInfo{}, fmt.Errorf("Error reading from reader: %v", err)
	}
	err = json.Unmarshal(BodyBytes, &device)
	if err != nil {
		return DeviceInfo{}, fmt.Errorf("Error unmarshaling response body: %v", err)
	}

	return device, nil
}

// client.GetBootboxes() retrieves all BootBoxes linked to organization
func (c *client) GetBootboxes() ([]DeviceInfo, error) {
	// Refresh tokens if needed
	claims, err := c.RefreshTokens()
	if err != nil {
		return nil, fmt.Errorf("Failed to refresh tokens, please re-login.%w", err)
	}

	// Call API to get devices
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/bootbox/%s", c.baseURL, claims.OrgID), nil)
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
		return nil, fmt.Errorf("Get bootboxes command failed, status code: %w", res.StatusCode)
	}

	// Unmarshal response body into DeviceInfo slice
	BodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("Error reading from reader: %v", err)
	}
	var devices []DeviceInfo
	err = json.Unmarshal(BodyBytes, &devices)
	if err != nil {
		return nil, fmt.Errorf("Error unmarshaling response body: %w", err)
	}

	return devices, nil
}

// client.GetBootbox() retrieves information for a specific BootBox by AgentID
func (c *client) GetBootbox(devID string) (DeviceInfo, error) {
	// Refresh tokens if needed
	claims, err := c.RefreshTokens()
	if err != nil {
		return DeviceInfo{}, fmt.Errorf("Failed to refresh tokens, please re-login.%w", err)
	}

	// Call API to get devices
	var device DeviceInfo
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/bootbox/%s/%s",
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
		return device, fmt.Errorf("Get bootbox command failed, status code: %v", res.StatusCode)
	}

	// Unmarshal response body into DeviceInfo slice
	BodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return DeviceInfo{}, fmt.Errorf("Error reading from reader: %v", err)
	}
	err = json.Unmarshal(BodyBytes, &device)
	if err != nil {
		return DeviceInfo{}, fmt.Errorf("Error unmarshaling response body: %v", err)
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
		log.Fatalf("Command failed.\n%v\n", res.StatusCode)
	}

	return nil
}
