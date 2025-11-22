package api

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"prov/internal/config"
)

type Client interface {
	GetUser(ctx context.Context, id string) (*User, error)
	ListUsers(ctx context.Context) ([]User, error)
	Login(ctx context.Context, creds Credentials) (config.Tokens, error)
}

type client struct {
	baseURL string
	token   config.Tokens
	http    *http.Client
}

func NewClient(baseURL string, tokens config.Tokens) Client {
	//tokens := config.LoadTokens()

	return &client{
		baseURL: baseURL,
		token:   tokens,
		http: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (c *client) GetUser(ctx context.Context, id string) (*User, error) {
	req, _ := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("%s/users/%s", c.baseURL, id), nil)
	c.addHeaders(req)

	res, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusNotFound {
		return nil, errors.New("user not found")
	}
	if res.StatusCode >= 400 {
		return nil, fmt.Errorf("API error: %s", res.Status)
	}

	var user User
	if err := json.NewDecoder(res.Body).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (c *client) ListUsers(ctx context.Context) ([]User, error) {
	req, _ := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("%s/users", c.baseURL), nil)
	c.addHeaders(req)

	res, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode >= 400 {
		return nil, fmt.Errorf("API error: %s", res.Status)
	}

	var users []User
	if err := json.NewDecoder(res.Body).Decode(&users); err != nil {
		return nil, err
	}
	return users, nil
}

func (c *client) addHeaders(req *http.Request) {
	if c.token != (config.Tokens{}) {
		req.Header.Set("Authorization", "Bearer "+c.token.AccessToken)
	}
	req.Header.Set("Accept", "application/json")
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
		return config.Tokens{}, fmt.Errorf("API error: %s", res.Status)
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

	// Process tokens from response
	var tokens config.Tokens
	if err := json.NewDecoder(res.Body).Decode(&tokens); err != nil {
		fmt.Errorf("failed to parse login response: %v", err)
	}

	return tokens, nil
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
