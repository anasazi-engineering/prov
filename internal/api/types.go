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

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Credentials struct {
	Username string
	Password string
	OrgID    string
}

type Device struct {
	AgentID     string `json:"agent_id"`
	AssgnRecipe string `json:"assigned_recipe"`
	CreatedAt   int64  `json:"created_at"`
}

type jwtClaims struct {
	Username  string `json:"user_id"`
	OrgID     string `json:"org_id"`
	ExpiresAt int64  `json:"exp"`
}

type DeviceInfo struct {
	OrgID          string `json:"org_id"`
	AgentID        string `json:"agent_id"`
	FriendlyName   string `json:"friendly_name"`
	AssdRecipe     string `json:"assigned_recipe"`
	AssdRecipeAt   int64  `json:"assigned_recipe_at"`
	RecipeProgress int64  `json:"recipe_progress"`
	CreatedAt      int64  `json:"created_at"`
	LastSeen       int64  `json:"last_seen"`
}
