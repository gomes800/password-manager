package repository

import (
	"context"
	"database/sql"

	"github.com/gomes800/password-manager/internal/model"
)

type CredentialRepository interface {
	CreateTable(ctx context.Context) error
	Save(ctx context.Context, c *model.Credential) error
	GetByID(ctx context.Context, id string) (model.Credential, error)
}

type sqliteCredentialRepo struct {
	db *sql.DB
}

func NewCredentialRepository(db *sql.DB) CredentialRepository {
	return &sqliteCredentialRepo{db: db}
}

func (r *sqliteCredentialRepo) CreateTable(ctx context.Context) error {
	query := `
    CREATE TABLE IF NOT EXISTS credentials (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        user_id INTEGER NOT NULL,
        service_name TEXT NOT NULL,
        username TEXT NOT NULL,
        ciphertext BLOB NOT NULL,
        nonce BLOB NOT NULL,
        salt BLOB NOT NULL
    );`
	_, err := r.db.ExecContext(ctx, query)
	return err
}

func (r *sqliteCredentialRepo) Save(ctx context.Context, c *model.Credential) error {
	query := "INSERT INTO credentials (user_id, service_name, username, ciphertext, nonce, salt) VALUES (?, ?, ?, ?, ?, ?) RETURNING id;"
	return r.db.QueryRowContext(ctx, query, c.UserId, c.ServiceName, c.Username, c.Ciphertext, c.Nonce, c.Salt).Scan(&c.ID)
}

func (r *sqliteCredentialRepo) GetByID(ctx context.Context, id string) (model.Credential, error) {
	var c model.Credential
	query := "SELECT id, user_id, service_name, username, ciphertext, nonce, salt FROM credentials WHERE id = ?"
	err := r.db.QueryRowContext(ctx, query, id).Scan(&c.ID, &c.UserId, &c.ServiceName, &c.Username, &c.Ciphertext, &c.Nonce, &c.Salt)
	return c, err
}
