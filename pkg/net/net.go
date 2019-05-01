// Package net provides simple networking helper functions
package net

import (
	"io"
	"net"
	"net/http"
	"net/url"
	"os"
	"path"
	"time"

	"github.com/phR0ze/n/pkg/sys"
)

const (
	// TCP is a protocol string constant
	TCP = "tcp"

	// UDP is a protocol string constant
	UDP = "udp"
)

// DownloadFile from the given URL to the given destination
// returning the full path to the resulting downloaded file
func DownloadFile(url, dst string, perms ...uint32) (result string, err error) {
	if result, err = sys.Abs(dst); err != nil {
		return
	}

	perm := uint32(0644)
	if len(perms) > 0 {
		perm = perms[0]
	}

	// Create the destination path if it doesn't exist
	sys.MkdirP(path.Dir(result))

	// Open destination truncating if it exists
	var f *os.File
	flags := os.O_CREATE | os.O_TRUNC | os.O_WRONLY
	if f, err = os.OpenFile(result, flags, os.FileMode(perm)); err != nil {

	}
	defer f.Close()

	// Download the file to memory via GET
	var res *http.Response
	if res, err = http.Get(url); err != nil {
		return
	}
	defer res.Body.Close()

	// Write streamed http bits to the file
	_, err = io.Copy(f, res.Body)

	return
}

// Ping simply checks if the given protocol, address is accessible
// and listening. An error will be returned if the ping was not successful.
// optional timeout in seconds defaults to 1
func Ping(proto, addr string, timeout ...int) (err error) {
	_timeout := 1
	if len(timeout) > 0 {
		_timeout = timeout[0]
	}

	var conn net.Conn
	dialer := net.Dialer{Timeout: time.Duration(_timeout) * time.Second}
	if conn, err = dialer.Dial(proto, addr); err == nil {
		conn.Close()
		return
	}
	return
}

// DisableProxy unsets the http_proxy env var and sets the http.DefaultTransport to not use a proxy
func DisableProxy(proxy *url.URL) {
	if proxy != nil {
		os.Unsetenv("http_proxy")
		http.DefaultTransport = &http.Transport{Proxy: nil}
	}
}

// EnableProxy sets the http_proxy env var and sets the http.DefaultTransport to use a proxy
func EnableProxy(proxy *url.URL) {
	if proxy != nil {
		os.Setenv("http_proxy", proxy.String())
		http.DefaultTransport = &http.Transport{Proxy: http.ProxyURL(proxy)}
	}
}
