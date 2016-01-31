package mountebank

import (
"strconv"
	. "github.com/durmaze/gobank/builders"
	"github.com/parnurzeal/gorequest"
)

type Client struct{
	mountebankUri string
	impostersUri string
}

func NewClient(mountebankUri string) *Client {
	return &Client{
		mountebankUri: mountebankUri,
		impostersUri: mountebankUri + "/imposters",
	}
}

func (c *Client) CreateImposter(imposter Imposter) {
	gorequest.New().Post(c.impostersUri).Send(imposter).End()
}

func (c *Client) DeleteImposter(imposter Imposter) {
	imposterUri := c.impostersUri + "/" + strconv.Itoa(imposter.Port)

	gorequest.New().Delete(imposterUri).End()
}

func (c *Client) DeleteAllImposters() {
	gorequest.New().Delete(c.impostersUri).End()
}