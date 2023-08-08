package mdomain

import "time"

type AccountID string

type Account struct {
	ID          AccountID
	Name        string
	DisplayName string
	Picture     string
	Banner      string
	Website     string
	About       string
	Lud06       string
	CreatedAt   time.Time
	LastStatsAt *time.Time

	StatusesCount  int
	FollowersCount int
	FollowingCount int
}
