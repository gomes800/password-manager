package service

import (
	"context"

	"github.com/gomes800/password-manager/internal/model"
	"github.com/gomes800/password-manager/internal/repository"
)

type CredentialService struct {
	repo repository.CredentialRepository // Depende da Interface
}

func NewCredentialService(repo repository.CredentialRepository) *CredentialService {
	return &CredentialService{repo: repo}
}

func (s *CredentialService) CreateCredential(ctx context.Context, c *model.Credential) error {
	return s.repo.Save(ctx, c)
}

func (s *CredentialService) GetCredential(ctx context.Context, id string) (model.Credential, error) {
	return s.repo.GetByID(ctx, id)
}
