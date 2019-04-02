package main

import (
	"bytes"
	"errors"

	"github.com/BurntSushi/toml"
)

var (
	ErrSkipped = errors.New("skipped")
)

func readConfig() (*Config, error) {
	var config Config
	_, err := toml.DecodeFile("Repositories.toml", &config)
	return &config, err
}

func url(vcs, baseUrl, repo string) string {
	var buf bytes.Buffer
	// TODO: filepath.Join() transforms "ssh://stuff" into "ssh:/stuff". Why?
	buf.WriteString(baseUrl)
	buf.WriteString(repo)
	if vcs == "git" {
		buf.WriteString(".git")
	}
	return buf.String()
}
