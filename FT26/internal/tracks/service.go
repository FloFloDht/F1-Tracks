package tracks

import (
	"FT26/internal/tracks/models"
	"context"
	"errors"
)

// Service définit les méthodes disponibles pour manipuler les circuits
type Service interface {
	GetAll(ctx context.Context) ([]*models.Track, error)
	GetByID(ctx context.Context, id int) (*models.Track, error)
	Create(ctx context.Context, track *models.Track) error
	Update(ctx context.Context, id int, input *models.Track) error
	Delete(ctx context.Context, id int) error
}

// service implémente l'interface Service
type service struct {
	repo Repository
}

// NewService crée un nouveau service
func NewService(repo Repository) Service {
	return &service{repo: repo}
}

// GetAll retourne tous les circuits
func (s *service) GetAll(ctx context.Context) ([]*models.Track, error) {
	return s.repo.GetAll(ctx)
}

// GetByID retourne un circuit par son ID
func (s *service) GetByID(ctx context.Context, id int) (*models.Track, error) {
	return s.repo.GetByID(ctx, id)
}

// Create ajoute un nouveau circuit
func (s *service) Create(ctx context.Context, track *models.Track) error {
	if track.Name == "" {
		return errors.New("track name cannot be empty")
	}
	return s.repo.Create(ctx, track)
}

// Update modifie un circuit existant
func (s *service) Update(ctx context.Context, id int, input *models.Track) error {
	existing, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	if input.Name != "" {
		existing.Name = input.Name
	}
	if input.Country != "" {
		existing.Country = input.Country
	}
	if input.LengthKm != 0 {
		existing.LengthKm = input.LengthKm
	}

	return s.repo.Update(ctx, existing)
}

// Delete supprime un circuit
func (s *service) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}
