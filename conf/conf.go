// +build !js

package conf

import "flag"

func Conf() (addr string, worldName string, dev bool, secure bool, dummy bool) {
	flag.BoolVar(&dummy, "dummy", false, "create a dummy player who walks around randomly, mostly for development purposes")
	flag.BoolVar(&dev, "dev", false, "start the development server on local machine")
	flag.BoolVar(&dev, "secure", false, "enable TLS")
	flag.StringVar(&addr, "addr", "", "address to remote server: default: run local mode")
	flag.StringVar(&worldName, "world", "", "name of the world to play on")
	flag.Parse()
	return
}
