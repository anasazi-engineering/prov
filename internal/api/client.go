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
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"prov/internal/config"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

type Client interface {
	Login(ctx context.Context, creds Credentials) (config.Tokens, error)
	Logout(ctx context.Context) error
	GetWorkers() ([]DeviceInfo, error)
	GetWorker(devID string) (DeviceInfo, error)
	GetBootboxes() ([]DeviceInfo, error)
	GetBootbox(devID string) (DeviceInfo, error)
	AuthBootBox(ctx context.Context, otp string) error
	ApplyRecipe(ctx context.Context, agegntID string, url string) error
}

type client struct {
	baseURL string
	token   config.Tokens
	http    *http.Client
}

func NewClient(baseURL string, tokens config.Tokens) Client {
	return &client{
		baseURL: baseURL,
		token:   tokens,
		http: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (c *client) addHeaders(req *http.Request) {
	if c.token != (config.Tokens{}) {
		req.Header.Set("Authorization", "Bearer "+c.token.AccessToken)
	}
	req.Header.Set("Accept", "application/json")
}

// client.RefreshTokens() refreshes the access and refresh tokens using the
// current refresh token.
func (c *client) RefreshTokens() (jwtClaims, error) {
	// Get JWT claims from access token
	var claims jwtClaims
	now := time.Now().Unix()
	token, _, err := jwt.NewParser().ParseUnverified(c.token.AccessToken, jwt.MapClaims{})
	if err != nil {
		return jwtClaims{}, fmt.Errorf("failed to parse JWT: %w", err)
	}
	rawClaims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return jwtClaims{}, fmt.Errorf("unable to convert claims to MapClaims")
	}
	claims.Username, _ = rawClaims.GetSubject()
	claims.OrgID, _ = rawClaims["oid"].(string)
	expAt, _ := rawClaims.GetExpirationTime()
	claims.ExpiresAt = expAt.Unix()

	// If access token is still valid, return existing claims
	if claims.ExpiresAt > now {
		return claims, nil
	}

	// Return with error if refresh token is expired
	token, _, err = jwt.NewParser().ParseUnverified(c.token.RefreshToken, jwt.MapClaims{})
	if err != nil {
		return jwtClaims{}, fmt.Errorf("failed to parse JWT: %w", err)
	}
	rawClaims, ok = token.Claims.(jwt.MapClaims)
	if !ok {
		return jwtClaims{}, fmt.Errorf("unable to convert claims to MapClaims")
	}
	expAt, _ = rawClaims.GetExpirationTime()
	claims.ExpiresAt = expAt.Unix()

	// Check if refresh token is expired
	if claims.ExpiresAt < now {
		return claims, fmt.Errorf("refresh token expired, need to re-login")
	}

	// Configure request
	req, _ := http.NewRequest("POST", fmt.Sprintf("%s/auth/refresh", c.baseURL), nil)

	// Add refresh token to cookie
	req.AddCookie(&http.Cookie{
		Name:  "refresh_token",
		Value: c.token.RefreshToken,
	})

	// Make refresh API request
	res, err := c.http.Do(req)
	if err != nil {
		return jwtClaims{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(res.Body)
		return jwtClaims{}, fmt.Errorf("refresh error: %s", string(bodyBytes))
	}

	// Get new tokens from response
	var tokens config.Tokens
	if err := json.NewDecoder(res.Body).Decode(&tokens); err != nil {
		return jwtClaims{}, fmt.Errorf("failed to parse refresh response: %v", err)
	}
	// Get refresh token from Server response cookie
	var refreshToken string
	for _, cookie := range res.Cookies() {
		if cookie.Name == "refresh_token" {
			refreshToken = cookie.Value
			break
		}
	}
	if refreshToken == "" {
		return jwtClaims{}, fmt.Errorf("Refresh Error: refresh token not found in refresh response")
	}
	tokens.RefreshToken = refreshToken

	c.token = tokens
	viper.Set("access_token", tokens.AccessToken)
	viper.Set("refresh_token", tokens.RefreshToken)
	viper.WriteConfig()

	return claims, nil
}

// client.Login() performs user login, including 2FA, and returns tokens on success
func (c *client) Login(ctx context.Context, creds Credentials) (config.Tokens, error) {

	// Configure request
	url := fmt.Sprintf("%s/auth/login", c.baseURL)
	req, err := http.NewRequestWithContext(ctx, "POST", url, nil)
	if err != nil {
		return config.Tokens{}, fmt.Errorf("failed to create login request: %v", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	// Add payload body to request
	var reqPayload struct {
		UserID   string `json:"user_id"`
		Password string `json:"password"`
	}
	reqPayload.UserID = creds.Username
	reqPayload.Password = creds.Password
	jsonBody, err := json.Marshal(reqPayload)
	if err != nil {
		return config.Tokens{}, fmt.Errorf("failed to create login request: %v", err)
	}
	req.Body = io.NopCloser(bytes.NewReader(jsonBody))

	// Make login API request
	res, err := c.http.Do(req)
	if err != nil {
		return config.Tokens{}, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusAccepted {
		fmt.Println(res.Body)
		return config.Tokens{}, fmt.Errorf("Login failed: %s", res.Status)
	}

	// Configure TOTP request
	url = fmt.Sprintf("%s/auth/verify-totp", c.baseURL)
	req, err = http.NewRequestWithContext(ctx, "POST", url, nil)
	if err != nil {
		return config.Tokens{}, fmt.Errorf("failed to create login request: %v", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	// Get TOTP from user input from the console here
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	totp, err := ReadString(ctx, "Enter Authenticator TOTP: ")
	if err != nil {
		fmt.Println("Error:", err)
		return config.Tokens{}, fmt.Errorf("failed to read TOTP from input: %v", err)
	}

	// Add payload to body
	var totpPayload struct {
		Username string `json:"username"`
		Role     string `json:"role"`
		TOTP     string `json:"totp"`
		OrgID    string `json:"org_id"`
	}
	totpPayload.Username = creds.Username
	totpPayload.Role = "admin"
	totpPayload.TOTP = totp
	totpPayload.OrgID = creds.OrgID
	jsonBody, err = json.Marshal(totpPayload)
	if err != nil {
		return config.Tokens{}, fmt.Errorf("failed to create login request: %v", err)
	}
	req.Body = io.NopCloser(bytes.NewReader(jsonBody))

	// Make TOTP verification API request
	res, err = c.http.Do(req)
	if err != nil {
		return config.Tokens{}, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		fmt.Println(res.Body)
		return config.Tokens{}, fmt.Errorf("TOTP verification failed: %s", res.Status)
	}

	// Process tokens from response
	var tokens config.Tokens
	if err := json.NewDecoder(res.Body).Decode(&tokens); err != nil {
		return config.Tokens{}, fmt.Errorf("failed to parse login response: %v", err)
	}

	// Get refresh token from Server response cookie
	var refreshToken string
	for _, cookie := range res.Cookies() {
		if cookie.Name == "refresh_token" {
			refreshToken = cookie.Value
			break
		}
	}
	if refreshToken == "" {
		return config.Tokens{}, fmt.Errorf("Login Error: refresh token not found in login response")
	}
	tokens.RefreshToken = refreshToken

	return tokens, nil
}

func (c *client) Logout(ctx context.Context) error {
	req, _ := http.NewRequestWithContext(ctx, "POST", fmt.Sprintf("%s/auth/logout", c.baseURL), nil)
	c.addHeaders(req)

	res, err := c.http.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	// Clear local tokens
	viper.Set("access_token", "")
	viper.Set("refresh_token", "")
	viper.WriteConfig()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("Logout error: \n%s", res.Status)
	}

	return nil
}

// Client.ApplyRecipe() is used to apply a recipe to a Worker device.
func (c *client) ApplyRecipe(ctx context.Context, agegntID string, url string) error {
	// Refresh tokens if needed
	claims, err := c.RefreshTokens()
	if err != nil {
		log.Fatalf("Failed to refresh tokens, please re-login.\n%v\n", err)
	}

	// Configure request
	req, _ := http.NewRequestWithContext(ctx, "POST", fmt.Sprintf("%s/device/%s/%s/command",
		c.baseURL, claims.OrgID, agegntID), nil)
	c.addHeaders(req)

	//curl -X POST -H "Content-Type: application/json"
	//   -d '{"method":"APPLY_RECIPE", "payload":{"url":"http://www.yayteam.com"}}'
	//   -H 'Authorization: {"username":"mgmillsa", "org_id":"org1"}'
	//   localhost:8000/api/v1/device/org1/1230a3b5-5d5d-4f6a-b87e-19cf1ce08511/command

	// Create request payload
	var reqPayload struct {
		Method  string `json:"method"`
		Payload struct {
			URL string `json:"url"`
		} `json:"payload"`
	}
	reqPayload.Method = "APPLY_RECIPE"
	reqPayload.Payload.URL = url
	jsonBody, err := json.Marshal(reqPayload)
	if err != nil {
		return fmt.Errorf("failed to create ApplyRecipe request: %v", err)
	}
	req.Body = io.NopCloser(bytes.NewReader(jsonBody))

	// Make API request
	res, err := c.http.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("provisioning error: \n%s", res.Status)
	}

	return nil
}

// ReadString() prompts the user for input and returns a sanitized string.
func ReadString(ctx context.Context, prompt string) (string, error) {
	if prompt != "" {
		fmt.Print(prompt)
	}

	reader := bufio.NewReader(os.Stdin)

	inputCh := make(chan string)
	errCh := make(chan error)

	go func() {
		line, err := reader.ReadString('\n')
		if err != nil {
			errCh <- err
			return
		}
		inputCh <- strings.TrimSpace(line)
	}()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case err := <-errCh:
		return "", fmt.Errorf("failed to read input: %w", err)
	case line := <-inputCh:
		if line == "" {
			return "", errors.New("input cannot be empty")
		}
		return line, nil
	}
}
