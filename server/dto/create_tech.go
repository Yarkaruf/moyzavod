package dto

import "myzavod/models"

// CreateTech ...
type CreateTech struct {
	Name string  `json:"name"`
	Cost float32 `json:"cost"`
}

// Technology ...
func (d CreateTech) Technology() models.Technology {
	return models.Technology{
		Name: d.Name,
		Cost: d.Cost,
	}
}

// CreateVariant ...
type CreateVariant struct {
	TechnologyID uint    `json:"technology_id"`
	Name         string  `json:"name"`
	Cost         float32 `json:"cost"`
}

// Variation ...
func (d CreateVariant) Variation() models.Variation {
	return models.Variation{
		TechnologyID: d.TechnologyID,
		Name:         d.Name,
		Cost:         d.Cost,
	}
}
