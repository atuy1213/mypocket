package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"os"

	"github.com/skratchdot/open-golang/open"
)

func main() {
	Login()
}

func Login() {

	fmt.Print("please type consumer_key and enter > ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	consumerKey := scanner.Text()
	fmt.Println(consumerKey)

	// code, err := getAuthCode()
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// if _, err := getAuthorize(code); err != nil {
	// 	log.Fatalln(err)
	// }

	// accessToken, err := getAccessToken(code)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// log.Println(accessToken)
}

func getAuthorize(code string) (string, error) {

	l, err := net.Listen("tcp", "localhost:8989")
	if err != nil {
		return "", err
	}
	defer l.Close()

	err = open.Start(fmt.Sprintf("https://getpocket.com/auth/authorize?request_token=%s&redirect_uri=http://localhost:8989/oauth/idpresponse", code))
	if err != nil {
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

func getAuthCode() (string, error) {

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

func getAccessToken(code string) (string, error) {

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
