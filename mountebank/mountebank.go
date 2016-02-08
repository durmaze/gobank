package mountebank

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	. "github.com/durmaze/gobank/builders"
	"github.com/parnurzeal/gorequest"
)

type Client struct {
	mountebankUri string
	impostersUri  string
}

func NewClient(mountebankUri string) *Client {
	return &Client{
		mountebankUri: mountebankUri,
		impostersUri:  mountebankUri + "/imposters",
	}
}

func (c *Client) CreateImposter(imposter Imposter) (map[string]interface{}, error) {
	resp, body, errs := gorequest.New().Post(c.impostersUri).Send(imposter).EndBytes()

	if len(errs) > 0 {
		return nil, errs[0]
	}

	if resp.StatusCode == http.StatusCreated {
		var created map[string]interface{}
		json.Unmarshal(body, &created)

		return created, nil
	}

	return nil, errors.New("Cannot create the imposter")
}

func (c *Client) DeleteImposter(port int) (map[string]interface{}, error) {
	imposterUri := c.impostersUri + "/" + strconv.Itoa(port)

	resp, body, errs := gorequest.New().Delete(imposterUri).EndBytes()

	if len(errs) > 0 {
		return nil, errs[0]
	}

	if resp.StatusCode == http.StatusOK {
		var deleted map[string]interface{}
		json.Unmarshal(body, &deleted)

		return deleted, nil
	}

	return nil, errors.New("Cannot delete the imposter")
}

func (c *Client) DeleteAllImposters() (map[string]interface{}, error) {
	resp, body, errs := gorequest.New().Delete(c.impostersUri).EndBytes()

	if len(errs) > 0 {
		return nil, errs[0]
	}

	if resp.StatusCode == http.StatusOK {
		var allDeleted map[string]interface{}
		json.Unmarshal(body, &allDeleted)

		return allDeleted, nil
	}

	return nil, errors.New("Cannot delete all of the imposters")
}
