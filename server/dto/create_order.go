package dto

import (
	"myzavod/models"
)

// CreateOrder ...
type CreateOrder struct {
	Technology uint    `json:"technology"`
	Plastic    uint    `json:"plastic"`
	Quality    float32 `json:"filling"`
	Color      string  `json:"color"`
	Comment    string  `json:"comment"`
}

// Order returns Order model.
// Password will be hashed and email lowered.
func (dto *CreateOrder) Order() models.Order {
	return models.Order{
		Technology: dto.Technology,
		Plastic:    dto.Plastic,
		Quality:    dto.Quality,
		Color:      dto.Color,
		Comment:    dto.Comment,
	}
}
