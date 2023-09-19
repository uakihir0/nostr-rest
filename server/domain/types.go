package domain

import "time"

type PagingOptions struct {
	MaxResults int
	SinceTime  *time.Time
	UntilTime  *time.Time
}


