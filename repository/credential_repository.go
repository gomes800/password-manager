package repository

import (
	"context"
	"database/sql"

	"github.com/gomes800/password-manager/model"
)

type CredentialRepository struct {
	db *sql.DB
}

func NewCredentialRepository(db *sql.DB) *CredentialRepository {
	return &CredentialRepository{db: db}
}

func (r *CredentialRepository) CreateTable(ctx context.Context) error {
	query := `
	CREATE TABLE IF NOT EXISTS credentials (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	service_name TEXT NOT NUll,
	username TEXT NOT NULL,
	ciphertext BLOB NOT NULL,
	nonce BLOB NOT NULL,
	salt BLOB NOT NULL
	);`
	_, err := r.db.ExecContext(ctx, query)
	return err
}

func (r *CredentialRepository) Save(ctx context.Context, c *model.Credential) error {
	query := "INSERT INTO credentials (service_name, username, ciphertext, nonce, salt) VALUES (?, ?, ?, ?, ?) RETURNING id;"
	return r.db.QueryRowContext(ctx, query, c.ServiceName, c.Username, c.Ciphertext, c.Nonce, c.Salt).Scan(&c.ID)
}

func (r *CredentialRepository) GetByID(ctx context.Context, id string) (model.Credential, error) {
	var c model.Credential
	query := "SELECT id, service_name, username, ciphertext, nonce, salt FROM credentials WHERE id = ?"
	err := r.db.QueryRowContext(ctx, query, id).Scan(&c.ID, &c.ServiceName, &c.Username, &c.Ciphertext, &c.Nonce, &c.Salt)

	return c, err
}
