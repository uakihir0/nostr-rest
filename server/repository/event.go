package repository

import (
	"encoding/json"
	"github.com/nbd-wtf/go-nostr"
	"github.com/samber/lo"
	"github.com/uakihir0/nostr-rest/server/domain"
	"time"
)

type Event struct {
	ID        string
	PubKey    string
	Content   string
	CreatedAt int64
}

// --------------------------------------------------------------------- //
// USER
// --------------------------------------------------------------------- //

type UserEvent struct {
	Event
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
	About       string `json:"about"`
	Picture     string `json:"picture"`
	Banner      string `json:"banner"`
	Website     string `json:"website"`
	Lud06       string `json:"lud06"`
}

// lud16
// nip05

func (e *UserEvent) ToUser() *domain.User {
	return &domain.User{
		PubKey:      domain.UserPubKey(e.PubKey),
		Name:        e.Name,
		DisplayName: e.DisplayName,
		About:       e.About,
		Picture:     e.Picture,
		Banner:      e.Banner,
		Website:     e.Website,
		Lud06:       e.Lud06,
		CreatedAt:   e.CreatedAt,
	}
}

func MarshalUserEvent(e *nostr.Event) (*UserEvent, error) {
	event := &UserEvent{}
	if err := json.Unmarshal([]byte(e.Content), event); err != nil {
		return nil, err
	}
	event.CreatedAt = int64(e.CreatedAt)
	event.PubKey = e.PubKey
	return event, nil
}

// --------------------------------------------------------------------- //
// POST
// --------------------------------------------------------------------- //

type PostEvent struct {
	Event
	RootPostID           *string
	ReplyPostID          *string
	MentionedUserPubKeys []string
}

func (e *PostEvent) ToPost() domain.Post {

	var rootPostID *domain.PostID = nil
	if e.RootPostID != nil {
		rootPostID = lo.ToPtr(domain.PostID(*e.RootPostID))
	}
	var replyPostID *domain.PostID = nil
	if e.ReplyPostID != nil {
		replyPostID = lo.ToPtr(domain.PostID(*e.ReplyPostID))
	}

	return domain.Post{
		ID:         domain.PostID(e.ID),
		UserPubKey: domain.UserPubKey(e.PubKey),
		Content:    e.Content,
		CreatedAt:  time.Unix(e.CreatedAt, 0),

		RootPostID:  rootPostID,
		ReplyPostID: replyPostID,
		MentionedUserPubKeys: lo.Map(e.MentionedUserPubKeys,
			func(item string, index int) domain.UserPubKey {
				return domain.UserPubKey(item)
			}),
	}
}

func MarshalPostEvent(e *nostr.Event) (*PostEvent, error) {
	event := &PostEvent{}
	event.CreatedAt = int64(e.CreatedAt)
	event.PubKey = e.PubKey
	event.Content = e.Content
	event.ID = e.ID

	root := FilterTagWithCond(e.Tags,
		map[int]string{0: "e", 3: "root"})
	if len(root) > 0 {
		event.RootPostID = &root[0][1]
	}

	replay := FilterTagWithCond(e.Tags,
		map[int]string{0: "e", 3: "replay"})
	if len(replay) > 0 {
		event.ReplyPostID = &replay[0][1]
	}

	event.MentionedUserPubKeys = ExtTagValue(
		FilterTag(e, "p"),
	)

	return event, nil
}

// --------------------------------------------------------------------- //
// REACTION
// --------------------------------------------------------------------- //

type ReactionEvent struct {
	Event
	ToPost []string
	ToUser []string
}

func (e *ReactionEvent) ToReaction() domain.Reaction {
	return domain.Reaction{
		ID:         domain.ReactionID(e.ID),
		UserPubKey: domain.UserPubKey(e.PubKey),
		Content:    e.Content,
		CreatedAt:  time.Unix(e.CreatedAt, 0),
	}
}

func MarshalReactionEvent(e *nostr.Event) (*ReactionEvent, error) {
	event := &ReactionEvent{}
	event.CreatedAt = int64(e.CreatedAt)
	event.PubKey = e.PubKey
	event.Content = e.Content
	event.ID = e.ID
	event.ToPost = ExtTagValue(FilterTag(e, "e"))
	event.ToUser = ExtTagValue(FilterTag(e, "p"))
	return event, nil
}

// --------------------------------------------------------------------- //
// REPOST
// --------------------------------------------------------------------- //

type RepostEvent struct {
	Event
	ToPost []string
	ToUser []string
}

func (e *RepostEvent) ToRepost() domain.Repost {
	return domain.Repost{
		ID:         domain.RepostID(e.ID),
		UserPubKey: domain.UserPubKey(e.PubKey),
		Content:    e.Content,
		CreatedAt:  time.Unix(e.CreatedAt, 0),
	}
}

func MarshalRepostEvent(e *nostr.Event) (*RepostEvent, error) {
	event := &RepostEvent{}
	event.CreatedAt = int64(e.CreatedAt)
	event.PubKey = e.PubKey
	event.Content = e.Content
	event.ID = e.ID
	event.ToPost = ExtTagValue(FilterTag(e, "e"))
	event.ToUser = ExtTagValue(FilterTag(e, "p"))
	return event, nil
}

// --------------------------------------------------------------------- //
// OTHER
// --------------------------------------------------------------------- //

func FilterTag(
	e *nostr.Event,
	tag string,
) nostr.Tags {
	return lo.Filter(e.Tags,
		func(t nostr.Tag, _ int) bool {
			return t.Key() == tag
		})
}

func ExtTagValue(
	tags nostr.Tags,
) []string {
	return lo.Map(tags,
		func(t nostr.Tag, _ int) string {
			return t.Value()
		})
}

// ExtractTagWithCond
func FilterTagWithCond(
	tags nostr.Tags,
	conditions map[int]string,
) nostr.Tags {
	return lo.Filter(tags,
		func(t nostr.Tag, _ int) bool {
			for _, i := range lo.Keys(conditions) {
				if len(t) > i && t[i] != conditions[i] {
					return false
				}
			}
			return true
		})
}
