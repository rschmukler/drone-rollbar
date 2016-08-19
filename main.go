package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/drone/drone-go/drone"
	"github.com/drone/drone-go/plugin"
)

const ROLLBAR_URL = "https://api.rollbar.com/api/1/deploy/"

type DroneConfig struct {
	AccessToken string `json:"access_token"`
	Environment string `json:"environment"`
}

type RollbarArgs struct {
	Comment       string `json:"comment"`
	AccessToken   string `json:"access_token"`
	Environment   string `json:"environment"`
	Revision      string `json:"revision"`
	LocalUsername string `json:"local_username"`
}

func main() {
	var (
		repo  = new(drone.Repo)
		build = new(drone.Build)
		sys   = new(drone.System)
		cfg   = new(DroneConfig)
	)

	plugin.Param("build", build)
	plugin.Param("repo", repo)
	plugin.Param("system", sys)
	plugin.Param("vargs", cfg)

	err := plugin.Parse()
	exitIfErr(err)

	payload := &RollbarArgs{
		AccessToken:   cfg.AccessToken,
		Environment:   cfg.Environment,
		Revision:      build.Commit,
		LocalUsername: build.Author,
		Comment:       build.Message,
	}

	body, err := json.Marshal(payload)
	exitIfErr(err)

	req, err := http.NewRequest("POST", ROLLBAR_URL, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Do(req)
	defer resp.Body.Close()
	exitIfErr(err)
}

func exitIfErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
