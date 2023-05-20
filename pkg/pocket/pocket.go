package pocket

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/url"

	"github.com/skratchdot/open-golang/open"
	"github.com/spf13/viper"
)

type PocketClientInterface interface {
	GetAuthCode(ctx context.Context, consumerKey string) (string, error)
	Authorize(ctx context.Context, code string) (string, error)
	GetAccessToken(ctx context.Context, consumerKey string, code string) (string, error)
	AddArticle(ctx context.Context, URL string) error
}

type PocketClient struct{}

func NewPocketClient() PocketClientInterface {
	return &PocketClient{}
}

func (u *PocketClient) GetAuthCode(ctx context.Context, consumerKey string) (string, error) {

	endpoint := "https://getpocket.com/v3/oauth/request"
	data := []byte(`{"consumer_key": "107363-f2dbdc562815cd3b57ecefb", "redirect_uri": "http://localhost:8989/oauth/idpresponse"}`)

	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(data))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	values, err := url.ParseQuery(string(body))
	if err != nil {
		return "", err
	}

	code := values.Get("code")
	if code == "" {
		return "", errors.New("code is empty")
	}

	return code, nil
}

func (u *PocketClient) Authorize(ctx context.Context, code string) (string, error) {

	l, err := net.Listen("tcp", "localhost:8989")
	if err != nil {
		log.Fatalln(err)
	}
	defer l.Close()

	if err := open.Start(
		fmt.Sprintf("https://getpocket.com/auth/authorize?request_token=%s&redirect_uri=http://localhost:8989/oauth/idpresponse", code),
	); err != nil {
		return "", err
	}

	quit := make(chan string)
	go http.Serve(l, http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if req.URL.Path == "/" {
			w.Write([]byte(`<script>location.href = "/close?" + location.hash.substring(1);</script>`))
		} else {
			w.Write([]byte(`<script>window.open("about:blank","_self").close()</script>`))
			w.(http.Flusher).Flush()
			quit <- "quit"
		}
	}))

	return <-quit, nil
}

func (u *PocketClient) GetAccessToken(ctx context.Context, consumerKey, code string) (string, error) {

	endpoint := "https://getpocket.com/v3/oauth/authorize"
	data := []byte(fmt.Sprintf(`{"consumer_key": "107363-f2dbdc562815cd3b57ecefb", "code": "%s"}`, code))

	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(data))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	values, err := url.ParseQuery(string(body))
	if err != nil {
		return "", err
	}

	accessToken := values.Get("access_token")
	if code == "" {
		return "", errors.New("access_token is empty")
	}

	return accessToken, nil
}

func (u *PocketClient) AddArticle(ctx context.Context, URL string) error {

	endpoint := "https://getpocket.com/v3/add"

	param := struct {
		URL         string `json:"url"`
		ConsumerKey string `json:"consumer_key"`
		AccessToken string `json:"access_token"`
	}{
		URL:         URL,
		ConsumerKey: viper.GetString("consumer_key"),
		AccessToken: viper.GetString("access_token"),
	}

	data, err := json.Marshal(param)
	if err != nil {
		return err
	}

	fmt.Println(string(data))

	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("pocket api response status code: %d", resp.StatusCode)
	}

	return nil
}
