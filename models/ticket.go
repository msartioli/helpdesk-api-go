package models

import "time"

type CreateTicketRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Requester   string `json:"requester"`
	Category    string `json:"category"`
	Priority    string `json:"priority"`
}

type Ticket struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Requester   string    `json:"requester"`
	Category    string    `json:"category"`
	Priority    string    `json:"priority"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
}
