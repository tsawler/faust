package clientmodels

import (
	"time"
)

// FTMember describes a member
type FTMember struct {
	ID        int
	Email     string
	Voted     int
	CreatedAt time.Time
	UpdatedAt time.Time
}

// PTMember describes a pt member
type PTMember struct {
	ID        int
	Email     string
	Voted     int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type VoteTotals struct {
	ID        int
	Yes       int
	No        int
	CreatedAt time.Time
	UpdatedAt time.Time
}
