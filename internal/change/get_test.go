package change

import "testing"

func TestParseGithubURL(t *testing.T) {
	urls := []struct {
		url   string
		owner string
		repo  string
	}{
		{"https://github.com/go-xt/xt.git", "go-xt", "xt"},
		{"https://github.com/go-xt/xt", "go-xt", "xt"},
		{"git@github.com:go-xt/xt.git", "go-xt", "xt"},
		{"https://github.com/go-xt/go-xt.dev.git", "go-xt", "go-xt.dev"},
	}
	for _, url := range urls {
		owner, repo := ParseGithubURL(url.url)
		if owner != url.owner {
			t.Fatalf("owner want: %s, got: %s", owner, url.owner)
		}
		if repo != url.repo {
			t.Fatalf("repo want: %s, got: %s", repo, url.repo)
		}
	}
}
