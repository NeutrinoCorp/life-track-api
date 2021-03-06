package category

import (
	"context"

	"github.com/neutrinocorp/life-track-api/internal/domain/aggregate"
	"github.com/neutrinocorp/life-track-api/internal/domain/event"
	"github.com/neutrinocorp/life-track-api/internal/domain/repository"
)

// AddCommand requests an aggregate.Category creation
type AddCommand struct {
	Ctx    context.Context
	UserID string
	Name   string
}

// AddCommandHandler handles AddCommand(s) requests
type AddCommandHandler struct {
	repo repository.Category
	bus  event.Bus
}

// NewAddCommandHandler creates a new AddCommandHandler
func NewAddCommandHandler(r repository.Category, b event.Bus) *AddCommandHandler {
	return &AddCommandHandler{
		repo: r,
		bus:  b,
	}
}

func (h AddCommandHandler) Invoke(cmd AddCommand) (string, error) {
	category, err := aggregate.NewCategory(cmd.UserID, cmd.Name)
	if err != nil {
		return "", err
	} else if err = h.persist(cmd.Ctx, category); err != nil {
		return "", err
	}

	return category.ID(), nil
}

func (h AddCommandHandler) persist(ctx context.Context, category *aggregate.Category) error {
	if err := h.repo.Save(ctx, *category); err != nil {
		return err
	}

	return h.pushEvents(ctx, category)
}

func (h AddCommandHandler) pushEvents(ctx context.Context, category *aggregate.Category) error {
	if err := h.bus.Publish(ctx, category.PullEvents()...); err != nil {
		// rollback
		if errR := h.repo.Remove(ctx, category.ID()); errR != nil {
			return errR
		}
		return err
	}

	return nil
}
