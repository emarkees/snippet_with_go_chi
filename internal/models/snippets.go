package models

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

type Snippet struct {
	ID        int
	Title     string
	Content   string
	CreatedAt time.Time
	ExpiresAt time.Time
}

/**
Define a SnippetModel type which wraps a sql.DB connection pool.
*/

type SnippetModel struct {
	DB *pgxpool.Pool
}

// func NewSnippetModel(db *pgxpool.Pool) *SnippetModel {
// 	return &SnippetModel{
// 		DB: db,
// 	}
// }

var ErrNoRecord = errors.New("models: no matching record found")

func (m *SnippetModel) Insert(title string, content string, expiresAt int) (int, error) {
	query := `
		INSERT INTO snippets (title, content, expires_at)
		VALUES ($1, $2, CURRENT_TIMESTAMP + ($3 * INTERVAL '1 day'))
		RETURNING id
	`

	var id int

	err := m.DB.QueryRow(context.Background(), query, title, content, expiresAt).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (m *SnippetModel) Get(ctx context.Context, id int) (*Snippet, error) {
	query := `
		SELECT id, title, content, created_at, expires_at
		FROM snippets
		WHERE id = $1
	`

	s := &Snippet{}

	// ctx, cancel := context.WithTimeouts()
	err := m.DB.QueryRow(ctx, query, id).Scan(&s.ID, &s.Title, &s.Content, &s.CreatedAt, &s.ExpiresAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNoRecord
		}
		return nil, err
	}

	return s, nil
}

func (m *SnippetModel) Latest() ([]*Snippet, error) {
	query := `
		SELECT id, title, content, created_at, expires_at
		FROM snippets
		WHERE expires_at > NOW()
		ORDER BY created_at DESC
		LIMIT 10
	`
	rows, err := m.DB.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	// Initialize an empty slice to hold the Snippet structs.
	snippets := []*Snippet{}

	for rows.Next() {

		// Create a pointer to a new zeroed Snippet struct.
		s := &Snippet{}

		err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.CreatedAt, &s.ExpiresAt)
		if err != nil {
			return nil, err
		}

		snippets = append(snippets, s)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return snippets, nil
}
