package main

import (
	"flag"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
	"net/http"
	"os/exec"
)

var config struct {
	GithubUsername  string
	GithubAuthToken string
}

const (
	BuildCloning = iota
	BuildReadingTravis
	BuildExecutingBeforeInstall
	BuildExecutingInstall
	BuildExecutingScript
)

type Build struct {
	Status  BuildStatus
	travis  *TravisYml
	workDir string
}

func (b *Build) runCmd(cmdArgs string) {
	cmd := exec.Command("bash", "-c", cmdArgs)

	cmd.Dir = b.workDir
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		fmt.Println(err)
		return err
	}
}

func (b *Build) Clone() bool {
	b.Status = BuildClone

	return true
}

func (b *Build) ReadTravis() bool {
	b.Status = BuildReadingTravis

	return true
}

func (b *Build) ExecuteTravis() bool {
	for i := range travis.BeforeInstall {
		b.Status = BuildExecutingBeforeInstall
		b.runCmd(i)
	}

	for i := range travis.Install {
		b.Status = BuildExecutingInstall
		b.runCmd(i)
	}

	for i := range travis.Script {
		b.Status = BuildExecutingScript
		b.runCmd(i)
	}
}

func newPullRequestHandler(c web.C, w http.ResponseWriter, r *http.Request) {
	// Git clone shallow on the Pull Request
	// Open the travis.yml file

	http.ServeFile(w, r, "public/index.html")
}

func init() {
	flag.StringVar(&config.GithubUsername, "github-username", "", "Github Username")
	flag.StringVar(&config.GithubAuthToken, "github-auth-token", "", "Github Auth Token")

	// TODO: Check for Git binary
}

func main() {
	flags.Parse()

	// TODO: Install the netrc

	goji.Post("/pull_requests", newPullRequestHandler)

}
