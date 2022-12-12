package usecase

import (
	"context"
	"time"
)

import (
	"github.com/yrcs/nicehouse/pkg/repo"
)

type Usecase interface {
	Create(ctx context.Context, o E) (E, error)
	Get(ctx context.Context, conds ...any) (E, error)
	List(ctx context.Context, conds ...any) ([]E, error)
	ListByPage(ctx context.Context, offset int, limit int, conds map[string]any, orderBy map[string]string) ([]E, int, error)
	Update(ctx context.Context, column string, value any, conds ...any) (E, error)
	Updates(ctx context.Context, values map[string]any) (E, error)
	Delete(ctx context.Context, ids []string, query any, conds ...any) error
}

type BaseDO struct {
	Id        string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type E *any
type T *any

type BaseUsecase[E, T any] struct {
	Repo repo.Repo[E, T]
}

var _ Usecase = (*BaseUsecase[E, T])(nil)

func (c *BaseUsecase[E, T]) Create(ctx context.Context, o E) (E, error) {
	var do E
	if err := c.Repo.Create(ctx, o); err != nil {
		return do, err
	}
	return o, nil
}

func (c *BaseUsecase[E, T]) Get(ctx context.Context, conds ...any) (E, error) {
	return c.Repo.FindOne(ctx, conds)
}

func (c *BaseUsecase[E, T]) List(ctx context.Context, conds ...any) ([]E, error) {
	return c.Repo.Find(ctx, conds)
}

func (c *BaseUsecase[E, T]) ListByPage(ctx context.Context, offset int, limit int, conds map[string]any, orderBy map[string]string) ([]E, int, error) {
	return c.Repo.FindByPage(ctx, offset, limit, conds, orderBy)
}

func (c *BaseUsecase[E, T]) Update(ctx context.Context, column string, value any, conds ...any) (E, error) {
	return c.Repo.Update(ctx, column, value, conds)
}

func (c *BaseUsecase[E, T]) Updates(ctx context.Context, values map[string]any) (E, error) {
	id := values["Id"]
	delete(values, "Id")
	var do E
	o, err := c.Repo.Updates(ctx, values, "id = ?", id)
	if err != nil {
		return do, err
	}
	return o, nil
}

func (c *BaseUsecase[E, T]) Delete(ctx context.Context, ids []string, query any, conds ...any) error {
	return c.Repo.Delete(ctx, ids, query, conds)
}
