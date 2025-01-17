package base

import (
	"net"
	"strings"
	"testing"
)

func TestParseVCSUrl(t *testing.T) {
	repos := []string{
		// ssh://[user@]host.xz[:port]/path/to/repo.git/
		"ssh://git@github.com:7875/go-xt/xt.git",
		// git://host.xz[:port]/path/to/repo.git/
		"git://github.com:7875/go-xt/xt.git",
		// http[s]://host.xz[:port]/path/to/repo.git/
		"https://github.com:7875/go-xt/xt.git",
		// ftp[s]://host.xz[:port]/path/to/repo.git/
		"ftps://github.com:7875/go-xt/xt.git",
		//[user@]host.xz:path/to/repo.git/
		"git@github.com:go-xt/xt.git",
		// ssh://[user@]host.xz[:port]/~[user]/path/to/repo.git/
		"ssh://git@github.com:7875/go-xt/xt.git",
		// git://host.xz[:port]/~[user]/path/to/repo.git/
		"git://github.com:7875/go-xt/xt.git",
		//[user@]host.xz:/~[user]/path/to/repo.git/
		"git@github.com:go-xt/xt.git",
		///path/to/repo.git/
		"~/go-xt/xt.git",
		// file:///path/to/repo.git/
		"file://~/go-xt/xt.git",
	}
	for _, repo := range repos {
		url, err := ParseVCSUrl(repo)
		if err != nil {
			t.Fatal(repo, err)
		}
		urlPath := strings.TrimLeft(url.Path, "/")
		if urlPath != "go-xt/xt.git" {
			t.Fatal(repo, "parse url failed", urlPath)
		}
	}
}

func TestParseSsh(t *testing.T) {
	repo := "ssh://git@github.com:7875/go-xt/xt.git"
	url, err := ParseVCSUrl(repo)
	if err != nil {
		t.Fatal(err)
	}
	host, _, err := net.SplitHostPort(url.Host)
	if err != nil {
		host = url.Host
	}
	t.Log(host, url.Path)
}
