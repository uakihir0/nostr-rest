package mdomain

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
