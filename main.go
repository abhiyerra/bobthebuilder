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

	http.ServeFile(w, r, "public/index.html")
}

func init() {
	flags.StringVar(&config.GithubUsername, "github-username", "", "Github Username")
	flags.StringVar(&config.GithubAuthToken, "github-auth-token", "", "Github Auth Token")
}

func main() {
	flags.Parse()

	goji.Post("/pull_requests", newPullRequestHandler)

}
