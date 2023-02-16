package gorm

import (
	"github.com/nullc4t/crud-rest-api/pkg/repo"
	"gorm.io/gorm"
)

func Get[T any](db *gorm.DB, offset, limit *uint32, sort *repo.Sort) ([]*T, error) {
	var (
		v    []*T
		stmt = db
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

func Create[T any](db *gorm.DB, v T) (*T, error) {
	if err := db.Create(&v).Error; err != nil {
		return nil, err
	}

	return &v, nil
}

func DeleteByID[T any](db *gorm.DB, id uint) error {
	var v T
	return db.Delete(&v, id).Error
}

func GetByID[T any](db *gorm.DB, id uint) (*T, error) {
	var v T
	if err := db.Take(&v, id).Error; err != nil {
		return nil, err
	}

	return &v, nil
}

func Update[T any](db *gorm.DB, id uint, m map[string]any) error {
	var v T
	return db.Model(&v).Where("id = ?", id).Updates(m).Error
}

func Replace[T any](db *gorm.DB, id uint, v T) error {
	return db.Model(&v).Where("id = ?", id).Select("*").Updates(v).Error
}
