package dto

type RegisterRes struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
}

type LoginRes struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token,omitempty"`
}

type Response struct {
	ID        uint   `json:"id"`
	Name      string `json:"name" `
	Email     string `json:"email"  `
	Token     string `json:"token,omitempty"`
	CreatedAt string `json:"created_at"`
}
