package interceptor

import (
	"github.com/spf13/viper"
)

type AuthInterceptorInterface interface {
	IsAuthenticated() bool
}

type AuthInterceptor struct{}

func NewAuthInterceptor() AuthInterceptorInterface {
	return &AuthInterceptor{}
}

func (u *AuthInterceptor) IsAuthenticated() bool {
	return viper.GetString("consumer_key") != "" && viper.GetString("access_token") != ""
}
