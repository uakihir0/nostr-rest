package repository

import (
	"encoding/json"
	"github.com/nbd-wtf/go-nostr"
	"github.com/uakihir0/nostr-rest/server/domain"
)

type UserEvent struct {
	PubKey      string
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
	About       string `json:"about"`
	Picture     string `json:"picture"`
	Banner      string `json:"banner"`
	Website     string `json:"website"`
}

func (e *UserEvent) ToUser() *domain.User {
	return &domain.User{
		PubKey:      domain.UserPubKey(e.PubKey),
		Name:        e.Name,
		DisplayName: e.DisplayName,
		About:       e.About,
		Picture:     e.Picture,
		Website:     e.Website,
	}
}

func MarshalUserEvent(e *nostr.Event) (*UserEvent, error) {
	event := &UserEvent{}
	if err := json.Unmarshal([]byte(e.Content), event); err != nil {
		return nil, err
	}
	event.PubKey = e.PubKey
	return event, nil
}
