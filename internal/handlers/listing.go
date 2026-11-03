package handlers

import (
	"database/sql"
	"encoding/json"
	"example/backend/internal/httpx"
	"example/backend/internal/middleware"
	"log/slog"
	"net/http"
	"time"
)

type listing struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Price       int64     `json:"price"`
	City        string    `json:"city"`
	CreatedAt   time.Time `json:"created_at"`
}

type ListingHandlers struct {
	db     *sql.DB
	logger *slog.Logger
}

// dependency injection handler
func NewListingHandler(db *sql.DB, logger *slog.Logger) *ListingHandlers {
	return &ListingHandlers{
		db:     db,
		logger: logger,
	}
}

func (lh *ListingHandlers) List(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	xRequestID := middleware.GetRequestID(ctx)
	rows, err := lh.db.QueryContext(ctx,
		`SELECT id, title, description, price, city, created_at 
		    FROM listings
			ORDER BY created_at DESC
			LIMIT 100`)
	if err != nil {
		lh.logger.Error("rows.db.QueryContext", "err", err, "x-request-id", xRequestID)
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	listings := []listing{}

	for rows.Next() {
		var l listing
		if err := rows.Scan(&l.ID, &l.Title, &l.Description, &l.Price, &l.City, &l.CreatedAt); err != nil {
			lh.logger.Error("rows.Scan", "err", err, "x-request-id", xRequestID)
			http.Error(w, "internal error", http.StatusInternalServerError)
			return
		}
		listings = append(listings, l)
	}

	if err := rows.Err(); err != nil {
		lh.logger.Error("rows.Err", "err", err, "x-request-id", xRequestID)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	lh.logger.Info("List", "x-request-id", xRequestID)

	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(http.StatusOK)

	_ = json.NewEncoder(w).Encode(listings)
}

func (lh *ListingHandlers) Delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := r.PathValue("id")
	xRequestID := middleware.GetRequestID(ctx)
	_, err := lh.db.ExecContext(ctx, `DELETE FROM listings WHERE id = $1`, id)
	if err != nil {
		lh.logger.Error("Delete List Error", "listing id", id, "error", err, "x-request-id", xRequestID)
		httpx.Error(w, http.StatusInternalServerError, "Something went wrong", httpx.CodeInvalidID)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (lh *ListingHandlers) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	xRequestID := middleware.GetRequestID(ctx)
	var req CreateListingRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		lh.logger.Error("Failed to decode", "err", err, "x-request-id", xRequestID)
		httpx.Error(w, http.StatusBadRequest, "invalid body", httpx.CodeInvalidBody)
		return
	}

	if err := req.Validate(); err != nil {
		httpx.Error(w, http.StatusUnprocessableEntity, err.Error(), httpx.CodeValidationFailed)
		return
	}

	row := lh.db.QueryRowContext(ctx, `
	   INSERT INTO listings (title, description, price, city) VALUES ($1, $2, $3, $4) RETURNING id, title, created_at
	 `, req.Title, req.Description, req.Price, req.City)

	var out CreateListingResponse
	if err := row.Scan(&out.ID, &out.Title, &out.CreatedAt); err != nil {
		lh.logger.Error("Failed to insert", "err", err, "requestID", xRequestID)
		httpx.Error(w, http.StatusBadRequest, "something went wrong", httpx.CodeInternalServerError)
		return
	}

	lh.logger.Info("listing created", "id", &out.ID, "x-request-id", xRequestID)

	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(http.StatusCreated)

	_ = json.NewEncoder(w).Encode(out)
}
