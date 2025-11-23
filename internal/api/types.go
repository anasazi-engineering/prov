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
