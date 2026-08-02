package controllers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/msartioli/helpdesk-api-go/models"
)

func Tickets(db *pgxpool.Pool) {
	http.HandleFunc("POST /tickets", func(w http.ResponseWriter, r *http.Request) {
		createTicketHandler(w, r, db)
	})
}

func createTicketHandler(
	w http.ResponseWriter,
	r *http.Request,
	db *pgxpool.Pool,
) {
	w.Header().Set("Content-Type", "application/json")

	var input models.CreateTicketRequest

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&input); err != nil {
		writeJSONError(w, http.StatusBadRequest, "JSON inválido")
		return
	}

	input.Title = strings.TrimSpace(input.Title)
	input.Description = strings.TrimSpace(input.Description)
	input.Requester = strings.TrimSpace(input.Requester)
	input.Category = strings.TrimSpace(input.Category)
	input.Priority = strings.TrimSpace(input.Priority)

	switch {
	case input.Title == "":
		writeJSONError(w, http.StatusBadRequest, "title é obrigatório")
		return

	case input.Description == "":
		writeJSONError(w, http.StatusBadRequest, "description é obrigatório")
		return

	case input.Requester == "":
		writeJSONError(w, http.StatusBadRequest, "requester é obrigatório")
		return

	case input.Category == "":
		writeJSONError(w, http.StatusBadRequest, "category é obrigatório")
		return

	case input.Priority == "":
		writeJSONError(w, http.StatusBadRequest, "priority é obrigatório")
		return
	}

	const query = `
		INSERT INTO tickets (
			titulo,
			descricao,
			solicitante,
			categoria,
			prioridade
		)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING
			id,
			titulo,
			descricao,
			solicitante,
			categoria,
			prioridade,
			status,
			data_criacao;
	`

	var ticket models.Ticket

	err := db.QueryRow(
		r.Context(),
		query,
		input.Title,
		input.Description,
		input.Requester,
		input.Category,
		input.Priority,
	).Scan(
		&ticket.ID,
		&ticket.Title,
		&ticket.Description,
		&ticket.Requester,
		&ticket.Category,
		&ticket.Priority,
		&ticket.Status,
		&ticket.CreatedAt,
	)

	if err != nil {
		writeJSONError(
			w,
			http.StatusInternalServerError,
			"erro ao salvar chamado",
		)
		return
	}

	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(ticket); err != nil {
		return
	}
}

func writeJSONError(
	w http.ResponseWriter,
	status int,
	message string,
) {
	w.WriteHeader(status)

	_ = json.NewEncoder(w).Encode(map[string]string{
		"error": message,
	})
}
