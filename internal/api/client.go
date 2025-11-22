package api

import (
    "context"
    "encoding/json"
    "errors"
    "fmt"
    "net/http"
    "time"
)

type Client interface {
    GetUser(ctx context.Context, id string) (*User, error)
    ListUsers(ctx context.Context) ([]User, error)
}

type client struct {
    baseURL string
    token   string
    http    *http.Client
}

func NewClient(baseURL, token string) Client {
    return &client{
        baseURL: baseURL,
        token:   token,
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
    if c.token != "" {
        req.Header.Set("Authorization", "Bearer "+c.token)
    }
    req.Header.Set("Accept", "application/json")
}
