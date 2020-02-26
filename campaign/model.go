package campaign

import "time"

// Campaign - model - entity - schema of the domain
type Campaign struct {
	ID        string    `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	Code      string    `json:"code" db:"code"`
	Method    string    `json:"method" db:"method"`
	Start     string    `json:"start" db:"start"`
	End       string    `json:"end" db:"end"`
	CreatedAt time.Time `json:"createdAt" db:"createAt"`
	UpdatedAt time.Time `json:"updatedAt" db:"updatedAt"`
}
