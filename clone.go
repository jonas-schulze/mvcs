package main

import (
	"os"
	"os/exec"
	"sync"

	"github.com/urfave/cli"
)

func cloneRepositories(c *cli.Context) error {
	config, err := readConfig()
	if err != nil {
		return err
	}

	var wg sync.WaitGroup
	for _, remote := range config.Remotes {
		wg.Add(len(remote.Repositories))
		for _, repo := range remote.Repositories {
			url := url(config.VCS, remote.BaseUrl, repo)
			go func() {
				cloneRepository(config.VCS, url)
				wg.Done()
			}()
		}
	}

	wg.Wait()
	return nil
}

func cloneRepository(vcs string, url string) {
	cmd := exec.Command(vcs, "clone", url)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
