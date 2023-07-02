package mapi

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/uakihir0/nostr-rest/server/domain"
	minjection "github.com/uakihir0/nostr-rest/server/mastodon/injection"
	"net/http"
)

// GetMApiV1AccountsVerifyCredentials
func (h *MastodonHandler) GetApiV1AccountsVerifyCredentials(c echo.Context) error {
	accountService := minjection.AccountService()

	pk := c.(*domain.Context).PubKey
	if pk == nil {
		return errors.New("needs user public key")
	}

	account, err := accountService.GetAccount(*pk)
	if err != nil {
		return err
	}

	return c.JSON(
		http.StatusOK,
		ToCredentialAccount(*account),
	)
}

// GetMApiV1AccountsId
func (h *MastodonHandler) GetApiV1AccountsUid(c echo.Context, uid string) error {
	accountService := minjection.AccountService()

	pk, err := domain.ToUserPubKey(uid)
	if err != nil {
		return err
	}

	account, err := accountService.GetAccount(*pk)
	if err != nil {
		return err
	}

	return c.JSON(
		http.StatusOK,
		ToCredentialAccount(*account),
	)
}
