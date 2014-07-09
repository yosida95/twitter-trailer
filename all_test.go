package trailer_test

import (
	. "github.com/otiai10/mint"
	"github.com/yosida95/twitter-trailer"
	"testing"
)

const (
	TOKEN  = "YOUR_ACCESS_TOKEN"
	SECRET = "YOUR_ACCESS_TOKEN_SECRET"
)

func TestNewClient(t *testing.T) {
	client := trailer.NewClient(TOKEN, SECRET)
	Expect(t, client).TypeOf("*trailer.Client")
}
