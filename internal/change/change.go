package change

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// CmdChange is xt change log tool
var CmdChange = &cobra.Command{
	Use:   "changelog",
	Short: "Get a xt change log",
	Long:  "Get a xt release or commits info. Example: xt changelog dev or xt changelog {version}",
	Run:   run,
}

var (
	token   string
	repoURL string
)

func init() {
	if repoURL = os.Getenv("XT_REPO"); repoURL == "" {
		repoURL = "https://github.com/go-xt/xt.git"
	}
	CmdChange.Flags().StringVarP(&repoURL, "repo-url", "r", repoURL, "github repo")
	token = os.Getenv("GITHUB_TOKEN")
}

func run(_ *cobra.Command, args []string) {
	owner, repo := ParseGithubURL(repoURL)
	api := GithubAPI{Owner: owner, Repo: repo, Token: token}
	version := "latest"
	if len(args) > 0 {
		version = args[0]
	}
	if version == "dev" {
		info := api.GetCommitsInfo()
		fmt.Print(ParseCommitsInfo(info))
		return
	}
	info := api.GetReleaseInfo(version)
	fmt.Print(ParseReleaseInfo(info))
}
