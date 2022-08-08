package dto

import (
	"myzavod/models"
	"myzavod/pkg/tools"
	"strings"
)

type CreateUser struct {
	Email      string `json:"email"`
	Name       string `json:"name"`
	Phone      string `json:"phone"`
	Password   string `json:"password"`
	PayService string `json:"pay_service"`
}

// User returns User model.
// Password will be hashed and email lowered.
func (dto *CreateUser) User() models.User {
	return models.User{
		Email:        strings.ToLower(dto.Email),
		Name:         dto.Name,
		Phone:        dto.Phone,
		PayService:   dto.PayService,
		PasswordHash: tools.PasswordHash(dto.Password),
	}
}
