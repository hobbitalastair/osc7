package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestEncodePath(t *testing.T) {
	tests := map[string]string{
		"/home/ash":              "/home/ash",
		"/home/ash/has spaces":   "/home/ash/has%20spaces",
		"/tmp/percent%literal":   "/tmp/percent%25literal",
		"/tmp/query?fragment#x":  "/tmp/query%3Ffragment%23x",
		"/tmp/control\nchar\x07": "/tmp/control%0Achar%07",
	}

	for input, want := range tests {
		if got := encodePath(input); got != want {
			t.Fatalf("encodePath(%q) = %q, want %q", input, got, want)
		}
	}
}

func TestRunWithExplicitDirectory(t *testing.T) {
	var out bytes.Buffer

	if err := run([]string{"/tmp/has spaces"}, &out); err != nil {
		t.Fatal(err)
	}

	got := out.String()
	if !strings.HasPrefix(got, "\x1b]7;file://") {
		t.Fatalf("output prefix = %q", got)
	}
	if !strings.HasSuffix(got, "/tmp/has%20spaces\x07") {
		t.Fatalf("output suffix = %q", got)
	}
	if strings.HasSuffix(got, "\n") {
		t.Fatal("output unexpectedly ended with newline")
	}
}

func TestRunRejectsRelativeDirectory(t *testing.T) {
	var out bytes.Buffer

	if err := run([]string{"relative"}, &out); err == nil {
		t.Fatal("expected relative path to be rejected")
	}
}

func TestRunRejectsExtraArguments(t *testing.T) {
	var out bytes.Buffer

	if err := run([]string{"/tmp", "/var"}, &out); err == nil {
		t.Fatal("expected extra arguments to be rejected")
	}
}

func TestValidHost(t *testing.T) {
	if err := validHost("local-host"); err != nil {
		t.Fatal(err)
	}

	for _, host := range []string{"", "bad/host", "bad;host", "bad%host", "bad\nhost"} {
		if err := validHost(host); err == nil {
			t.Fatalf("expected host %q to be rejected", host)
		}
	}
}
