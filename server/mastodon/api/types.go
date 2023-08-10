package mapi

import (
	"github.com/jinzhu/copier"
	"github.com/nbd-wtf/go-nostr/nip19"
	"github.com/samber/lo"
	"github.com/uakihir0/nostr-rest/server/mastodon/domain"
	"github.com/uakihir0/nostr-rest/server/mastodon/openapi"
)

const (
	TimeLayout = "2006-01-02T15:04:05.999Z"
)

type CredentialAccount struct {
	mopenapi.Account
	Role   mopenapi.Role
	Source mopenapi.Source
}

func ToCredentialAccount(
	acc mdomain.Account,
) mopenapi.CredentialAccount {
	account := ToAccount(acc)
	response := mopenapi.CredentialAccount{}

	_ = copier.Copy(&response, &account)

	// Set dummy source data
	response.Source = mopenapi.Source{
		Fields:              []mopenapi.Field{},
		Language:            "",
		Note:                account.Note,
		Privacy:             "public",
		Sensitive:           false,
		FollowRequestsCount: 0,
	}
	// Set dummy role data
	response.Role = mopenapi.Role{
		Id:          0,
		Name:        "guest",
		Color:       "",
		Highlighted: false,
		Permissions: 0,
	}

	return response
}

func ToAccount(
	acc mdomain.Account,
) mopenapi.Account {

	fields := make([]mopenapi.Field, 0)
	fields = append(fields,
		mopenapi.Field{
			Name:       "website",
			Value:      acc.Website,
			VerifiedAt: nil,
		})

	// Lightning Network
	if len(acc.Lud06) > 0 {
		fields = append(fields,
			mopenapi.Field{
				Name:       "lud06",
				Value:      acc.Lud06,
				VerifiedAt: nil,
			})
	}

	// Use the encoded public key hash (npub...)
	acct, _ := nip19.EncodePublicKey(string(acc.ID))

	account := mopenapi.Account{
		Id:           string(acc.ID),
		Acct:         acct,
		Username:     acct,
		Avatar:       acc.Picture,
		AvatarStatic: acc.Picture,
		Header:       acc.Banner,
		HeaderStatic: acc.Banner,
		DisplayName:  acc.DisplayName,
		Note:         acc.About,
		CreatedAt:    acc.CreatedAt.Format(TimeLayout),
		Url:          "https://snort.social/p/" + acct,

		FollowersCount: acc.FollowersCount,
		FollowingCount: acc.FollowingCount,
		StatusesCount:  acc.StatusesCount,

		Bot:          false,
		Group:        false,
		Locked:       false,
		Fields:       fields,
		Emojis:       []mopenapi.CustomEmoji{},
		Discoverable: lo.ToPtr(true),
	}

	if acc.LastStatsAt != nil {
		format := (*acc.LastStatsAt).Format(TimeLayout)
		account.LastStatusAt = &format
	}

	return account
}

func ToStatus(
	st mdomain.Status,
) mopenapi.Status {

	card := &mopenapi.Status_Card{}
	_ = card.FromStatusCard1(nil)

	poll := &mopenapi.Status_Poll{}
	_ = poll.FromStatusPoll1(nil)

	// non-authenticated information
	status := mopenapi.Status{
		Id:        string(st.ID),
		Uri:       "https://",
		Emojis:    []mopenapi.CustomEmoji{},
		Account:   ToAccount(st.Account),
		Content:   st.Text,
		Text:      &st.Text,
		CreatedAt: st.CreatedAt.Format(TimeLayout),
		Card:      *card,
		Poll:      *poll,

		FavouritesCount:  st.FavouritesCount,
		ReblogsCount:     st.ReblogsCount,
		RepliesCount:     0,
		Sensitive:        false,
		MediaAttachments: []mopenapi.MediaAttachment{},
		Mentions:         []mopenapi.StatusMention{},
		Tags:             []mopenapi.StatusTag{},
		Visibility:       "public",
	}

	// authenticated information
	status.Bookmarked = lo.ToPtr(false)
	status.Pinned = lo.ToPtr(false)

	return status
}
