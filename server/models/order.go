package models

import (
	"myzavod/pkg/tools"
	"time"
)

// Order ...
type Order struct {
	tools.Model

	File     string
	FileName string
	Preview  string
	UserID   uint
	Status   int64
	Count    int64 `json:"count"`

	Technology uint    `json:"technology"`
	Plastic    uint    `json:"plastic"`
	Quality    float32 `json:"filling"`

	Cost uint `json:"cost"`

	Color   string    `json:"color"`
	Comment string    `json:"comment"`
	Date    time.Time `json:"date"`
}
