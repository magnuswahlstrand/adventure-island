// +build !js

package conf

func Conf(addr string, worldName string, dev bool, secure bool, dummy bool) (string, string, bool, bool, bool) {
	return addr, worldName, dev, secure, dummy
}
