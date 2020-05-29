package clientmodels

import (
	"time"
)

// FTMember describes a member
type FTMember struct {
	ID        int
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// PTMember describes a pt member
type PTMember struct {
	ID        int
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type VoteResult struct {
	ID        int
	MemberID  int
	Vote      int
	CreatedAt time.Time
	UpdatedAt time.Time
}
