package common

import (
	"gorm.io/gorm"
	"time"
)

type Model struct {
	Id        *uint64    `json:"id,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

func EncodeModel(from gorm.Model) Model {
	id := uint64(from.ID)
	return Model{
		Id:        &id,
		CreatedAt: &from.CreatedAt,
		UpdatedAt: &from.UpdatedAt,
		DeletedAt: &from.DeletedAt.Time,
	}
}

func DecodeModel(from Model) gorm.Model {
	var (
		id        uint64
		createdAt time.Time
		updatedAt time.Time
		deletedAt gorm.DeletedAt
	)

	if from.Id != nil {
		id = *from.Id
	}
	if from.CreatedAt != nil {
		createdAt = *from.CreatedAt
	}
	if from.UpdatedAt != nil {
		updatedAt = *from.UpdatedAt
	}
	if from.DeletedAt != nil {
		deletedAt = gorm.DeletedAt{Time: *from.DeletedAt, Valid: true}
	}

	return gorm.Model{
		ID:        uint(id),
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		DeletedAt: deletedAt,
	}
}
