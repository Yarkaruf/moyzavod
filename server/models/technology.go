package models

import "myzavod/pkg/tools"

// Service ...
type Service struct {
	tools.Model

	Cost float32 `json:"cost"`
}

// Technology ...
type Technology struct {
	tools.Model

	PlasticVariations []Variation `json:"plastic_variations"`
	Name              string      `json:"name"`
	Cost              float32     `json:"cost"`
}

// Variation ...
type Variation struct {
	tools.Model

	TechnologyID uint    `json:"technology_id"`
	Name         string  `json:"name"`
	Cost         float32 `json:"cost"`
}
