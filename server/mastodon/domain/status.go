package mdomain

import (
	"errors"
	"fmt"
	"github.com/samber/lo"
	"strconv"
	"strings"
	"time"
)

type StatusID string

func (i StatusID) ToDate() (*time.Time, error) {
	elements := strings.Split(string(i), ".")
	msec, err := strconv.Atoi(elements[1])
	if err != nil {
		return nil, errors.New("illegal status id format")
	}
	return lo.ToPtr(time.UnixMilli(int64(msec))), nil
}

func (i StatusID) ToNostrID() string {
	return strings.Split(string(i), ".")[0]
}

func NewStatusID(id string, createdAt time.Time) StatusID {
	return StatusID(fmt.Sprintf("%s.%d", id, createdAt.UnixMilli()))
}

type Status struct {
	ID              StatusID
	Text            string
	Account         Account
	CreatedAt       time.Time
	FavouritesCount int
	ReblogsCount    int
}
