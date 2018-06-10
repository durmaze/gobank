package gobank

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/parnurzeal/gorequest"
)

type Client struct {
	mountebankURI string
	impostersURI  string
}

func NewClient(mountebankURI string) *Client {
	return &Client{
		mountebankURI: mountebankURI,
		impostersURI:  mountebankURI + "/imposters",
	}
}

func (c *Client) CreateImposter(imposter ImposterElement) (map[string]interface{}, error) {
	resp, body, errs := gorequest.New().Post(c.impostersURI).Send(imposter).EndBytes()

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
	imposterURI := c.impostersURI + "/" + strconv.Itoa(port)

	resp, body, errs := gorequest.New().Delete(imposterURI).EndBytes()

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
	resp, body, errs := gorequest.New().Delete(c.impostersURI).EndBytes()

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