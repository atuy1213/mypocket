package usecase

import (
	"context"
	"errors"

	"github.com/atuy1213/mypocket/pkg/interceptor"
	"github.com/atuy1213/mypocket/pkg/pocket"
)

type AddExecutor struct {
	pocket          pocket.PocketClientInterface
	authInterceptor interceptor.AuthInterceptorInterface
}

func NewAddExecutor(
	pocket pocket.PocketClientInterface,
	authInterceptor interceptor.AuthInterceptorInterface,
) *AddExecutor {
	return &AddExecutor{
		pocket:          pocket,
		authInterceptor: authInterceptor,
	}
}

func (u *AddExecutor) Add(ctx context.Context, URL string) error {

	if !u.authInterceptor.IsAuthenticated() {
		return errors.New("not authorized, execute command `mypocket login` and login your pocket")
	}

	if err := u.pocket.AddArticle(ctx, URL); err != nil {
		return err
	}
	return nil
}
