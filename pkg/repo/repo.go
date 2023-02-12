package repo

import (
	"context"
	"gorm.io/gorm"
)

type (
	Sort map[string]string

	CRUD[T any] struct {
		db *gorm.DB
	}
)

func New[T any](db *gorm.DB) *CRUD[T] {
	return &CRUD[T]{db: db}
}

func (c CRUD[T]) Get(ctx context.Context, offset, limit *uint32, sort *Sort) ([]*T, error) {
	var (
		v    []*T
		stmt = c.db.WithContext(ctx)
	)

	if offset != nil {
		stmt = stmt.Offset(int(*offset))
	}

	if limit != nil {
		stmt = stmt.Limit(int(*limit))
	}

	if sort != nil {
		for col, dir := range *sort {
			stmt = stmt.Order(col + " " + dir)
		}
	}

	if err := stmt.Find(&v).Error; err != nil {
		return nil, err
	}

	return v, nil
}

func (c CRUD[T]) Create(ctx context.Context, v T) (*T, error) {
	if err := c.db.WithContext(ctx).Create(&v).Error; err != nil {
		return nil, err
	}

	return &v, nil
}

func (c CRUD[T]) DeleteByID(ctx context.Context, id uint) error {
	var v T
	return c.db.WithContext(ctx).Delete(&v, id).Error
}

func (c CRUD[T]) GetByID(ctx context.Context, id uint) (*T, error) {
	var v T
	if err := c.db.WithContext(ctx).Take(&v, id).Error; err != nil {
		return nil, err
	}

	return &v, nil
}

func (c CRUD[T]) Update(ctx context.Context, id uint, m map[string]any) error {
	var v T
	return c.db.WithContext(ctx).Model(&v).Where("id = ?", id).Updates(m).Error
}

func (c CRUD[T]) Replace(ctx context.Context, id uint, v T) error {
	return c.db.WithContext(ctx).Model(&v).Where("id = ?", id).Select("*").Updates(v).Error
}
