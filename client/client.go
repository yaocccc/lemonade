package client

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	log "github.com/inconshreveable/log15"

	"github.com/lemonade-command/lemonade/lemon"
)

type client struct {
	host       string
	port       int
	lineEnding string
	logger     log.Logger
	timeout    time.Duration
}

func New(c *lemon.CLI, logger log.Logger) *client {
	return &client{
		host:       c.Host,
		port:       c.Port,
		lineEnding: c.LineEnding,
		logger:     logger,
	}
}

func (c *client) Paste() (string, error) {
	request, _ := http.NewRequest("GET", fmt.Sprintf("http://%s:%d/paste", c.host, c.port), nil)
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return "", err
	}
	b, _ := ioutil.ReadAll(resp.Body)
	return lemon.ConvertLineEnding(string(b), c.lineEnding), nil
}

func (c *client) Copy(text string) error {
	c.logger.Debug("Sending: " + text)
	request, _ := http.NewRequest("POST", fmt.Sprintf("http://%s:%d/copy", c.host, c.port), bytes.NewReader([]byte(text)))
	_, err := http.DefaultClient.Do(request)
	return err
}
