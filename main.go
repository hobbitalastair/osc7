package main

import (
	"errors"
	"fmt"
	"net/url"
	"os"
	"strings"
)

const (
	esc = "\x1b"
	bel = "\x07"
)

func main() {
	if err := run(os.Args[1:], os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "osc7: %v\n", err)
		os.Exit(1)
	}
}

func run(args []string, out interface{ Write([]byte) (int, error) }) error {
	if len(args) > 1 {
		return errors.New("usage: osc7 [directory]")
	}

	host, err := os.Hostname()
	if err != nil {
		return err
	}
	if err := validHost(host); err != nil {
		return err
	}

	path := ""
	if len(args) == 1 {
		path = args[0]
	} else {
		path, err = os.Getwd()
		if err != nil {
			return err
		}
	}

	if !strings.HasPrefix(path, "/") {
		return errors.New("directory must be an absolute path")
	}

	_, err = fmt.Fprintf(out, "%s]7;file://%s%s%s", esc, host, encodePath(path), bel)
	return err
}

func validHost(host string) error {
	if host == "" {
		return errors.New("hostname is empty")
	}

	for _, r := range host {
		if r <= 0x20 || r == 0x7f || r == '/' || r == ';' || r == '%' {
			return fmt.Errorf("hostname contains unsupported character %q", r)
		}
	}

	return nil
}

func encodePath(path string) string {
	u := url.URL{Path: path}
	return u.EscapedPath()
}
