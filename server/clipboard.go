package server

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/atotto/clipboard"
	"github.com/lemonade-command/lemonade/lemon"
)

type Clipboard struct{}

func (_ *Clipboard) Copy(w http.ResponseWriter, r *http.Request) {
	b, _ := ioutil.ReadAll(r.Body)
	clipboard.WriteAll(lemon.ConvertLineEnding(string(b), LineEndingOpt))
}

func (_ *Clipboard) Paste(w http.ResponseWriter, r *http.Request) {
	t, _ := clipboard.ReadAll()
	fmt.Fprint(w, t)
}
