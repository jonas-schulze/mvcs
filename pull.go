package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"sync"

	"github.com/urfave/cli"
)

func pullRepositories(c *cli.Context) error {
	config, err := readConfig()
	if err != nil {
		return err
	}

	var wg sync.WaitGroup
	for _, remote := range config.Remotes {
		wg.Add(len(remote.Repositories))
		for _, repo := range remote.Repositories {
			dir := filepath.Base(repo)
			go func() {
				pullRepository(config.VCS, dir)
				wg.Done()
			}()
		}
	}

	wg.Wait()
	return nil
}

func pullRepository(vcs, dir string) {
	cmd := exec.Command(vcs, "pull")
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
