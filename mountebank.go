package gobank

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/parnurzeal/gorequest"
)

type Client struct {
	mountebankURI string
	impostersURI  string
}

type Log struct {
	level     string
	message   string
	timestamp string
}

func NewClient(mountebankURI string) *Client {
	return &Client{
		mountebankURI: mountebankURI,
		impostersURI:  mountebankURI + "/imposters",
	}
}

func (c *Client) Logs() ([]Log, error) {

	_, body, errs := gorequest.New().Get(c.mountebankURI + "/logs").EndBytes()
	if len(errs) > 0 {
		return nil, errs[0]
	}

	var logsStructure map[string][]map[string]string
	json.NewDecoder(strings.NewReader(string(body))).Decode(&logsStructure)

	var logs []Log
	for _, log := range logsStructure["logs"] {
		logs = append(logs, Log{level: log["level"], message: log["message"], timestamp: log["timestamp"]})
	}

	return logs, nil
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
