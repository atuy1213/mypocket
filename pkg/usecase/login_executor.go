package usecase

import (
	"bufio"
	"context"
	"fmt"
	"os"

	"github.com/atuy1213/mypocket/pkg/pocket"
	"github.com/spf13/viper"
)

type LoginExecutor struct {
	pocket pocket.PocketClientInterface
}

func NewLoginExecutor(pocket pocket.PocketClientInterface) *LoginExecutor {
	return &LoginExecutor{
		pocket: pocket,
	}
}

func (u *LoginExecutor) Login(ctx context.Context) error {
	fmt.Print("CONSUMER KEY > ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	consumerKey := scanner.Text()

	code, err := u.pocket.GetAuthCode(ctx, consumerKey)
	if err != nil {
		return err
	}

	if _, err := u.pocket.Authorize(ctx, code); err != nil {
		return err
	}

	accessToken, err := u.pocket.GetAccessToken(ctx, consumerKey, code)
	if err != nil {
		return err
	}

	viper.Set("consumer_key", consumerKey)
	viper.Set("access_token", accessToken)
	if err := viper.WriteConfig(); err != nil {
		return err
	}
	return nil
}
