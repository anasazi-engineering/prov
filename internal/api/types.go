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
