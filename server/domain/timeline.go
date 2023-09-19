package domain

type Timeline struct {
	Post   *Post
	Repost *Repost
}

type TimelineRepository interface {

	// GetTimelines
	GetTimelines(
		pks []UserPubKey,
		options PagingOptions,
	) ([]Timeline, error)
}
