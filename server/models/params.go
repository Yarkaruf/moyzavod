package models

import "myzavod/pkg/tools"

// Params ...
type Params struct {
	tools.Model

	Params string `json:"params"`
}
