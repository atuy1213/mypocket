package usecase

import (
	"context"

	"github.com/atuy1213/mypocket/pkg/pocket"
)

type AddExecutor struct {
	pocket pocket.PocketClientInterface
}

func NewAddExecutor(pocket pocket.PocketClientInterface) *AddExecutor {
	return &AddExecutor{
		pocket: pocket,
	}
}

func (u *AddExecutor) Add(ctx context.Context, URL string) error {
	if err := u.pocket.AddArticle(ctx, URL); err != nil {
		return err
	}
	return nil
}
