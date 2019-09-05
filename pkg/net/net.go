// Package net provides simple networking helper functions
package net

import (
	"io"
	"net"
	"net/http"
	"net/url"
	"os"
	"path"
	"regexp"
	"strings"
	"time"

	"github.com/phR0ze/go-errors"
	"github.com/phR0ze/n/pkg/sys"
)

const (
	// TCP is a protocol string constant
	TCP = "tcp"

	// UDP is a protocol string constant
	UDP = "udp"
)

var (
	gRXURLGetSchema     = regexp.MustCompile(`(?m)^(.*//).*$`)
	gRXURLReplaceSchema = regexp.MustCompile(`(?m)^.*//`)
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
	var fw *os.File
	flags := os.O_CREATE | os.O_TRUNC | os.O_WRONLY
	if fw, err = os.OpenFile(result, flags, os.FileMode(perm)); err != nil {

	}
	defer func() {
		if e := fw.Close(); e != nil {
			if err == nil {
				err = e
			}
			err = errors.Wrap(err, "failed to close file writer")
		}
	}()

	// Download the file to memory via GET
	var res *http.Response
	if res, err = http.Get(url); err != nil {
		return
	}
	defer func() {
		if e := res.Body.Close(); e != nil {
			if err == nil {
				err = e
			}
			err = errors.Wrap(err, "failed to close http response")
		}
	}()

	// Write streamed http bits to the file
	_, err = io.Copy(fw, res.Body)

	return
}

// DirURL behaves much like the path.Dir only it doesn't garble the schema
func DirURL(uri string) (result string) {
	return sys.SlicePath(uri, 0, -2)
}

// JoinURL will join and append the given paths to the first element which
// is assumed will be the schema of the URL. Will also normalize the schema.
func JoinURL(elems ...string) (result string) {

	// Drop empty strings
	for i := len(elems) - 1; i >= 0; i-- {
		if elems[i] == "" {
			if i+1 < len(elems) {
				elems = append(elems[:i], elems[i+1:]...)
			} else {
				elems = elems[:i]
			}
		}
	}

	// Normalize schema and join with / skipping the path.Join's Clean call
	if len(elems) > 0 {
		elems[0] = NormalizeURL(elems[0])

		// Drop path absolute slash
		if len(elems) > 1 && strings.HasPrefix(elems[1], "/") {
			elems[1] = elems[1][1:]
		}
		result = strings.Join(elems, "/")
	}
	return
}

// NormalizeURL parses the URL and corrects case and other oddities.
// 1. properly formats the schema
func NormalizeURL(uri string) (result string) {
	if matches := gRXURLGetSchema.FindStringSubmatch(uri); len(matches) > 1 {
		schema := strings.ToLower(matches[1])
		result = gRXURLReplaceSchema.ReplaceAllString(uri, schema)
	}
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

// SplitURL will split the URL into pieces much as strings.Split(url, "/") would
// except it keeps the schema intact if it exists.
func SplitURL(uri string) (pieces []string) {
	uri = NormalizeURL(uri)
	if matches := gRXURLGetSchema.FindStringSubmatch(uri); len(matches) > 1 {
		schema := matches[1]
		uri = strings.Replace(uri, schema, "", 1)
		pieces = strings.Split(uri, "/")
		pieces = append([]string{schema}, pieces...)
	} else {
		pieces = strings.Split(uri, "/")
	}
	return
}
