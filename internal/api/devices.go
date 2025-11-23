package api

import (
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// client.GetDevices() retrieves all devices linked to organization
func (c *client) GetDevices() ([]Device, error) {
	var claims jwtClaims
	token, _, err := jwt.NewParser().ParseUnverified(c.token.AccessToken, jwt.MapClaims{})
	if err != nil {
		return nil, fmt.Errorf("failed to parse JWT: %w", err)
	}
	rawClaims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("unable to convert claims to MapClaims")
	}
	claims.Username, _ = rawClaims.GetSubject()
	claims.OrgID, _ = rawClaims["oid"].(string)
	expAt, _ := rawClaims.GetExpirationTime()
	claims.ExpiresAt = expAt.Unix()

	// TODO: need to also check if Refresh token is expired

	// Refresh tokens if needed
	now := time.Now().Unix()
	if claims.ExpiresAt < now {
		log.Println("Access token expired, refreshing tokens...") // TODO: remove
		err := c.RefreshTokens()
		if err != nil {
			return nil, fmt.Errorf("Failed to refresh tokens: %w", err)
		}
	} else {
		log.Println("Access token valid, proceeding with API call...") // TODO: remove
	}

	var devs []Device
	return devs, nil
}
