package server

import (
	"fmt"
	"net/http"

	log "github.com/inconshreveable/log15"

	"github.com/lemonade-command/lemonade/lemon"
)

var LineEndingOpt string

func Serve(c *lemon.CLI, logger log.Logger) {
	LineEndingOpt = c.LineEnding
	clipboard := &Clipboard{}
	http.HandleFunc("/paste", clipboard.Paste)
	http.HandleFunc("/copy", clipboard.Copy)
	http.ListenAndServe(fmt.Sprintf(":%d", c.Port), nil)
}
