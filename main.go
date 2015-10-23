package main

import (
	"flags"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
	"net/http"
)

var config struct {
	GithubUsername  string
	GithubAuthToken string
}

func newPullRequestHandler(c web.C, w http.ResponseWriter, r *http.Request) {
	// Git clone shallow on the Pull Request
	// Open the travis.yml file

	http.ServeFile(w, r, "public/index.html")
}

func init() {
	flags.StringVar(&config.GithubUsername, "github-username", "", "Github Username")
	flags.StringVar(&config.GithubAuthToken, "github-auth-token", "", "Github Auth Token")

	// TODO: Check for Git binary
}

func main() {
	flags.Parse()

	// TODO: Install the netrc

	goji.Post("/pull_requests", newPullRequestHandler)

}
