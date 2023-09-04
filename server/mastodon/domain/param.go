package mdomain

import (
	"time"

	"github.com/samber/lo"
)

type TimelineOptions struct {
	MaxId   *StatusID
	SinceId *StatusID
	MinId   *StatusID
	Limit   *int

	OnlyMedia      *bool
	ExcludeReplies *bool
	ExcludeReblogs *bool
	Pinned         *bool

	Tagged *string
}

func (o TimelineOptions) GetLimit(def int) int {
	if o.Limit != nil {
		return *o.Limit
	}
	return def
}

func (o TimelineOptions) GetSinceTime() *time.Time {
	if o.SinceId != nil {
		date, err := o.SinceId.ToDate()
		if err != nil {
			return nil
		}
		// Need to add 1 millis convert Max to Until
		return lo.ToPtr(time.UnixMilli(date.UnixMilli() + 1))
	}
	return nil
}

func (o TimelineOptions) GetUntilTime() *time.Time {
	if o.MaxId != nil {
		date, err := o.MaxId.ToDate()
		if err != nil {
			return nil
		}
		// Need to sub 1 millis convert Max to Until
		return lo.ToPtr(time.UnixMilli(date.UnixMilli() - 1))
	}
	return nil
}
