# Lessons learnt

[<- Back to index](README.md)

## 1. Structures for handling byte arrays in a simple way

Egon recommended this simple setup, after I complained about that byte slices were cumbersome to work with:

```go
package main

type Coord struct{ X, Y int }

type Tile byte

const (
	Invalid Tile = iota
	Grass
	Water
)

type Map struct {
	tiles  []Tile
	width  int
	height int
}

func (m *Map) At(p Coord) Tile {
	// do bounds checks here
	return m.tiles[p.Y*m.width+p.X]
}

func (m *Map) At(p Coord, t Tile) {
	// do bounds checks here
	m.tiles[p.Y*m.width+p.X] = tile
}

```

## 2. Empty message for gRPC

There is a ready made [Empty](https://github.com/protocolbuffers/protobuf/blob/master/src/google/protobuf/empty.proto) message for gRPC that can be used to empty return values or no parameters in gRPC services.

```protobuf
import "google/protobuf/empty.proto";


service Backend {
    rpc NewPlayer(google.protobuf.Empty) returns (PlayerID) {}
    rpc EntityStream(google.protobuf.Empty) returns (EntityResponse) {}
}

```

## 3. Simple TLS python http server

```
# taken from http://www.piware.de/2011/01/creating-an-https-server-in-python/
# generate server.xml with the following command:
#    openssl req -new -x509 -keyout server.pem -out server.pem -days 365 -nodes
# run as follows:
#    python simple-https-server.py
# then in your browser, visit:
#    https://localhost:4443

import BaseHTTPServer, SimpleHTTPServer
import ssl

httpd = BaseHTTPServer.HTTPServer(('localhost', 4443), SimpleHTTPServer.SimpleHTTPRequestHandler)
httpd.socket = ssl.wrap_socket (httpd.socket, certfile='./server.pem', server_side=True)
httpd.serve_forever()
```

# TLS Certificates for localhost

Peter Hellberg introduced me to this [blog post](https://blog.filippo.io/mkcert-valid-https-certificates-for-localhost/) on using `mkcert`to generate certificates for developing of HTTPS on localhost. This paragraph summarizes the problem:

> The web is moving to HTTPS, preventing network attackers from observing or injecting page contents. But HTTPS needs TLS certificates, and while deployment is increasingly a solved issue thanks to the ACME protocol and Let's Encrypt, development still mostly ends up happening over HTTP because no one can get an universally valid certificate for localhost.

The tool [mkcert](https://github.com/FiloSottile/mkcert) generates a certificate that can be used like this

```golang
package main

import (
	"fmt"
	"net/http"
)

func main() {
	hs := &http.Server{
		Addr:    ":4443",
		Handler: http.HandlerFunc(hello),
	}

	hs.ListenAndServeTLS("localhost+1.pem", "localhost+1-key.pem")
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello TLS world!")
}
```

## Allow insecure on localhost

There is a flag to allow HTTPS even though invalid certificate is presented.
`chrome://flags/#allow-insecure-localhost`

> Allows requests to localhost over HTTPS even when an invalid certificate is presented. – Mac, Windows, Linux, Chrome OS, Android
