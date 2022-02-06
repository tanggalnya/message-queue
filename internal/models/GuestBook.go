package models

type GuestBook struct {
	From      string `json:"from"`
	IsPublic  bool   `json:"is_public"`
	Name      string `json:"name"`
	UpdatedAt string `json:"updated_at"`
	CreatedAt string `json:"created_at"`
	ID        string `json:"id"`
	Message   string `json:"message"`
	EventID   string `json:"event_id"`
}
