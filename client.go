package trailer

import (
	"bufio"
	"encoding/json"
	"log"
	"net/http"
	"net/url"

	"github.com/garyburd/go-oauth/oauth"
)

type messageType int

type message struct {
	raw         []byte
	messageType messageType
}

const (
	onStatus messageType = 1 << iota
	onDelete messageType = 1 << iota
	onEvent  messageType = 1 << iota
)

var (
	messageTypeTable = map[string]messageType{
		"in_reply_to_status_id": onStatus,
		"delete":                onDelete,
		"event":                 onEvent,
	}
)

type Client struct {
	oauth      oauth.Client
	credential *oauth.Credentials

	messages chan *message
}

func NewClient(token, secret string) *Client {
	return &Client{
		oauth: oauth.Client{
			TemporaryCredentialRequestURI: "https://api.twitter.com/oauth/request_token",
			ResourceOwnerAuthorizationURI: "https://api.twitter.com/oauth/authorize",
			TokenRequestURI:               "https://api.twitter.com/oauth/access_token",
			Credentials: oauth.Credentials{
				Token:  token,
				Secret: secret,
			},
		},

		messages: make(chan *message, 20),
	}
}

func (client *Client) SetAccessToken(token, secret string) {
	client.credential = &oauth.Credentials{
		Token:  token,
		Secret: secret,
	}
}

func (client *Client) onStatus(handler Handler, raw []byte) (err error) {
	status := new(Tweet)
	if err = json.Unmarshal(raw, status); err != nil {
		return
	}

	handler.OnStatus(status)
	return
}

func (client *Client) onDelete(handler Handler, raw []byte) (err error) {
	event := new(DeleteEvent)
	if err = json.Unmarshal(raw, event); err != nil {
		return
	}

	handler.OnDelete(event)
	return
}

func (client *Client) onEvent(handler Handler, raw []byte) (err error) {
	event := new(Event)
	if err = json.Unmarshal(raw, event); err != nil {
		return
	}

	handler.OnEvent(event)
	return
}

func (client *Client) messageListener(handler Handler) {
	for message := range client.messages {
		switch message.messageType {
		case onStatus:
			client.onStatus(handler, message.raw)
		case onDelete:
			client.onDelete(handler, message.raw)
		case onEvent:
			client.onEvent(handler, message.raw)
		}
	}
}

func (client *Client) connect(method, endpoint string, form url.Values) (resp *http.Response, err error) {
	switch method {
	case "GET":
		resp, err = client.oauth.Get(
			http.DefaultClient,
			client.credential,
			endpoint,
			form)
	case "POST":
		resp, err = client.oauth.Post(
			http.DefaultClient,
			client.credential,
			endpoint,
			form)
	default:
		err = Errorf("Unsupported HTTP Method:", method)
	}

	if err == nil && resp.StatusCode != 200 {
		resp.Body.Close()
		err = Errorf(resp.Status)
	}

	return
}

func (client *Client) handleStream(resp *http.Response, handler Handler, numGoroutine int) (err error) {
	defer resp.Body.Close()

	for i := 0; i < numGoroutine; i++ {
		go client.messageListener(handler)
	}

	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		raw := scanner.Bytes()
		marshaled := make(map[string]interface{})
		if err := json.Unmarshal(raw, &marshaled); err != nil {
			log.Println(err)
			continue
		}

		for key := range messageTypeTable {
			if _, ok := marshaled[key]; ok {
				client.messages <- &message{
					raw:         raw,
					messageType: messageTypeTable[key],
				}
				break
			}
		}
	}

	return scanner.Err()
}

func (client *Client) UserStream(handler Handler, numGoroutine int) (err error) {
	const (
		METHOD   = "GET"
		ENDPOINT = "https://userstream.twitter.com/1.1/user.json"
	)

	resp, err := client.connect(METHOD, ENDPOINT, nil)
	if err != nil {
		return err
	}

	return client.handleStream(resp, handler, numGoroutine)
}

func (client *Client) Sample(handler Handler, numGoroutine int) (err error) {
	const (
		METHOD   = "GET"
		ENDPOINT = "https://stream.twitter.com/1.1/statuses/sample.json"
	)

	resp, err := client.connect(METHOD, ENDPOINT, nil)
	if err != nil {
		return
	}

	return client.handleStream(resp, handler, numGoroutine)
}
