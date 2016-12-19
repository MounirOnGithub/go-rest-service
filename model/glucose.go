package model

import "time"

type Glucose struct {
	Result    float64
	CreatedAt time.Time
	IsHyper   bool
	IsHypo    bool
}
