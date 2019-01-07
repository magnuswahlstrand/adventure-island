// +build js

package conf

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/apex/log"
	"honnef.co/go/js/dom"
)

var document = dom.GetWindow().Document().(dom.HTMLDocument)

func Conf() (addr string, worldName string, dev bool, secure bool, dummy bool) {

	u, err := url.Parse(document.DocumentURI())
	if err != nil {
		log.Errorf("unexpected error parsing URI: %s", err)
		return
	}

	if u.Query().Get("addr") != "" {
		addr = u.Query().Get("addr")

		if !strings.Contains(addr, "http") {
			addr = "https://" + addr
		}
	}
	fmt.Println("addr:", addr)

	worldName = u.Query().Get("world")
	fmt.Println("world:", worldName)

	switch u.Query().Get("dummy") {
	case "true", "1":
		dummy = true
	default:
		dummy = false
	}
	fmt.Println("dummy:", dummy)

	dev = false

	switch u.Query().Get("secure") {
	case "true", "1":
		secure = true
	default:
		secure = false
	}
	fmt.Println("TLS (secure):", secure)
	return
}
