package gorm

import (
	"context"
	"github.com/nullc4t/crud-rest-api/pkg/repo"
	"gorm.io/gorm"
)

type impl[T any] struct {
	db *gorm.DB
}

func New[T any](db *gorm.DB) repo.Interface[T] {
	return impl[T]{db: db}
}

func (i impl[T]) Atomic(fn func(repo repo.Interface[T]) error) error {
	return i.db.Transaction(func(tx *gorm.DB) error {
		return fn(impl[T]{db: tx})
	})
}

func (i impl[T]) Get(ctx context.Context, offset, limit *uint32, sort *repo.Sort) ([]*T, error) {
	return Get[T](i.db.WithContext(ctx), offset, limit, sort)
}

func (i impl[T]) Create(ctx context.Context, v T) (*T, error) {
	return Create[T](i.db.WithContext(ctx), v)
}

func (i impl[T]) DeleteByID(ctx context.Context, id uint) error {
	return DeleteByID[T](i.db.WithContext(ctx), id)
}

func (i impl[T]) GetByID(ctx context.Context, id uint) (*T, error) {
	return GetByID[T](i.db.WithContext(ctx), id)
}

func (i impl[T]) Update(ctx context.Context, id uint, m map[string]any) error {
	return Update[T](i.db.WithContext(ctx), id, m)
}

func (i impl[T]) Replace(ctx context.Context, id uint, v T) error {
	return Replace[T](i.db.WithContext(ctx), id, v)
}

var _ repo.Interface[int] = (*impl[int])(nil)
