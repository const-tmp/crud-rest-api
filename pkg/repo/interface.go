package repo

import "context"

type (
	Sort map[string]string

	Interface[T any] interface {
		Atomic(fn func(repo Interface[T]) error) error
		Get(ctx context.Context, offset, limit *uint32, sort *Sort) ([]*T, error)
		Create(ctx context.Context, v T) (*T, error)
		DeleteByID(ctx context.Context, id uint) error
		GetByID(ctx context.Context, id uint) (*T, error)
		Update(ctx context.Context, id uint, m map[string]any) error
		Replace(ctx context.Context, id uint, v T) error
	}
)
